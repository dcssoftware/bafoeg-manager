package rag

import (
	"context"

	"github.com/dcssoftware/bafoeg-manager/src/helper/database/connstrbuilder"
	"github.com/tmc/langchaingo/vectorstores/pgvector"
)

func CreatePGVectorConnection(opts ...pgvector.Option) (*pgvector.Store, error) {
	connStr := connstrbuilder.GetSQLConnectionString()

	var pgVectorArguments []pgvector.Option

	pgVectorArguments = append(pgVectorArguments, pgvector.WithConnectionURL(connStr))
	pgVectorArguments = append(pgVectorArguments, opts...)

	pgvectorStore, pgvectorStoreErr := pgvector.New(
		context.Background(),
		pgVectorArguments...,
	)
	if pgvectorStoreErr != nil {
		return nil, pgvectorStoreErr
	}

	return &pgvectorStore, nil
}
