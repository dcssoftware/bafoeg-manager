package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	custombadrequestconstraints "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/bad-request-constraints"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicationsService) DeleteApplicationFile(tx *sqlx.Tx, applicationID, fileID string) customerrors.ErrorInterface {

	var txStarted bool
	if tx == nil {
		var txErr error
		tx, txErr = s.storage.StartTx()
		if txErr != nil {
			return customerrors.NewDatabaseTransactionBeginError(txErr, "")
		}
		txStarted = true
	}

	file, fileErr := s.storage.GetApplicationFileByFileID(tx, fileID)
	if fileErr != nil {
		return fileErr
	}

	if applicationID != file.ApplicationID.String() {
		return customerrors.NewBadRequestError(custombadrequestconstraints.BadRequest_InformationMismatch)
	}

	_, deletedFileErr := s.storage.DeleteApplicationFileByFileID(tx, fileID)
	if deletedFileErr != nil {
		return deletedFileErr
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

	// if it fails, it can be cleaned up by a daily running cleanup cronjob
	// As long as it's deleted in the database, it's fine 👍🏻
	// s.storageS3.DeleteApplicationFileByFileID(
	// 	deletedFileApplicationID.String(),
	// 	file.ID.String(),
	// )

	return nil
}
