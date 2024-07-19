package schema

import (
	"database/sql/driver"
	"encoding/binary"
	"encoding/json"
	"errors"
	"math"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/twpayne/go-geom/encoding/ewkbhex"
	"github.com/twpayne/go-geom/encoding/geojson"
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

		field.Other("location", &GeoJson{}).
			SchemaType(map[string]string{
				dialect.Postgres: "geometry(point, 4326)",
			}).StorageKey("location").Optional(),

		field.String("type"),
		field.String("description"),
		field.Bytes("image").Nillable().Default(nil),

		field.Int("vote").Default(0),
		field.Bytes("vote_filter").Nillable().Default(nil),

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

type GeoJson struct {
	*geojson.Geometry
}

func (t *GeoJson) Value() (driver.Value, error) {
	geometry, err := t.Decode()
	if err != nil {
		return nil, err
	}

	encodedGeometry, err := ewkbhex.Encode(geometry, binary.LittleEndian)
	if err != nil {
		return nil, err
	}

	return encodedGeometry, nil
}

func (t *GeoJson) Scan(value interface{}) error {
	// handle nil
	if value == nil {
		t = nil
		return nil
	}

	// parse as string
	stringValue, ok := value.(string)
	if !ok {
		return errors.New("value is no string")
	}

	geometry, err := ewkbhex.Decode(stringValue)
	if err != nil {
		return err
	}

	geometryAsBytes, err := geojson.Marshal(geometry)
	if err != nil {
		return err
	}

	var geoJson GeoJson
	if err := json.Unmarshal(geometryAsBytes, &geoJson); err != nil {
		return err
	}

	*t = geoJson

	return nil
}
