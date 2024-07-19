package incident

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/geojson"

	"github.com/DeltaLaboratory/incident-api/internal/ent"
	"github.com/DeltaLaboratory/incident-api/internal/ent/incident"
	"github.com/DeltaLaboratory/incident-api/internal/ent/schema"
	"github.com/DeltaLaboratory/incident-api/server/common"
)

// Report godoc
//
//	@Summary		Report an incident
//	@Description	Report an incident
//	@Tags			Incident
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ReportRequest	true	"Report request"
//	@Success		200		{object}	ReportResponse
//	@Failure		400		{object}	common.ErrorResponse
//	@Failure		500		{object}	common.ErrorResponse
//
//	@Router			/incident/report [post]
func (i *Incident) Report(ctx *fiber.Ctx) error {
	req := new(ReportRequest)
	if err := common.ParseBody(ctx, req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Code:    common.GeneralBadRequest,
			Message: fmt.Errorf("failed to parse request: %w", err).Error(),
		})
	}

	encodedGeo, err := geojson.Encode(geom.NewPoint(geom.XY).MustSetCoords([]float64{float64(req.Latitude), float64(req.Longitude)}).SetSRID(4326))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.ErrorResponse{
			Code:    common.GeneralInternalError,
			Message: fmt.Errorf("failed to encode location: %w", err).Error(),
		})
	}

	qb := i.DB.Incident.Create()
	qb.SetIdempotencyKey(req.UserKey)
	qb.SetReporter(req.UserID)
	qb.SetType(req.Type)
	qb.SetLocation(&schema.GeoJson{Geometry: encodedGeo})
	qb.SetDescription(req.Description)

	if req.Image != nil {
		decoded, err := base64.StdEncoding.DecodeString(*req.Image)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
				Code:    common.GeneralBadRequest,
				Message: fmt.Errorf("failed to decode image: %w", err).Error(),
			})
		}

		if http.DetectContentType(decoded) != "image/jpeg" {
			i.Logger.Warn().Str("content_type", http.DetectContentType(decoded)).Msg("invalid image content type")
			return ctx.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
				Code:    common.GeneralBadRequest,
				Message: "image must be jpeg",
			})
		}

		qb.SetImage(decoded)
	} else {
		qb.SetImage(nil)
	}

	inci, err := qb.Save(ctx.Context())
	if err != nil {
		i.Logger.Error().Err(err).Msg("failed to save incident")
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.ErrorResponse{
			Code:    common.GeneralInternalError,
			Message: "failed to save incident",
		})
	}

	return ctx.JSON(ReportResponse{
		IncidentID: inci.ID,
	})
}

// DeleteReport godoc
//
//	@Summary		Delete an incident report
//	@Description	Delete an incident report
//	@Tags			Incident
//	@Accept			json
//	@Param			request	body	DeleteReportRequest	true	"Delete report request"
//	@Success		204
//	@Failure		400	{object}	common.ErrorResponse
//	@Failure		404	{object}	common.ErrorResponse
//	@Failure		500	{object}	common.ErrorResponse
//
//	@Router			/incident/report [delete]
func (i *Incident) DeleteReport(ctx *fiber.Ctx) error {
	req := new(DeleteReportRequest)
	if err := common.ParseBody(ctx, req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Code:    common.GeneralBadRequest,
			Message: fmt.Errorf("failed to parse request: %w", err).Error(),
		})
	}

	inci, err := i.DB.Incident.Query().Where(incident.ID(req.IncidentID)).Only(ctx.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return ctx.Status(fiber.StatusNotFound).JSON(common.ErrorResponse{
				Code:    "incident:not_found",
				Message: "incident not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.ErrorResponse{
			Code:    common.GeneralInternalError,
			Message: "failed to query incident",
		})
	}

	if inci.Reporter != req.UserID {
		return ctx.Status(fiber.StatusUnauthorized).JSON(common.ErrorResponse{
			Code:    "incident:unauthorized",
			Message: "unauthorized",
		})
	}

	if err := i.DB.Incident.DeleteOneID(req.IncidentID).Exec(ctx.Context()); err != nil {
		i.Logger.Error().Err(err).Msg("failed to delete incident")
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.ErrorResponse{
			Code:    common.GeneralInternalError,
			Message: "failed to delete incident",
		})
	}

	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

