package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RecentInvoices holds the schema definition for the RecentInvoices entity.
type RecentInvoices struct {
	ent.Schema
}

// Fields of the RecentInvoices.
func (RecentInvoices) Fields() []ent.Field {
	return []ent.Field{
		field.String("status"),
		field.Int64("invoiceID").Unique(),
	}
}

// Edges of the RecentInvoices.
func (RecentInvoices) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("recentInvoices").
			Unique().
			Required(),
	}
}
