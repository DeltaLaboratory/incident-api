// Code generated by ent, DO NOT EDIT.

package help

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/DeltaLaboratory/incident-api/internal/ent/predicate"
	"github.com/DeltaLaboratory/incident-api/internal/ent/schema"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Help {
	return predicate.Help(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Help {
	return predicate.Help(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Help {
	return predicate.Help(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Help {
	return predicate.Help(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Help {
	return predicate.Help(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Help {
	return predicate.Help(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Help {
	return predicate.Help(sql.FieldLTE(FieldID, id))
}

// IdempotencyKey applies equality check predicate on the "idempotency_key" field. It's identical to IdempotencyKeyEQ.
func IdempotencyKey(v string) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldIdempotencyKey, v))
}

// Reporter applies equality check predicate on the "reporter" field. It's identical to ReporterEQ.
func Reporter(v uuid.UUID) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldReporter, v))
}

// Location applies equality check predicate on the "location" field. It's identical to LocationEQ.
func Location(v *schema.GeoJson) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldLocation, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldDescription, v))
}

// HeartRate applies equality check predicate on the "heart_rate" field. It's identical to HeartRateEQ.
func HeartRate(v int) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldHeartRate, v))
}

// BloodPressure applies equality check predicate on the "blood_pressure" field. It's identical to BloodPressureEQ.
func BloodPressure(v int) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldBloodPressure, v))
}

// BodyTemperature applies equality check predicate on the "body_temperature" field. It's identical to BodyTemperatureEQ.
func BodyTemperature(v int) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldBodyTemperature, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldCreatedAt, v))
}

// IdempotencyKeyEQ applies the EQ predicate on the "idempotency_key" field.
func IdempotencyKeyEQ(v string) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldIdempotencyKey, v))
}

// IdempotencyKeyNEQ applies the NEQ predicate on the "idempotency_key" field.
func IdempotencyKeyNEQ(v string) predicate.Help {
	return predicate.Help(sql.FieldNEQ(FieldIdempotencyKey, v))
}

// IdempotencyKeyIn applies the In predicate on the "idempotency_key" field.
func IdempotencyKeyIn(vs ...string) predicate.Help {
	return predicate.Help(sql.FieldIn(FieldIdempotencyKey, vs...))
}

// IdempotencyKeyNotIn applies the NotIn predicate on the "idempotency_key" field.
func IdempotencyKeyNotIn(vs ...string) predicate.Help {
	return predicate.Help(sql.FieldNotIn(FieldIdempotencyKey, vs...))
}

// IdempotencyKeyGT applies the GT predicate on the "idempotency_key" field.
func IdempotencyKeyGT(v string) predicate.Help {
	return predicate.Help(sql.FieldGT(FieldIdempotencyKey, v))
}

// IdempotencyKeyGTE applies the GTE predicate on the "idempotency_key" field.
func IdempotencyKeyGTE(v string) predicate.Help {
	return predicate.Help(sql.FieldGTE(FieldIdempotencyKey, v))
}

// IdempotencyKeyLT applies the LT predicate on the "idempotency_key" field.
func IdempotencyKeyLT(v string) predicate.Help {
	return predicate.Help(sql.FieldLT(FieldIdempotencyKey, v))
}

// IdempotencyKeyLTE applies the LTE predicate on the "idempotency_key" field.
func IdempotencyKeyLTE(v string) predicate.Help {
	return predicate.Help(sql.FieldLTE(FieldIdempotencyKey, v))
}

// IdempotencyKeyContains applies the Contains predicate on the "idempotency_key" field.
func IdempotencyKeyContains(v string) predicate.Help {
	return predicate.Help(sql.FieldContains(FieldIdempotencyKey, v))
}

// IdempotencyKeyHasPrefix applies the HasPrefix predicate on the "idempotency_key" field.
func IdempotencyKeyHasPrefix(v string) predicate.Help {
	return predicate.Help(sql.FieldHasPrefix(FieldIdempotencyKey, v))
}

// IdempotencyKeyHasSuffix applies the HasSuffix predicate on the "idempotency_key" field.
func IdempotencyKeyHasSuffix(v string) predicate.Help {
	return predicate.Help(sql.FieldHasSuffix(FieldIdempotencyKey, v))
}

// IdempotencyKeyEqualFold applies the EqualFold predicate on the "idempotency_key" field.
func IdempotencyKeyEqualFold(v string) predicate.Help {
	return predicate.Help(sql.FieldEqualFold(FieldIdempotencyKey, v))
}

// IdempotencyKeyContainsFold applies the ContainsFold predicate on the "idempotency_key" field.
func IdempotencyKeyContainsFold(v string) predicate.Help {
	return predicate.Help(sql.FieldContainsFold(FieldIdempotencyKey, v))
}

