package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/rag/service/models"
	"github.com/go-sqlx/sqlx"
)

func (s *RAGService) GetDocumentSchülerByID(tx *sqlx.Tx, id string) ([]byte, *models.DocumentModel, customerrors.ErrorInterface) {
	documentModel, documentModelErr := s.storage.GetDocumentSchülerByID(tx, id)
	if documentModelErr != nil {
		return nil, nil, documentModelErr
	}

	document, documentErr := s.storageS3.DownloadDocumentSchülerByID(documentModel.ID.String())
	if documentErr != nil {
		return nil, nil, documentErr
	}

	return document, documentModel, nil
}
