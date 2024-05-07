// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Simplewallethq/tg-bot/ent/user"
	"github.com/Simplewallethq/tg-bot/ent/userstate"
)

// UserState is the model entity for the UserState schema.
type UserState struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// State holds the value of the "state" field.
	State string `json:"state,omitempty"`
	// Data holds the value of the "data" field.
	Data []byte `json:"data,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserStateQuery when eager-loading is set.
	Edges        UserStateEdges `json:"edges"`
	user_state   *int64
	selectValues sql.SelectValues
}

// UserStateEdges holds the relations/edges for other nodes in the graph.
type UserStateEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserStateEdges) OwnerOrErr() (*User, error) {
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
func (*UserState) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userstate.FieldData:
			values[i] = new([]byte)
		case userstate.FieldID:
			values[i] = new(sql.NullInt64)
		case userstate.FieldState:
			values[i] = new(sql.NullString)
		case userstate.ForeignKeys[0]: // user_state
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserState fields.
func (us *UserState) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userstate.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			us.ID = int(value.Int64)
		case userstate.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				us.State = value.String
			}
		case userstate.FieldData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value != nil {
				us.Data = *value
			}
		case userstate.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_state", value)
			} else if value.Valid {
				us.user_state = new(int64)
				*us.user_state = int64(value.Int64)
			}
		default:
			us.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserState.
// This includes values selected through modifiers, order, etc.
func (us *UserState) Value(name string) (ent.Value, error) {
	return us.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the UserState entity.
func (us *UserState) QueryOwner() *UserQuery {
	return NewUserStateClient(us.config).QueryOwner(us)
}

// Update returns a builder for updating this UserState.
// Note that you need to call UserState.Unwrap() before calling this method if this UserState
// was returned from a transaction, and the transaction was committed or rolled back.
func (us *UserState) Update() *UserStateUpdateOne {
	return NewUserStateClient(us.config).UpdateOne(us)
}

// Unwrap unwraps the UserState entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (us *UserState) Unwrap() *UserState {
	_tx, ok := us.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserState is not a transactional entity")
	}
	us.config.driver = _tx.drv
	return us
}

// String implements the fmt.Stringer.
func (us *UserState) String() string {
	var builder strings.Builder
	builder.WriteString("UserState(")
	builder.WriteString(fmt.Sprintf("id=%v, ", us.ID))
	builder.WriteString("state=")
	builder.WriteString(us.State)
	builder.WriteString(", ")
	builder.WriteString("data=")
	builder.WriteString(fmt.Sprintf("%v", us.Data))
	builder.WriteByte(')')
	return builder.String()
}

// UserStates is a parsable slice of UserState.
type UserStates []*UserState