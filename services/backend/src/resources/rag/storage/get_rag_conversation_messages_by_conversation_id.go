package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/rag/service/models"
	"github.com/dcssoftware/bafoeg-manager/src/resources/rag/storage/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *RAGStorage) GetRagConversationMessagesByConversationID(tx *sqlx.Tx, page uint, conversationID uuid.UUID) ([]serviceModel.ConversationMessage, customerrors.ErrorInterface) {
	var limit uint = configuration.Webserver.Display.MaxResponseEntityCount
	offset := (limit * page) - limit

	sqlquerybuilder := squirrel.
		Select(`*`).
		From("rag_conversation_messages_overview").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"conversation_id": conversationID}).
		Limit(uint64(limit)).
		Offset(uint64(offset))

	sqlquery, sqlArgs, sqlErr := sqlquerybuilder.ToSql()
	if sqlErr != nil {
		return nil, customerrors.NewDatabaseError(sqlErr, "", "cannot build sql query", "", nil)
	}

	var rows *sqlx.Rows
	var messages []serviceModel.ConversationMessage
	var err error

	if tx != nil {
		rows, err = tx.Queryx(sqlquery, sqlArgs...)
	} else {
		rows, err = s.db.Queryx(sqlquery, sqlArgs...)
	}

	sqlErrorData := customerrors.SQLData{}

	if err != nil {

		switch err {

		case sql.ErrNoRows:
			return nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return nil, customerrors.NewDatabaseError(err, "", "Cannot get rag conversation messages", sqlquery, sqlErrorData)
		}
	}

	defer rows.Close()

	for rows.Next() {
		var message models.ConversationMessage
		if err := rows.StructScan(&message); err != nil {
			return nil, customerrors.NewDatabaseError(err, "", "Could not read row from database", "sqlquery", sqlErrorData)
		}

		messages = append(messages, *message.ToServiceConversationMessageModel())
	}

	return messages, nil
}
