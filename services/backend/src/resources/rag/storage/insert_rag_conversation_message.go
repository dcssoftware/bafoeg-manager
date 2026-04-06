package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	storageModels "github.com/dcssoftware/bafoeg-manager/src/resources/applications/storage/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *RAGStorage) InsertRagConversationMessage(tx *sqlx.Tx, conversationID uuid.UUID, messageContent string, isUserInputMessage bool) (uuid.UUID, customerrors.ErrorInterface) {
	sqlquery := `
		INSERT INTO rag_conversation_messages (conversation_id, message, sender)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	var role string = "SYSTEM"
	if isUserInputMessage {
		role = "USER"
	}

	var result storageModels.IDModel
	var row *sqlx.Row

	if tx != nil {
		row = tx.QueryRowx(sqlquery, conversationID.String(), messageContent, role)
	} else {
		row = s.db.QueryRowx(sqlquery, conversationID.String(), messageContent, role)
	}

	err := row.StructScan(&result)

	sqlErrorData := customerrors.SQLData{
		"conversation_id": conversationID.String(),
		"message":         messageContent,
		"role":            role,
	}

	if err != nil {
		switch err {

		case sql.ErrNoRows:
			return uuid.Nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return uuid.Nil, customerrors.NewDatabaseError(err, "", "Cannot create rag conversation message", sqlquery, sqlErrorData)
		}
	}

	return result.ID, nil
}
