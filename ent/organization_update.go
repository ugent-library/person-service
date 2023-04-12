// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/ugent-library/people/ent/organization"
	"github.com/ugent-library/people/ent/organizationperson"
	"github.com/ugent-library/people/ent/person"
	"github.com/ugent-library/people/ent/predicate"
	"github.com/ugent-library/people/ent/schema"
)

// OrganizationUpdate is the builder for updating Organization entities.
type OrganizationUpdate struct {
	config
	hooks    []Hook
	mutation *OrganizationMutation
}

// Where appends a list predicates to the OrganizationUpdate builder.
func (ou *OrganizationUpdate) Where(ps ...predicate.Organization) *OrganizationUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetDateUpdated sets the "date_updated" field.
func (ou *OrganizationUpdate) SetDateUpdated(t time.Time) *OrganizationUpdate {
	ou.mutation.SetDateUpdated(t)
	return ou
}

// SetType sets the "type" field.
func (ou *OrganizationUpdate) SetType(s string) *OrganizationUpdate {
	ou.mutation.SetType(s)
	return ou
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ou *OrganizationUpdate) SetNillableType(s *string) *OrganizationUpdate {
	if s != nil {
		ou.SetType(*s)
	}
	return ou
}

// SetNameDut sets the "name_dut" field.
func (ou *OrganizationUpdate) SetNameDut(s string) *OrganizationUpdate {
	ou.mutation.SetNameDut(s)
	return ou
}

// SetNillableNameDut sets the "name_dut" field if the given value is not nil.
func (ou *OrganizationUpdate) SetNillableNameDut(s *string) *OrganizationUpdate {
	if s != nil {
		ou.SetNameDut(*s)
	}
	return ou
}

// ClearNameDut clears the value of the "name_dut" field.
func (ou *OrganizationUpdate) ClearNameDut() *OrganizationUpdate {
	ou.mutation.ClearNameDut()
	return ou
}

// SetNameEng sets the "name_eng" field.
func (ou *OrganizationUpdate) SetNameEng(s string) *OrganizationUpdate {
	ou.mutation.SetNameEng(s)
	return ou
}

// SetNillableNameEng sets the "name_eng" field if the given value is not nil.
func (ou *OrganizationUpdate) SetNillableNameEng(s *string) *OrganizationUpdate {
	if s != nil {
		ou.SetNameEng(*s)
	}
	return ou
}

// ClearNameEng clears the value of the "name_eng" field.
func (ou *OrganizationUpdate) ClearNameEng() *OrganizationUpdate {
	ou.mutation.ClearNameEng()
	return ou
}

// SetOtherID sets the "other_id" field.
func (ou *OrganizationUpdate) SetOtherID(sr []schema.IdRef) *OrganizationUpdate {
	ou.mutation.SetOtherID(sr)
	return ou
}

// AppendOtherID appends sr to the "other_id" field.
func (ou *OrganizationUpdate) AppendOtherID(sr []schema.IdRef) *OrganizationUpdate {
	ou.mutation.AppendOtherID(sr)
	return ou
}

// ClearOtherID clears the value of the "other_id" field.
func (ou *OrganizationUpdate) ClearOtherID() *OrganizationUpdate {
	ou.mutation.ClearOtherID()
	return ou
}

// SetParentID sets the "parent_id" field.
func (ou *OrganizationUpdate) SetParentID(i int) *OrganizationUpdate {
	ou.mutation.SetParentID(i)
	return ou
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (ou *OrganizationUpdate) SetNillableParentID(i *int) *OrganizationUpdate {
	if i != nil {
		ou.SetParentID(*i)
	}
	return ou
}

// ClearParentID clears the value of the "parent_id" field.
func (ou *OrganizationUpdate) ClearParentID() *OrganizationUpdate {
	ou.mutation.ClearParentID()
	return ou
}

// AddPersonIDs adds the "people" edge to the Person entity by IDs.
func (ou *OrganizationUpdate) AddPersonIDs(ids ...int) *OrganizationUpdate {
	ou.mutation.AddPersonIDs(ids...)
	return ou
}

// AddPeople adds the "people" edges to the Person entity.
func (ou *OrganizationUpdate) AddPeople(p ...*Person) *OrganizationUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ou.AddPersonIDs(ids...)
}

// SetParent sets the "parent" edge to the Organization entity.
func (ou *OrganizationUpdate) SetParent(o *Organization) *OrganizationUpdate {
	return ou.SetParentID(o.ID)
}

