package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upDirDbMigrations, downDirDbMigrations)
}

func upDirDbMigrations(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func downDirDbMigrations(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
