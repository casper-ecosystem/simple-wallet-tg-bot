package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Swaps holds the schema definition for the Swaps entity.
type Swaps struct {
	ent.Schema
}

// Fields of the Swaps.
func (Swaps) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").Optional(),
		field.String("to_address").Optional(),
		field.String("from_currency").Optional(),
		field.String("to_currency").Optional(),
		field.String("to_network").Optional(),
		field.String("from_network").Optional(),
		field.String("amount").Optional(),
		field.String("amountRecive").Optional(),
		field.String("refund_address").Optional(),
		field.String("swap_id").Optional(),
		field.String("extra_id").Optional(),
		field.Int64("invoiceID").Optional(),
	}
}

// Edges of the Swaps.
func (Swaps) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("swaps").
			Unique().
			Required(),
	}
}
