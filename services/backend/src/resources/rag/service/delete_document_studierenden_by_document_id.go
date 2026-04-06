package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *RAGService) DeleteRagDocumentStudierenden(tx *sqlx.Tx, documentID uuid.UUID) customerrors.ErrorInterface {

	var txStarted bool
	if tx == nil {
		var txErr error
		tx, txErr = s.storage.StartTx()
		if txErr != nil {
			return customerrors.NewDatabaseTransactionBeginError(txErr, "")
		}
		txStarted = true
	}

	fileRefErr := s.storage.DeleteRagFileStudierenden(tx, documentID)
	if fileRefErr != nil {
		return fileRefErr
	}

	// delete file from s3 // todo later, not important now

	vectorErr := s.storage.DeleteRagVectorStudierenden(tx, documentID)
	if vectorErr != nil {
		return vectorErr
	}

	if txStarted {
		commitErr := tx.Commit()
		if commitErr != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				return customerrors.NewInternalServerError(rollbackErr, "", "")
			}
			return customerrors.NewDatabaseTransactionCommitError(commitErr, "")
		}
	}

	return nil
}
