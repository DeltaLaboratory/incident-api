package incident

import (
	"github.com/gofiber/fiber/v2"

	"github.com/DeltaLaboratory/incident-api/server/common"
)

type Incident struct {
	common.Router
}

func (i *Incident) Route(route *fiber.App, prefix ...string) {
	rpx := "/crawl"
	if len(prefix) > 0 {
		rpx = prefix[0]
	}

	r := route.Group(rpx)
	r.Post("/report", i.Report)
	r.Delete("/report", i.DeleteReport)
	r.Post("/query", i.Query)
	r.Get("/image/:incident_id", i.Query)
	r.Post("/help", i.Help)
	r.Post("/help/query", i.QueryHelp)
}