// Query godoc
//
//	@Summary		Query incidents
//	@Description	Query incidents
//	@Tags			Incident
//	@Accept			json
//	@Produce		json
//	@Param			request	body		QueryRequest	true	"Query request"
//	@Success		200		{object}	[]QueryResponse
//	@Failure		400		{object}	common.ErrorResponse
//	@Failure		500		{object}	common.ErrorResponse
//
//	@Router			/incident/query [post]
func (i *Incident) Query(ctx *fiber.Ctx) error {
	req := new(QueryRequest)
	if err := common.ParseBody(ctx, req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Code:    common.GeneralBadRequest,
			Message: fmt.Errorf("failed to parse request: %w", err).Error(),
		})
	}

	qb := i.DB.Incident.Query()

	if req.IncidentID != nil {
		qb = qb.Where(incident.ID(*req.IncidentID))
	} else {
		qb = qb.Where(func(selector *sql.Selector) {
			selector.Where(sql.ExprP(fmt.Sprintf("ST_DWithin(location, ST_GeomFromText('POINT (%f %f)',4326), %d)", req.Longitude, req.Latitude, req.Radius)))
		})
	}

	incidents, err := qb.Select(incident.FieldID, incident.FieldType, incident.FieldReporter, incident.FieldLocation, incident.FieldDescription, incident.FieldVote, incident.FieldCreatedAt).All(ctx.Context())
	if err != nil {
		i.Logger.Error().Err(err).Msg("failed to query incidents")
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.ErrorResponse{
			Code:    common.GeneralInternalError,
			Message: "failed to query incidents",
		})
	}

	response := make([]QueryResponse, len(incidents))
	for idx, inc := range incidents {
		decodedGeo, err := inc.Location.Decode()
		if err != nil {
			i.Logger.Error().Err(err).Msg("failed to decode location")
			return ctx.Status(fiber.StatusInternalServerError).JSON(common.ErrorResponse{
				Code:    common.GeneralInternalError,
				Message: "failed to decode location",
			})
		}

		response[idx] = QueryResponse{
			IncidentID:  inc.ID,
			Reporter:    inc.Reporter,
			Latitude:    schema.Coordinate(decodedGeo.FlatCoords()[0]),
			Longitude:   schema.Coordinate(decodedGeo.FlatCoords()[1]),
			Type:        inc.Type,
			Description: inc.Description,
			Vote:        inc.Vote,
			CreatedAt:   inc.CreatedAt,
		}
	}

	return ctx.JSON(response)
}

// QueryImage godoc
//
//	@Summary		Query incident image
//	@Description	Query incident image
//	@Tags			Incident
//	@Param			incident_id	path		string	true	"Incident ID"
//	@Success		200			{string}	image/jpeg
//	@Failure		400			{object}	common.ErrorResponse
//	@Failure		404			{object}	common.ErrorResponse
//	@Failure		500			{object}	common.ErrorResponse
//
//	@Router			/incident/image/{incident_id} [get]
func (i *Incident) QueryImage(ctx *fiber.Ctx) error {
	incidentID, err := uuid.Parse(ctx.Params("incident_id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Code:    common.GeneralBadRequest,
			Message: "invalid incident_id",
		})
	}

	inc, err := i.DB.Incident.Query().Where(incident.ID(incidentID)).Only(ctx.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return ctx.Status(fiber.StatusNotFound).JSON(common.ErrorResponse{
				Code:    "incident:not_found",
				Message: "incident not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.ErrorResponse{
			Code:    common.GeneralInternalError,
			Message: "failed to query incident",
		})
	}

	if inc.Image == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(common.ErrorResponse{
			Code:    "incident:image_not_found",
			Message: "image not found",
		})
	}

	ctx.Set("Content-Type", "image/jpeg")
	return ctx.Send(*inc.Image)
}

