package integrationtestsetup

import (
	"testing"

	"github.com/dcssoftware/bafoeg-manager/src/helper/integration-test-setup/database"
	"github.com/go-sqlx/sqlx"
	"github.com/stretchr/testify/assert"
)

func CreateIntegrationDatabase(t *testing.T, db *sqlx.DB, testHash uint64) string {
	var databaseName string = database.CreateintegrationDatabaseName(testHash)
	_, err := db.Exec("CREATE DATABASE $1 TEMPLATE postgres;", databaseName)
	assert.NoError(t, err, "could not create unit test template database")

	return databaseName
}
