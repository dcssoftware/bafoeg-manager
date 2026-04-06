package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *RAGService) StartRagConversationStudierende(tx *sqlx.Tx, userID uuid.UUID) (uuid.UUID, customerrors.ErrorInterface) {
	return s.storage.InsertRagConversation(tx, userID, "STUDIERENDENBAFOEG")
}
