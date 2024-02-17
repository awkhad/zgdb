package migrations

import (
	"github.com/zdb/dbx"
)

// Cleanup dangling deleted collections references
// (see https://https://github.com/awkhad/zserver/discussions/2570).
func init() {
	AppMigrations.Register(func(db dbx.Builder) error {
		_, err := db.NewQuery(`
			DELETE FROM {{_externalAuths}}
			WHERE [[collectionId]] NOT IN (SELECT [[id]] FROM {{_collections}})
		`).Execute()

		return err
	}, func(db dbx.Builder) error {
		return nil
	})
}
