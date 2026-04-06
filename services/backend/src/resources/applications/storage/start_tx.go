package storage

import (
	"context"

	"github.com/go-sqlx/sqlx"
)

func (s *ApplicationsStorage) StartTx() (*sqlx.Tx, error) {
	return s.db.BeginTxx(context.Background(), nil)
}
