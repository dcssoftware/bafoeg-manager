package storage

import (
	"context"

	"github.com/go-sqlx/sqlx"
)

type FileStorage struct {
	db *sqlx.DB
}

func NewFileStorage(db *sqlx.DB) *FileStorage {
	return &FileStorage{
		db: db,
	}
}

func (s *FileStorage) StartTx() (*sqlx.Tx, error) {
	return s.db.BeginTxx(context.Background(), nil)
}
