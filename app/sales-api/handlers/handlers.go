// Package handlers contains full set of handlers and routes
package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/baybaraandrey/advertising/business/data/advert"
	"github.com/baybaraandrey/advertising/business/data/category"
	"github.com/baybaraandrey/advertising/business/data/user"
	"github.com/baybaraandrey/advertising/business/mid"
	"github.com/baybaraandrey/advertising/foundation/web"
	"github.com/jmoiron/sqlx"
)

// API constructs an http.Handler with all application routes defined
func API(build string, shutdown chan os.Signal, log *log.Logger, db *sqlx.DB) *web.App {

	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Panics(log), mid.Metrics())

	check := checkGroup{
		log: log,
		db:  db,
	}
	app.Handle(http.MethodGet, "/v1/rediness/", check.readiness)

	// Register user management and authentication endpoints.
	ug := userGroup{
		user: user.New(log, db),
	}
	app.Handle(http.MethodGet, "/v1/users/", ug.query)
	app.Handle(http.MethodPost, "/v1/users/", ug.create)
	app.Handle(http.MethodGet, "/v1/users/:id/", ug.queryByID)

	// Register category management endpoints
	cg := categoryGroup{
		category: category.New(log, db),
	}
	app.Handle(http.MethodGet, "/v1/categories/", cg.query)
	app.Handle(http.MethodPost, "/v1/categories/", cg.create)

	// Register adverts management endpoints
	ag := advertGroup{
		advert: advert.New(log, db),
	}
	app.Handle(http.MethodGet, "/v1/adverts/", ag.query)

	return app
}
