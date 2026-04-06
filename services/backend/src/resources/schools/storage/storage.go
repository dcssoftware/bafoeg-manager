package storage

import "github.com/go-sqlx/sqlx"

type SchoolStorage struct {
	db *sqlx.DB
}

func NewSchoolStorage(db *sqlx.DB) *SchoolStorage {
	return &SchoolStorage{
		db,
	}
}
