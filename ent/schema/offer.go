package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Offer — модель предложения.
type Offer struct {
	ent.Schema
}

func (Offer) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("master_id", uuid.UUID{}).Comment("ID исполнителя который предлагает заказ"),
		field.UUID("order_id", uuid.UUID{}).Comment("ID заказа"),
		field.Enum("status").Values("active", "accepted", "cancelled", "rejected").Default("active"),
		field.Float32("price").Default(0).Comment("Цена"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (Offer) Edges() []ent.Edge {
	return nil
}
