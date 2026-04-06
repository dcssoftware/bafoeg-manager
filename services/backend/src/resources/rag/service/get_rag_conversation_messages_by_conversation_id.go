package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/rag/service/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *RAGService) GetRagConversationMessagesByConversationID(tx *sqlx.Tx, page uint, conversationID uuid.UUID) ([]models.ConversationMessage, customerrors.ErrorInterface) {
	return s.storage.GetRagConversationMessagesByConversationID(tx, page, conversationID)
}
