package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	storageModels "github.com/dcssoftware/bafoeg-manager/src/resources/applications/storage/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *RAGStorage) InsertDocumentStudierende(tx *sqlx.Tx, fileName, fileType string, fileSize uint, fileHash string, createdFromUserID string) (uuid.UUID, customerrors.ErrorInterface) {
	var result storageModels.IDModel
	var row *sqlx.Row

	sqlquery := `
		INSERT INTO pgvector_rag_studierendenbafoeg_files(
			file_name,
			file_type,
			file_size,
			file_hash,
			status,
			processed_error,
			created_from
		) VALUES 
		 ($1,$2,$3,$4, (SELECT id FROM rag_document_process_status WHERE identifier = 'IN_PROGRESS'), '', $5)
		 RETURNING id
	`

	if tx != nil {
		row = tx.QueryRowx(
			sqlquery,
			fileName,
			fileType,
			fileSize,
			fileHash,
			createdFromUserID,
		)
	} else {
		row = s.db.QueryRowx(
			sqlquery,
			fileName,
			fileType,
			fileSize,
			fileHash,
			createdFromUserID,
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
