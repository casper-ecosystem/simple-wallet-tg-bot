// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/settings"
)

// SettingsCreate is the builder for creating a Settings entity.
type SettingsCreate struct {
	config
	mutation *SettingsMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetLastScannedBlockNotificator sets the "last_scanned_block_notificator" field.
func (sc *SettingsCreate) SetLastScannedBlockNotificator(i int64) *SettingsCreate {
	sc.mutation.SetLastScannedBlockNotificator(i)
	return sc
}

// SetNillableLastScannedBlockNotificator sets the "last_scanned_block_notificator" field if the given value is not nil.
func (sc *SettingsCreate) SetNillableLastScannedBlockNotificator(i *int64) *SettingsCreate {
	if i != nil {
		sc.SetLastScannedBlockNotificator(*i)
	}
	return sc
}

// SetLastScannedEraValidators sets the "last_scanned_era_validators" field.
func (sc *SettingsCreate) SetLastScannedEraValidators(i int64) *SettingsCreate {
	sc.mutation.SetLastScannedEraValidators(i)
	return sc
}

// SetNillableLastScannedEraValidators sets the "last_scanned_era_validators" field if the given value is not nil.
func (sc *SettingsCreate) SetNillableLastScannedEraValidators(i *int64) *SettingsCreate {
	if i != nil {
		sc.SetLastScannedEraValidators(*i)
	}
	return sc
}

// Mutation returns the SettingsMutation object of the builder.
func (sc *SettingsCreate) Mutation() *SettingsMutation {
	return sc.mutation
}

// Save creates the Settings in the database.
func (sc *SettingsCreate) Save(ctx context.Context) (*Settings, error) {
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SettingsCreate) SaveX(ctx context.Context) *Settings {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SettingsCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SettingsCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SettingsCreate) check() error {
	return nil
}

func (sc *SettingsCreate) sqlSave(ctx context.Context) (*Settings, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SettingsCreate) createSpec() (*Settings, *sqlgraph.CreateSpec) {
	var (
		_node = &Settings{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(settings.Table, sqlgraph.NewFieldSpec(settings.FieldID, field.TypeInt))
	)
	_spec.OnConflict = sc.conflict
	if value, ok := sc.mutation.LastScannedBlockNotificator(); ok {
		_spec.SetField(settings.FieldLastScannedBlockNotificator, field.TypeInt64, value)
		_node.LastScannedBlockNotificator = value
	}
	if value, ok := sc.mutation.LastScannedEraValidators(); ok {
		_spec.SetField(settings.FieldLastScannedEraValidators, field.TypeInt64, value)
		_node.LastScannedEraValidators = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Settings.Create().
//		SetLastScannedBlockNotificator(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SettingsUpsert) {
//			SetLastScannedBlockNotificator(v+v).
//		}).
//		Exec(ctx)
func (sc *SettingsCreate) OnConflict(opts ...sql.ConflictOption) *SettingsUpsertOne {
	sc.conflict = opts
	return &SettingsUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Settings.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *SettingsCreate) OnConflictColumns(columns ...string) *SettingsUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &SettingsUpsertOne{
		create: sc,
	}
}

type (
	// SettingsUpsertOne is the builder for "upsert"-ing
	//  one Settings node.
	SettingsUpsertOne struct {
		create *SettingsCreate
	}

	// SettingsUpsert is the "OnConflict" setter.
	SettingsUpsert struct {
		*sql.UpdateSet
	}
)

// SetLastScannedBlockNotificator sets the "last_scanned_block_notificator" field.
func (u *SettingsUpsert) SetLastScannedBlockNotificator(v int64) *SettingsUpsert {
	u.Set(settings.FieldLastScannedBlockNotificator, v)
	return u
}

// UpdateLastScannedBlockNotificator sets the "last_scanned_block_notificator" field to the value that was provided on create.
func (u *SettingsUpsert) UpdateLastScannedBlockNotificator() *SettingsUpsert {
	u.SetExcluded(settings.FieldLastScannedBlockNotificator)
	return u
}

// AddLastScannedBlockNotificator adds v to the "last_scanned_block_notificator" field.
func (u *SettingsUpsert) AddLastScannedBlockNotificator(v int64) *SettingsUpsert {
	u.Add(settings.FieldLastScannedBlockNotificator, v)
	return u
}

// ClearLastScannedBlockNotificator clears the value of the "last_scanned_block_notificator" field.
func (u *SettingsUpsert) ClearLastScannedBlockNotificator() *SettingsUpsert {
	u.SetNull(settings.FieldLastScannedBlockNotificator)
	return u
}

// SetLastScannedEraValidators sets the "last_scanned_era_validators" field.
func (u *SettingsUpsert) SetLastScannedEraValidators(v int64) *SettingsUpsert {
	u.Set(settings.FieldLastScannedEraValidators, v)
	return u
}

// UpdateLastScannedEraValidators sets the "last_scanned_era_validators" field to the value that was provided on create.
func (u *SettingsUpsert) UpdateLastScannedEraValidators() *SettingsUpsert {
	u.SetExcluded(settings.FieldLastScannedEraValidators)
	return u
}

// AddLastScannedEraValidators adds v to the "last_scanned_era_validators" field.
func (u *SettingsUpsert) AddLastScannedEraValidators(v int64) *SettingsUpsert {
	u.Add(settings.FieldLastScannedEraValidators, v)
	return u
}

// ClearLastScannedEraValidators clears the value of the "last_scanned_era_validators" field.
func (u *SettingsUpsert) ClearLastScannedEraValidators() *SettingsUpsert {
	u.SetNull(settings.FieldLastScannedEraValidators)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Settings.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *SettingsUpsertOne) UpdateNewValues() *SettingsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Settings.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SettingsUpsertOne) Ignore() *SettingsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SettingsUpsertOne) DoNothing() *SettingsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SettingsCreate.OnConflict
// documentation for more info.
func (u *SettingsUpsertOne) Update(set func(*SettingsUpsert)) *SettingsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SettingsUpsert{UpdateSet: update})
	}))
	return u
}

