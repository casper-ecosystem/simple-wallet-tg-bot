// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent"
)

// The AdressBookFunc type is an adapter to allow the use of ordinary
// function as AdressBook mutator.
type AdressBookFunc func(context.Context, *ent.AdressBookMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AdressBookFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AdressBookMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AdressBookMutation", m)
}

// The BalancesFunc type is an adapter to allow the use of ordinary
// function as Balances mutator.
type BalancesFunc func(context.Context, *ent.BalancesMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BalancesFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.BalancesMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BalancesMutation", m)
}

// The DelegatesFunc type is an adapter to allow the use of ordinary
// function as Delegates mutator.
type DelegatesFunc func(context.Context, *ent.DelegatesMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DelegatesFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.DelegatesMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DelegatesMutation", m)
}

// The InvoiceFunc type is an adapter to allow the use of ordinary
// function as Invoice mutator.
type InvoiceFunc func(context.Context, *ent.InvoiceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f InvoiceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.InvoiceMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.InvoiceMutation", m)
}

// The Invoices_paymentsFunc type is an adapter to allow the use of ordinary
// function as Invoices_payments mutator.
type Invoices_paymentsFunc func(context.Context, *ent.InvoicesPaymentsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f Invoices_paymentsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.InvoicesPaymentsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.InvoicesPaymentsMutation", m)
}

// The PrivateKeysFunc type is an adapter to allow the use of ordinary
// function as PrivateKeys mutator.
type PrivateKeysFunc func(context.Context, *ent.PrivateKeysMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PrivateKeysFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PrivateKeysMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PrivateKeysMutation", m)
}

// The RecentInvoicesFunc type is an adapter to allow the use of ordinary
// function as RecentInvoices mutator.
type RecentInvoicesFunc func(context.Context, *ent.RecentInvoicesMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RecentInvoicesFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.RecentInvoicesMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RecentInvoicesMutation", m)
}

// The RewardsDataFunc type is an adapter to allow the use of ordinary
// function as RewardsData mutator.
type RewardsDataFunc func(context.Context, *ent.RewardsDataMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RewardsDataFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.RewardsDataMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RewardsDataMutation", m)
}

// The SettingsFunc type is an adapter to allow the use of ordinary
// function as Settings mutator.
type SettingsFunc func(context.Context, *ent.SettingsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SettingsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SettingsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SettingsMutation", m)
}

// The SwapsFunc type is an adapter to allow the use of ordinary
// function as Swaps mutator.
type SwapsFunc func(context.Context, *ent.SwapsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SwapsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SwapsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SwapsMutation", m)
}

// The TaskFunc type is an adapter to allow the use of ordinary
// function as Task mutator.
type TaskFunc func(context.Context, *ent.TaskMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TaskFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.TaskMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TaskMutation", m)
}

// The TransfersFunc type is an adapter to allow the use of ordinary
// function as Transfers mutator.
type TransfersFunc func(context.Context, *ent.TransfersMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TransfersFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.TransfersMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TransfersMutation", m)
}

// The UndelegatesFunc type is an adapter to allow the use of ordinary
// function as Undelegates mutator.
type UndelegatesFunc func(context.Context, *ent.UndelegatesMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UndelegatesFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UndelegatesMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UndelegatesMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
}

// The UserStateFunc type is an adapter to allow the use of ordinary
// function as UserState mutator.
type UserStateFunc func(context.Context, *ent.UserStateMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserStateFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserStateMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserStateMutation", m)
}

// The ValidatorsFunc type is an adapter to allow the use of ordinary
// function as Validators mutator.
type ValidatorsFunc func(context.Context, *ent.ValidatorsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ValidatorsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ValidatorsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ValidatorsMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
