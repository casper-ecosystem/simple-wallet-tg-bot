package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.String("name"),
		field.Time("created_at").Default(time.Now()),
		field.Bytes("data"),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return nil
}
