package storage

import (
	"github.com/go-sqlx/sqlx"
)

type ApplicationsStorage struct {
	db *sqlx.DB
}

func NewApplicationsStorage(db *sqlx.DB) *ApplicationsStorage {
	return &ApplicationsStorage{
		db: db,
	}
}
