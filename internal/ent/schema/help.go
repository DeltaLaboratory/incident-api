package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Help holds the schema definition for the Help entity.
type Help struct {
	ent.Schema
}

// Fields of the Help.
func (Help) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.Nil).Default(uuid.New),
		field.String("idempotency_key").Unique(),

		field.UUID("reporter", uuid.Nil),
		field.Other("location", &GeoJson{}).
			SchemaType(map[string]string{
				dialect.Postgres: "geometry(point, 4326)",
			}).StorageKey("location").Optional(),
		field.String("description"),

		field.Int("heart_rate").Nillable(),
		field.Int("blood_pressure").Nillable(),
		field.Int("body_temperature").Nillable(),

		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Help.
func (Help) Edges() []ent.Edge {
	return nil
}