// AddChildIDs adds the "children" edge to the Organization entity by IDs.
func (ou *OrganizationUpdate) AddChildIDs(ids ...int) *OrganizationUpdate {
	ou.mutation.AddChildIDs(ids...)
	return ou
}

// AddChildren adds the "children" edges to the Organization entity.
func (ou *OrganizationUpdate) AddChildren(o ...*Organization) *OrganizationUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ou.AddChildIDs(ids...)
}

// AddOrganizationPersonIDs adds the "organization_person" edge to the OrganizationPerson entity by IDs.
func (ou *OrganizationUpdate) AddOrganizationPersonIDs(ids ...int) *OrganizationUpdate {
	ou.mutation.AddOrganizationPersonIDs(ids...)
	return ou
}

// AddOrganizationPerson adds the "organization_person" edges to the OrganizationPerson entity.
func (ou *OrganizationUpdate) AddOrganizationPerson(o ...*OrganizationPerson) *OrganizationUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ou.AddOrganizationPersonIDs(ids...)
}

// Mutation returns the OrganizationMutation object of the builder.
func (ou *OrganizationUpdate) Mutation() *OrganizationMutation {
	return ou.mutation
}

// ClearPeople clears all "people" edges to the Person entity.
func (ou *OrganizationUpdate) ClearPeople() *OrganizationUpdate {
	ou.mutation.ClearPeople()
	return ou
}

// RemovePersonIDs removes the "people" edge to Person entities by IDs.
func (ou *OrganizationUpdate) RemovePersonIDs(ids ...int) *OrganizationUpdate {
	ou.mutation.RemovePersonIDs(ids...)
	return ou
}

// RemovePeople removes "people" edges to Person entities.
func (ou *OrganizationUpdate) RemovePeople(p ...*Person) *OrganizationUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ou.RemovePersonIDs(ids...)
}

// ClearParent clears the "parent" edge to the Organization entity.
func (ou *OrganizationUpdate) ClearParent() *OrganizationUpdate {
	ou.mutation.ClearParent()
	return ou
}

// ClearChildren clears all "children" edges to the Organization entity.
func (ou *OrganizationUpdate) ClearChildren() *OrganizationUpdate {
	ou.mutation.ClearChildren()
	return ou
}

// RemoveChildIDs removes the "children" edge to Organization entities by IDs.
func (ou *OrganizationUpdate) RemoveChildIDs(ids ...int) *OrganizationUpdate {
	ou.mutation.RemoveChildIDs(ids...)
	return ou
}

// RemoveChildren removes "children" edges to Organization entities.
func (ou *OrganizationUpdate) RemoveChildren(o ...*Organization) *OrganizationUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ou.RemoveChildIDs(ids...)
}

// ClearOrganizationPerson clears all "organization_person" edges to the OrganizationPerson entity.
func (ou *OrganizationUpdate) ClearOrganizationPerson() *OrganizationUpdate {
	ou.mutation.ClearOrganizationPerson()
	return ou
}

// RemoveOrganizationPersonIDs removes the "organization_person" edge to OrganizationPerson entities by IDs.
func (ou *OrganizationUpdate) RemoveOrganizationPersonIDs(ids ...int) *OrganizationUpdate {
	ou.mutation.RemoveOrganizationPersonIDs(ids...)
	return ou
}

