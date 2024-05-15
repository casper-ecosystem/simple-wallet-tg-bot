// Code generated by ent, DO NOT EDIT.

package transfers

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Transfers {
	return predicate.Transfers(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Transfers {
	return predicate.Transfers(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Transfers {
	return predicate.Transfers(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Transfers {
	return predicate.Transfers(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Transfers {
	return predicate.Transfers(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Transfers {
	return predicate.Transfers(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Transfers {
	return predicate.Transfers(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Transfers {
	return predicate.Transfers(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Transfers {
	return predicate.Transfers(sql.FieldLTE(FieldID, id))
}

// FromPubkey applies equality check predicate on the "from_pubkey" field. It's identical to FromPubkeyEQ.
func FromPubkey(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldEQ(FieldFromPubkey, v))
}

// ToPubkey applies equality check predicate on the "to_pubkey" field. It's identical to ToPubkeyEQ.
func ToPubkey(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldEQ(FieldToPubkey, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldEQ(FieldName, v))
}

// SenderBalance applies equality check predicate on the "sender_balance" field. It's identical to SenderBalanceEQ.
func SenderBalance(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldEQ(FieldSenderBalance, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldEQ(FieldAmount, v))
}

// MemoID applies equality check predicate on the "memo_id" field. It's identical to MemoIDEQ.
func MemoID(v uint64) predicate.Transfers {
	return predicate.Transfers(sql.FieldEQ(FieldMemoID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Transfers {
	return predicate.Transfers(sql.FieldEQ(FieldCreatedAt, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldEQ(FieldStatus, v))
}

// Deploy applies equality check predicate on the "Deploy" field. It's identical to DeployEQ.
func Deploy(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldEQ(FieldDeploy, v))
}

// AdditionalType applies equality check predicate on the "AdditionalType" field. It's identical to AdditionalTypeEQ.
func AdditionalType(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldEQ(FieldAdditionalType, v))
}

// InvoiceID applies equality check predicate on the "invoiceID" field. It's identical to InvoiceIDEQ.
func InvoiceID(v int64) predicate.Transfers {
	return predicate.Transfers(sql.FieldEQ(FieldInvoiceID, v))
}

// FromPubkeyEQ applies the EQ predicate on the "from_pubkey" field.
func FromPubkeyEQ(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldEQ(FieldFromPubkey, v))
}

// FromPubkeyNEQ applies the NEQ predicate on the "from_pubkey" field.
func FromPubkeyNEQ(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldNEQ(FieldFromPubkey, v))
}

// FromPubkeyIn applies the In predicate on the "from_pubkey" field.
func FromPubkeyIn(vs ...string) predicate.Transfers {
	return predicate.Transfers(sql.FieldIn(FieldFromPubkey, vs...))
}

// FromPubkeyNotIn applies the NotIn predicate on the "from_pubkey" field.
func FromPubkeyNotIn(vs ...string) predicate.Transfers {
	return predicate.Transfers(sql.FieldNotIn(FieldFromPubkey, vs...))
}

// FromPubkeyGT applies the GT predicate on the "from_pubkey" field.
func FromPubkeyGT(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldGT(FieldFromPubkey, v))
}

// FromPubkeyGTE applies the GTE predicate on the "from_pubkey" field.
func FromPubkeyGTE(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldGTE(FieldFromPubkey, v))
}

// FromPubkeyLT applies the LT predicate on the "from_pubkey" field.
func FromPubkeyLT(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldLT(FieldFromPubkey, v))
}

// FromPubkeyLTE applies the LTE predicate on the "from_pubkey" field.
func FromPubkeyLTE(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldLTE(FieldFromPubkey, v))
}

// FromPubkeyContains applies the Contains predicate on the "from_pubkey" field.
func FromPubkeyContains(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldContains(FieldFromPubkey, v))
}

// FromPubkeyHasPrefix applies the HasPrefix predicate on the "from_pubkey" field.
func FromPubkeyHasPrefix(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldHasPrefix(FieldFromPubkey, v))
}

// FromPubkeyHasSuffix applies the HasSuffix predicate on the "from_pubkey" field.
func FromPubkeyHasSuffix(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldHasSuffix(FieldFromPubkey, v))
}

// FromPubkeyEqualFold applies the EqualFold predicate on the "from_pubkey" field.
func FromPubkeyEqualFold(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldEqualFold(FieldFromPubkey, v))
}

// FromPubkeyContainsFold applies the ContainsFold predicate on the "from_pubkey" field.
func FromPubkeyContainsFold(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldContainsFold(FieldFromPubkey, v))
}

// ToPubkeyEQ applies the EQ predicate on the "to_pubkey" field.
func ToPubkeyEQ(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldEQ(FieldToPubkey, v))
}

// ToPubkeyNEQ applies the NEQ predicate on the "to_pubkey" field.
func ToPubkeyNEQ(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldNEQ(FieldToPubkey, v))
}

// ToPubkeyIn applies the In predicate on the "to_pubkey" field.
func ToPubkeyIn(vs ...string) predicate.Transfers {
	return predicate.Transfers(sql.FieldIn(FieldToPubkey, vs...))
}

// ToPubkeyNotIn applies the NotIn predicate on the "to_pubkey" field.
func ToPubkeyNotIn(vs ...string) predicate.Transfers {
	return predicate.Transfers(sql.FieldNotIn(FieldToPubkey, vs...))
}

// ToPubkeyGT applies the GT predicate on the "to_pubkey" field.
func ToPubkeyGT(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldGT(FieldToPubkey, v))
}

// ToPubkeyGTE applies the GTE predicate on the "to_pubkey" field.
func ToPubkeyGTE(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldGTE(FieldToPubkey, v))
}

// ToPubkeyLT applies the LT predicate on the "to_pubkey" field.
func ToPubkeyLT(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldLT(FieldToPubkey, v))
}

// ToPubkeyLTE applies the LTE predicate on the "to_pubkey" field.
func ToPubkeyLTE(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldLTE(FieldToPubkey, v))
}

