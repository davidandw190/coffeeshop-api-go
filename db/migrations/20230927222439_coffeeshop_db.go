package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCoffeeshopDb, downCoffeeshopDb)
}

func upCoffeeshopDb(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func downCoffeeshopDb(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
