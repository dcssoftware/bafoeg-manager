package storage

import (
	"context"

	"github.com/go-sqlx/sqlx"
)

type EakteStorage struct {
	db *sqlx.DB
}

func NewEakteStorage(db *sqlx.DB) *EakteStorage {
	return &EakteStorage{
		db: db,
	}
}

func (s *EakteStorage) StartTx() (*sqlx.Tx, error) {
	return s.db.BeginTxx(context.Background(), nil)
}
