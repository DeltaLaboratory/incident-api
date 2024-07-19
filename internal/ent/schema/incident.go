package schema

import (
	"math"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Incident holds the schema definition for the Incident entity.
type Incident struct {
	ent.Schema
}

// Fields of the Incident.
func (Incident) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.Nil).Default(uuid.New),
		field.String("idempotency_key").Unique(),

		field.UUID("reporter", uuid.Nil),

		field.Float("latitude").GoType(Coordinate(0)).SchemaType(map[string]string{dialect.Postgres: "numeric"}),
		field.Float("longitude").GoType(Coordinate(0)).SchemaType(map[string]string{dialect.Postgres: "numeric"}),

		field.String("type"),
		field.String("description"),
		field.Bytes("image").Nillable(),

		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Incident.
func (Incident) Edges() []ent.Edge {
	return nil
}

type Coordinate float64

func (c Coordinate) Radians() float64 {
	return float64(c) * math.Pi / 180
}
