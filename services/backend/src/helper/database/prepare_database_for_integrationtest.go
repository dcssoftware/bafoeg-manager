package database

import (
	"fmt"
	"testing"

	"github.com/dcssoftware/bafoeg-manager/src/helper/integration-test-setup/database"
	"github.com/stretchr/testify/assert"
)

func PrepareDatabaseForIntegrationTest(t *testing.T, testHash uint64) {
	setupDB, setupDBerr := CreateDatabaseConnectionByConfig()
	assert.NoError(t, setupDBerr, "could not create database connection to prepare integration test")

	if setupDB != nil {
		defer setupDB.Close()
	} else {
		t.Fatal("database connection is nil")
	}

	testDatabaseName := database.CreateintegrationDatabaseName(testHash)
	query := fmt.Sprintf("CREATE DATABASE %s WITH TEMPLATE unittest_template", testDatabaseName)
	_, err := setupDB.Exec(query)
	assert.NoError(t, err, "could not create new database")
}
