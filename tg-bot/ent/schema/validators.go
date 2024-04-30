package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Validators holds the schema definition for the Validators entity.
type Validators struct {
	ent.Schema
}

// Fields of the Validators.
func (Validators) Fields() []ent.Field {
	return []ent.Field{
		field.String("address").Unique(),
		field.String("name").Optional(),
		field.Int8("fee").Optional(),
		field.Int64("delegators").Optional(),
		field.Bool("active").Optional(),
	}
}

// Edges of the Validators.
func (Validators) Edges() []ent.Edge {
	return nil
}
