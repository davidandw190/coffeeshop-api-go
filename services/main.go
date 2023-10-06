package services

import (
	"database/sql"
	"time"
)

var db *sql.DB

const dbTimeout = time.Second * 3

// Models contains instances of data models and services.
type Models struct {
	Coffee       Coffee
	JsonResponse JsonResponse
}

// New creates a new Models instance with the given database connection pool.
func New(dbPool *sql.DB) Models {
	db = dbPool
	return Models{}
}
