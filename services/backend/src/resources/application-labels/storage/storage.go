package storage

import "github.com/go-sqlx/sqlx"

type ApplicationLabelsStorage struct {
	db *sqlx.DB
}

func NewApplicationLabelsStorage(dbConn *sqlx.DB) *ApplicationLabelsStorage {
	return &ApplicationLabelsStorage{
		db: dbConn,
	}
}
