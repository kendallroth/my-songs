package main

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	_ "github.com/kendallroth/my-songs/migrations"
)

func main() {
	app := pocketbase.New()

	migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
		// Automatically create migrations when making schema changes
		Automigrate: true,
	})

	// Add custom endpoint for user stats
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/api/stats",
			Handler: func(c echo.Context) error {
				authRecord, _ := c.Get(apis.ContextAuthRecordKey).(*models.Record)

				stats := struct {
					Groups int `json:"groups"`
					Lists  int `json:"lists"`
					Songs  int `json:"songs"`
				}{
					Groups: 0,
					Lists:  0,
					Songs:  0,
				}

				err := app.Dao().DB().Select("count(*)").From("groups").Where(dbx.HashExp{"account": authRecord.Id}).Row(&stats.Groups)
				if err != nil {
					apis.NewBadRequestError("Failed to fetch stats", err)
				}

				err = app.Dao().DB().Select("count(*)").From("lists").Where(dbx.HashExp{"account": authRecord.Id}).Row(&stats.Lists)
				if err != nil {
					apis.NewBadRequestError("Failed to fetch stats", err)
				}

				err = app.Dao().DB().Select("count(*)").From("songs").Where(dbx.HashExp{"account": authRecord.Id}).Row(&stats.Songs)
				if err != nil {
					apis.NewBadRequestError("Failed to fetch stats", err)
				}

				return c.JSON(http.StatusOK, stats)
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
				apis.RequireRecordAuth("accounts"),
			},
		})

		return nil
	})

	// Set last login time on successful login
	app.OnRecordAfterAuthWithPasswordRequest().Add(func(e *core.RecordAuthWithPasswordEvent) error {
		e.Record.Set("lastLoginAt", time.Now())

		if err := app.Dao().SaveRecord(e.Record); err != nil {
			return err
		}

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
