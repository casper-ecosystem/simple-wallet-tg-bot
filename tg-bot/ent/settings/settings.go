// Code generated by ent, DO NOT EDIT.

package settings

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the settings type in the database.
	Label = "settings"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLastScannedBlockNotificator holds the string denoting the last_scanned_block_notificator field in the database.
	FieldLastScannedBlockNotificator = "last_scanned_block_notificator"
	// FieldLastScannedEraValidators holds the string denoting the last_scanned_era_validators field in the database.
	FieldLastScannedEraValidators = "last_scanned_era_validators"
	// Table holds the table name of the settings in the database.
	Table = "settings"
)

// Columns holds all SQL columns for settings fields.
var Columns = []string{
	FieldID,
	FieldLastScannedBlockNotificator,
	FieldLastScannedEraValidators,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Settings queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByLastScannedBlockNotificator orders the results by the last_scanned_block_notificator field.
func ByLastScannedBlockNotificator(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastScannedBlockNotificator, opts...).ToFunc()
}

// ByLastScannedEraValidators orders the results by the last_scanned_era_validators field.
func ByLastScannedEraValidators(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastScannedEraValidators, opts...).ToFunc()
}