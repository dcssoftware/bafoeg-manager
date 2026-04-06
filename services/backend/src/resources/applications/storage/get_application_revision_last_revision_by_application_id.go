package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	storageModel "github.com/dcssoftware/bafoeg-manager/src/resources/applications/storage/models"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicationsStorage) GetApplicationRevisionLatestRevisionByApplicationID(tx *sqlx.Tx, applicationID string) (*serviceModel.ApplicationRevisionModel, customerrors.ErrorInterface) {
	sqlquery := `SELECT * FROM application_revisions WHERE application_id = $1 ORDER BY created`

	var model storageModel.ApplicationRevisionModel
	var err error

	if tx != nil {
		err = tx.QueryRowx(sqlquery, applicationID).StructScan(&model)
	} else {
		err = s.db.QueryRowx(sqlquery, applicationID).StructScan(&model)
	}

	if err != nil {
		data := customerrors.SQLData{}

		switch err {

		case sql.ErrNoRows:
			return nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, data)

		default:
			return nil, customerrors.NewDatabaseError(err, "", "Cannot get application count", sqlquery, data)
		}
	}

	return nil, nil
}
