package migrations

import (
	"github.com/zdb/dbx"
)

// Renormalizes old single and multiple values of MultiValuer fields (file, select, relation)
// (see https://https://github.com/awkhad/zserver/issues/2930).
func init() {
	AppMigrations.Register(func(db dbx.Builder) error {
		return normalizeMultivaluerFields(db)
	}, func(db dbx.Builder) error {
		return nil
	})
}
