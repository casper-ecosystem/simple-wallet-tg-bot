// Code generated by ent, DO NOT EDIT.

package transfers

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the transfers type in the database.
	Label = "transfers"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFromPubkey holds the string denoting the from_pubkey field in the database.
	FieldFromPubkey = "from_pubkey"
	// FieldToPubkey holds the string denoting the to_pubkey field in the database.
	FieldToPubkey = "to_pubkey"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSenderBalance holds the string denoting the sender_balance field in the database.
	FieldSenderBalance = "sender_balance"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldMemoID holds the string denoting the memo_id field in the database.
	FieldMemoID = "memo_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldDeploy holds the string denoting the deploy field in the database.
	FieldDeploy = "deploy"
	// FieldAdditionalType holds the string denoting the additionaltype field in the database.
	FieldAdditionalType = "additional_type"
	// FieldInvoiceID holds the string denoting the invoiceid field in the database.
	FieldInvoiceID = "invoice_id"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the transfers in the database.
	Table = "transfers"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "transfers"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_transfers"
)

// Columns holds all SQL columns for transfers fields.
var Columns = []string{
	FieldID,
	FieldFromPubkey,
	FieldToPubkey,
	FieldName,
	FieldSenderBalance,
	FieldAmount,
	FieldMemoID,
	FieldCreatedAt,
	FieldStatus,
	FieldDeploy,
	FieldAdditionalType,
	FieldInvoiceID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "transfers"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_transfers",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Transfers queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByFromPubkey orders the results by the from_pubkey field.
func ByFromPubkey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFromPubkey, opts...).ToFunc()
}

// ByToPubkey orders the results by the to_pubkey field.
func ByToPubkey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldToPubkey, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// BySenderBalance orders the results by the sender_balance field.
func BySenderBalance(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSenderBalance, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByMemoID orders the results by the memo_id field.
func ByMemoID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemoID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByDeploy orders the results by the Deploy field.
func ByDeploy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeploy, opts...).ToFunc()
}

// ByAdditionalType orders the results by the AdditionalType field.
func ByAdditionalType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAdditionalType, opts...).ToFunc()
}

// ByInvoiceID orders the results by the invoiceID field.
func ByInvoiceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInvoiceID, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}