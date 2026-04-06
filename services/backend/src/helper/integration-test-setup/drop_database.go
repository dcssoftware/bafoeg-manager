package integrationtestsetup

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/integration-test-setup/database"
	"github.com/go-sqlx/sqlx"
)

func DropIntegrationDatabase(db *sqlx.DB, testHash uint64) {
	var databaseName string = database.CreateintegrationDatabaseName(testHash)
	db.Exec("DROP DATABASE IF EXISTS $1 WITH (FORCE); ;", databaseName)
}
