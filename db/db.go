package db

import (
	"database/sql"
	"log"
	"time"
)

// DB represents the database connection.
type DB struct {
	DB *sql.DB
}

var dbConn = &DB{}

const (
	maxOpenDbConn = 10
	maxIdleDbConn = 5
	maxDbLifetime = 5 * time.Minute
)

// ConnectPostgres establishes a connection to a PostgreSQL database.
func ConnectPostgres(dsn string) (*DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	// Database connection pool settings
	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDbConn)
	db.SetConnMaxLifetime(maxDbLifetime)

	if err = testDB(db); err != nil {
		return nil, err
	}

	dbConn.DB = db

	return dbConn, nil
}

// testDB pings the database to ensure the connection is active.
func testDB(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		log.Println("DB: Error", err)
		return err
	}

	log.Println("**** Database pinged successfully! ****")

	return nil
}
