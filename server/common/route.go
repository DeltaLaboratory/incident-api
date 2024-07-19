package common

import (
	"github.com/rs/zerolog"

	"github.com/DeltaLaboratory/incident-api/internal/ent"
)

type Router struct {
	DB     *ent.Client
	Logger zerolog.Logger
}
