package incident

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/DeltaLaboratory/incident-api/internal/ent/help"
	"github.com/DeltaLaboratory/incident-api/internal/ent/schema"
	"github.com/DeltaLaboratory/incident-api/server/common"
)

// Help godoc
//
//	@Summary		Help
//	@Description	Help
//	@Tags			Incident
//	@Accept			json
//	@Param			request	body	HelpRequest	true	"Help request"
//	@Success		201
//	@Failure		400	{object}	common.ErrorResponse
//	@Failure		500	{object}	common.ErrorResponse
//	@Router			/incident/help [post]
func (i *Incident) Help(ctx *fiber.Ctx) error {
	req := new(HelpRequest)
	if err := common.ParseBody(ctx, req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Code:    common.GeneralBadRequest,
			Message: err.Error(),
		})
	}

	qb := i.DB.Help.Create()
	qb.SetReporter(req.UserID)
	qb.SetIdempotencyKey(req.IdempotencyKey)
	qb.SetLatitude(req.Latitude)
	qb.SetLongitude(req.Longitude)
	qb.SetDescription(req.Description)

	if req.HeartRate != nil {
		qb.SetHeartRate(*req.HeartRate)
	}

	if req.BloodPressure != nil {
		qb.SetBloodPressure(*req.BloodPressure)
	}

	if req.BodyTemperature != nil {
		qb.SetBodyTemperature(*req.BodyTemperature)
	}

	if _, err := qb.Save(ctx.Context()); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.ErrorResponse{
			Code:    common.GeneralInternalError,
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).Send(nil)
}

// QueryHelp godoc
//
//	@Summary		Query help
//	@Description	Query help
//	@Tags			Incident
//	@Accept			json
//	@Produce		json
//	@Param			request	body		QueryHelpRequest	true	"Query help request"
//	@Success		200		{object}	[]QueryHelpResponse
//	@Failure		400		{object}	common.ErrorResponse
//	@Failure		500		{object}	common.ErrorResponse
//	@Router			/incident/help/query [post]
func (i *Incident) QueryHelp(ctx *fiber.Ctx) error {
	req := new(QueryHelpRequest)
	if err := common.ParseBody(ctx, req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Code:    common.GeneralBadRequest,
			Message: err.Error(),
		})
	}

	qb := i.DB.Help.Query()

	if req.HelpID != nil {
		qb = qb.Where(help.ID(*req.HelpID))
	} else {
		// Convert LatLong to radians
		lat1 := req.Latitude.Radians()
		lon1 := req.Longitude.Radians()

		qb = qb.Where(func(s *sql.Selector) {
			s.Where(sql.ExprP(fmt.Sprintf(
				"6371000 * 2 * ASIN(SQRT(POWER(SIN((%f - RADIANS(latitude)) / 2), 2) + COS(%f) * COS(RADIANS(latitude)) * POWER(SIN((%f - RADIANS(longitude)) / 2), 2))) <= ?",
				lat1, lat1, lon1,
			), req.Radius))
		})
	}

	helps, err := qb.All(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.ErrorResponse{
			Code:    common.GeneralInternalError,
			Message: err.Error(),
		})
	}

	response := make([]QueryHelpResponse, len(helps))
	for i, h := range helps {
		response[i] = QueryHelpResponse{
			Latitude:        h.Latitude,
			Longitude:       h.Longitude,
			Description:     h.Description,
			HeartRate:       h.HeartRate,
			BloodPressure:   h.BloodPressure,
			BodyTemperature: h.BodyTemperature,
			CreatedAt:       h.CreatedAt,
		}
	}

	return ctx.JSON(response)
}

type HelpRequest struct {
	UserID         uuid.UUID `json:"user_id"`
	IdempotencyKey string    `json:"idempotency_key"`

	Latitude    schema.Coordinate `json:"latitude"`
	Longitude   schema.Coordinate `json:"longitude"`
	Description string            `json:"description"`

	HeartRate       *int `json:"heart_rate"`
	BloodPressure   *int `json:"blood_pressure"`
	BodyTemperature *int `json:"body_temperature"`
}

type QueryHelpRequest struct {
	Latitude  schema.Coordinate `json:"latitude"`
	Longitude schema.Coordinate `json:"longitude"`

	// Radius in meters
	Radius int `json:"radius"`

	HelpID *uuid.UUID `json:"help_id"`
}

type QueryHelpResponse struct {
	Reporter    uuid.UUID         `json:"reporter"`
	Latitude    schema.Coordinate `json:"latitude"`
	Longitude   schema.Coordinate `json:"longitude"`
	Description string            `json:"description"`

	HeartRate       *int `json:"heart_rate"`
	BloodPressure   *int `json:"blood_pressure"`
	BodyTemperature *int `json:"body_temperature"`

	CreatedAt time.Time `json:"created_at"`
}
