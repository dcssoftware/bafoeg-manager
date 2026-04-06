package database

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/database/connstrbuilder"
	"github.com/go-sqlx/sqlx"
	_ "github.com/lib/pq"
)

func CreateDatabaseConnectionByConfig() (*sqlx.DB, error) {
	connStr := connstrbuilder.GetSQLConnectionString()
	return createDatabaseConnection(connStr)
}

func CreateDatabaseConnectionByConnectionString(connStr string) (*sqlx.DB, error) {
	return createDatabaseConnection(connStr)
}

func CreateDatabaseConnectionByConfigWithIntegrationTestID(testHash uint64) (*sqlx.DB, error) {
	connStr := connstrbuilder.GetSQLConnectionStringIntegrationTest(testHash)
	return createDatabaseConnection(connStr)
}

func createDatabaseConnection(connStr string) (*sqlx.DB, error) {
	db, dbErr := sqlx.Connect("postgres", connStr)
	if dbErr != nil {

		return nil, dbErr
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}

	return db, nil
}
