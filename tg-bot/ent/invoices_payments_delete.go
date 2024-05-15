// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/invoices_payments"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/predicate"
)

// InvoicesPaymentsDelete is the builder for deleting a Invoices_payments entity.
type InvoicesPaymentsDelete struct {
	config
	hooks    []Hook
	mutation *InvoicesPaymentsMutation
}

// Where appends a list predicates to the InvoicesPaymentsDelete builder.
func (ipd *InvoicesPaymentsDelete) Where(ps ...predicate.Invoices_payments) *InvoicesPaymentsDelete {
	ipd.mutation.Where(ps...)
	return ipd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ipd *InvoicesPaymentsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ipd.sqlExec, ipd.mutation, ipd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ipd *InvoicesPaymentsDelete) ExecX(ctx context.Context) int {
	n, err := ipd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ipd *InvoicesPaymentsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(invoices_payments.Table, sqlgraph.NewFieldSpec(invoices_payments.FieldID, field.TypeInt))
	if ps := ipd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ipd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ipd.mutation.done = true
	return affected, err
}

// InvoicesPaymentsDeleteOne is the builder for deleting a single Invoices_payments entity.
type InvoicesPaymentsDeleteOne struct {
	ipd *InvoicesPaymentsDelete
}

// Where appends a list predicates to the InvoicesPaymentsDelete builder.
func (ipdo *InvoicesPaymentsDeleteOne) Where(ps ...predicate.Invoices_payments) *InvoicesPaymentsDeleteOne {
	ipdo.ipd.mutation.Where(ps...)
	return ipdo
}

// Exec executes the deletion query.
func (ipdo *InvoicesPaymentsDeleteOne) Exec(ctx context.Context) error {
	n, err := ipdo.ipd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{invoices_payments.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ipdo *InvoicesPaymentsDeleteOne) ExecX(ctx context.Context) {
	if err := ipdo.Exec(ctx); err != nil {
		panic(err)
	}
}
