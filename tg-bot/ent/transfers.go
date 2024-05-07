// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Simplewallethq/tg-bot/ent/transfers"
	"github.com/Simplewallethq/tg-bot/ent/user"
)

// Transfers is the model entity for the Transfers schema.
type Transfers struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FromPubkey holds the value of the "from_pubkey" field.
	FromPubkey string `json:"from_pubkey,omitempty"`
	// ToPubkey holds the value of the "to_pubkey" field.
	ToPubkey string `json:"to_pubkey,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// SenderBalance holds the value of the "sender_balance" field.
	SenderBalance string `json:"sender_balance,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount string `json:"amount,omitempty"`
	// MemoID holds the value of the "memo_id" field.
	MemoID uint64 `json:"memo_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// Deploy holds the value of the "Deploy" field.
	Deploy string `json:"Deploy,omitempty"`
	// AdditionalType holds the value of the "AdditionalType" field.
	AdditionalType string `json:"AdditionalType,omitempty"`
	// InvoiceID holds the value of the "invoiceID" field.
	InvoiceID int64 `json:"invoiceID,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TransfersQuery when eager-loading is set.
	Edges          TransfersEdges `json:"edges"`
	user_transfers *int64
	selectValues   sql.SelectValues
}

// TransfersEdges holds the relations/edges for other nodes in the graph.
type TransfersEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransfersEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Transfers) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case transfers.FieldID, transfers.FieldMemoID, transfers.FieldInvoiceID:
			values[i] = new(sql.NullInt64)
		case transfers.FieldFromPubkey, transfers.FieldToPubkey, transfers.FieldName, transfers.FieldSenderBalance, transfers.FieldAmount, transfers.FieldStatus, transfers.FieldDeploy, transfers.FieldAdditionalType:
			values[i] = new(sql.NullString)
		case transfers.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case transfers.ForeignKeys[0]: // user_transfers
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Transfers fields.
func (t *Transfers) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case transfers.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case transfers.FieldFromPubkey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field from_pubkey", values[i])
			} else if value.Valid {
				t.FromPubkey = value.String
			}
		case transfers.FieldToPubkey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field to_pubkey", values[i])
			} else if value.Valid {
				t.ToPubkey = value.String
			}
		case transfers.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case transfers.FieldSenderBalance:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sender_balance", values[i])
			} else if value.Valid {
				t.SenderBalance = value.String
			}
		case transfers.FieldAmount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				t.Amount = value.String
			}
		case transfers.FieldMemoID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field memo_id", values[i])
			} else if value.Valid {
				t.MemoID = uint64(value.Int64)
			}
		case transfers.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case transfers.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				t.Status = value.String
			}
		case transfers.FieldDeploy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Deploy", values[i])
			} else if value.Valid {
				t.Deploy = value.String
			}
		case transfers.FieldAdditionalType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field AdditionalType", values[i])
			} else if value.Valid {
				t.AdditionalType = value.String
			}
		case transfers.FieldInvoiceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field invoiceID", values[i])
			} else if value.Valid {
				t.InvoiceID = value.Int64
			}
		case transfers.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_transfers", value)
			} else if value.Valid {
				t.user_transfers = new(int64)
				*t.user_transfers = int64(value.Int64)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Transfers.
// This includes values selected through modifiers, order, etc.
func (t *Transfers) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Transfers entity.
func (t *Transfers) QueryOwner() *UserQuery {
	return NewTransfersClient(t.config).QueryOwner(t)
}

// Update returns a builder for updating this Transfers.
// Note that you need to call Transfers.Unwrap() before calling this method if this Transfers
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Transfers) Update() *TransfersUpdateOne {
	return NewTransfersClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Transfers entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Transfers) Unwrap() *Transfers {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Transfers is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Transfers) String() string {
	var builder strings.Builder
	builder.WriteString("Transfers(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("from_pubkey=")
	builder.WriteString(t.FromPubkey)
	builder.WriteString(", ")
	builder.WriteString("to_pubkey=")
	builder.WriteString(t.ToPubkey)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("sender_balance=")
	builder.WriteString(t.SenderBalance)
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(t.Amount)
	builder.WriteString(", ")
	builder.WriteString("memo_id=")
	builder.WriteString(fmt.Sprintf("%v", t.MemoID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(t.Status)
	builder.WriteString(", ")
	builder.WriteString("Deploy=")
	builder.WriteString(t.Deploy)
	builder.WriteString(", ")
	builder.WriteString("AdditionalType=")
	builder.WriteString(t.AdditionalType)
	builder.WriteString(", ")
	builder.WriteString("invoiceID=")
	builder.WriteString(fmt.Sprintf("%v", t.InvoiceID))
	builder.WriteByte(')')
	return builder.String()
}

// TransfersSlice is a parsable slice of Transfers.
type TransfersSlice []*Transfers