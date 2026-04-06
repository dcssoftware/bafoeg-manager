package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/rag/service/models"
	"github.com/go-sqlx/sqlx"
)

func (s *RAGService) GetDocumentsStudierenden(tx *sqlx.Tx, page uint, filterResult string) ([]models.DocumentModel, customerrors.ErrorInterface) {
	return s.storage.GetDocumentsStudierenden(tx, page, filterResult)
}
