// Code generated by ent, DO NOT EDIT.

package recentinvoices

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Simplewallethq/tg-bot/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldLTE(FieldID, id))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldEQ(FieldStatus, v))
}

// InvoiceID applies equality check predicate on the "invoiceID" field. It's identical to InvoiceIDEQ.
func InvoiceID(v int64) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldEQ(FieldInvoiceID, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldContainsFold(FieldStatus, v))
}

// InvoiceIDEQ applies the EQ predicate on the "invoiceID" field.
func InvoiceIDEQ(v int64) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldEQ(FieldInvoiceID, v))
}

// InvoiceIDNEQ applies the NEQ predicate on the "invoiceID" field.
func InvoiceIDNEQ(v int64) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldNEQ(FieldInvoiceID, v))
}

// InvoiceIDIn applies the In predicate on the "invoiceID" field.
func InvoiceIDIn(vs ...int64) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldIn(FieldInvoiceID, vs...))
}

// InvoiceIDNotIn applies the NotIn predicate on the "invoiceID" field.
func InvoiceIDNotIn(vs ...int64) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldNotIn(FieldInvoiceID, vs...))
}

// InvoiceIDGT applies the GT predicate on the "invoiceID" field.
func InvoiceIDGT(v int64) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldGT(FieldInvoiceID, v))
}

// InvoiceIDGTE applies the GTE predicate on the "invoiceID" field.
func InvoiceIDGTE(v int64) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldGTE(FieldInvoiceID, v))
}

// InvoiceIDLT applies the LT predicate on the "invoiceID" field.
func InvoiceIDLT(v int64) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldLT(FieldInvoiceID, v))
}

// InvoiceIDLTE applies the LTE predicate on the "invoiceID" field.
func InvoiceIDLTE(v int64) predicate.RecentInvoices {
	return predicate.RecentInvoices(sql.FieldLTE(FieldInvoiceID, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.RecentInvoices {
	return predicate.RecentInvoices(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.RecentInvoices {
	return predicate.RecentInvoices(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RecentInvoices) predicate.RecentInvoices {
	return predicate.RecentInvoices(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RecentInvoices) predicate.RecentInvoices {
	return predicate.RecentInvoices(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.RecentInvoices) predicate.RecentInvoices {
	return predicate.RecentInvoices(func(s *sql.Selector) {
		p(s.Not())
	})
}