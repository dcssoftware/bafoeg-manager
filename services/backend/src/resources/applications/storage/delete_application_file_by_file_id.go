package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	storageModels "github.com/dcssoftware/bafoeg-manager/src/resources/applications/storage/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *ApplicationsStorage) DeleteApplicationFileByFileID(tx *sqlx.Tx, fileID string) (uuid.UUID, customerrors.ErrorInterface) {
	var result storageModels.IDModel
	var row *sqlx.Row

	sqlquery := `DELETE FROM application_files WHERE id = $1 RETURNING application_id;`

	if tx != nil {
		row = tx.QueryRowx(
			sqlquery,
			fileID,
		)
	} else {
		row = s.db.QueryRowx(
			sqlquery,
			fileID,
		)
	}

	err := row.StructScan(&result)

	sqlErrorData := customerrors.SQLData{}
	if err != nil {
		switch err {

		case sql.ErrNoRows:
			return uuid.Nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return uuid.Nil, customerrors.NewDatabaseError(err, "", "Cannot get application by id", sqlquery, sqlErrorData)
		}
	}

	return result.ID, nil

}