// ToPubkeyContains applies the Contains predicate on the "to_pubkey" field.
func ToPubkeyContains(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldContains(FieldToPubkey, v))
}

// ToPubkeyHasPrefix applies the HasPrefix predicate on the "to_pubkey" field.
func ToPubkeyHasPrefix(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldHasPrefix(FieldToPubkey, v))
}

// ToPubkeyHasSuffix applies the HasSuffix predicate on the "to_pubkey" field.
func ToPubkeyHasSuffix(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldHasSuffix(FieldToPubkey, v))
}

// ToPubkeyIsNil applies the IsNil predicate on the "to_pubkey" field.
func ToPubkeyIsNil() predicate.Transfers {
	return predicate.Transfers(sql.FieldIsNull(FieldToPubkey))
}

// ToPubkeyNotNil applies the NotNil predicate on the "to_pubkey" field.
func ToPubkeyNotNil() predicate.Transfers {
	return predicate.Transfers(sql.FieldNotNull(FieldToPubkey))
}

// ToPubkeyEqualFold applies the EqualFold predicate on the "to_pubkey" field.
func ToPubkeyEqualFold(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldEqualFold(FieldToPubkey, v))
}

// ToPubkeyContainsFold applies the ContainsFold predicate on the "to_pubkey" field.
func ToPubkeyContainsFold(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldContainsFold(FieldToPubkey, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Transfers {
	return predicate.Transfers(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Transfers {
	return predicate.Transfers(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.Transfers {
	return predicate.Transfers(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.Transfers {
	return predicate.Transfers(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldContainsFold(FieldName, v))
}

// SenderBalanceEQ applies the EQ predicate on the "sender_balance" field.
func SenderBalanceEQ(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldEQ(FieldSenderBalance, v))
}

// SenderBalanceNEQ applies the NEQ predicate on the "sender_balance" field.
func SenderBalanceNEQ(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldNEQ(FieldSenderBalance, v))
}

// SenderBalanceIn applies the In predicate on the "sender_balance" field.
func SenderBalanceIn(vs ...string) predicate.Transfers {
	return predicate.Transfers(sql.FieldIn(FieldSenderBalance, vs...))
}

// SenderBalanceNotIn applies the NotIn predicate on the "sender_balance" field.
func SenderBalanceNotIn(vs ...string) predicate.Transfers {
	return predicate.Transfers(sql.FieldNotIn(FieldSenderBalance, vs...))
}

// SenderBalanceGT applies the GT predicate on the "sender_balance" field.
func SenderBalanceGT(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldGT(FieldSenderBalance, v))
}

// SenderBalanceGTE applies the GTE predicate on the "sender_balance" field.
func SenderBalanceGTE(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldGTE(FieldSenderBalance, v))
}

// SenderBalanceLT applies the LT predicate on the "sender_balance" field.
func SenderBalanceLT(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldLT(FieldSenderBalance, v))
}

// SenderBalanceLTE applies the LTE predicate on the "sender_balance" field.
func SenderBalanceLTE(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldLTE(FieldSenderBalance, v))
}

// SenderBalanceContains applies the Contains predicate on the "sender_balance" field.
func SenderBalanceContains(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldContains(FieldSenderBalance, v))
}

// SenderBalanceHasPrefix applies the HasPrefix predicate on the "sender_balance" field.
func SenderBalanceHasPrefix(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldHasPrefix(FieldSenderBalance, v))
}

