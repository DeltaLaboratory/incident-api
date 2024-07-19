package incident

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/geojson"

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

	encodedGeo, err := geojson.Encode(geom.NewPoint(geom.XY).MustSetCoords([]float64{float64(req.Longitude), float64(req.Latitude)}).SetSRID(4326))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.ErrorResponse{
			Code:    common.GeneralInternalError,
			Message: err.Error(),
		})
	}

	qb := i.DB.Help.Create()
	qb.SetReporter(req.UserID)
	qb.SetIdempotencyKey(req.IdempotencyKey)
	qb.SetLocation(&schema.GeoJson{
		Geometry: encodedGeo,
	})
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
		// use postgis function to calculate distance
		qb = qb.Where(func(selector *sql.Selector) {
			selector.Where(sql.ExprP(fmt.Sprintf("ST_DWithin(location, 'POINT(%f %f)', %d)", req.Longitude, req.Latitude, req.Radius)))
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
		decodedGeo, err := h.Location.Decode()
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(common.ErrorResponse{
				Code:    common.GeneralInternalError,
				Message: err.Error(),
			})
		}

		response[i] = QueryHelpResponse{
			Latitude:        schema.Coordinate(decodedGeo.FlatCoords()[0]),
			Longitude:       schema.Coordinate(decodedGeo.FlatCoords()[1]),
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