// ReporterEQ applies the EQ predicate on the "reporter" field.
func ReporterEQ(v uuid.UUID) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldReporter, v))
}

// ReporterNEQ applies the NEQ predicate on the "reporter" field.
func ReporterNEQ(v uuid.UUID) predicate.Help {
	return predicate.Help(sql.FieldNEQ(FieldReporter, v))
}

// ReporterIn applies the In predicate on the "reporter" field.
func ReporterIn(vs ...uuid.UUID) predicate.Help {
	return predicate.Help(sql.FieldIn(FieldReporter, vs...))
}

// ReporterNotIn applies the NotIn predicate on the "reporter" field.
func ReporterNotIn(vs ...uuid.UUID) predicate.Help {
	return predicate.Help(sql.FieldNotIn(FieldReporter, vs...))
}

// ReporterGT applies the GT predicate on the "reporter" field.
func ReporterGT(v uuid.UUID) predicate.Help {
	return predicate.Help(sql.FieldGT(FieldReporter, v))
}

// ReporterGTE applies the GTE predicate on the "reporter" field.
func ReporterGTE(v uuid.UUID) predicate.Help {
	return predicate.Help(sql.FieldGTE(FieldReporter, v))
}

// ReporterLT applies the LT predicate on the "reporter" field.
func ReporterLT(v uuid.UUID) predicate.Help {
	return predicate.Help(sql.FieldLT(FieldReporter, v))
}

// ReporterLTE applies the LTE predicate on the "reporter" field.
func ReporterLTE(v uuid.UUID) predicate.Help {
	return predicate.Help(sql.FieldLTE(FieldReporter, v))
}

// LocationEQ applies the EQ predicate on the "location" field.
func LocationEQ(v *schema.GeoJson) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldLocation, v))
}

// LocationNEQ applies the NEQ predicate on the "location" field.
func LocationNEQ(v *schema.GeoJson) predicate.Help {
	return predicate.Help(sql.FieldNEQ(FieldLocation, v))
}

// LocationIn applies the In predicate on the "location" field.
func LocationIn(vs ...*schema.GeoJson) predicate.Help {
	return predicate.Help(sql.FieldIn(FieldLocation, vs...))
}

// LocationNotIn applies the NotIn predicate on the "location" field.
func LocationNotIn(vs ...*schema.GeoJson) predicate.Help {
	return predicate.Help(sql.FieldNotIn(FieldLocation, vs...))
}

// LocationGT applies the GT predicate on the "location" field.
func LocationGT(v *schema.GeoJson) predicate.Help {
	return predicate.Help(sql.FieldGT(FieldLocation, v))
}

// LocationGTE applies the GTE predicate on the "location" field.
func LocationGTE(v *schema.GeoJson) predicate.Help {
	return predicate.Help(sql.FieldGTE(FieldLocation, v))
}

// LocationLT applies the LT predicate on the "location" field.
func LocationLT(v *schema.GeoJson) predicate.Help {
	return predicate.Help(sql.FieldLT(FieldLocation, v))
}

// LocationLTE applies the LTE predicate on the "location" field.
func LocationLTE(v *schema.GeoJson) predicate.Help {
	return predicate.Help(sql.FieldLTE(FieldLocation, v))
}

// LocationIsNil applies the IsNil predicate on the "location" field.
func LocationIsNil() predicate.Help {
	return predicate.Help(sql.FieldIsNull(FieldLocation))
}

