package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AdressBook holds the schema definition for the AdressBook entity.
type AdressBook struct {
	ent.Schema
}

// Fields of the AdressBook.
func (AdressBook) Fields() []ent.Field {
	return []ent.Field{
		field.String("address"),
		field.String("name"),
		field.Time("created_at"),
		field.Bool("InUpdate"),
	}
}

// Edges of the AdressBook.
func (AdressBook) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("address_book").
			Unique().
			Required(),
	}
}
