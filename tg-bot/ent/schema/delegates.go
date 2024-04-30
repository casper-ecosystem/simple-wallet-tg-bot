package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Delegates holds the schema definition for the Delegates entity.
type Delegates struct {
	ent.Schema
}

// Fields of the Delegates.
func (Delegates) Fields() []ent.Field {
	return []ent.Field{
		field.String("delegator"),
		field.String("validator").Optional(),
		field.String("name").Optional(),
		field.String("user_balance").Optional(),
		field.String("amount").Optional(),
		field.Time("created_at").Optional(),
		field.String("status").Optional(),
		field.String("Deploy").Optional(),
	}
}

// Edges of the Delegates.
func (Delegates) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("delegates").
			Unique().
			Required(),
	}
}
