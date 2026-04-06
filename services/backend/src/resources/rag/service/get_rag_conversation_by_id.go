package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/rag/service/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *RAGService) GetRagConversationByID(tx *sqlx.Tx, conversationID uuid.UUID) (*models.Conversation, customerrors.ErrorInterface) {
	return s.storage.GetRagConversationByID(tx, conversationID)
}
