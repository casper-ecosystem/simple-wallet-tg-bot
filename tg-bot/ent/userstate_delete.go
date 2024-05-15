// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/predicate"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/userstate"
)

// UserStateDelete is the builder for deleting a UserState entity.
type UserStateDelete struct {
	config
	hooks    []Hook
	mutation *UserStateMutation
}

// Where appends a list predicates to the UserStateDelete builder.
func (usd *UserStateDelete) Where(ps ...predicate.UserState) *UserStateDelete {
	usd.mutation.Where(ps...)
	return usd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (usd *UserStateDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, usd.sqlExec, usd.mutation, usd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (usd *UserStateDelete) ExecX(ctx context.Context) int {
	n, err := usd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (usd *UserStateDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(userstate.Table, sqlgraph.NewFieldSpec(userstate.FieldID, field.TypeInt))
	if ps := usd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, usd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	usd.mutation.done = true
	return affected, err
}

// UserStateDeleteOne is the builder for deleting a single UserState entity.
type UserStateDeleteOne struct {
	usd *UserStateDelete
}

// Where appends a list predicates to the UserStateDelete builder.
func (usdo *UserStateDeleteOne) Where(ps ...predicate.UserState) *UserStateDeleteOne {
	usdo.usd.mutation.Where(ps...)
	return usdo
}

// Exec executes the deletion query.
func (usdo *UserStateDeleteOne) Exec(ctx context.Context) error {
	n, err := usdo.usd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userstate.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (usdo *UserStateDeleteOne) ExecX(ctx context.Context) {
	if err := usdo.Exec(ctx); err != nil {
		panic(err)
	}
}
