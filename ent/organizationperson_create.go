// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ugent-library/people-service/ent/organization"
	"github.com/ugent-library/people-service/ent/organizationperson"
	"github.com/ugent-library/people-service/ent/person"
)

// OrganizationPersonCreate is the builder for creating a OrganizationPerson entity.
type OrganizationPersonCreate struct {
	config
	mutation *OrganizationPersonMutation
	hooks    []Hook
}

// SetDateCreated sets the "date_created" field.
func (opc *OrganizationPersonCreate) SetDateCreated(t time.Time) *OrganizationPersonCreate {
	opc.mutation.SetDateCreated(t)
	return opc
}

// SetNillableDateCreated sets the "date_created" field if the given value is not nil.
func (opc *OrganizationPersonCreate) SetNillableDateCreated(t *time.Time) *OrganizationPersonCreate {
	if t != nil {
		opc.SetDateCreated(*t)
	}
	return opc
}

// SetDateUpdated sets the "date_updated" field.
func (opc *OrganizationPersonCreate) SetDateUpdated(t time.Time) *OrganizationPersonCreate {
	opc.mutation.SetDateUpdated(t)
	return opc
}

// SetNillableDateUpdated sets the "date_updated" field if the given value is not nil.
func (opc *OrganizationPersonCreate) SetNillableDateUpdated(t *time.Time) *OrganizationPersonCreate {
	if t != nil {
		opc.SetDateUpdated(*t)
	}
	return opc
}

// SetOrganizationID sets the "organization_id" field.
func (opc *OrganizationPersonCreate) SetOrganizationID(i int) *OrganizationPersonCreate {
	opc.mutation.SetOrganizationID(i)
	return opc
}

// SetPersonID sets the "person_id" field.
func (opc *OrganizationPersonCreate) SetPersonID(i int) *OrganizationPersonCreate {
	opc.mutation.SetPersonID(i)
	return opc
}

// SetFrom sets the "from" field.
func (opc *OrganizationPersonCreate) SetFrom(t time.Time) *OrganizationPersonCreate {
	opc.mutation.SetFrom(t)
	return opc
}

// SetNillableFrom sets the "from" field if the given value is not nil.
func (opc *OrganizationPersonCreate) SetNillableFrom(t *time.Time) *OrganizationPersonCreate {
	if t != nil {
		opc.SetFrom(*t)
	}
	return opc
}

// SetUntil sets the "until" field.
func (opc *OrganizationPersonCreate) SetUntil(t time.Time) *OrganizationPersonCreate {
	opc.mutation.SetUntil(t)
	return opc
}

// SetNillableUntil sets the "until" field if the given value is not nil.
func (opc *OrganizationPersonCreate) SetNillableUntil(t *time.Time) *OrganizationPersonCreate {
	if t != nil {
		opc.SetUntil(*t)
	}
	return opc
}

// SetPeopleID sets the "people" edge to the Person entity by ID.
func (opc *OrganizationPersonCreate) SetPeopleID(id int) *OrganizationPersonCreate {
	opc.mutation.SetPeopleID(id)
	return opc
}

// SetPeople sets the "people" edge to the Person entity.
func (opc *OrganizationPersonCreate) SetPeople(p *Person) *OrganizationPersonCreate {
	return opc.SetPeopleID(p.ID)
}

// SetOrganizationsID sets the "organizations" edge to the Organization entity by ID.
func (opc *OrganizationPersonCreate) SetOrganizationsID(id int) *OrganizationPersonCreate {
	opc.mutation.SetOrganizationsID(id)
	return opc
}

// SetOrganizations sets the "organizations" edge to the Organization entity.
func (opc *OrganizationPersonCreate) SetOrganizations(o *Organization) *OrganizationPersonCreate {
	return opc.SetOrganizationsID(o.ID)
}

// Mutation returns the OrganizationPersonMutation object of the builder.
func (opc *OrganizationPersonCreate) Mutation() *OrganizationPersonMutation {
	return opc.mutation
}

