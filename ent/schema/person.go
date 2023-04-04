package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Person holds the schema definition for the Person entity.
type Person struct {
	ent.Schema
}

type IdRef struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// TODO validate type
var IdRefTypes = []string{
	"ugent_id",
	"ugent_barcode",
	"ugent_username",
	"historic_ugent_id",
	"ugent_memorialis_id",
	"uzgent_id",
}

func (Person) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "person"},
	}
}

// Fields of the Person.
func (Person) Fields() []ent.Field {
	return []ent.Field{
		field.String("primary_id").Immutable().Unique(),
		field.Bool("active").Default(false),
		field.String("birth_date").Optional(),
		field.String("email").Optional(),
		field.JSON("other_id", []IdRef{}).Optional(),
		field.String("first_name").Optional(),
		field.String("full_name").Optional(),
		field.String("last_name").Optional(),
		field.Strings("job_category").Optional(),
		field.String("orcid").Optional(),
		field.String("orcid_token").Optional(),
		field.String("preferred_first_name").Optional(),
		field.String("preferred_last_name").Optional(),
		field.String("title").Optional(),
	}
}

func (Person) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the Person.
func (Person) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organizations", Organization.Type).Ref("people"),
	}
}

func (Person) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("primary_id"),
		index.Fields("active"),
		index.Fields("orcid"),
		index.Fields("email"),
		index.Fields("first_name"),
		index.Fields("last_name"),
		index.Fields("full_name"),
	}
}
