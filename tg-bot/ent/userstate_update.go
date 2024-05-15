// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/predicate"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/user"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/userstate"
)

// UserStateUpdate is the builder for updating UserState entities.
type UserStateUpdate struct {
	config
	hooks    []Hook
	mutation *UserStateMutation
}

// Where appends a list predicates to the UserStateUpdate builder.
func (usu *UserStateUpdate) Where(ps ...predicate.UserState) *UserStateUpdate {
	usu.mutation.Where(ps...)
	return usu
}

// SetState sets the "state" field.
func (usu *UserStateUpdate) SetState(s string) *UserStateUpdate {
	usu.mutation.SetState(s)
	return usu
}

// SetData sets the "data" field.
func (usu *UserStateUpdate) SetData(b []byte) *UserStateUpdate {
	usu.mutation.SetData(b)
	return usu
}

// ClearData clears the value of the "data" field.
func (usu *UserStateUpdate) ClearData() *UserStateUpdate {
	usu.mutation.ClearData()
	return usu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (usu *UserStateUpdate) SetOwnerID(id int64) *UserStateUpdate {
	usu.mutation.SetOwnerID(id)
	return usu
}

// SetOwner sets the "owner" edge to the User entity.
func (usu *UserStateUpdate) SetOwner(u *User) *UserStateUpdate {
	return usu.SetOwnerID(u.ID)
}

// Mutation returns the UserStateMutation object of the builder.
func (usu *UserStateUpdate) Mutation() *UserStateMutation {
	return usu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (usu *UserStateUpdate) ClearOwner() *UserStateUpdate {
	usu.mutation.ClearOwner()
	return usu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (usu *UserStateUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, usu.sqlSave, usu.mutation, usu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (usu *UserStateUpdate) SaveX(ctx context.Context) int {
	affected, err := usu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (usu *UserStateUpdate) Exec(ctx context.Context) error {
	_, err := usu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usu *UserStateUpdate) ExecX(ctx context.Context) {
	if err := usu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (usu *UserStateUpdate) check() error {
	if _, ok := usu.mutation.OwnerID(); usu.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserState.owner"`)
	}
	return nil
}

func (usu *UserStateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := usu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(userstate.Table, userstate.Columns, sqlgraph.NewFieldSpec(userstate.FieldID, field.TypeInt))
	if ps := usu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := usu.mutation.State(); ok {
		_spec.SetField(userstate.FieldState, field.TypeString, value)
	}
	if value, ok := usu.mutation.Data(); ok {
		_spec.SetField(userstate.FieldData, field.TypeBytes, value)
	}
	if usu.mutation.DataCleared() {
		_spec.ClearField(userstate.FieldData, field.TypeBytes)
	}
	if usu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   userstate.OwnerTable,
			Columns: []string{userstate.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := usu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   userstate.OwnerTable,
			Columns: []string{userstate.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, usu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userstate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	usu.mutation.done = true
	return n, nil
}

// UserStateUpdateOne is the builder for updating a single UserState entity.
type UserStateUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserStateMutation
}

// SetState sets the "state" field.
func (usuo *UserStateUpdateOne) SetState(s string) *UserStateUpdateOne {
	usuo.mutation.SetState(s)
	return usuo
}

// SetData sets the "data" field.
func (usuo *UserStateUpdateOne) SetData(b []byte) *UserStateUpdateOne {
	usuo.mutation.SetData(b)
	return usuo
}

// ClearData clears the value of the "data" field.
func (usuo *UserStateUpdateOne) ClearData() *UserStateUpdateOne {
	usuo.mutation.ClearData()
	return usuo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (usuo *UserStateUpdateOne) SetOwnerID(id int64) *UserStateUpdateOne {
	usuo.mutation.SetOwnerID(id)
	return usuo
}

// SetOwner sets the "owner" edge to the User entity.
func (usuo *UserStateUpdateOne) SetOwner(u *User) *UserStateUpdateOne {
	return usuo.SetOwnerID(u.ID)
}

// Mutation returns the UserStateMutation object of the builder.
func (usuo *UserStateUpdateOne) Mutation() *UserStateMutation {
	return usuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (usuo *UserStateUpdateOne) ClearOwner() *UserStateUpdateOne {
	usuo.mutation.ClearOwner()
	return usuo
}

// Where appends a list predicates to the UserStateUpdate builder.
func (usuo *UserStateUpdateOne) Where(ps ...predicate.UserState) *UserStateUpdateOne {
	usuo.mutation.Where(ps...)
	return usuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (usuo *UserStateUpdateOne) Select(field string, fields ...string) *UserStateUpdateOne {
	usuo.fields = append([]string{field}, fields...)
	return usuo
}

// Save executes the query and returns the updated UserState entity.
func (usuo *UserStateUpdateOne) Save(ctx context.Context) (*UserState, error) {
	return withHooks(ctx, usuo.sqlSave, usuo.mutation, usuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (usuo *UserStateUpdateOne) SaveX(ctx context.Context) *UserState {
	node, err := usuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (usuo *UserStateUpdateOne) Exec(ctx context.Context) error {
	_, err := usuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usuo *UserStateUpdateOne) ExecX(ctx context.Context) {
	if err := usuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (usuo *UserStateUpdateOne) check() error {
	if _, ok := usuo.mutation.OwnerID(); usuo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserState.owner"`)
	}
	return nil
}

func (usuo *UserStateUpdateOne) sqlSave(ctx context.Context) (_node *UserState, err error) {
	if err := usuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(userstate.Table, userstate.Columns, sqlgraph.NewFieldSpec(userstate.FieldID, field.TypeInt))
	id, ok := usuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserState.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := usuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userstate.FieldID)
		for _, f := range fields {
			if !userstate.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userstate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := usuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := usuo.mutation.State(); ok {
		_spec.SetField(userstate.FieldState, field.TypeString, value)
	}
	if value, ok := usuo.mutation.Data(); ok {
		_spec.SetField(userstate.FieldData, field.TypeBytes, value)
	}
	if usuo.mutation.DataCleared() {
		_spec.ClearField(userstate.FieldData, field.TypeBytes)
	}
	if usuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   userstate.OwnerTable,
			Columns: []string{userstate.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := usuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   userstate.OwnerTable,
			Columns: []string{userstate.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserState{config: usuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, usuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userstate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	usuo.mutation.done = true
	return _node, nil
}
