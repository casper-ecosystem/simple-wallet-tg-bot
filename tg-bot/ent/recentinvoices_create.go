// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/recentinvoices"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/user"
)

// RecentInvoicesCreate is the builder for creating a RecentInvoices entity.
type RecentInvoicesCreate struct {
	config
	mutation *RecentInvoicesMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetStatus sets the "status" field.
func (ric *RecentInvoicesCreate) SetStatus(s string) *RecentInvoicesCreate {
	ric.mutation.SetStatus(s)
	return ric
}

// SetInvoiceID sets the "invoiceID" field.
func (ric *RecentInvoicesCreate) SetInvoiceID(i int64) *RecentInvoicesCreate {
	ric.mutation.SetInvoiceID(i)
	return ric
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (ric *RecentInvoicesCreate) SetOwnerID(id int64) *RecentInvoicesCreate {
	ric.mutation.SetOwnerID(id)
	return ric
}

// SetOwner sets the "owner" edge to the User entity.
func (ric *RecentInvoicesCreate) SetOwner(u *User) *RecentInvoicesCreate {
	return ric.SetOwnerID(u.ID)
}

// Mutation returns the RecentInvoicesMutation object of the builder.
func (ric *RecentInvoicesCreate) Mutation() *RecentInvoicesMutation {
	return ric.mutation
}

// Save creates the RecentInvoices in the database.
func (ric *RecentInvoicesCreate) Save(ctx context.Context) (*RecentInvoices, error) {
	return withHooks(ctx, ric.sqlSave, ric.mutation, ric.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ric *RecentInvoicesCreate) SaveX(ctx context.Context) *RecentInvoices {
	v, err := ric.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ric *RecentInvoicesCreate) Exec(ctx context.Context) error {
	_, err := ric.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ric *RecentInvoicesCreate) ExecX(ctx context.Context) {
	if err := ric.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ric *RecentInvoicesCreate) check() error {
	if _, ok := ric.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "RecentInvoices.status"`)}
	}
	if _, ok := ric.mutation.InvoiceID(); !ok {
		return &ValidationError{Name: "invoiceID", err: errors.New(`ent: missing required field "RecentInvoices.invoiceID"`)}
	}
	if _, ok := ric.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "RecentInvoices.owner"`)}
	}
	return nil
}

