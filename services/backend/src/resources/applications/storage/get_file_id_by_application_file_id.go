package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	storageModels "github.com/dcssoftware/bafoeg-manager/src/resources/applications/storage/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *ApplicationsStorage) GetFileIDByApplicationFileID(tx *sqlx.Tx, applicationFileID string) (uuid.UUID, customerrors.ErrorInterface) {

	var result storageModels.FileIDModel
	var row *sqlx.Row

	sqlquery := `
	SELECT file_id FROM application_files
	WHERE id = $1 LIMIT 1;
	`

	if tx != nil {
		row = tx.QueryRowx(
			sqlquery,
			applicationFileID,
		)
	} else {
		row = s.db.QueryRowx(
			sqlquery,
			applicationFileID,
		)
	}

	err := row.StructScan(&result)

	sqlErrorData := customerrors.SQLData{}
	if err != nil {
		switch err {

		case sql.ErrNoRows:
			return uuid.Nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return uuid.Nil, customerrors.NewDatabaseError(err, "", "Cannot get file id by application file id", sqlquery, sqlErrorData)
		}
	}

	return result.ID, nil
}
