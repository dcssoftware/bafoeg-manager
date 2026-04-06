package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
)

func (s *EakteService) GetDocumentFileMetadataByDocumentID(tx *sqlx.Tx, documentID string) (any, customerrors.ErrorInterface) {
	return nil, nil
}
