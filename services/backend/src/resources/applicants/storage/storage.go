package storage

import "github.com/go-sqlx/sqlx"

type ApplicantStorage struct {
	db *sqlx.DB
}

func NewApplicantStorage(dbConn *sqlx.DB) *ApplicantStorage {
	return &ApplicantStorage{
		db: dbConn,
	}
}