// SetLastScannedBlockNotificator sets the "last_scanned_block_notificator" field.
func (u *SettingsUpsertOne) SetLastScannedBlockNotificator(v int64) *SettingsUpsertOne {
	return u.Update(func(s *SettingsUpsert) {
		s.SetLastScannedBlockNotificator(v)
	})
}

// AddLastScannedBlockNotificator adds v to the "last_scanned_block_notificator" field.
func (u *SettingsUpsertOne) AddLastScannedBlockNotificator(v int64) *SettingsUpsertOne {
	return u.Update(func(s *SettingsUpsert) {
		s.AddLastScannedBlockNotificator(v)
	})
}

// UpdateLastScannedBlockNotificator sets the "last_scanned_block_notificator" field to the value that was provided on create.
func (u *SettingsUpsertOne) UpdateLastScannedBlockNotificator() *SettingsUpsertOne {
	return u.Update(func(s *SettingsUpsert) {
		s.UpdateLastScannedBlockNotificator()
	})
}

// ClearLastScannedBlockNotificator clears the value of the "last_scanned_block_notificator" field.
func (u *SettingsUpsertOne) ClearLastScannedBlockNotificator() *SettingsUpsertOne {
	return u.Update(func(s *SettingsUpsert) {
		s.ClearLastScannedBlockNotificator()
	})
}

// SetLastScannedEraValidators sets the "last_scanned_era_validators" field.
func (u *SettingsUpsertOne) SetLastScannedEraValidators(v int64) *SettingsUpsertOne {
	return u.Update(func(s *SettingsUpsert) {
		s.SetLastScannedEraValidators(v)
	})
}

// AddLastScannedEraValidators adds v to the "last_scanned_era_validators" field.
func (u *SettingsUpsertOne) AddLastScannedEraValidators(v int64) *SettingsUpsertOne {
	return u.Update(func(s *SettingsUpsert) {
		s.AddLastScannedEraValidators(v)
	})
}

// UpdateLastScannedEraValidators sets the "last_scanned_era_validators" field to the value that was provided on create.
func (u *SettingsUpsertOne) UpdateLastScannedEraValidators() *SettingsUpsertOne {
	return u.Update(func(s *SettingsUpsert) {
		s.UpdateLastScannedEraValidators()
	})
}

// ClearLastScannedEraValidators clears the value of the "last_scanned_era_validators" field.
func (u *SettingsUpsertOne) ClearLastScannedEraValidators() *SettingsUpsertOne {
	return u.Update(func(s *SettingsUpsert) {
		s.ClearLastScannedEraValidators()
	})
}

