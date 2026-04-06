package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	storageModel "github.com/dcssoftware/bafoeg-manager/src/resources/applications/storage/models"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicationsStorage) GetApplicationsMetricsInProgress(tx *sqlx.Tx) (uint, customerrors.ErrorInterface) {
	sqlquery := `
		SELECT COUNT(*) AS "count" FROM applications
		WHERE status IN (
			SELECT id FROM application_status
			WHERE identifier = 'IN_PROGRESS'
		)
	`

	var model storageModel.CountModel
	var err error

	if tx != nil {
		err = tx.QueryRowx(sqlquery).StructScan(&model)
	} else {
		err = s.db.QueryRowx(sqlquery).StructScan(&model)
	}

	if err != nil {
		data := customerrors.SQLData{}

		switch err {

		case sql.ErrNoRows:
			return 0, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, data)

		default:
			return 0, customerrors.NewDatabaseError(err, "", "Cannot get application count", sqlquery, data)
		}
	}

	return model.Count, nil
}
