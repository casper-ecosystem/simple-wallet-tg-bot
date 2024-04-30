package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Balances holds the schema definition for the Balances entity.
type Balances struct {
	ent.Schema
}

// Fields of the Balances.
func (Balances) Fields() []ent.Field {
	return []ent.Field{
		// field.Int64("id").Unique(),
		// field.String("public_key"),
		field.Float("balance"),
		field.Uint64("height"),
	}
}

// Edges of the Balances.
func (Balances) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("balance").
			Unique().
			Required(),
	}
}
