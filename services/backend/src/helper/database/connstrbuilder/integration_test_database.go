package connstrbuilder

import (
	"fmt"
	"net"
	"strconv"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/integration-test-setup/database"
)

func GetSQLConnectionString() string {
	return getSQLConnectionString(configuration.Database.Database)
}

func GetSQLConnectionStringIntegrationTest(testHash uint64) string {
	databaseName := database.CreateintegrationDatabaseName(testHash)
	return getSQLConnectionString(databaseName)
}

func getSQLConnectionString(databaseName string) string {
	sslMode := "disable"
	if configuration.Database.Ssl {
		sslMode = "require"
	}

	addr := net.JoinHostPort(
		configuration.Database.Addr,
		strconv.Itoa(configuration.Database.Port),
	)

	return fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s",
		configuration.Database.Username,
		configuration.Database.Password,
		addr,
		databaseName,
		sslMode,
	)
}
