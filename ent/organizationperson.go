// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ugent-library/people/ent/organization"
	"github.com/ugent-library/people/ent/organizationperson"
	"github.com/ugent-library/people/ent/person"
)

// OrganizationPerson is the model entity for the OrganizationPerson schema.
type OrganizationPerson struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DateCreated holds the value of the "date_created" field.
	DateCreated time.Time `json:"date_created,omitempty"`
	// DateUpdated holds the value of the "date_updated" field.
	DateUpdated time.Time `json:"date_updated,omitempty"`
	// OrganizationID holds the value of the "organization_id" field.
	OrganizationID int `json:"organization_id,omitempty"`
	// PersonID holds the value of the "person_id" field.
	PersonID int `json:"person_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrganizationPersonQuery when eager-loading is set.
	Edges OrganizationPersonEdges `json:"edges"`
}

// OrganizationPersonEdges holds the relations/edges for other nodes in the graph.
type OrganizationPersonEdges struct {
	// People holds the value of the people edge.
	People *Person `json:"people,omitempty"`
	// Organizations holds the value of the organizations edge.
	Organizations *Organization `json:"organizations,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PeopleOrErr returns the People value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrganizationPersonEdges) PeopleOrErr() (*Person, error) {
	if e.loadedTypes[0] {
		if e.People == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: person.Label}
		}
		return e.People, nil
	}
	return nil, &NotLoadedError{edge: "people"}
}

// OrganizationsOrErr returns the Organizations value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrganizationPersonEdges) OrganizationsOrErr() (*Organization, error) {
	if e.loadedTypes[1] {
		if e.Organizations == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: organization.Label}
		}
		return e.Organizations, nil
	}
	return nil, &NotLoadedError{edge: "organizations"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrganizationPerson) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case organizationperson.FieldID, organizationperson.FieldOrganizationID, organizationperson.FieldPersonID:
			values[i] = new(sql.NullInt64)
		case organizationperson.FieldDateCreated, organizationperson.FieldDateUpdated:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OrganizationPerson", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrganizationPerson fields.
func (op *OrganizationPerson) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case organizationperson.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			op.ID = int(value.Int64)
		case organizationperson.FieldDateCreated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date_created", values[i])
			} else if value.Valid {
				op.DateCreated = value.Time
			}
		case organizationperson.FieldDateUpdated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date_updated", values[i])
			} else if value.Valid {
				op.DateUpdated = value.Time
			}
		case organizationperson.FieldOrganizationID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value.Valid {
				op.OrganizationID = int(value.Int64)
			}
		case organizationperson.FieldPersonID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field person_id", values[i])
			} else if value.Valid {
				op.PersonID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryPeople queries the "people" edge of the OrganizationPerson entity.
func (op *OrganizationPerson) QueryPeople() *PersonQuery {
	return NewOrganizationPersonClient(op.config).QueryPeople(op)
}

// QueryOrganizations queries the "organizations" edge of the OrganizationPerson entity.
func (op *OrganizationPerson) QueryOrganizations() *OrganizationQuery {
	return NewOrganizationPersonClient(op.config).QueryOrganizations(op)
}

// Update returns a builder for updating this OrganizationPerson.
// Note that you need to call OrganizationPerson.Unwrap() before calling this method if this OrganizationPerson
// was returned from a transaction, and the transaction was committed or rolled back.
func (op *OrganizationPerson) Update() *OrganizationPersonUpdateOne {
	return NewOrganizationPersonClient(op.config).UpdateOne(op)
}

// Unwrap unwraps the OrganizationPerson entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (op *OrganizationPerson) Unwrap() *OrganizationPerson {
	_tx, ok := op.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrganizationPerson is not a transactional entity")
	}
	op.config.driver = _tx.drv
	return op
}

// String implements the fmt.Stringer.
func (op *OrganizationPerson) String() string {
	var builder strings.Builder
	builder.WriteString("OrganizationPerson(")
	builder.WriteString(fmt.Sprintf("id=%v, ", op.ID))
	builder.WriteString("date_created=")
	builder.WriteString(op.DateCreated.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("date_updated=")
	builder.WriteString(op.DateUpdated.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("organization_id=")
	builder.WriteString(fmt.Sprintf("%v", op.OrganizationID))
	builder.WriteString(", ")
	builder.WriteString("person_id=")
	builder.WriteString(fmt.Sprintf("%v", op.PersonID))
	builder.WriteByte(')')
	return builder.String()
}

// OrganizationPersons is a parsable slice of OrganizationPerson.
type OrganizationPersons []*OrganizationPerson
