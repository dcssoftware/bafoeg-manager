package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *RAGService) InsertRagConversationMessage(tx *sqlx.Tx, conversationID uuid.UUID, messageContent string, isUserInputMessage bool) (uuid.UUID, customerrors.ErrorInterface) {
	return s.storage.InsertRagConversationMessage(tx, conversationID, messageContent, isUserInputMessage)
}
