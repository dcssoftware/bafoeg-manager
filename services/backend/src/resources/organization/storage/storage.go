package storage

import "github.com/go-sqlx/sqlx"

type OrganizationStorage struct {
	db *sqlx.DB
}

func NewOrganizationStorage(db *sqlx.DB) *OrganizationStorage {
	return &OrganizationStorage{db: db}
}