// SenderBalanceHasSuffix applies the HasSuffix predicate on the "sender_balance" field.
func SenderBalanceHasSuffix(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldHasSuffix(FieldSenderBalance, v))
}

// SenderBalanceIsNil applies the IsNil predicate on the "sender_balance" field.
func SenderBalanceIsNil() predicate.Transfers {
	return predicate.Transfers(sql.FieldIsNull(FieldSenderBalance))
}

// SenderBalanceNotNil applies the NotNil predicate on the "sender_balance" field.
func SenderBalanceNotNil() predicate.Transfers {
	return predicate.Transfers(sql.FieldNotNull(FieldSenderBalance))
}

// SenderBalanceEqualFold applies the EqualFold predicate on the "sender_balance" field.
func SenderBalanceEqualFold(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldEqualFold(FieldSenderBalance, v))
}

// SenderBalanceContainsFold applies the ContainsFold predicate on the "sender_balance" field.
func SenderBalanceContainsFold(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldContainsFold(FieldSenderBalance, v))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...string) predicate.Transfers {
	return predicate.Transfers(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...string) predicate.Transfers {
	return predicate.Transfers(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldLTE(FieldAmount, v))
}

// AmountContains applies the Contains predicate on the "amount" field.
func AmountContains(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldContains(FieldAmount, v))
}

// AmountHasPrefix applies the HasPrefix predicate on the "amount" field.
func AmountHasPrefix(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldHasPrefix(FieldAmount, v))
}

// AmountHasSuffix applies the HasSuffix predicate on the "amount" field.
func AmountHasSuffix(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldHasSuffix(FieldAmount, v))
}

// AmountIsNil applies the IsNil predicate on the "amount" field.
func AmountIsNil() predicate.Transfers {
	return predicate.Transfers(sql.FieldIsNull(FieldAmount))
}

// AmountNotNil applies the NotNil predicate on the "amount" field.
func AmountNotNil() predicate.Transfers {
	return predicate.Transfers(sql.FieldNotNull(FieldAmount))
}

// AmountEqualFold applies the EqualFold predicate on the "amount" field.
func AmountEqualFold(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldEqualFold(FieldAmount, v))
}

// AmountContainsFold applies the ContainsFold predicate on the "amount" field.
func AmountContainsFold(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldContainsFold(FieldAmount, v))
}

// MemoIDEQ applies the EQ predicate on the "memo_id" field.
func MemoIDEQ(v uint64) predicate.Transfers {
	return predicate.Transfers(sql.FieldEQ(FieldMemoID, v))
}

// MemoIDNEQ applies the NEQ predicate on the "memo_id" field.
func MemoIDNEQ(v uint64) predicate.Transfers {
	return predicate.Transfers(sql.FieldNEQ(FieldMemoID, v))
}

// MemoIDIn applies the In predicate on the "memo_id" field.
func MemoIDIn(vs ...uint64) predicate.Transfers {
	return predicate.Transfers(sql.FieldIn(FieldMemoID, vs...))
}

// MemoIDNotIn applies the NotIn predicate on the "memo_id" field.
func MemoIDNotIn(vs ...uint64) predicate.Transfers {
	return predicate.Transfers(sql.FieldNotIn(FieldMemoID, vs...))
}

// MemoIDGT applies the GT predicate on the "memo_id" field.
func MemoIDGT(v uint64) predicate.Transfers {
	return predicate.Transfers(sql.FieldGT(FieldMemoID, v))
}

// MemoIDGTE applies the GTE predicate on the "memo_id" field.
func MemoIDGTE(v uint64) predicate.Transfers {
	return predicate.Transfers(sql.FieldGTE(FieldMemoID, v))
}

// MemoIDLT applies the LT predicate on the "memo_id" field.
func MemoIDLT(v uint64) predicate.Transfers {
	return predicate.Transfers(sql.FieldLT(FieldMemoID, v))
}

// MemoIDLTE applies the LTE predicate on the "memo_id" field.
func MemoIDLTE(v uint64) predicate.Transfers {
	return predicate.Transfers(sql.FieldLTE(FieldMemoID, v))
}

// MemoIDIsNil applies the IsNil predicate on the "memo_id" field.
func MemoIDIsNil() predicate.Transfers {
	return predicate.Transfers(sql.FieldIsNull(FieldMemoID))
}