// Vote godoc
//
//	@Summary		Vote an incident
//	@Description	Vote an incident
//	@Tags			Incident
//	@Accept			json
//	@Param			request	body		VoteRequest	true	"Vote request"
//	@Success		204
//	@Failure		400	{object}	common.ErrorResponse
//	@Failure		404	{object}	common.ErrorResponse
//	@Failure		409	{object}	common.ErrorResponse
//	@Failure		500	{object}	common.ErrorResponse
//
//	@Router			/incident/vote [post]
func (i *Incident) Vote(ctx *fiber.Ctx) error {
	req := new(VoteRequest)
	if err := common.ParseBody(ctx, req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Code:    common.GeneralBadRequest,
			Message: fmt.Errorf("failed to parse request: %w", err).Error(),
		})
	}

	tx, err := i.DB.Tx(ctx.Context())
	if err != nil {
		i.Logger.Error().Err(err).Msg("failed to start transaction")
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.ErrorResponse{
			Code:    common.GeneralInternalError,
			Message: "failed to start transaction",
		})
	}

	inci, err := tx.Incident.Query().Where(incident.ID(req.IncidentID)).Select(incident.FieldVoteFilter).Only(ctx.Context())
	if err != nil {
		if err := tx.Rollback(); err != nil {
			i.Logger.Warn().Err(err).Msg("failed to rollback transaction")
		}

		if ent.IsNotFound(err) {
			return ctx.Status(fiber.StatusNotFound).JSON(common.ErrorResponse{
				Code:    "incident:not_found",
				Message: "incident not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.ErrorResponse{
			Code:    common.GeneralInternalError,
			Message: "failed to query incident",
		})
	}

	var g = &bloom.BloomFilter{}
	if inci.VoteFilter != nil {
		if _, err := g.ReadFrom(bytes.NewReader(*inci.VoteFilter)); err != nil {
			if err := tx.Rollback(); err != nil {
				i.Logger.Warn().Err(err).Msg("failed to rollback transaction")
			}

			i.Logger.Error().Err(err).Msg("failed to read bloom filter")
			return ctx.Status(fiber.StatusInternalServerError).JSON(common.ErrorResponse{
				Code:    common.GeneralInternalError,
				Message: "failed to read bloom filter",
			})
		}
	} else {
		g = bloom.New(1000000, 5)
	}

	if g.Test(req.UserID[:]) {
		if err := tx.Rollback(); err != nil {
			i.Logger.Warn().Err(err).Msg("failed to rollback transaction")
		}

		return ctx.Status(fiber.StatusConflict).JSON(common.ErrorResponse{
			Code:    "incident:voted",
			Message: "already voted",
		})
	}

	g.Add(req.UserID[:])
	var buf bytes.Buffer
	qb := tx.Incident.UpdateOneID(req.IncidentID)
	if _, err := g.WriteTo(&buf); err != nil {
		if err := tx.Rollback(); err != nil {
			i.Logger.Warn().Err(err).Msg("failed to rollback transaction")
		}

		i.Logger.Error().Err(err).Msg("failed to write bloom filter")
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.ErrorResponse{
			Code:    common.GeneralInternalError,
			Message: "failed to write bloom filter",
		})
	}

	qb.SetVoteFilter(buf.Bytes())
	if req.VotePositive {
		qb.AddVote(1)
	} else {
		qb.AddVote(-1)
	}

	if _, err := qb.Save(ctx.Context()); err != nil {
		if err := tx.Rollback(); err != nil {
			i.Logger.Warn().Err(err).Msg("failed to rollback transaction")
		}

		i.Logger.Error().Err(err).Msg("failed to save vote")
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.ErrorResponse{
			Code:    common.GeneralInternalError,
			Message: "failed to save vote",
		})
	}

	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			i.Logger.Warn().Err(err).Msg("failed to rollback transaction")
		}

		i.Logger.Error().Err(err).Msg("failed to commit transaction")
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.ErrorResponse{
			Code:    common.GeneralInternalError,
			Message: "failed to commit transaction",
		})
	}

	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

