package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Invoice holds the schema definition for the Invoice entity.
type Invoice struct {
	ent.Schema
}

// Fields of the Invoice.
func (Invoice) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.String("address").Optional(),
		field.String("amount").Optional(),
		field.String("currency").Optional(), //now only cspr
		field.String("comment").Optional(),
		field.Bool("active").Optional(),
		field.Int("repeatability").Optional(), // 0 for infinity
		field.Int("paid").Optional(),          //how many times paid
		field.String("short").Optional(),
		field.Uint64("memo").Optional(), //for cspr transfer
	}
}

// Edges of the Invoice.
func (Invoice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("invoices").
			Unique().
			Required(),
		edge.To("payments", Invoices_payments.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
