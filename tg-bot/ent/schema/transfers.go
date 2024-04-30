package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Transfers holds the schema definition for the Transfers entity.
type Transfers struct {
	ent.Schema
}

// Fields of the Transfers.
func (Transfers) Fields() []ent.Field {
	return []ent.Field{
		field.String("from_pubkey"),
		field.String("to_pubkey").Optional(),
		field.String("name").Optional(),
		field.String("sender_balance").Optional(),
		field.String("amount").Optional(),
		field.Uint64("memo_id").Optional(),
		field.Time("created_at").Optional(),
		field.String("status").Optional(),
		field.String("Deploy").Optional(),
		field.String("AdditionalType").Optional(),
		field.Int64("invoiceID").Optional(),
	}
}

// Edges of the Transfers.
func (Transfers) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("transfers").
			Unique().
			Required(),
	}
}