type ReportRequest struct {
	// user id in uuid format
	UserID uuid.UUID `json:"user_id"`
	// 32 byte ascii string
	UserKey string `json:"user_key"`

	Latitude  schema.Coordinate `json:"latitude"`
	Longitude schema.Coordinate `json:"longitude"`

	Type        string `json:"type"`
	Description string `json:"description"`

	// Base64 encoded jpeg image
	Image *string `json:"image"`
}

func (r *ReportRequest) Validate() error {
	if r.UserID == uuid.Nil {
		return fmt.Errorf("user_id is required")
	}

	if r.UserKey == "" {
		return fmt.Errorf("user_key is required")
	}

	if len(r.UserKey) != 32 {
		return fmt.Errorf("length of user_key must be 32")
	}

	if r.Latitude == 0 {
		return fmt.Errorf("latitude is required")
	}

	if r.Longitude == 0 {
		return fmt.Errorf("longitude is required")
	}

	if r.Type == "" {
		return fmt.Errorf("type is required")
	}

	if len(r.Type) > 50 {
		return fmt.Errorf("max length of type is 50")
	}

	if r.Description == "" {
		return fmt.Errorf("description is required")
	}

	if len(r.Description) > 1500 {
		return fmt.Errorf("max length of description is 1500")
	}

	if r.Image != nil {
		if len(*r.Image) > 1024*1024*50 {
			return fmt.Errorf("max size of image is 50MB")
		}
	}

	return nil

}

type ReportResponse struct {
	IncidentID uuid.UUID `json:"incident_id"`
}

type DeleteReportRequest struct {
	IncidentID uuid.UUID `json:"incident_id"`
	UserID     uuid.UUID `json:"user_id"`
	UserKey    string    `json:"user_key"`
}

func (r *DeleteReportRequest) Validate() error {
	if r.IncidentID == uuid.Nil {
		return fmt.Errorf("incident_id is required")
	}

	if r.UserID == uuid.Nil {
		return fmt.Errorf("user_id is required")
	}

	if r.UserKey == "" {
		return fmt.Errorf("user_key is required")
	}

	if len(r.UserKey) != 32 {
		return fmt.Errorf("length of user_key must be 32")
	}

	return nil
}

type QueryRequest struct {
	Latitude  schema.Coordinate `json:"latitude"`
	Longitude schema.Coordinate `json:"longitude"`

	// Radius in meters
	Radius int `json:"radius"`

	IncidentID *uuid.UUID `json:"incident_id"`
}

func (q *QueryRequest) Validate() error {
	if q.IncidentID != nil {
		if *q.IncidentID == uuid.Nil {
			return fmt.Errorf("incident_id is required")
		}
		return nil
	}

	if q.Latitude == 0 {
		return fmt.Errorf("latitude is required")
	}

	if q.Longitude == 0 {
		return fmt.Errorf("longitude is required")
	}

	if q.Radius == 0 {
		return fmt.Errorf("radius is required")
	}

	if q.Radius > 10000 {
		return fmt.Errorf("max radius is 10km")
	}

	return nil
}

type QueryResponse struct {
	IncidentID uuid.UUID `json:"incident_id"`

	Reporter    uuid.UUID         `json:"reporter"`
	Latitude    schema.Coordinate `json:"latitude"`
	Longitude   schema.Coordinate `json:"longitude"`
	Type        string            `json:"type"`
	Description string            `json:"description"`

	Vote int `json:"vote"`

	CreatedAt time.Time `json:"created_at"`
}

type VoteRequest struct {
	UserID       uuid.UUID `json:"user_id"`
	IncidentID   uuid.UUID `json:"incident_id"`
	VotePositive bool      `json:"vote_positive"`
}