func (ric *RecentInvoicesCreate) sqlSave(ctx context.Context) (*RecentInvoices, error) {
	if err := ric.check(); err != nil {
		return nil, err
	}
	_node, _spec := ric.createSpec()
	if err := sqlgraph.CreateNode(ctx, ric.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ric.mutation.id = &_node.ID
	ric.mutation.done = true
	return _node, nil
}

func (ric *RecentInvoicesCreate) createSpec() (*RecentInvoices, *sqlgraph.CreateSpec) {
	var (
		_node = &RecentInvoices{config: ric.config}
		_spec = sqlgraph.NewCreateSpec(recentinvoices.Table, sqlgraph.NewFieldSpec(recentinvoices.FieldID, field.TypeInt))
	)
	_spec.OnConflict = ric.conflict
	if value, ok := ric.mutation.Status(); ok {
		_spec.SetField(recentinvoices.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := ric.mutation.InvoiceID(); ok {
		_spec.SetField(recentinvoices.FieldInvoiceID, field.TypeInt64, value)
		_node.InvoiceID = value
	}
	if nodes := ric.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recentinvoices.OwnerTable,
			Columns: []string{recentinvoices.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_recent_invoices = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.RecentInvoices.Create().
//		SetStatus(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RecentInvoicesUpsert) {
//			SetStatus(v+v).
//		}).
//		Exec(ctx)
func (ric *RecentInvoicesCreate) OnConflict(opts ...sql.ConflictOption) *RecentInvoicesUpsertOne {
	ric.conflict = opts
	return &RecentInvoicesUpsertOne{
		create: ric,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.RecentInvoices.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ric *RecentInvoicesCreate) OnConflictColumns(columns ...string) *RecentInvoicesUpsertOne {
	ric.conflict = append(ric.conflict, sql.ConflictColumns(columns...))
	return &RecentInvoicesUpsertOne{
		create: ric,
	}
}

type (
	// RecentInvoicesUpsertOne is the builder for "upsert"-ing
	//  one RecentInvoices node.
	RecentInvoicesUpsertOne struct {
		create *RecentInvoicesCreate
	}

	// RecentInvoicesUpsert is the "OnConflict" setter.
	RecentInvoicesUpsert struct {
		*sql.UpdateSet
	}
)

// SetStatus sets the "status" field.
func (u *RecentInvoicesUpsert) SetStatus(v string) *RecentInvoicesUpsert {
	u.Set(recentinvoices.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *RecentInvoicesUpsert) UpdateStatus() *RecentInvoicesUpsert {
	u.SetExcluded(recentinvoices.FieldStatus)
	return u
}

// SetInvoiceID sets the "invoiceID" field.
func (u *RecentInvoicesUpsert) SetInvoiceID(v int64) *RecentInvoicesUpsert {
	u.Set(recentinvoices.FieldInvoiceID, v)
	return u
}

// UpdateInvoiceID sets the "invoiceID" field to the value that was provided on create.
func (u *RecentInvoicesUpsert) UpdateInvoiceID() *RecentInvoicesUpsert {
	u.SetExcluded(recentinvoices.FieldInvoiceID)
	return u
}

// AddInvoiceID adds v to the "invoiceID" field.
func (u *RecentInvoicesUpsert) AddInvoiceID(v int64) *RecentInvoicesUpsert {
	u.Add(recentinvoices.FieldInvoiceID, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.RecentInvoices.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *RecentInvoicesUpsertOne) UpdateNewValues() *RecentInvoicesUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.RecentInvoices.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *RecentInvoicesUpsertOne) Ignore() *RecentInvoicesUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RecentInvoicesUpsertOne) DoNothing() *RecentInvoicesUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RecentInvoicesCreate.OnConflict
// documentation for more info.
func (u *RecentInvoicesUpsertOne) Update(set func(*RecentInvoicesUpsert)) *RecentInvoicesUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RecentInvoicesUpsert{UpdateSet: update})
	}))
	return u
}

// SetStatus sets the "status" field.
func (u *RecentInvoicesUpsertOne) SetStatus(v string) *RecentInvoicesUpsertOne {
	return u.Update(func(s *RecentInvoicesUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *RecentInvoicesUpsertOne) UpdateStatus() *RecentInvoicesUpsertOne {
	return u.Update(func(s *RecentInvoicesUpsert) {
		s.UpdateStatus()
	})
}

// SetInvoiceID sets the "invoiceID" field.
func (u *RecentInvoicesUpsertOne) SetInvoiceID(v int64) *RecentInvoicesUpsertOne {
	return u.Update(func(s *RecentInvoicesUpsert) {
		s.SetInvoiceID(v)
	})
}

// AddInvoiceID adds v to the "invoiceID" field.
func (u *RecentInvoicesUpsertOne) AddInvoiceID(v int64) *RecentInvoicesUpsertOne {
	return u.Update(func(s *RecentInvoicesUpsert) {
		s.AddInvoiceID(v)
	})
}

// UpdateInvoiceID sets the "invoiceID" field to the value that was provided on create.
func (u *RecentInvoicesUpsertOne) UpdateInvoiceID() *RecentInvoicesUpsertOne {
	return u.Update(func(s *RecentInvoicesUpsert) {
		s.UpdateInvoiceID()
	})
}

// Exec executes the query.
func (u *RecentInvoicesUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RecentInvoicesCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RecentInvoicesUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *RecentInvoicesUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *RecentInvoicesUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// RecentInvoicesCreateBulk is the builder for creating many RecentInvoices entities in bulk.
type RecentInvoicesCreateBulk struct {
	config
	builders []*RecentInvoicesCreate
	conflict []sql.ConflictOption
}

// Save creates the RecentInvoices entities in the database.
func (ricb *RecentInvoicesCreateBulk) Save(ctx context.Context) ([]*RecentInvoices, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ricb.builders))
	nodes := make([]*RecentInvoices, len(ricb.builders))
	mutators := make([]Mutator, len(ricb.builders))
	for i := range ricb.builders {
		func(i int, root context.Context) {
			builder := ricb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RecentInvoicesMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ricb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ricb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ricb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ricb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ricb *RecentInvoicesCreateBulk) SaveX(ctx context.Context) []*RecentInvoices {
	v, err := ricb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ricb *RecentInvoicesCreateBulk) Exec(ctx context.Context) error {
	_, err := ricb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ricb *RecentInvoicesCreateBulk) ExecX(ctx context.Context) {
	if err := ricb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.RecentInvoices.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RecentInvoicesUpsert) {
//			SetStatus(v+v).
//		}).
//		Exec(ctx)
func (ricb *RecentInvoicesCreateBulk) OnConflict(opts ...sql.ConflictOption) *RecentInvoicesUpsertBulk {
	ricb.conflict = opts
	return &RecentInvoicesUpsertBulk{
		create: ricb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.RecentInvoices.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ricb *RecentInvoicesCreateBulk) OnConflictColumns(columns ...string) *RecentInvoicesUpsertBulk {
	ricb.conflict = append(ricb.conflict, sql.ConflictColumns(columns...))
	return &RecentInvoicesUpsertBulk{
		create: ricb,
	}
}

// RecentInvoicesUpsertBulk is the builder for "upsert"-ing
// a bulk of RecentInvoices nodes.
type RecentInvoicesUpsertBulk struct {
	create *RecentInvoicesCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.RecentInvoices.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *RecentInvoicesUpsertBulk) UpdateNewValues() *RecentInvoicesUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.RecentInvoices.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *RecentInvoicesUpsertBulk) Ignore() *RecentInvoicesUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RecentInvoicesUpsertBulk) DoNothing() *RecentInvoicesUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RecentInvoicesCreateBulk.OnConflict
// documentation for more info.
func (u *RecentInvoicesUpsertBulk) Update(set func(*RecentInvoicesUpsert)) *RecentInvoicesUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RecentInvoicesUpsert{UpdateSet: update})
	}))
	return u
}