// LocationNotNil applies the NotNil predicate on the "location" field.
func LocationNotNil() predicate.Help {
	return predicate.Help(sql.FieldNotNull(FieldLocation))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Help {
	return predicate.Help(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Help {
	return predicate.Help(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Help {
	return predicate.Help(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Help {
	return predicate.Help(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Help {
	return predicate.Help(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Help {
	return predicate.Help(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Help {
	return predicate.Help(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Help {
	return predicate.Help(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Help {
	return predicate.Help(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Help {
	return predicate.Help(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Help {
	return predicate.Help(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Help {
	return predicate.Help(sql.FieldContainsFold(FieldDescription, v))
}

// HeartRateEQ applies the EQ predicate on the "heart_rate" field.
func HeartRateEQ(v int) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldHeartRate, v))
}

// HeartRateNEQ applies the NEQ predicate on the "heart_rate" field.
func HeartRateNEQ(v int) predicate.Help {
	return predicate.Help(sql.FieldNEQ(FieldHeartRate, v))
}

// HeartRateIn applies the In predicate on the "heart_rate" field.
func HeartRateIn(vs ...int) predicate.Help {
	return predicate.Help(sql.FieldIn(FieldHeartRate, vs...))
}

// HeartRateNotIn applies the NotIn predicate on the "heart_rate" field.
func HeartRateNotIn(vs ...int) predicate.Help {
	return predicate.Help(sql.FieldNotIn(FieldHeartRate, vs...))
}

// HeartRateGT applies the GT predicate on the "heart_rate" field.
func HeartRateGT(v int) predicate.Help {
	return predicate.Help(sql.FieldGT(FieldHeartRate, v))
}

// HeartRateGTE applies the GTE predicate on the "heart_rate" field.
func HeartRateGTE(v int) predicate.Help {
	return predicate.Help(sql.FieldGTE(FieldHeartRate, v))
}

// HeartRateLT applies the LT predicate on the "heart_rate" field.
func HeartRateLT(v int) predicate.Help {
	return predicate.Help(sql.FieldLT(FieldHeartRate, v))
}

// HeartRateLTE applies the LTE predicate on the "heart_rate" field.
func HeartRateLTE(v int) predicate.Help {
	return predicate.Help(sql.FieldLTE(FieldHeartRate, v))
}

// BloodPressureEQ applies the EQ predicate on the "blood_pressure" field.
func BloodPressureEQ(v int) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldBloodPressure, v))
}

// BloodPressureNEQ applies the NEQ predicate on the "blood_pressure" field.
func BloodPressureNEQ(v int) predicate.Help {
	return predicate.Help(sql.FieldNEQ(FieldBloodPressure, v))
}

// BloodPressureIn applies the In predicate on the "blood_pressure" field.
func BloodPressureIn(vs ...int) predicate.Help {
	return predicate.Help(sql.FieldIn(FieldBloodPressure, vs...))
}

// BloodPressureNotIn applies the NotIn predicate on the "blood_pressure" field.
func BloodPressureNotIn(vs ...int) predicate.Help {
	return predicate.Help(sql.FieldNotIn(FieldBloodPressure, vs...))
}

// BloodPressureGT applies the GT predicate on the "blood_pressure" field.
func BloodPressureGT(v int) predicate.Help {
	return predicate.Help(sql.FieldGT(FieldBloodPressure, v))
}

// BloodPressureGTE applies the GTE predicate on the "blood_pressure" field.
func BloodPressureGTE(v int) predicate.Help {
	return predicate.Help(sql.FieldGTE(FieldBloodPressure, v))
}

// BloodPressureLT applies the LT predicate on the "blood_pressure" field.
func BloodPressureLT(v int) predicate.Help {
	return predicate.Help(sql.FieldLT(FieldBloodPressure, v))
}

// BloodPressureLTE applies the LTE predicate on the "blood_pressure" field.
func BloodPressureLTE(v int) predicate.Help {
	return predicate.Help(sql.FieldLTE(FieldBloodPressure, v))
}

// BodyTemperatureEQ applies the EQ predicate on the "body_temperature" field.
func BodyTemperatureEQ(v int) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldBodyTemperature, v))
}

// BodyTemperatureNEQ applies the NEQ predicate on the "body_temperature" field.
func BodyTemperatureNEQ(v int) predicate.Help {
	return predicate.Help(sql.FieldNEQ(FieldBodyTemperature, v))
}

// BodyTemperatureIn applies the In predicate on the "body_temperature" field.
func BodyTemperatureIn(vs ...int) predicate.Help {
	return predicate.Help(sql.FieldIn(FieldBodyTemperature, vs...))
}

// BodyTemperatureNotIn applies the NotIn predicate on the "body_temperature" field.
func BodyTemperatureNotIn(vs ...int) predicate.Help {
	return predicate.Help(sql.FieldNotIn(FieldBodyTemperature, vs...))
}

// BodyTemperatureGT applies the GT predicate on the "body_temperature" field.
func BodyTemperatureGT(v int) predicate.Help {
	return predicate.Help(sql.FieldGT(FieldBodyTemperature, v))
}

// BodyTemperatureGTE applies the GTE predicate on the "body_temperature" field.
func BodyTemperatureGTE(v int) predicate.Help {
	return predicate.Help(sql.FieldGTE(FieldBodyTemperature, v))
}

// BodyTemperatureLT applies the LT predicate on the "body_temperature" field.
func BodyTemperatureLT(v int) predicate.Help {
	return predicate.Help(sql.FieldLT(FieldBodyTemperature, v))
}

// BodyTemperatureLTE applies the LTE predicate on the "body_temperature" field.
func BodyTemperatureLTE(v int) predicate.Help {
	return predicate.Help(sql.FieldLTE(FieldBodyTemperature, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Help {
	return predicate.Help(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Help {
	return predicate.Help(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Help {
	return predicate.Help(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Help {
	return predicate.Help(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Help {
	return predicate.Help(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Help {
	return predicate.Help(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Help {
	return predicate.Help(sql.FieldLTE(FieldCreatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Help) predicate.Help {
	return predicate.Help(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Help) predicate.Help {
	return predicate.Help(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Help) predicate.Help {
	return predicate.Help(sql.NotPredicates(p))
}
