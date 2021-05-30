package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/baybaraandrey/advertising/foundation/database"
	"github.com/baybaraandrey/advertising/foundation/web"
	"github.com/jmoiron/sqlx"
)

type checkGroup struct {
	log *log.Logger
	db  *sqlx.DB
}

type health struct {
	Status string `json:"status"`
}

// @Summary check health
// @Description check health
// @Produce  json
// @Success 200 {object} health
// @Router /v1/rediness/ [get]
// @Tags health
func (c checkGroup) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := "ok"
	code := http.StatusOK

	if err := database.StatusCheck(context.Background(), c.db); err != nil {
		status = "db not ready"
		code = http.StatusInternalServerError
	}
	health := health{
		Status: status,
	}

	return web.Respond(ctx, w, health, code)
}
