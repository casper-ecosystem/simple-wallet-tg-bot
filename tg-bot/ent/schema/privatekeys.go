package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PrivateKeys holds the schema definition for the PrivateKeys entity.
type PrivateKeys struct {
	ent.Schema
}

// Fields of the PrivateKeys.
func (PrivateKeys) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("private_key").Optional(),
	}
}

// Edges of the PrivateKeys.
func (PrivateKeys) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("PrivateKey").
			Unique().
			Required(),
	}
}
