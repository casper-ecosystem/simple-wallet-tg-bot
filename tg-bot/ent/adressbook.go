// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/adressbook"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/user"
)

// AdressBook is the model entity for the AdressBook schema.
type AdressBook struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// InUpdate holds the value of the "InUpdate" field.
	InUpdate bool `json:"InUpdate,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AdressBookQuery when eager-loading is set.
	Edges             AdressBookEdges `json:"edges"`
	user_address_book *int64
	selectValues      sql.SelectValues
}

// AdressBookEdges holds the relations/edges for other nodes in the graph.
type AdressBookEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AdressBookEdges) OwnerOrErr() (*User, error) {
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
func (*AdressBook) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case adressbook.FieldInUpdate:
			values[i] = new(sql.NullBool)
		case adressbook.FieldID:
			values[i] = new(sql.NullInt64)
		case adressbook.FieldAddress, adressbook.FieldName:
			values[i] = new(sql.NullString)
		case adressbook.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case adressbook.ForeignKeys[0]: // user_address_book
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AdressBook fields.
func (ab *AdressBook) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case adressbook.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ab.ID = int(value.Int64)
		case adressbook.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				ab.Address = value.String
			}
		case adressbook.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ab.Name = value.String
			}
		case adressbook.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ab.CreatedAt = value.Time
			}
		case adressbook.FieldInUpdate:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field InUpdate", values[i])
			} else if value.Valid {
				ab.InUpdate = value.Bool
			}
		case adressbook.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_address_book", value)
			} else if value.Valid {
				ab.user_address_book = new(int64)
				*ab.user_address_book = int64(value.Int64)
			}
		default:
			ab.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AdressBook.
// This includes values selected through modifiers, order, etc.
func (ab *AdressBook) Value(name string) (ent.Value, error) {
	return ab.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the AdressBook entity.
func (ab *AdressBook) QueryOwner() *UserQuery {
	return NewAdressBookClient(ab.config).QueryOwner(ab)
}

// Update returns a builder for updating this AdressBook.
// Note that you need to call AdressBook.Unwrap() before calling this method if this AdressBook
// was returned from a transaction, and the transaction was committed or rolled back.
func (ab *AdressBook) Update() *AdressBookUpdateOne {
	return NewAdressBookClient(ab.config).UpdateOne(ab)
}

// Unwrap unwraps the AdressBook entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ab *AdressBook) Unwrap() *AdressBook {
	_tx, ok := ab.config.driver.(*txDriver)
	if !ok {
		panic("ent: AdressBook is not a transactional entity")
	}
	ab.config.driver = _tx.drv
	return ab
}

// String implements the fmt.Stringer.
func (ab *AdressBook) String() string {
	var builder strings.Builder
	builder.WriteString("AdressBook(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ab.ID))
	builder.WriteString("address=")
	builder.WriteString(ab.Address)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ab.Name)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ab.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("InUpdate=")
	builder.WriteString(fmt.Sprintf("%v", ab.InUpdate))
	builder.WriteByte(')')
	return builder.String()
}

// AdressBooks is a parsable slice of AdressBook.
type AdressBooks []*AdressBook
