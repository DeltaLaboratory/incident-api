// Code generated by ent, DO NOT EDIT.

package help

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the help type in the database.
	Label = "help"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIdempotencyKey holds the string denoting the idempotency_key field in the database.
	FieldIdempotencyKey = "idempotency_key"
	// FieldReporter holds the string denoting the reporter field in the database.
	FieldReporter = "reporter"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldHeartRate holds the string denoting the heart_rate field in the database.
	FieldHeartRate = "heart_rate"
	// FieldBloodPressure holds the string denoting the blood_pressure field in the database.
	FieldBloodPressure = "blood_pressure"
	// FieldBodyTemperature holds the string denoting the body_temperature field in the database.
	FieldBodyTemperature = "body_temperature"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the help in the database.
	Table = "helps"
)

// Columns holds all SQL columns for help fields.
var Columns = []string{
	FieldID,
	FieldIdempotencyKey,
	FieldReporter,
	FieldLocation,
	FieldDescription,
	FieldHeartRate,
	FieldBloodPressure,
	FieldBodyTemperature,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Help queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByIdempotencyKey orders the results by the idempotency_key field.
func ByIdempotencyKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIdempotencyKey, opts...).ToFunc()
}

// ByReporter orders the results by the reporter field.
func ByReporter(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReporter, opts...).ToFunc()
}

// ByLocation orders the results by the location field.
func ByLocation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocation, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByHeartRate orders the results by the heart_rate field.
func ByHeartRate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHeartRate, opts...).ToFunc()
}

// ByBloodPressure orders the results by the blood_pressure field.
func ByBloodPressure(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBloodPressure, opts...).ToFunc()
}

// ByBodyTemperature orders the results by the body_temperature field.
func ByBodyTemperature(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBodyTemperature, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}
