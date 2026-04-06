package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/rag/service/models"
	"github.com/dcssoftware/bafoeg-manager/src/resources/rag/storage/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *RAGStorage) GetRagConversationByID(tx *sqlx.Tx, conversationID uuid.UUID) (*serviceModel.Conversation, customerrors.ErrorInterface) {
	sqlquerybuilder := squirrel.Select("*").
		From("rag_conversations").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"id": conversationID,
		})

	sqlquery, sqlArgs, sqlErr := sqlquerybuilder.ToSql()
	if sqlErr != nil {
		return nil, customerrors.NewDatabaseError(sqlErr, "", "cannot build sql query", "", nil)
	}

	var row *sqlx.Row
	var err error

	if tx != nil {
		row = tx.QueryRowx(sqlquery, sqlArgs...)
	} else {
		row = s.db.QueryRowx(sqlquery, sqlArgs...)
	}

	var conversationModel models.Conversation
	err = row.StructScan(&conversationModel)

	sqlErrorData := customerrors.SQLData{
		"conversation_id": conversationID,
	}

	if err != nil {

		switch err {

		case sql.ErrNoRows:
			break

		default:
			return nil, customerrors.NewDatabaseError(err, "", "Cannot get conversation by conversation id", sqlquery, sqlErrorData)
		}
	}

	return conversationModel.ToServiceConversationModel(), nil
}
