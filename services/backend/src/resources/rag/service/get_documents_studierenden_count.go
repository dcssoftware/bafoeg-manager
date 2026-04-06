package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
)

func (s *RAGService) GetDocumentsStudierendenCount(tx *sqlx.Tx, filterResult string) (uint, customerrors.ErrorInterface) {
	return s.storage.GetDocumentsStudierendenCount(tx, filterResult)
}
