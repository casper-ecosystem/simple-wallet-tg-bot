// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/adressbook"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/predicate"
)

// AdressBookDelete is the builder for deleting a AdressBook entity.
type AdressBookDelete struct {
	config
	hooks    []Hook
	mutation *AdressBookMutation
}

// Where appends a list predicates to the AdressBookDelete builder.
func (abd *AdressBookDelete) Where(ps ...predicate.AdressBook) *AdressBookDelete {
	abd.mutation.Where(ps...)
	return abd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (abd *AdressBookDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, abd.sqlExec, abd.mutation, abd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (abd *AdressBookDelete) ExecX(ctx context.Context) int {
	n, err := abd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (abd *AdressBookDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(adressbook.Table, sqlgraph.NewFieldSpec(adressbook.FieldID, field.TypeInt))
	if ps := abd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, abd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	abd.mutation.done = true
	return affected, err
}

// AdressBookDeleteOne is the builder for deleting a single AdressBook entity.
type AdressBookDeleteOne struct {
	abd *AdressBookDelete
}

// Where appends a list predicates to the AdressBookDelete builder.
func (abdo *AdressBookDeleteOne) Where(ps ...predicate.AdressBook) *AdressBookDeleteOne {
	abdo.abd.mutation.Where(ps...)
	return abdo
}

// Exec executes the deletion query.
func (abdo *AdressBookDeleteOne) Exec(ctx context.Context) error {
	n, err := abdo.abd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{adressbook.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (abdo *AdressBookDeleteOne) ExecX(ctx context.Context) {
	if err := abdo.Exec(ctx); err != nil {
		panic(err)
	}
}
