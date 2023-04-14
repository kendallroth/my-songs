package utils

import (
	"time"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

// Set last login time on successful login
func AddLoginLogging(app *pocketbase.PocketBase) {
	app.OnRecordAfterAuthWithPasswordRequest().Add(func(e *core.RecordAuthWithPasswordEvent) error {
		e.Record.Set("lastLoginAt", time.Now())

		if err := app.Dao().SaveRecord(e.Record); err != nil {
			return err
		}

		return nil
	})
}
