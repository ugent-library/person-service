// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ugent-library/people/ent/organization"
	"github.com/ugent-library/people/ent/person"
)

// OrganizationCreate is the builder for creating a Organization entity.
type OrganizationCreate struct {
	config
	mutation *OrganizationMutation
	hooks    []Hook
}

// SetPrimaryID sets the "primary_id" field.
func (oc *OrganizationCreate) SetPrimaryID(s string) *OrganizationCreate {
	oc.mutation.SetPrimaryID(s)
	return oc
}

// SetName sets the "name" field.
func (oc *OrganizationCreate) SetName(s string) *OrganizationCreate {
	oc.mutation.SetName(s)
	return oc
}

// AddPersonIDs adds the "people" edge to the Person entity by IDs.
func (oc *OrganizationCreate) AddPersonIDs(ids ...int) *OrganizationCreate {
	oc.mutation.AddPersonIDs(ids...)
	return oc
}

// AddPeople adds the "people" edges to the Person entity.
func (oc *OrganizationCreate) AddPeople(p ...*Person) *OrganizationCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return oc.AddPersonIDs(ids...)
}

// Mutation returns the OrganizationMutation object of the builder.
func (oc *OrganizationCreate) Mutation() *OrganizationMutation {
	return oc.mutation
}

// Save creates the Organization in the database.
func (oc *OrganizationCreate) Save(ctx context.Context) (*Organization, error) {
	return withHooks[*Organization, OrganizationMutation](ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrganizationCreate) SaveX(ctx context.Context) *Organization {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OrganizationCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OrganizationCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OrganizationCreate) check() error {
	if _, ok := oc.mutation.PrimaryID(); !ok {
		return &ValidationError{Name: "primary_id", err: errors.New(`ent: missing required field "Organization.primary_id"`)}
	}
	if _, ok := oc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Organization.name"`)}
	}
	return nil
}

func (oc *OrganizationCreate) sqlSave(ctx context.Context) (*Organization, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OrganizationCreate) createSpec() (*Organization, *sqlgraph.CreateSpec) {
	var (
		_node = &Organization{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(organization.Table, sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt))
	)
	if value, ok := oc.mutation.PrimaryID(); ok {
		_spec.SetField(organization.FieldPrimaryID, field.TypeString, value)
		_node.PrimaryID = value
	}
	if value, ok := oc.mutation.Name(); ok {
		_spec.SetField(organization.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := oc.mutation.PeopleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   organization.PeopleTable,
			Columns: organization.PeoplePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: person.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrganizationCreateBulk is the builder for creating many Organization entities in bulk.
type OrganizationCreateBulk struct {
	config
	builders []*OrganizationCreate
}

// Save creates the Organization entities in the database.
func (ocb *OrganizationCreateBulk) Save(ctx context.Context) ([]*Organization, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Organization, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrganizationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OrganizationCreateBulk) SaveX(ctx context.Context) []*Organization {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OrganizationCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OrganizationCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}
