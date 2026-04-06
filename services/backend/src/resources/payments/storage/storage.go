package storage

import "github.com/go-sqlx/sqlx"

type PaymentStorage struct {
	db *sqlx.DB
}

func NewPaymentStorage(dbConn *sqlx.DB) *PaymentStorage {
	return &PaymentStorage{
		db: dbConn,
	}
}