// RemoveOrganizationPerson removes "organization_person" edges to OrganizationPerson entities.
func (ou *OrganizationUpdate) RemoveOrganizationPerson(o ...*OrganizationPerson) *OrganizationUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ou.RemoveOrganizationPersonIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrganizationUpdate) Save(ctx context.Context) (int, error) {
	ou.defaults()
	return withHooks[int, OrganizationMutation](ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrganizationUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrganizationUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrganizationUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ou *OrganizationUpdate) defaults() {
	if _, ok := ou.mutation.DateUpdated(); !ok {
		v := organization.UpdateDefaultDateUpdated()
		ou.mutation.SetDateUpdated(v)
	}
}

func (ou *OrganizationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(organization.Table, organization.Columns, sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt))
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.DateUpdated(); ok {
		_spec.SetField(organization.FieldDateUpdated, field.TypeTime, value)
	}
	if value, ok := ou.mutation.GetType(); ok {
		_spec.SetField(organization.FieldType, field.TypeString, value)
	}
	if value, ok := ou.mutation.NameDut(); ok {
		_spec.SetField(organization.FieldNameDut, field.TypeString, value)
	}
	if ou.mutation.NameDutCleared() {
		_spec.ClearField(organization.FieldNameDut, field.TypeString)
	}
	if value, ok := ou.mutation.NameEng(); ok {
		_spec.SetField(organization.FieldNameEng, field.TypeString, value)
	}
	if ou.mutation.NameEngCleared() {
		_spec.ClearField(organization.FieldNameEng, field.TypeString)
	}
	if value, ok := ou.mutation.OtherID(); ok {
		_spec.SetField(organization.FieldOtherID, field.TypeJSON, value)
	}
	if value, ok := ou.mutation.AppendedOtherID(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, organization.FieldOtherID, value)
		})
	}
	if ou.mutation.OtherIDCleared() {
		_spec.ClearField(organization.FieldOtherID, field.TypeJSON)
	}
	if ou.mutation.PeopleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
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
		createE := &OrganizationPersonCreate{config: ou.config, mutation: newOrganizationPersonMutation(ou.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedPeopleIDs(); len(nodes) > 0 && !ou.mutation.PeopleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
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
		createE := &OrganizationPersonCreate{config: ou.config, mutation: newOrganizationPersonMutation(ou.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.PeopleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
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
		createE := &OrganizationPersonCreate{config: ou.config, mutation: newOrganizationPersonMutation(ou.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ou.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organization.ParentTable,
			Columns: []string{organization.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organization.ParentTable,
			Columns: []string{organization.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ou.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.ChildrenTable,
			Columns: []string{organization.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !ou.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.ChildrenTable,
			Columns: []string{organization.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.ChildrenTable,
			Columns: []string{organization.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ou.mutation.OrganizationPersonCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   organization.OrganizationPersonTable,
			Columns: []string{organization.OrganizationPersonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organizationperson.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedOrganizationPersonIDs(); len(nodes) > 0 && !ou.mutation.OrganizationPersonCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   organization.OrganizationPersonTable,
			Columns: []string{organization.OrganizationPersonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organizationperson.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.OrganizationPersonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   organization.OrganizationPersonTable,
			Columns: []string{organization.OrganizationPersonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organizationperson.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organization.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OrganizationUpdateOne is the builder for updating a single Organization entity.
type OrganizationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrganizationMutation
}

// SetDateUpdated sets the "date_updated" field.
func (ouo *OrganizationUpdateOne) SetDateUpdated(t time.Time) *OrganizationUpdateOne {
	ouo.mutation.SetDateUpdated(t)
	return ouo
}

// SetType sets the "type" field.
func (ouo *OrganizationUpdateOne) SetType(s string) *OrganizationUpdateOne {
	ouo.mutation.SetType(s)
	return ouo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ouo *OrganizationUpdateOne) SetNillableType(s *string) *OrganizationUpdateOne {
	if s != nil {
		ouo.SetType(*s)
	}
	return ouo
}

// SetNameDut sets the "name_dut" field.
func (ouo *OrganizationUpdateOne) SetNameDut(s string) *OrganizationUpdateOne {
	ouo.mutation.SetNameDut(s)
	return ouo
}

// SetNillableNameDut sets the "name_dut" field if the given value is not nil.
func (ouo *OrganizationUpdateOne) SetNillableNameDut(s *string) *OrganizationUpdateOne {
	if s != nil {
		ouo.SetNameDut(*s)
	}
	return ouo
}

// ClearNameDut clears the value of the "name_dut" field.
func (ouo *OrganizationUpdateOne) ClearNameDut() *OrganizationUpdateOne {
	ouo.mutation.ClearNameDut()
	return ouo
}

// SetNameEng sets the "name_eng" field.
func (ouo *OrganizationUpdateOne) SetNameEng(s string) *OrganizationUpdateOne {
	ouo.mutation.SetNameEng(s)
	return ouo
}

// SetNillableNameEng sets the "name_eng" field if the given value is not nil.
func (ouo *OrganizationUpdateOne) SetNillableNameEng(s *string) *OrganizationUpdateOne {
	if s != nil {
		ouo.SetNameEng(*s)
	}
	return ouo
}

// ClearNameEng clears the value of the "name_eng" field.
func (ouo *OrganizationUpdateOne) ClearNameEng() *OrganizationUpdateOne {
	ouo.mutation.ClearNameEng()
	return ouo
}

// SetOtherID sets the "other_id" field.
func (ouo *OrganizationUpdateOne) SetOtherID(sr []schema.IdRef) *OrganizationUpdateOne {
	ouo.mutation.SetOtherID(sr)
	return ouo
}

// AppendOtherID appends sr to the "other_id" field.
func (ouo *OrganizationUpdateOne) AppendOtherID(sr []schema.IdRef) *OrganizationUpdateOne {
	ouo.mutation.AppendOtherID(sr)
	return ouo
}

// ClearOtherID clears the value of the "other_id" field.
func (ouo *OrganizationUpdateOne) ClearOtherID() *OrganizationUpdateOne {
	ouo.mutation.ClearOtherID()
	return ouo
}

// SetParentID sets the "parent_id" field.
func (ouo *OrganizationUpdateOne) SetParentID(i int) *OrganizationUpdateOne {
	ouo.mutation.SetParentID(i)
	return ouo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (ouo *OrganizationUpdateOne) SetNillableParentID(i *int) *OrganizationUpdateOne {
	if i != nil {
		ouo.SetParentID(*i)
	}
	return ouo
}

// ClearParentID clears the value of the "parent_id" field.
func (ouo *OrganizationUpdateOne) ClearParentID() *OrganizationUpdateOne {
	ouo.mutation.ClearParentID()
	return ouo
}

// AddPersonIDs adds the "people" edge to the Person entity by IDs.
func (ouo *OrganizationUpdateOne) AddPersonIDs(ids ...int) *OrganizationUpdateOne {
	ouo.mutation.AddPersonIDs(ids...)
	return ouo
}

// AddPeople adds the "people" edges to the Person entity.
func (ouo *OrganizationUpdateOne) AddPeople(p ...*Person) *OrganizationUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ouo.AddPersonIDs(ids...)
}

// SetParent sets the "parent" edge to the Organization entity.
func (ouo *OrganizationUpdateOne) SetParent(o *Organization) *OrganizationUpdateOne {
	return ouo.SetParentID(o.ID)
}

// AddChildIDs adds the "children" edge to the Organization entity by IDs.
func (ouo *OrganizationUpdateOne) AddChildIDs(ids ...int) *OrganizationUpdateOne {
	ouo.mutation.AddChildIDs(ids...)
	return ouo
}

// AddChildren adds the "children" edges to the Organization entity.
func (ouo *OrganizationUpdateOne) AddChildren(o ...*Organization) *OrganizationUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouo.AddChildIDs(ids...)
}

// AddOrganizationPersonIDs adds the "organization_person" edge to the OrganizationPerson entity by IDs.
func (ouo *OrganizationUpdateOne) AddOrganizationPersonIDs(ids ...int) *OrganizationUpdateOne {
	ouo.mutation.AddOrganizationPersonIDs(ids...)
	return ouo
}

// AddOrganizationPerson adds the "organization_person" edges to the OrganizationPerson entity.
func (ouo *OrganizationUpdateOne) AddOrganizationPerson(o ...*OrganizationPerson) *OrganizationUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouo.AddOrganizationPersonIDs(ids...)
}

// Mutation returns the OrganizationMutation object of the builder.
func (ouo *OrganizationUpdateOne) Mutation() *OrganizationMutation {
	return ouo.mutation
}

// ClearPeople clears all "people" edges to the Person entity.
func (ouo *OrganizationUpdateOne) ClearPeople() *OrganizationUpdateOne {
	ouo.mutation.ClearPeople()
	return ouo
}

// RemovePersonIDs removes the "people" edge to Person entities by IDs.
func (ouo *OrganizationUpdateOne) RemovePersonIDs(ids ...int) *OrganizationUpdateOne {
	ouo.mutation.RemovePersonIDs(ids...)
	return ouo
}

// RemovePeople removes "people" edges to Person entities.
func (ouo *OrganizationUpdateOne) RemovePeople(p ...*Person) *OrganizationUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ouo.RemovePersonIDs(ids...)
}

// ClearParent clears the "parent" edge to the Organization entity.
func (ouo *OrganizationUpdateOne) ClearParent() *OrganizationUpdateOne {
	ouo.mutation.ClearParent()
	return ouo
}

// ClearChildren clears all "children" edges to the Organization entity.
func (ouo *OrganizationUpdateOne) ClearChildren() *OrganizationUpdateOne {
	ouo.mutation.ClearChildren()
	return ouo
}

// RemoveChildIDs removes the "children" edge to Organization entities by IDs.
func (ouo *OrganizationUpdateOne) RemoveChildIDs(ids ...int) *OrganizationUpdateOne {
	ouo.mutation.RemoveChildIDs(ids...)
	return ouo
}

// RemoveChildren removes "children" edges to Organization entities.
func (ouo *OrganizationUpdateOne) RemoveChildren(o ...*Organization) *OrganizationUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouo.RemoveChildIDs(ids...)
}

// ClearOrganizationPerson clears all "organization_person" edges to the OrganizationPerson entity.
func (ouo *OrganizationUpdateOne) ClearOrganizationPerson() *OrganizationUpdateOne {
	ouo.mutation.ClearOrganizationPerson()
	return ouo
}

// RemoveOrganizationPersonIDs removes the "organization_person" edge to OrganizationPerson entities by IDs.
func (ouo *OrganizationUpdateOne) RemoveOrganizationPersonIDs(ids ...int) *OrganizationUpdateOne {
	ouo.mutation.RemoveOrganizationPersonIDs(ids...)
	return ouo
}

// RemoveOrganizationPerson removes "organization_person" edges to OrganizationPerson entities.
func (ouo *OrganizationUpdateOne) RemoveOrganizationPerson(o ...*OrganizationPerson) *OrganizationUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouo.RemoveOrganizationPersonIDs(ids...)
}

// Where appends a list predicates to the OrganizationUpdate builder.
func (ouo *OrganizationUpdateOne) Where(ps ...predicate.Organization) *OrganizationUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrganizationUpdateOne) Select(field string, fields ...string) *OrganizationUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Organization entity.
func (ouo *OrganizationUpdateOne) Save(ctx context.Context) (*Organization, error) {
	ouo.defaults()
	return withHooks[*Organization, OrganizationMutation](ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrganizationUpdateOne) SaveX(ctx context.Context) *Organization {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrganizationUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrganizationUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouo *OrganizationUpdateOne) defaults() {
	if _, ok := ouo.mutation.DateUpdated(); !ok {
		v := organization.UpdateDefaultDateUpdated()
		ouo.mutation.SetDateUpdated(v)
	}
}

func (ouo *OrganizationUpdateOne) sqlSave(ctx context.Context) (_node *Organization, err error) {
	_spec := sqlgraph.NewUpdateSpec(organization.Table, organization.Columns, sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt))
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Organization.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organization.FieldID)
		for _, f := range fields {
			if !organization.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != organization.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.DateUpdated(); ok {
		_spec.SetField(organization.FieldDateUpdated, field.TypeTime, value)
	}
	if value, ok := ouo.mutation.GetType(); ok {
		_spec.SetField(organization.FieldType, field.TypeString, value)
	}
	if value, ok := ouo.mutation.NameDut(); ok {
		_spec.SetField(organization.FieldNameDut, field.TypeString, value)
	}
	if ouo.mutation.NameDutCleared() {
		_spec.ClearField(organization.FieldNameDut, field.TypeString)
	}
	if value, ok := ouo.mutation.NameEng(); ok {
		_spec.SetField(organization.FieldNameEng, field.TypeString, value)
	}
	if ouo.mutation.NameEngCleared() {
		_spec.ClearField(organization.FieldNameEng, field.TypeString)
	}
	if value, ok := ouo.mutation.OtherID(); ok {
		_spec.SetField(organization.FieldOtherID, field.TypeJSON, value)
	}
	if value, ok := ouo.mutation.AppendedOtherID(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, organization.FieldOtherID, value)
		})
	}
	if ouo.mutation.OtherIDCleared() {
		_spec.ClearField(organization.FieldOtherID, field.TypeJSON)
	}
	if ouo.mutation.PeopleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
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
		createE := &OrganizationPersonCreate{config: ouo.config, mutation: newOrganizationPersonMutation(ouo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedPeopleIDs(); len(nodes) > 0 && !ouo.mutation.PeopleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
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
		createE := &OrganizationPersonCreate{config: ouo.config, mutation: newOrganizationPersonMutation(ouo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.PeopleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
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
		createE := &OrganizationPersonCreate{config: ouo.config, mutation: newOrganizationPersonMutation(ouo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organization.ParentTable,
			Columns: []string{organization.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organization.ParentTable,
			Columns: []string{organization.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.ChildrenTable,
			Columns: []string{organization.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !ouo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.ChildrenTable,
			Columns: []string{organization.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.ChildrenTable,
			Columns: []string{organization.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouo.mutation.OrganizationPersonCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   organization.OrganizationPersonTable,
			Columns: []string{organization.OrganizationPersonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organizationperson.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedOrganizationPersonIDs(); len(nodes) > 0 && !ouo.mutation.OrganizationPersonCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   organization.OrganizationPersonTable,
			Columns: []string{organization.OrganizationPersonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organizationperson.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.OrganizationPersonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   organization.OrganizationPersonTable,
			Columns: []string{organization.OrganizationPersonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organizationperson.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Organization{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organization.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}
