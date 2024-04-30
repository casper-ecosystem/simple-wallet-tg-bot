package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		//field.String("state").Optional(),
		field.String("public_key").Optional(),
		field.String("password").Optional(),
		field.Bool("logged_in").Default(false),
		field.Time("last_access").Optional(),
		field.Int64("lock_timeout").Optional(),
		field.Bool("locked_manual").Optional(),
		field.Bool("notify").Default(false),
		field.Int8("notify_time").Default(0),
		field.Time("notify_last_time").Default(time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC)),
		field.Bool("store_privat_key").Default(false),
		field.Bool("enable_logging").Default(true),
		field.Bool("registered").Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("balance", Balances.Type).
			Unique().Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("address_book", AdressBook.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("state", UserState.Type).
			Unique().Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("rewards_data", RewardsData.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("PrivateKey", PrivateKeys.Type).
			Unique().Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("transfers", Transfers.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("delegates", Delegates.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("undelegates", Undelegates.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("swaps", Swaps.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("invoices", Invoice.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("recentInvoices", RecentInvoices.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		// non-unique index.
		index.Fields("public_key"),
	}
}
