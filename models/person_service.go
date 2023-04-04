package models

import (
	"context"
	"database/sql"
	"errors"

	entdialect "entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	v1 "github.com/ugent-library/people/api/v1"
	"github.com/ugent-library/people/ent"
	entmigrate "github.com/ugent-library/people/ent/migrate"
	"github.com/ugent-library/people/ent/organization"
	"github.com/ugent-library/people/ent/person"
	"github.com/ugent-library/people/ent/schema"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var ErrNotFound = errors.New("not found")

type PersonConfig struct {
	DB string
}

type PersonService interface {
	Create(context.Context, *Person) (*Person, error)
	Update(context.Context, *Person) (*Person, error)
	Get(context.Context, string) (*Person, error)
	Delete(context.Context, string) error
	Each(context.Context, func(*Person) bool) error
}

type personService struct {
	db *ent.Client
}

func NewPersonService(cfg *PersonConfig) (PersonService, error) {
	db, err := sql.Open("pgx", cfg.DB)
	if err != nil {
		return nil, err
	}

	driver := entsql.OpenDB(entdialect.Postgres, db)
	client := ent.NewClient(ent.Driver(driver))

	err = client.Schema.Create(context.Background(),
		entmigrate.WithDropIndex(true),
	)
	if err != nil {
		return nil, err
	}

	return &personService{
		db: client,
	}, nil
}

func (ps *personService) Create(ctx context.Context, p *Person) (*Person, error) {
	// date fields filled by schema
	t := ps.db.Person.Create()

	// keep in order; copy to Update if it changes
	t.SetActive(p.Active)
	t.SetBirthDate(p.BirthDate)
	t.SetJobCategory(p.JobCategory)
	t.SetEmail(p.Email)
	t.SetFirstName(p.FirstName)
	t.SetFullName(p.FullName)
	t.SetPrimaryID(p.Id) // TODO: nil value overriden by entgo default function?
	t.SetTitle(p.Title)
	t.SetLastName(p.LastName)
	t.SetOrcid(p.Orcid)
	t.SetOrcidToken(p.OrcidToken)
	schemaOtherIds := make([]schema.IdRef, 0, len(p.OtherId))
	for _, refId := range p.OtherId {
		schemaOtherIds = append(schemaOtherIds, schema.IdRef{
			ID:   refId.Id,
			Type: refId.Type,
		})
	}
	t.SetOtherID(schemaOtherIds)
	t.SetPreferredFirstName(p.PreferredFirstName)
	t.SetPreferredLastName(p.PreferredLastName)

	// TODO: test
	if p.OrganizationId != nil && len(p.OrganizationId) > 0 {
		// TODO: crashes with segmentation violation error when org does not exist
		orgs, err := ps.db.Organization.Query().Where(organization.PrimaryIDIn(p.OrganizationId...)).All(ctx)
		if err != nil {
			return nil, err
		}
		t.AddOrganizations(orgs...)
	}

	row, err := t.Save(ctx)
	if err != nil {
		return nil, err
	}

	// collect entgo managed fields
	p.DateCreated = timestamppb.New(row.DateCreated)
	p.DateUpdated = timestamppb.New(row.DateUpdated)
	p.Id = row.PrimaryID

	return p, nil
}

func (ps *personService) Update(ctx context.Context, p *Person) (*Person, error) {
	t := ps.db.Person.Update().Where(person.PrimaryIDEQ(p.Id))

	// keep in order; copy to Update if it changes
	t.SetActive(p.Active)
	t.SetBirthDate(p.BirthDate)
	t.SetJobCategory(p.JobCategory)
	t.SetEmail(p.Email)
	t.SetFirstName(p.FirstName)
	t.SetFullName(p.FullName)
	t.SetTitle(p.Title)
	t.SetLastName(p.LastName)
	t.SetOrcid(p.Orcid)
	t.SetOrcidToken(p.OrcidToken)
	schemaOtherIds := make([]schema.IdRef, 0, len(p.OtherId))
	for _, refId := range p.OtherId {
		schemaOtherIds = append(schemaOtherIds, schema.IdRef{
			ID:   refId.Id,
			Type: refId.Type,
		})
	}
	t.SetOtherID(schemaOtherIds)
	t.SetPreferredFirstName(p.PreferredFirstName)
	t.SetPreferredLastName(p.PreferredLastName)

	t.ClearOrganizations()
	if p.OrganizationId != nil && len(p.OrganizationId) > 0 {
		// TODO: crashes with segmentation violation error when org does not exist
		orgs, err := ps.db.Organization.Query().Where(organization.PrimaryIDIn(p.OrganizationId...)).All(ctx)
		if err != nil {
			return nil, err
		}
		t.AddOrganizations(orgs...)
	}

	_, err := t.Save(ctx)
	if err != nil {
		return nil, err
	}

	return ps.Get(ctx, p.Id)
}

func (ps *personService) Get(ctx context.Context, id string) (*Person, error) {
	row, err := ps.db.Person.Query().WithOrganizations().Where(person.PrimaryIDEQ(id)).First(ctx)
	if err != nil {
		var e *ent.NotFoundError
		if errors.As(err, &e) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return personUnwrap(row), nil
}

func (ps *personService) Delete(ctx context.Context, id string) error {
	_, err := ps.db.Person.Delete().Where(person.PrimaryIDEQ(id)).Exec(ctx)
	return err
}

func (ps *personService) Each(ctx context.Context, cb func(*Person) bool) error {

	// TODO: find a better way to do this (no cursors possible)
	var offset int = 0
	var limit int = 500
	for {
		rows, err := ps.db.Person.Query().WithOrganizations().Offset(offset).Limit(limit).Order(ent.Asc(person.FieldDateCreated)).All(ctx)
		if err != nil {
			return err
		}
		// entgo returns no error on empty results
		if len(rows) == 0 {
			break
		}
		for _, row := range rows {
			if !cb(personUnwrap(row)) {
				return nil
			}
		}
		offset += limit
	}

	return nil
}

func personUnwrap(e *ent.Person) *Person {
	refIds := make([]*v1.IdRef, 0, len(e.OtherID))
	for _, schemaOtherId := range e.OtherID {
		refIds = append(refIds, &v1.IdRef{
			Id:   schemaOtherId.ID,
			Type: schemaOtherId.Type,
		})
	}
	orgIds := make([]string, 0)
	for _, org := range e.Edges.Organizations {
		orgIds = append(orgIds, org.PrimaryID)
	}
	p := &Person{
		Person: v1.Person{
			Active:             e.Active,
			BirthDate:          e.BirthDate,
			DateCreated:        timestamppb.New(e.DateCreated),
			DateUpdated:        timestamppb.New(e.DateUpdated),
			Email:              e.Email,
			OtherId:            refIds,
			FirstName:          e.FirstName,
			FullName:           e.FullName,
			Id:                 e.PrimaryID,
			LastName:           e.LastName,
			JobCategory:        e.JobCategory,
			Orcid:              e.Orcid,
			OrcidToken:         e.OrcidToken,
			OrganizationId:     orgIds,
			PreferredLastName:  e.PreferredLastName,
			PreferredFirstName: e.PreferredFirstName,
			Title:              e.Title,
		},
	}
	return p
}
