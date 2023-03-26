package main

import (
	"log"
	"time"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	_ "github.com/kendallroth/my-songs/migrations"
)

func main() {
	app := pocketbase.New()

	migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
		// Automatically create migrations when making schema changes
		Automigrate: true,
	})

	app.OnRecordAfterAuthWithPasswordRequest().Add(func(e *core.RecordAuthWithPasswordEvent) error {
		// Set last login time on successful login
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
