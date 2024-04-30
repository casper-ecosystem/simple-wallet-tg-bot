package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Invoices_payments holds the schema definition for the Invoices_payments entity.
type Invoices_payments struct {
	ent.Schema
}

// Fields of the Invoices_payments.
func (Invoices_payments) Fields() []ent.Field {
	return []ent.Field{
		field.String("from").Optional(),
		field.String("amount").Optional(),
		field.Bool("correct").Optional(),
	}
}

// Edges of the Invoices_payments.
func (Invoices_payments) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("invoice", Invoice.Type).
			Ref("payments").
			Unique().
			Required(),
	}
}
