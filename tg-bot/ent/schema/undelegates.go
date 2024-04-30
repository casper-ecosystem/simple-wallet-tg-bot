package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Undelegates holds the schema definition for the Undelegates entity.
type Undelegates struct {
	ent.Schema
}

// Fields of the Undelegates.
func (Undelegates) Fields() []ent.Field {
	return []ent.Field{
		field.String("delegator"),
		field.String("validator").Optional(),
		field.String("name").Optional(),
		field.String("staked_balance").Optional(),
		field.String("amount").Optional(),
		field.Time("created_at").Optional(),
		field.String("status").Optional(),
		field.String("Deploy").Optional(),
	}
}

// Edges of the Undelegates.
func (Undelegates) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("undelegates").
			Unique().
			Required(),
	}
}
