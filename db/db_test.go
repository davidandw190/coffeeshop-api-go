package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func TestConnectPostgres(t *testing.T) {
	t.Parallel()

	t.Run("Successful Connection", func(t *testing.T) {
		// Test connecting to a PostgreSQL database with valid credentials.
		testDSN := os.Getenv("TEST_DSN")
		db, err := ConnectPostgres(testDSN)

		if err != nil {
			t.Errorf("ConnectPostgres() error = %v, want nil", err)
		}

		if db == nil || db.DB == nil {
			t.Error("ConnectPostgres() returned nil database connection")
		}

		defer db.DB.Close()
	})

	t.Run("Invalid Connection String", func(t *testing.T) {
		// Test connecting with an invalid connection string.
		invalidDSN := "invalid_connnection_string"
		_, err := ConnectPostgres(invalidDSN)

		if err == nil {
			t.Error("ConnectPostgres() expected an error, got nil")
		}
	})

	t.Run("Database Ping Error", func(t *testing.T) {
		// Test when the database ping fails.
		mockDB := &sql.DB{}
		if err := testDB(mockDB); err == nil {
			t.Error("testDB() expected an error, got nil")
		}
	})
}
