package storage

import (
	"context"

	"github.com/go-sqlx/sqlx"
)

type RAGStorage struct {
	db *sqlx.DB
}

func NewRAGStorage(db *sqlx.DB) *RAGStorage {
	return &RAGStorage{
		db: db,
	}
}

func (s *RAGStorage) StartTx() (*sqlx.Tx, error) {
	return s.db.BeginTxx(context.Background(), nil)
}