// Save creates the OrganizationPerson in the database.
func (opc *OrganizationPersonCreate) Save(ctx context.Context) (*OrganizationPerson, error) {
	opc.defaults()
	return withHooks[*OrganizationPerson, OrganizationPersonMutation](ctx, opc.sqlSave, opc.mutation, opc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (opc *OrganizationPersonCreate) SaveX(ctx context.Context) *OrganizationPerson {
	v, err := opc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (opc *OrganizationPersonCreate) Exec(ctx context.Context) error {
	_, err := opc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opc *OrganizationPersonCreate) ExecX(ctx context.Context) {
	if err := opc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (opc *OrganizationPersonCreate) defaults() {
	if _, ok := opc.mutation.DateCreated(); !ok {
		v := organizationperson.DefaultDateCreated()
		opc.mutation.SetDateCreated(v)
	}
	if _, ok := opc.mutation.DateUpdated(); !ok {
		v := organizationperson.DefaultDateUpdated()
		opc.mutation.SetDateUpdated(v)
	}
	if _, ok := opc.mutation.From(); !ok {
		v := organizationperson.DefaultFrom()
		opc.mutation.SetFrom(v)
	}
	if _, ok := opc.mutation.Until(); !ok {
		v := organizationperson.DefaultUntil()
		opc.mutation.SetUntil(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (opc *OrganizationPersonCreate) check() error {
	if _, ok := opc.mutation.DateCreated(); !ok {
		return &ValidationError{Name: "date_created", err: errors.New(`ent: missing required field "OrganizationPerson.date_created"`)}
	}
	if _, ok := opc.mutation.DateUpdated(); !ok {
		return &ValidationError{Name: "date_updated", err: errors.New(`ent: missing required field "OrganizationPerson.date_updated"`)}
	}
	if _, ok := opc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization_id", err: errors.New(`ent: missing required field "OrganizationPerson.organization_id"`)}
	}
	if _, ok := opc.mutation.PersonID(); !ok {
		return &ValidationError{Name: "person_id", err: errors.New(`ent: missing required field "OrganizationPerson.person_id"`)}
	}
	if _, ok := opc.mutation.From(); !ok {
		return &ValidationError{Name: "from", err: errors.New(`ent: missing required field "OrganizationPerson.from"`)}
	}
	if _, ok := opc.mutation.Until(); !ok {
		return &ValidationError{Name: "until", err: errors.New(`ent: missing required field "OrganizationPerson.until"`)}
	}
	if _, ok := opc.mutation.PeopleID(); !ok {
		return &ValidationError{Name: "people", err: errors.New(`ent: missing required edge "OrganizationPerson.people"`)}
	}
	if _, ok := opc.mutation.OrganizationsID(); !ok {
		return &ValidationError{Name: "organizations", err: errors.New(`ent: missing required edge "OrganizationPerson.organizations"`)}
	}
	return nil
}

func (opc *OrganizationPersonCreate) sqlSave(ctx context.Context) (*OrganizationPerson, error) {
	if err := opc.check(); err != nil {
		return nil, err
	}
	_node, _spec := opc.createSpec()
	if err := sqlgraph.CreateNode(ctx, opc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	opc.mutation.id = &_node.ID
	opc.mutation.done = true
	return _node, nil
}

func (opc *OrganizationPersonCreate) createSpec() (*OrganizationPerson, *sqlgraph.CreateSpec) {
	var (
		_node = &OrganizationPerson{config: opc.config}
		_spec = sqlgraph.NewCreateSpec(organizationperson.Table, sqlgraph.NewFieldSpec(organizationperson.FieldID, field.TypeInt))
	)
	if value, ok := opc.mutation.DateCreated(); ok {
		_spec.SetField(organizationperson.FieldDateCreated, field.TypeTime, value)
		_node.DateCreated = value
	}
	if value, ok := opc.mutation.DateUpdated(); ok {
		_spec.SetField(organizationperson.FieldDateUpdated, field.TypeTime, value)
		_node.DateUpdated = value
	}
	if value, ok := opc.mutation.From(); ok {
		_spec.SetField(organizationperson.FieldFrom, field.TypeTime, value)
		_node.From = value
	}
	if value, ok := opc.mutation.Until(); ok {
		_spec.SetField(organizationperson.FieldUntil, field.TypeTime, value)
		_node.Until = value
	}
	if nodes := opc.mutation.PeopleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   organizationperson.PeopleTable,
			Columns: []string{organizationperson.PeopleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PersonID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := opc.mutation.OrganizationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   organizationperson.OrganizationsTable,
			Columns: []string{organizationperson.OrganizationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OrganizationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrganizationPersonCreateBulk is the builder for creating many OrganizationPerson entities in bulk.
type OrganizationPersonCreateBulk struct {
	config
	builders []*OrganizationPersonCreate
}

// Save creates the OrganizationPerson entities in the database.
func (opcb *OrganizationPersonCreateBulk) Save(ctx context.Context) ([]*OrganizationPerson, error) {
	specs := make([]*sqlgraph.CreateSpec, len(opcb.builders))
	nodes := make([]*OrganizationPerson, len(opcb.builders))
	mutators := make([]Mutator, len(opcb.builders))
	for i := range opcb.builders {
		func(i int, root context.Context) {
			builder := opcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrganizationPersonMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, opcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, opcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, opcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (opcb *OrganizationPersonCreateBulk) SaveX(ctx context.Context) []*OrganizationPerson {
	v, err := opcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (opcb *OrganizationPersonCreateBulk) Exec(ctx context.Context) error {
	_, err := opcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opcb *OrganizationPersonCreateBulk) ExecX(ctx context.Context) {
	if err := opcb.Exec(ctx); err != nil {
		panic(err)
	}
}
