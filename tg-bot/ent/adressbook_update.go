// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/adressbook"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/predicate"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/user"
)

// AdressBookUpdate is the builder for updating AdressBook entities.
type AdressBookUpdate struct {
	config
	hooks    []Hook
	mutation *AdressBookMutation
}

// Where appends a list predicates to the AdressBookUpdate builder.
func (abu *AdressBookUpdate) Where(ps ...predicate.AdressBook) *AdressBookUpdate {
	abu.mutation.Where(ps...)
	return abu
}

// SetAddress sets the "address" field.
func (abu *AdressBookUpdate) SetAddress(s string) *AdressBookUpdate {
	abu.mutation.SetAddress(s)
	return abu
}

// SetName sets the "name" field.
func (abu *AdressBookUpdate) SetName(s string) *AdressBookUpdate {
	abu.mutation.SetName(s)
	return abu
}

// SetCreatedAt sets the "created_at" field.
func (abu *AdressBookUpdate) SetCreatedAt(t time.Time) *AdressBookUpdate {
	abu.mutation.SetCreatedAt(t)
	return abu
}

// SetInUpdate sets the "InUpdate" field.
func (abu *AdressBookUpdate) SetInUpdate(b bool) *AdressBookUpdate {
	abu.mutation.SetInUpdate(b)
	return abu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (abu *AdressBookUpdate) SetOwnerID(id int64) *AdressBookUpdate {
	abu.mutation.SetOwnerID(id)
	return abu
}

// SetOwner sets the "owner" edge to the User entity.
func (abu *AdressBookUpdate) SetOwner(u *User) *AdressBookUpdate {
	return abu.SetOwnerID(u.ID)
}

// Mutation returns the AdressBookMutation object of the builder.
func (abu *AdressBookUpdate) Mutation() *AdressBookMutation {
	return abu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (abu *AdressBookUpdate) ClearOwner() *AdressBookUpdate {
	abu.mutation.ClearOwner()
	return abu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (abu *AdressBookUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, abu.sqlSave, abu.mutation, abu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (abu *AdressBookUpdate) SaveX(ctx context.Context) int {
	affected, err := abu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (abu *AdressBookUpdate) Exec(ctx context.Context) error {
	_, err := abu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (abu *AdressBookUpdate) ExecX(ctx context.Context) {
	if err := abu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (abu *AdressBookUpdate) check() error {
	if _, ok := abu.mutation.OwnerID(); abu.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AdressBook.owner"`)
	}
	return nil
}

func (abu *AdressBookUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := abu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(adressbook.Table, adressbook.Columns, sqlgraph.NewFieldSpec(adressbook.FieldID, field.TypeInt))
	if ps := abu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := abu.mutation.Address(); ok {
		_spec.SetField(adressbook.FieldAddress, field.TypeString, value)
	}
	if value, ok := abu.mutation.Name(); ok {
		_spec.SetField(adressbook.FieldName, field.TypeString, value)
	}
	if value, ok := abu.mutation.CreatedAt(); ok {
		_spec.SetField(adressbook.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := abu.mutation.InUpdate(); ok {
		_spec.SetField(adressbook.FieldInUpdate, field.TypeBool, value)
	}
	if abu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   adressbook.OwnerTable,
			Columns: []string{adressbook.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := abu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   adressbook.OwnerTable,
			Columns: []string{adressbook.OwnerColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, abu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adressbook.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	abu.mutation.done = true
	return n, nil
}

// AdressBookUpdateOne is the builder for updating a single AdressBook entity.
type AdressBookUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AdressBookMutation
}

// SetAddress sets the "address" field.
func (abuo *AdressBookUpdateOne) SetAddress(s string) *AdressBookUpdateOne {
	abuo.mutation.SetAddress(s)
	return abuo
}

// SetName sets the "name" field.
func (abuo *AdressBookUpdateOne) SetName(s string) *AdressBookUpdateOne {
	abuo.mutation.SetName(s)
	return abuo
}

// SetCreatedAt sets the "created_at" field.
func (abuo *AdressBookUpdateOne) SetCreatedAt(t time.Time) *AdressBookUpdateOne {
	abuo.mutation.SetCreatedAt(t)
	return abuo
}

// SetInUpdate sets the "InUpdate" field.
func (abuo *AdressBookUpdateOne) SetInUpdate(b bool) *AdressBookUpdateOne {
	abuo.mutation.SetInUpdate(b)
	return abuo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (abuo *AdressBookUpdateOne) SetOwnerID(id int64) *AdressBookUpdateOne {
	abuo.mutation.SetOwnerID(id)
	return abuo
}

// SetOwner sets the "owner" edge to the User entity.
func (abuo *AdressBookUpdateOne) SetOwner(u *User) *AdressBookUpdateOne {
	return abuo.SetOwnerID(u.ID)
}

// Mutation returns the AdressBookMutation object of the builder.
func (abuo *AdressBookUpdateOne) Mutation() *AdressBookMutation {
	return abuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (abuo *AdressBookUpdateOne) ClearOwner() *AdressBookUpdateOne {
	abuo.mutation.ClearOwner()
	return abuo
}

// Where appends a list predicates to the AdressBookUpdate builder.
func (abuo *AdressBookUpdateOne) Where(ps ...predicate.AdressBook) *AdressBookUpdateOne {
	abuo.mutation.Where(ps...)
	return abuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (abuo *AdressBookUpdateOne) Select(field string, fields ...string) *AdressBookUpdateOne {
	abuo.fields = append([]string{field}, fields...)
	return abuo
}

// Save executes the query and returns the updated AdressBook entity.
func (abuo *AdressBookUpdateOne) Save(ctx context.Context) (*AdressBook, error) {
	return withHooks(ctx, abuo.sqlSave, abuo.mutation, abuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (abuo *AdressBookUpdateOne) SaveX(ctx context.Context) *AdressBook {
	node, err := abuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (abuo *AdressBookUpdateOne) Exec(ctx context.Context) error {
	_, err := abuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (abuo *AdressBookUpdateOne) ExecX(ctx context.Context) {
	if err := abuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (abuo *AdressBookUpdateOne) check() error {
	if _, ok := abuo.mutation.OwnerID(); abuo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AdressBook.owner"`)
	}
	return nil
}

func (abuo *AdressBookUpdateOne) sqlSave(ctx context.Context) (_node *AdressBook, err error) {
	if err := abuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(adressbook.Table, adressbook.Columns, sqlgraph.NewFieldSpec(adressbook.FieldID, field.TypeInt))
	id, ok := abuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AdressBook.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := abuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, adressbook.FieldID)
		for _, f := range fields {
			if !adressbook.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != adressbook.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := abuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := abuo.mutation.Address(); ok {
		_spec.SetField(adressbook.FieldAddress, field.TypeString, value)
	}
	if value, ok := abuo.mutation.Name(); ok {
		_spec.SetField(adressbook.FieldName, field.TypeString, value)
	}
	if value, ok := abuo.mutation.CreatedAt(); ok {
		_spec.SetField(adressbook.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := abuo.mutation.InUpdate(); ok {
		_spec.SetField(adressbook.FieldInUpdate, field.TypeBool, value)
	}
	if abuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   adressbook.OwnerTable,
			Columns: []string{adressbook.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := abuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   adressbook.OwnerTable,
			Columns: []string{adressbook.OwnerColumn},
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
	_node = &AdressBook{config: abuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, abuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adressbook.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	abuo.mutation.done = true
	return _node, nil
}