// Exec executes the query.
func (u *SettingsUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SettingsCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SettingsUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SettingsUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SettingsUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SettingsCreateBulk is the builder for creating many Settings entities in bulk.
type SettingsCreateBulk struct {
	config
	builders []*SettingsCreate
	conflict []sql.ConflictOption
}

// Save creates the Settings entities in the database.
func (scb *SettingsCreateBulk) Save(ctx context.Context) ([]*Settings, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Settings, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SettingsMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SettingsCreateBulk) SaveX(ctx context.Context) []*Settings {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SettingsCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SettingsCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Settings.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SettingsUpsert) {
//			SetLastScannedBlockNotificator(v+v).
//		}).
//		Exec(ctx)
func (scb *SettingsCreateBulk) OnConflict(opts ...sql.ConflictOption) *SettingsUpsertBulk {
	scb.conflict = opts
	return &SettingsUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Settings.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *SettingsCreateBulk) OnConflictColumns(columns ...string) *SettingsUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &SettingsUpsertBulk{
		create: scb,
	}
}

// SettingsUpsertBulk is the builder for "upsert"-ing
// a bulk of Settings nodes.
type SettingsUpsertBulk struct {
	create *SettingsCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Settings.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *SettingsUpsertBulk) UpdateNewValues() *SettingsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Settings.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SettingsUpsertBulk) Ignore() *SettingsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SettingsUpsertBulk) DoNothing() *SettingsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SettingsCreateBulk.OnConflict
// documentation for more info.
func (u *SettingsUpsertBulk) Update(set func(*SettingsUpsert)) *SettingsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SettingsUpsert{UpdateSet: update})
	}))
	return u
}

// SetLastScannedBlockNotificator sets the "last_scanned_block_notificator" field.
func (u *SettingsUpsertBulk) SetLastScannedBlockNotificator(v int64) *SettingsUpsertBulk {
	return u.Update(func(s *SettingsUpsert) {
		s.SetLastScannedBlockNotificator(v)
	})
}

// AddLastScannedBlockNotificator adds v to the "last_scanned_block_notificator" field.
func (u *SettingsUpsertBulk) AddLastScannedBlockNotificator(v int64) *SettingsUpsertBulk {
	return u.Update(func(s *SettingsUpsert) {
		s.AddLastScannedBlockNotificator(v)
	})
}

// UpdateLastScannedBlockNotificator sets the "last_scanned_block_notificator" field to the value that was provided on create.
func (u *SettingsUpsertBulk) UpdateLastScannedBlockNotificator() *SettingsUpsertBulk {
	return u.Update(func(s *SettingsUpsert) {
		s.UpdateLastScannedBlockNotificator()
	})
}

// ClearLastScannedBlockNotificator clears the value of the "last_scanned_block_notificator" field.
func (u *SettingsUpsertBulk) ClearLastScannedBlockNotificator() *SettingsUpsertBulk {
	return u.Update(func(s *SettingsUpsert) {
		s.ClearLastScannedBlockNotificator()
	})
}

// SetLastScannedEraValidators sets the "last_scanned_era_validators" field.
func (u *SettingsUpsertBulk) SetLastScannedEraValidators(v int64) *SettingsUpsertBulk {
	return u.Update(func(s *SettingsUpsert) {
		s.SetLastScannedEraValidators(v)
	})
}

// AddLastScannedEraValidators adds v to the "last_scanned_era_validators" field.
func (u *SettingsUpsertBulk) AddLastScannedEraValidators(v int64) *SettingsUpsertBulk {
	return u.Update(func(s *SettingsUpsert) {
		s.AddLastScannedEraValidators(v)
	})
}

// UpdateLastScannedEraValidators sets the "last_scanned_era_validators" field to the value that was provided on create.
func (u *SettingsUpsertBulk) UpdateLastScannedEraValidators() *SettingsUpsertBulk {
	return u.Update(func(s *SettingsUpsert) {
		s.UpdateLastScannedEraValidators()
	})
}

// ClearLastScannedEraValidators clears the value of the "last_scanned_era_validators" field.
func (u *SettingsUpsertBulk) ClearLastScannedEraValidators() *SettingsUpsertBulk {
	return u.Update(func(s *SettingsUpsert) {
		s.ClearLastScannedEraValidators()
	})
}

// Exec executes the query.
func (u *SettingsUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SettingsCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SettingsCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SettingsUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
