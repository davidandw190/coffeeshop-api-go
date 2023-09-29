package db

import (
	"database/sql"
	"log"
	"time"
)

type DB struct {
	DB *sql.DB
}

var dbConn = &DB{}

const (
	maxOpenDbConn = 10
	maxIdleDbConn = 5
	maxDbLifetime = 5 * time.Minute
)

func ConnectPostgres(dsn string) (*DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDbConn)
	db.SetConnMaxLifetime(maxDbLifetime)

	if err = testDB(db); err != nil {
		return nil, err
	}

	dbConn.DB = db

	return dbConn, nil
}

func testDB(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		log.Println("DB: Error", err)
		return err
	}

	log.Println("**** Database pinged successfully! ****")

	return nil
}
