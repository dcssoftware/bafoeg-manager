package migrator

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"net/http"

	dbHelper "github.com/dcssoftware/bafoeg-manager/src/helper/database"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	_ "github.com/lib/pq"
)

//go:embed migration-files-app
var migrationAppDir embed.FS

type Migrator struct {
	db *sql.DB
}

func NewMigrator() *Migrator {
	dbConn, dbConnErr := dbHelper.CreateDatabaseConnectionByConfig()
	if dbConn == nil || dbConn.DB == nil || dbConnErr != nil {
		var errMsg = errors.New("could not connect to database")
		panic(fmt.Errorf("%w", errMsg))
	}
	return &Migrator{
		db: dbConn.DB,
	}
}

func (m *Migrator) MigrateUp() error {
	return m.migrateDatabase()
}

func (m *Migrator) migrateDatabase() error {
	driver, err := postgres.WithInstance(m.db, new(postgres.Config))
	if err != nil {
		return err
	}

	// useful for later when golang migrator implements named migrations for testing purposes
	migrationDirectory := migrationAppDir

	sourceInstance, err := httpfs.New(http.FS(migrationDirectory), "migration-files-app")
	if err != nil {
		return err
	}

	migrator, migratorErr := migrate.NewWithInstance("httpfs", sourceInstance, "postgres", driver)
	if migratorErr != nil {
		return migratorErr
	}

	err = migrator.Up()
	return err
}
