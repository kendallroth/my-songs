package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	dashboard_routes "mysongs/internal/routes"
	api_utils "mysongs/internal/utils"

	_ "mysongs/migrations"
)

func main() {
	app := pocketbase.New()

	migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
		// Automatically create migrations when making schema changes
		Automigrate: true,
	})

	dashboard_routes.AddRoutes(app)

	api_utils.AddLoginLogging(app)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
