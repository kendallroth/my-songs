package routes

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

// Add dashboard routes
func AddRoutes(app *pocketbase.PocketBase) {
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

				// NOTE: Previous approach ('Row' sets per variable)
				// err := app.Dao().DB().Select("count(*)").From("groups").Where(dbx.HashExp{"account": authRecord.Id}).Row(&stats.Groups)

				err := app.Dao().DB().Select("*").From("account_stats").Where(dbx.HashExp{"id": authRecord.Id}).One(&stats)
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
}
