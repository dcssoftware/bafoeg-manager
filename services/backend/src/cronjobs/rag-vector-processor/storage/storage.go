package storage

import "github.com/go-sqlx/sqlx"

type CronjobRagVectorProcessorStorage struct {
	db *sqlx.DB
}

func NewCronjobRagVectorProcessorStorage(db *sqlx.DB) *CronjobRagVectorProcessorStorage {
	return &CronjobRagVectorProcessorStorage{
		db: db,
	}
}
