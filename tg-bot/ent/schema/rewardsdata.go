package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RewardsData holds the schema definition for the RewardsData entity.
type RewardsData struct {
	ent.Schema
}

// Fields of the RewardsData.
func (RewardsData) Fields() []ent.Field {
	return []ent.Field{
		field.String("validator"),
		field.String("amount"),
		field.Time("last_reward"),
		field.Int64("first_era").Default(0),
		field.Int64("last_era").Default(0),
		field.String("first_era_timestamp"),
		field.String("last_era_timestamp"),
	}
}

// Edges of the RewardsData.
func (RewardsData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("rewards_data").
			Unique().
			Required(),
	}
}
