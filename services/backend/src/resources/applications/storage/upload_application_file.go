package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	storageModels "github.com/dcssoftware/bafoeg-manager/src/resources/applications/storage/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *ApplicationsStorage) UploadApplicationFile(tx *sqlx.Tx, applicationID string, fileID uuid.UUID) (uuid.UUID, customerrors.ErrorInterface) {
	var result storageModels.IDModel
	var row *sqlx.Row

	sqlquery := `
		INSERT INTO application_files (
			application_id,
			file_id
		) VALUES 
		 ($1,$2)
		 RETURNING id
	`

	if tx != nil {
		row = tx.QueryRowx(
			sqlquery,
			applicationID,
			fileID.String(),
		)
	} else {
		row = s.db.QueryRowx(
			sqlquery,
			applicationID,
			fileID.String(),
		)
	}

	err := row.StructScan(&result)

	sqlErrorData := customerrors.SQLData{}
	if err != nil {
		switch err {

		case sql.ErrNoRows:
			return uuid.Nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return uuid.Nil, customerrors.NewDatabaseError(err, "", "Cannot insert application file", sqlquery, sqlErrorData)
		}
	}

	return result.ID, nil
}
