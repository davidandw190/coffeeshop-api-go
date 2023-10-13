package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

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

func TestDatabaseIntegration(t *testing.T) {
	t.Parallel()

	testDSN := os.Getenv("TEST_DSN")
	db, err := ConnectPostgres(testDSN)
	if err != nil {
		t.Fatalf("ConnectPostgres() error = %v, want nil", err)
	}

	t.Run("Successful Ping", func(t *testing.T) {
		// Test successful ping of the connected database.
		err := testDB(db.DB)

		if err != nil {
			t.Errorf("testDB() error = %v, want nil", err)
		}
	})

	t.Run("Max Open Connections", func(t *testing.T) {
		// Test if the database enforces the maximum number of open connections.
		// REMINDER: Adjust maxOpenDbConn to a smaller value for this test to fail.

		maxConnections := db.DB.Stats().MaxOpenConnections
		if maxConnections != maxOpenDbConn {
			t.Errorf("MaxOpenConnections = %d, want %d", maxConnections, maxOpenDbConn)
		}
	})

	t.Run("Max Idle Connections", func(t *testing.T) {
		// Test if the database enforces the maximum number of idle connections.
		// REMINDER: Adjust maxIdleDbConn to a smaller value for this test to fail.
		idleConnections := db.DB.Stats().Idle
		if idleConnections != maxIdleDbConn {
			t.Errorf("Idle Connections = %d, want %d", idleConnections, maxIdleDbConn)
		}
	})

	t.Run("Connection Lifetime", func(t *testing.T) {
		// Test if the database enforces the maximum connection lifetime.
		// REMINDER: Adjust maxDbLifetime to a smaller value for this test to fail.
		lifetime := time.Duration(db.DB.Stats().MaxLifetimeClosed)
		expectedLifetime := maxDbLifetime
		if lifetime != expectedLifetime {
			t.Errorf("Connection Lifetime = %v, want %v", lifetime, expectedLifetime)
		}
	})
}
