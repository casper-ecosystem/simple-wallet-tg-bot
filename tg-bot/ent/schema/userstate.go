package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserState holds the schema definition for the UserState entity.
type UserState struct {
	ent.Schema
}

// Fields of the UserState.
func (UserState) Fields() []ent.Field {
	return []ent.Field{
		field.String("state"),
		field.Bytes("data").Optional(),
	}
}

// Edges of the UserState.
func (UserState) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("state").
			Unique().
			Required(),
	}
}
