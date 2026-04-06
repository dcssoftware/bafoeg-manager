package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	storageModels "github.com/dcssoftware/bafoeg-manager/src/resources/applications/storage/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *FileStorage) InsertFile(tx *sqlx.Tx, fileName, fileType string, fileSize uint, fileHash string) (uuid.UUID, customerrors.ErrorInterface) {
	var result storageModels.IDModel
	var row *sqlx.Row

	sqlquery := `
		INSERT INTO files (
			file_name,
			file_type,
			file_size,
			file_hash
		) VALUES 
		 ($1,$2,$3,$4)
		 RETURNING id
	`

	if fileType == "" {
		fileType = "plain/text"
	}

	if tx != nil {
		row = tx.QueryRowx(
			sqlquery,
			fileName,
			fileType,
			fileSize,
			fileHash,
		)
	} else {
		row = s.db.QueryRowx(
			sqlquery,
			fileName,
			fileType,
			fileSize,
			fileHash,
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
