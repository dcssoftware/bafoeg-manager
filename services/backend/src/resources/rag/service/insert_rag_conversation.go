package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/rag/service/consts"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *RAGService) InsertRagConversation(tx *sqlx.Tx, userID uuid.UUID, bafoegType consts.BafögType) (uuid.UUID, customerrors.ErrorInterface) {
	return s.storage.InsertRagConversation(tx, userID, bafoegType.String())
}