// MemoIDNotNil applies the NotNil predicate on the "memo_id" field.
func MemoIDNotNil() predicate.Transfers {
	return predicate.Transfers(sql.FieldNotNull(FieldMemoID))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Transfers {
	return predicate.Transfers(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Transfers {
	return predicate.Transfers(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Transfers {
	return predicate.Transfers(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Transfers {
	return predicate.Transfers(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Transfers {
	return predicate.Transfers(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Transfers {
	return predicate.Transfers(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Transfers {
	return predicate.Transfers(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Transfers {
	return predicate.Transfers(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.Transfers {
	return predicate.Transfers(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.Transfers {
	return predicate.Transfers(sql.FieldNotNull(FieldCreatedAt))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Transfers {
	return predicate.Transfers(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Transfers {
	return predicate.Transfers(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Transfers {
	return predicate.Transfers(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Transfers {
	return predicate.Transfers(sql.FieldNotNull(FieldStatus))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldContainsFold(FieldStatus, v))
}

// DeployEQ applies the EQ predicate on the "Deploy" field.
func DeployEQ(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldEQ(FieldDeploy, v))
}

// DeployNEQ applies the NEQ predicate on the "Deploy" field.
func DeployNEQ(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldNEQ(FieldDeploy, v))
}

// DeployIn applies the In predicate on the "Deploy" field.
func DeployIn(vs ...string) predicate.Transfers {
	return predicate.Transfers(sql.FieldIn(FieldDeploy, vs...))
}

// DeployNotIn applies the NotIn predicate on the "Deploy" field.
func DeployNotIn(vs ...string) predicate.Transfers {
	return predicate.Transfers(sql.FieldNotIn(FieldDeploy, vs...))
}

// DeployGT applies the GT predicate on the "Deploy" field.
func DeployGT(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldGT(FieldDeploy, v))
}

// DeployGTE applies the GTE predicate on the "Deploy" field.
func DeployGTE(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldGTE(FieldDeploy, v))
}

// DeployLT applies the LT predicate on the "Deploy" field.
func DeployLT(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldLT(FieldDeploy, v))
}

// DeployLTE applies the LTE predicate on the "Deploy" field.
func DeployLTE(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldLTE(FieldDeploy, v))
}

// DeployContains applies the Contains predicate on the "Deploy" field.
func DeployContains(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldContains(FieldDeploy, v))
}

// DeployHasPrefix applies the HasPrefix predicate on the "Deploy" field.
func DeployHasPrefix(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldHasPrefix(FieldDeploy, v))
}

// DeployHasSuffix applies the HasSuffix predicate on the "Deploy" field.
func DeployHasSuffix(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldHasSuffix(FieldDeploy, v))
}

// DeployIsNil applies the IsNil predicate on the "Deploy" field.
func DeployIsNil() predicate.Transfers {
	return predicate.Transfers(sql.FieldIsNull(FieldDeploy))
}

// DeployNotNil applies the NotNil predicate on the "Deploy" field.
func DeployNotNil() predicate.Transfers {
	return predicate.Transfers(sql.FieldNotNull(FieldDeploy))
}

// DeployEqualFold applies the EqualFold predicate on the "Deploy" field.
func DeployEqualFold(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldEqualFold(FieldDeploy, v))
}

// DeployContainsFold applies the ContainsFold predicate on the "Deploy" field.
func DeployContainsFold(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldContainsFold(FieldDeploy, v))
}

// AdditionalTypeEQ applies the EQ predicate on the "AdditionalType" field.
func AdditionalTypeEQ(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldEQ(FieldAdditionalType, v))
}

// AdditionalTypeNEQ applies the NEQ predicate on the "AdditionalType" field.
func AdditionalTypeNEQ(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldNEQ(FieldAdditionalType, v))
}

// AdditionalTypeIn applies the In predicate on the "AdditionalType" field.
func AdditionalTypeIn(vs ...string) predicate.Transfers {
	return predicate.Transfers(sql.FieldIn(FieldAdditionalType, vs...))
}

// AdditionalTypeNotIn applies the NotIn predicate on the "AdditionalType" field.
func AdditionalTypeNotIn(vs ...string) predicate.Transfers {
	return predicate.Transfers(sql.FieldNotIn(FieldAdditionalType, vs...))
}

// AdditionalTypeGT applies the GT predicate on the "AdditionalType" field.
func AdditionalTypeGT(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldGT(FieldAdditionalType, v))
}

// AdditionalTypeGTE applies the GTE predicate on the "AdditionalType" field.
func AdditionalTypeGTE(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldGTE(FieldAdditionalType, v))
}

// AdditionalTypeLT applies the LT predicate on the "AdditionalType" field.
func AdditionalTypeLT(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldLT(FieldAdditionalType, v))
}

// AdditionalTypeLTE applies the LTE predicate on the "AdditionalType" field.
func AdditionalTypeLTE(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldLTE(FieldAdditionalType, v))
}

// AdditionalTypeContains applies the Contains predicate on the "AdditionalType" field.
func AdditionalTypeContains(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldContains(FieldAdditionalType, v))
}

// AdditionalTypeHasPrefix applies the HasPrefix predicate on the "AdditionalType" field.
func AdditionalTypeHasPrefix(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldHasPrefix(FieldAdditionalType, v))
}

// AdditionalTypeHasSuffix applies the HasSuffix predicate on the "AdditionalType" field.
func AdditionalTypeHasSuffix(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldHasSuffix(FieldAdditionalType, v))
}

// AdditionalTypeIsNil applies the IsNil predicate on the "AdditionalType" field.
func AdditionalTypeIsNil() predicate.Transfers {
	return predicate.Transfers(sql.FieldIsNull(FieldAdditionalType))
}

// AdditionalTypeNotNil applies the NotNil predicate on the "AdditionalType" field.
func AdditionalTypeNotNil() predicate.Transfers {
	return predicate.Transfers(sql.FieldNotNull(FieldAdditionalType))
}

// AdditionalTypeEqualFold applies the EqualFold predicate on the "AdditionalType" field.
func AdditionalTypeEqualFold(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldEqualFold(FieldAdditionalType, v))
}

// AdditionalTypeContainsFold applies the ContainsFold predicate on the "AdditionalType" field.
func AdditionalTypeContainsFold(v string) predicate.Transfers {
	return predicate.Transfers(sql.FieldContainsFold(FieldAdditionalType, v))
}

// InvoiceIDEQ applies the EQ predicate on the "invoiceID" field.
func InvoiceIDEQ(v int64) predicate.Transfers {
	return predicate.Transfers(sql.FieldEQ(FieldInvoiceID, v))
}

// InvoiceIDNEQ applies the NEQ predicate on the "invoiceID" field.
func InvoiceIDNEQ(v int64) predicate.Transfers {
	return predicate.Transfers(sql.FieldNEQ(FieldInvoiceID, v))
}

// InvoiceIDIn applies the In predicate on the "invoiceID" field.
func InvoiceIDIn(vs ...int64) predicate.Transfers {
	return predicate.Transfers(sql.FieldIn(FieldInvoiceID, vs...))
}

// InvoiceIDNotIn applies the NotIn predicate on the "invoiceID" field.
func InvoiceIDNotIn(vs ...int64) predicate.Transfers {
	return predicate.Transfers(sql.FieldNotIn(FieldInvoiceID, vs...))
}

// InvoiceIDGT applies the GT predicate on the "invoiceID" field.
func InvoiceIDGT(v int64) predicate.Transfers {
	return predicate.Transfers(sql.FieldGT(FieldInvoiceID, v))
}

// InvoiceIDGTE applies the GTE predicate on the "invoiceID" field.
func InvoiceIDGTE(v int64) predicate.Transfers {
	return predicate.Transfers(sql.FieldGTE(FieldInvoiceID, v))
}

// InvoiceIDLT applies the LT predicate on the "invoiceID" field.
func InvoiceIDLT(v int64) predicate.Transfers {
	return predicate.Transfers(sql.FieldLT(FieldInvoiceID, v))
}

// InvoiceIDLTE applies the LTE predicate on the "invoiceID" field.
func InvoiceIDLTE(v int64) predicate.Transfers {
	return predicate.Transfers(sql.FieldLTE(FieldInvoiceID, v))
}

// InvoiceIDIsNil applies the IsNil predicate on the "invoiceID" field.
func InvoiceIDIsNil() predicate.Transfers {
	return predicate.Transfers(sql.FieldIsNull(FieldInvoiceID))
}

// InvoiceIDNotNil applies the NotNil predicate on the "invoiceID" field.
func InvoiceIDNotNil() predicate.Transfers {
	return predicate.Transfers(sql.FieldNotNull(FieldInvoiceID))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Transfers {
	return predicate.Transfers(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Transfers {
	return predicate.Transfers(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Transfers) predicate.Transfers {
	return predicate.Transfers(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Transfers) predicate.Transfers {
	return predicate.Transfers(func(s *sql.Selector) {
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
func Not(p predicate.Transfers) predicate.Transfers {
	return predicate.Transfers(func(s *sql.Selector) {
		p(s.Not())
	})
}
