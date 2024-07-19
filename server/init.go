package server

import (
	"context"
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/DeltaLaboratory/incident-api/internal/config"
	"github.com/DeltaLaboratory/incident-api/internal/ent"
)

func (s *Server) ConnectDB() error {
	db, err := sql.Open("pgx", config.C.DB.URI())
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	s.db = ent.NewClient(ent.Driver(drv))
	s.db = s.db.Debug()

	if !fiber.IsChild() {
		if err := s.db.Schema.Create(context.TODO(),
			schema.WithDropColumn(true),
			schema.WithDropIndex(true)); err != nil {
			return fmt.Errorf("failed to create schema: %w", err)
		}
	}

	return nil
}
