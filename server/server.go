package server

import (
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	_ "github.com/withuscpa/withuscpa-api/docs"

	"github.com/DeltaLaboratory/incident-api/internal/config"
	"github.com/DeltaLaboratory/incident-api/internal/ent"
	"github.com/DeltaLaboratory/incident-api/server/common"
	"github.com/DeltaLaboratory/incident-api/server/v0/incident"
)

type Server struct {
	app *fiber.App
	db  *ent.Client

	logger zerolog.Logger

	client *http.Client
}

func (s *Server) Run() {
	s.app = fiber.New()
	s.logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	s.client = &http.Client{}

	if err := config.LoadConfig(); err != nil {
		s.logger.Fatal().Err(err).Msg("failed to load config")
		return
	}

	if err := s.ConnectDB(); err != nil {
		s.logger.Fatal().Err(err).Msg("failed to connect to database")
		return
	}

	{
		router := common.Router{DB: s.db, Logger: s.logger}
		routeIncident := &incident.Incident{Router: router}

		routeIncident.Route(s.app, "/v0/incident")

		s.app.Get("/docs/*", swagger.HandlerDefault)
	}

	if err := s.app.Listen(config.C.ListenAddress); err != nil {
		s.logger.Fatal().Err(err).Msg("failed to start server")
	}
}
