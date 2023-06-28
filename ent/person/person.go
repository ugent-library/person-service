// Code generated by ent, DO NOT EDIT.

package person

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the person type in the database.
	Label = "person"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDateCreated holds the string denoting the date_created field in the database.
	FieldDateCreated = "date_created"
	// FieldDateUpdated holds the string denoting the date_updated field in the database.
	FieldDateUpdated = "date_updated"
	// FieldPublicID holds the string denoting the public_id field in the database.
	FieldPublicID = "public_id"
	// FieldGismoID holds the string denoting the gismo_id field in the database.
	FieldGismoID = "gismo_id"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// FieldBirthDate holds the string denoting the birth_date field in the database.
	FieldBirthDate = "birth_date"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldOtherID holds the string denoting the other_id field in the database.
	FieldOtherID = "other_id"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldFullName holds the string denoting the full_name field in the database.
	FieldFullName = "full_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldJobCategory holds the string denoting the job_category field in the database.
	FieldJobCategory = "job_category"
	// FieldOrcid holds the string denoting the orcid field in the database.
	FieldOrcid = "orcid"
	// FieldOrcidToken holds the string denoting the orcid_token field in the database.
	FieldOrcidToken = "orcid_token"
	// FieldPreferredFirstName holds the string denoting the preferred_first_name field in the database.
	FieldPreferredFirstName = "preferred_first_name"
	// FieldPreferredLastName holds the string denoting the preferred_last_name field in the database.
	FieldPreferredLastName = "preferred_last_name"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldSettings holds the string denoting the settings field in the database.
	FieldSettings = "settings"
	// EdgeOrganizations holds the string denoting the organizations edge name in mutations.
	EdgeOrganizations = "organizations"
	// EdgeOrganizationPerson holds the string denoting the organization_person edge name in mutations.
	EdgeOrganizationPerson = "organization_person"
	// Table holds the table name of the person in the database.
	Table = "person"
	// OrganizationsTable is the table that holds the organizations relation/edge. The primary key declared below.
	OrganizationsTable = "organization_person"
	// OrganizationsInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationsInverseTable = "organization"
	// OrganizationPersonTable is the table that holds the organization_person relation/edge.
	OrganizationPersonTable = "organization_person"
	// OrganizationPersonInverseTable is the table name for the OrganizationPerson entity.
	// It exists in this package in order to avoid circular dependency with the "organizationperson" package.
	OrganizationPersonInverseTable = "organization_person"
	// OrganizationPersonColumn is the table column denoting the organization_person relation/edge.
	OrganizationPersonColumn = "person_id"
)

// Columns holds all SQL columns for person fields.
var Columns = []string{
	FieldID,
	FieldDateCreated,
	FieldDateUpdated,
	FieldPublicID,
	FieldGismoID,
	FieldActive,
	FieldBirthDate,
	FieldEmail,
	FieldOtherID,
	FieldFirstName,
	FieldFullName,
	FieldLastName,
	FieldJobCategory,
	FieldOrcid,
	FieldOrcidToken,
	FieldPreferredFirstName,
	FieldPreferredLastName,
	FieldTitle,
	FieldRole,
	FieldSettings,
}

var (
	// OrganizationsPrimaryKey and OrganizationsColumn2 are the table columns denoting the
	// primary key for the organizations relation (M2M).
	OrganizationsPrimaryKey = []string{"person_id", "organization_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultDateCreated holds the default value on creation for the "date_created" field.
	DefaultDateCreated func() time.Time
	// DefaultDateUpdated holds the default value on creation for the "date_updated" field.
	DefaultDateUpdated func() time.Time
	// UpdateDefaultDateUpdated holds the default value on update for the "date_updated" field.
	UpdateDefaultDateUpdated func() time.Time
	// DefaultPublicID holds the default value on creation for the "public_id" field.
	DefaultPublicID func() string
	// DefaultActive holds the default value on creation for the "active" field.
	DefaultActive bool
)

// OrderOption defines the ordering options for the Person queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDateCreated orders the results by the date_created field.
func ByDateCreated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDateCreated, opts...).ToFunc()
}

// ByDateUpdated orders the results by the date_updated field.
func ByDateUpdated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDateUpdated, opts...).ToFunc()
}

// ByPublicID orders the results by the public_id field.
func ByPublicID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublicID, opts...).ToFunc()
}

// ByGismoID orders the results by the gismo_id field.
func ByGismoID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGismoID, opts...).ToFunc()
}

// ByActive orders the results by the active field.
func ByActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActive, opts...).ToFunc()
}

// ByBirthDate orders the results by the birth_date field.
func ByBirthDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBirthDate, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByFirstName orders the results by the first_name field.
func ByFirstName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstName, opts...).ToFunc()
}

// ByFullName orders the results by the full_name field.
func ByFullName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFullName, opts...).ToFunc()
}

// ByLastName orders the results by the last_name field.
func ByLastName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastName, opts...).ToFunc()
}

// ByOrcid orders the results by the orcid field.
func ByOrcid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrcid, opts...).ToFunc()
}

// ByOrcidToken orders the results by the orcid_token field.
func ByOrcidToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrcidToken, opts...).ToFunc()
}

// ByPreferredFirstName orders the results by the preferred_first_name field.
func ByPreferredFirstName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPreferredFirstName, opts...).ToFunc()
}

// ByPreferredLastName orders the results by the preferred_last_name field.
func ByPreferredLastName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPreferredLastName, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByOrganizationsCount orders the results by organizations count.
func ByOrganizationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOrganizationsStep(), opts...)
	}
}

// ByOrganizations orders the results by organizations terms.
func ByOrganizations(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrganizationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOrganizationPersonCount orders the results by organization_person count.
func ByOrganizationPersonCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOrganizationPersonStep(), opts...)
	}
}

// ByOrganizationPerson orders the results by organization_person terms.
func ByOrganizationPerson(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrganizationPersonStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newOrganizationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrganizationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, OrganizationsTable, OrganizationsPrimaryKey...),
	)
}
func newOrganizationPersonStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrganizationPersonInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, OrganizationPersonTable, OrganizationPersonColumn),
	)
}
