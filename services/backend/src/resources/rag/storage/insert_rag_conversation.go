package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	storageModels "github.com/dcssoftware/bafoeg-manager/src/resources/applications/storage/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *RAGStorage) InsertRagConversation(tx *sqlx.Tx, userID uuid.UUID, bafoegType string) (uuid.UUID, customerrors.ErrorInterface) {

	sqlquery := `INSERT INTO rag_conversations (user_id, bafoeg_type) VALUES ($1, $2) RETURNING id;`

	var result storageModels.IDModel
	var row *sqlx.Row

	if tx != nil {
		row = tx.QueryRowx(sqlquery, userID, bafoegType)
	} else {
		row = s.db.QueryRowx(sqlquery, userID, bafoegType)
	}

	err := row.StructScan(&result)

	sqlErrorData := customerrors.SQLData{
		"user_id":     userID,
		"bafoeg_type": bafoegType,
	}
	if err != nil {
		switch err {

		case sql.ErrNoRows:
			return uuid.Nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return uuid.Nil, customerrors.NewDatabaseError(err, "", "Cannot create rag conversation", sqlquery, sqlErrorData)
		}
	}

	return result.ID, nil
}