// SetStatus sets the "status" field.
func (u *RecentInvoicesUpsertBulk) SetStatus(v string) *RecentInvoicesUpsertBulk {
	return u.Update(func(s *RecentInvoicesUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *RecentInvoicesUpsertBulk) UpdateStatus() *RecentInvoicesUpsertBulk {
	return u.Update(func(s *RecentInvoicesUpsert) {
		s.UpdateStatus()
	})
}

// SetInvoiceID sets the "invoiceID" field.
func (u *RecentInvoicesUpsertBulk) SetInvoiceID(v int64) *RecentInvoicesUpsertBulk {
	return u.Update(func(s *RecentInvoicesUpsert) {
		s.SetInvoiceID(v)
	})
}

// AddInvoiceID adds v to the "invoiceID" field.
func (u *RecentInvoicesUpsertBulk) AddInvoiceID(v int64) *RecentInvoicesUpsertBulk {
	return u.Update(func(s *RecentInvoicesUpsert) {
		s.AddInvoiceID(v)
	})
}

// UpdateInvoiceID sets the "invoiceID" field to the value that was provided on create.
func (u *RecentInvoicesUpsertBulk) UpdateInvoiceID() *RecentInvoicesUpsertBulk {
	return u.Update(func(s *RecentInvoicesUpsert) {
		s.UpdateInvoiceID()
	})
}

// Exec executes the query.
func (u *RecentInvoicesUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the RecentInvoicesCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RecentInvoicesCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RecentInvoicesUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
