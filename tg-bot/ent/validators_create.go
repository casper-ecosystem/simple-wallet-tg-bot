// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Simplewallethq/tg-bot/ent/validators"
)

// ValidatorsCreate is the builder for creating a Validators entity.
type ValidatorsCreate struct {
	config
	mutation *ValidatorsMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAddress sets the "address" field.
func (vc *ValidatorsCreate) SetAddress(s string) *ValidatorsCreate {
	vc.mutation.SetAddress(s)
	return vc
}

// SetName sets the "name" field.
func (vc *ValidatorsCreate) SetName(s string) *ValidatorsCreate {
	vc.mutation.SetName(s)
	return vc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (vc *ValidatorsCreate) SetNillableName(s *string) *ValidatorsCreate {
	if s != nil {
		vc.SetName(*s)
	}
	return vc
}

// SetFee sets the "fee" field.
func (vc *ValidatorsCreate) SetFee(i int8) *ValidatorsCreate {
	vc.mutation.SetFee(i)
	return vc
}

// SetNillableFee sets the "fee" field if the given value is not nil.
func (vc *ValidatorsCreate) SetNillableFee(i *int8) *ValidatorsCreate {
	if i != nil {
		vc.SetFee(*i)
	}
	return vc
}

// SetDelegators sets the "delegators" field.
func (vc *ValidatorsCreate) SetDelegators(i int64) *ValidatorsCreate {
	vc.mutation.SetDelegators(i)
	return vc
}

// SetNillableDelegators sets the "delegators" field if the given value is not nil.
func (vc *ValidatorsCreate) SetNillableDelegators(i *int64) *ValidatorsCreate {
	if i != nil {
		vc.SetDelegators(*i)
	}
	return vc
}

// SetActive sets the "active" field.
func (vc *ValidatorsCreate) SetActive(b bool) *ValidatorsCreate {
	vc.mutation.SetActive(b)
	return vc
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (vc *ValidatorsCreate) SetNillableActive(b *bool) *ValidatorsCreate {
	if b != nil {
		vc.SetActive(*b)
	}
	return vc
}

// Mutation returns the ValidatorsMutation object of the builder.
func (vc *ValidatorsCreate) Mutation() *ValidatorsMutation {
	return vc.mutation
}

// Save creates the Validators in the database.
func (vc *ValidatorsCreate) Save(ctx context.Context) (*Validators, error) {
	return withHooks(ctx, vc.sqlSave, vc.mutation, vc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vc *ValidatorsCreate) SaveX(ctx context.Context) *Validators {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *ValidatorsCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *ValidatorsCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *ValidatorsCreate) check() error {
	if _, ok := vc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "Validators.address"`)}
	}
	return nil
}

func (vc *ValidatorsCreate) sqlSave(ctx context.Context) (*Validators, error) {
	if err := vc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	vc.mutation.id = &_node.ID
	vc.mutation.done = true
	return _node, nil
}

func (vc *ValidatorsCreate) createSpec() (*Validators, *sqlgraph.CreateSpec) {
	var (
		_node = &Validators{config: vc.config}
		_spec = sqlgraph.NewCreateSpec(validators.Table, sqlgraph.NewFieldSpec(validators.FieldID, field.TypeInt))
	)
	_spec.OnConflict = vc.conflict
	if value, ok := vc.mutation.Address(); ok {
		_spec.SetField(validators.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := vc.mutation.Name(); ok {
		_spec.SetField(validators.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := vc.mutation.Fee(); ok {
		_spec.SetField(validators.FieldFee, field.TypeInt8, value)
		_node.Fee = value
	}
	if value, ok := vc.mutation.Delegators(); ok {
		_spec.SetField(validators.FieldDelegators, field.TypeInt64, value)
		_node.Delegators = value
	}
	if value, ok := vc.mutation.Active(); ok {
		_spec.SetField(validators.FieldActive, field.TypeBool, value)
		_node.Active = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Validators.Create().
//		SetAddress(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ValidatorsUpsert) {
//			SetAddress(v+v).
//		}).
//		Exec(ctx)
func (vc *ValidatorsCreate) OnConflict(opts ...sql.ConflictOption) *ValidatorsUpsertOne {
	vc.conflict = opts
	return &ValidatorsUpsertOne{
		create: vc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Validators.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (vc *ValidatorsCreate) OnConflictColumns(columns ...string) *ValidatorsUpsertOne {
	vc.conflict = append(vc.conflict, sql.ConflictColumns(columns...))
	return &ValidatorsUpsertOne{
		create: vc,
	}
}

type (
	// ValidatorsUpsertOne is the builder for "upsert"-ing
	//  one Validators node.
	ValidatorsUpsertOne struct {
		create *ValidatorsCreate
	}

	// ValidatorsUpsert is the "OnConflict" setter.
	ValidatorsUpsert struct {
		*sql.UpdateSet
	}
)

// SetAddress sets the "address" field.
func (u *ValidatorsUpsert) SetAddress(v string) *ValidatorsUpsert {
	u.Set(validators.FieldAddress, v)
	return u
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *ValidatorsUpsert) UpdateAddress() *ValidatorsUpsert {
	u.SetExcluded(validators.FieldAddress)
	return u
}

// SetName sets the "name" field.
func (u *ValidatorsUpsert) SetName(v string) *ValidatorsUpsert {
	u.Set(validators.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ValidatorsUpsert) UpdateName() *ValidatorsUpsert {
	u.SetExcluded(validators.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *ValidatorsUpsert) ClearName() *ValidatorsUpsert {
	u.SetNull(validators.FieldName)
	return u
}

// SetFee sets the "fee" field.
func (u *ValidatorsUpsert) SetFee(v int8) *ValidatorsUpsert {
	u.Set(validators.FieldFee, v)
	return u
}

// UpdateFee sets the "fee" field to the value that was provided on create.
func (u *ValidatorsUpsert) UpdateFee() *ValidatorsUpsert {
	u.SetExcluded(validators.FieldFee)
	return u
}

// AddFee adds v to the "fee" field.
func (u *ValidatorsUpsert) AddFee(v int8) *ValidatorsUpsert {
	u.Add(validators.FieldFee, v)
	return u
}

// ClearFee clears the value of the "fee" field.
func (u *ValidatorsUpsert) ClearFee() *ValidatorsUpsert {
	u.SetNull(validators.FieldFee)
	return u
}

// SetDelegators sets the "delegators" field.
func (u *ValidatorsUpsert) SetDelegators(v int64) *ValidatorsUpsert {
	u.Set(validators.FieldDelegators, v)
	return u
}

// UpdateDelegators sets the "delegators" field to the value that was provided on create.
func (u *ValidatorsUpsert) UpdateDelegators() *ValidatorsUpsert {
	u.SetExcluded(validators.FieldDelegators)
	return u
}

// AddDelegators adds v to the "delegators" field.
func (u *ValidatorsUpsert) AddDelegators(v int64) *ValidatorsUpsert {
	u.Add(validators.FieldDelegators, v)
	return u
}

// ClearDelegators clears the value of the "delegators" field.
func (u *ValidatorsUpsert) ClearDelegators() *ValidatorsUpsert {
	u.SetNull(validators.FieldDelegators)
	return u
}

// SetActive sets the "active" field.
func (u *ValidatorsUpsert) SetActive(v bool) *ValidatorsUpsert {
	u.Set(validators.FieldActive, v)
	return u
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *ValidatorsUpsert) UpdateActive() *ValidatorsUpsert {
	u.SetExcluded(validators.FieldActive)
	return u
}

// ClearActive clears the value of the "active" field.
func (u *ValidatorsUpsert) ClearActive() *ValidatorsUpsert {
	u.SetNull(validators.FieldActive)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Validators.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ValidatorsUpsertOne) UpdateNewValues() *ValidatorsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Validators.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ValidatorsUpsertOne) Ignore() *ValidatorsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ValidatorsUpsertOne) DoNothing() *ValidatorsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ValidatorsCreate.OnConflict
// documentation for more info.
func (u *ValidatorsUpsertOne) Update(set func(*ValidatorsUpsert)) *ValidatorsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ValidatorsUpsert{UpdateSet: update})
	}))
	return u
}

// SetAddress sets the "address" field.
func (u *ValidatorsUpsertOne) SetAddress(v string) *ValidatorsUpsertOne {
	return u.Update(func(s *ValidatorsUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *ValidatorsUpsertOne) UpdateAddress() *ValidatorsUpsertOne {
	return u.Update(func(s *ValidatorsUpsert) {
		s.UpdateAddress()
	})
}

// SetName sets the "name" field.
func (u *ValidatorsUpsertOne) SetName(v string) *ValidatorsUpsertOne {
	return u.Update(func(s *ValidatorsUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ValidatorsUpsertOne) UpdateName() *ValidatorsUpsertOne {
	return u.Update(func(s *ValidatorsUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *ValidatorsUpsertOne) ClearName() *ValidatorsUpsertOne {
	return u.Update(func(s *ValidatorsUpsert) {
		s.ClearName()
	})
}

// SetFee sets the "fee" field.
func (u *ValidatorsUpsertOne) SetFee(v int8) *ValidatorsUpsertOne {
	return u.Update(func(s *ValidatorsUpsert) {
		s.SetFee(v)
	})
}

// AddFee adds v to the "fee" field.
func (u *ValidatorsUpsertOne) AddFee(v int8) *ValidatorsUpsertOne {
	return u.Update(func(s *ValidatorsUpsert) {
		s.AddFee(v)
	})
}

// UpdateFee sets the "fee" field to the value that was provided on create.
func (u *ValidatorsUpsertOne) UpdateFee() *ValidatorsUpsertOne {
	return u.Update(func(s *ValidatorsUpsert) {
		s.UpdateFee()
	})
}

// ClearFee clears the value of the "fee" field.
func (u *ValidatorsUpsertOne) ClearFee() *ValidatorsUpsertOne {
	return u.Update(func(s *ValidatorsUpsert) {
		s.ClearFee()
	})
}

// SetDelegators sets the "delegators" field.
func (u *ValidatorsUpsertOne) SetDelegators(v int64) *ValidatorsUpsertOne {
	return u.Update(func(s *ValidatorsUpsert) {
		s.SetDelegators(v)
	})
}

// AddDelegators adds v to the "delegators" field.
func (u *ValidatorsUpsertOne) AddDelegators(v int64) *ValidatorsUpsertOne {
	return u.Update(func(s *ValidatorsUpsert) {
		s.AddDelegators(v)
	})
}

// UpdateDelegators sets the "delegators" field to the value that was provided on create.
func (u *ValidatorsUpsertOne) UpdateDelegators() *ValidatorsUpsertOne {
	return u.Update(func(s *ValidatorsUpsert) {
		s.UpdateDelegators()
	})
}

// ClearDelegators clears the value of the "delegators" field.
func (u *ValidatorsUpsertOne) ClearDelegators() *ValidatorsUpsertOne {
	return u.Update(func(s *ValidatorsUpsert) {
		s.ClearDelegators()
	})
}

// SetActive sets the "active" field.
func (u *ValidatorsUpsertOne) SetActive(v bool) *ValidatorsUpsertOne {
	return u.Update(func(s *ValidatorsUpsert) {
		s.SetActive(v)
	})
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *ValidatorsUpsertOne) UpdateActive() *ValidatorsUpsertOne {
	return u.Update(func(s *ValidatorsUpsert) {
		s.UpdateActive()
	})
}

// ClearActive clears the value of the "active" field.
func (u *ValidatorsUpsertOne) ClearActive() *ValidatorsUpsertOne {
	return u.Update(func(s *ValidatorsUpsert) {
		s.ClearActive()
	})
}

// Exec executes the query.
func (u *ValidatorsUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ValidatorsCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ValidatorsUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ValidatorsUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ValidatorsUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ValidatorsCreateBulk is the builder for creating many Validators entities in bulk.
type ValidatorsCreateBulk struct {
	config
	builders []*ValidatorsCreate
	conflict []sql.ConflictOption
}

// Save creates the Validators entities in the database.
func (vcb *ValidatorsCreateBulk) Save(ctx context.Context) ([]*Validators, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Validators, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ValidatorsMutation)
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
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = vcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *ValidatorsCreateBulk) SaveX(ctx context.Context) []*Validators {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *ValidatorsCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *ValidatorsCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Validators.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ValidatorsUpsert) {
//			SetAddress(v+v).
//		}).
//		Exec(ctx)
func (vcb *ValidatorsCreateBulk) OnConflict(opts ...sql.ConflictOption) *ValidatorsUpsertBulk {
	vcb.conflict = opts
	return &ValidatorsUpsertBulk{
		create: vcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Validators.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (vcb *ValidatorsCreateBulk) OnConflictColumns(columns ...string) *ValidatorsUpsertBulk {
	vcb.conflict = append(vcb.conflict, sql.ConflictColumns(columns...))
	return &ValidatorsUpsertBulk{
		create: vcb,
	}
}

// ValidatorsUpsertBulk is the builder for "upsert"-ing
// a bulk of Validators nodes.
type ValidatorsUpsertBulk struct {
	create *ValidatorsCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Validators.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ValidatorsUpsertBulk) UpdateNewValues() *ValidatorsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Validators.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ValidatorsUpsertBulk) Ignore() *ValidatorsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ValidatorsUpsertBulk) DoNothing() *ValidatorsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ValidatorsCreateBulk.OnConflict
// documentation for more info.
func (u *ValidatorsUpsertBulk) Update(set func(*ValidatorsUpsert)) *ValidatorsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ValidatorsUpsert{UpdateSet: update})
	}))
	return u
}

// SetAddress sets the "address" field.
func (u *ValidatorsUpsertBulk) SetAddress(v string) *ValidatorsUpsertBulk {
	return u.Update(func(s *ValidatorsUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *ValidatorsUpsertBulk) UpdateAddress() *ValidatorsUpsertBulk {
	return u.Update(func(s *ValidatorsUpsert) {
		s.UpdateAddress()
	})
}

// SetName sets the "name" field.
func (u *ValidatorsUpsertBulk) SetName(v string) *ValidatorsUpsertBulk {
	return u.Update(func(s *ValidatorsUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ValidatorsUpsertBulk) UpdateName() *ValidatorsUpsertBulk {
	return u.Update(func(s *ValidatorsUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *ValidatorsUpsertBulk) ClearName() *ValidatorsUpsertBulk {
	return u.Update(func(s *ValidatorsUpsert) {
		s.ClearName()
	})
}

// SetFee sets the "fee" field.
func (u *ValidatorsUpsertBulk) SetFee(v int8) *ValidatorsUpsertBulk {
	return u.Update(func(s *ValidatorsUpsert) {
		s.SetFee(v)
	})
}

// AddFee adds v to the "fee" field.
func (u *ValidatorsUpsertBulk) AddFee(v int8) *ValidatorsUpsertBulk {
	return u.Update(func(s *ValidatorsUpsert) {
		s.AddFee(v)
	})
}

// UpdateFee sets the "fee" field to the value that was provided on create.
func (u *ValidatorsUpsertBulk) UpdateFee() *ValidatorsUpsertBulk {
	return u.Update(func(s *ValidatorsUpsert) {
		s.UpdateFee()
	})
}

// ClearFee clears the value of the "fee" field.
func (u *ValidatorsUpsertBulk) ClearFee() *ValidatorsUpsertBulk {
	return u.Update(func(s *ValidatorsUpsert) {
		s.ClearFee()
	})
}

// SetDelegators sets the "delegators" field.
func (u *ValidatorsUpsertBulk) SetDelegators(v int64) *ValidatorsUpsertBulk {
	return u.Update(func(s *ValidatorsUpsert) {
		s.SetDelegators(v)
	})
}

// AddDelegators adds v to the "delegators" field.
func (u *ValidatorsUpsertBulk) AddDelegators(v int64) *ValidatorsUpsertBulk {
	return u.Update(func(s *ValidatorsUpsert) {
		s.AddDelegators(v)
	})
}

// UpdateDelegators sets the "delegators" field to the value that was provided on create.
func (u *ValidatorsUpsertBulk) UpdateDelegators() *ValidatorsUpsertBulk {
	return u.Update(func(s *ValidatorsUpsert) {
		s.UpdateDelegators()
	})
}

// ClearDelegators clears the value of the "delegators" field.
func (u *ValidatorsUpsertBulk) ClearDelegators() *ValidatorsUpsertBulk {
	return u.Update(func(s *ValidatorsUpsert) {
		s.ClearDelegators()
	})
}

// SetActive sets the "active" field.
func (u *ValidatorsUpsertBulk) SetActive(v bool) *ValidatorsUpsertBulk {
	return u.Update(func(s *ValidatorsUpsert) {
		s.SetActive(v)
	})
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *ValidatorsUpsertBulk) UpdateActive() *ValidatorsUpsertBulk {
	return u.Update(func(s *ValidatorsUpsert) {
		s.UpdateActive()
	})
}

// ClearActive clears the value of the "active" field.
func (u *ValidatorsUpsertBulk) ClearActive() *ValidatorsUpsertBulk {
	return u.Update(func(s *ValidatorsUpsert) {
		s.ClearActive()
	})
}

// Exec executes the query.
func (u *ValidatorsUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ValidatorsCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ValidatorsCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ValidatorsUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}