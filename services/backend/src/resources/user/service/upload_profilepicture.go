package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/static/assets"
	"github.com/go-sqlx/sqlx"
)

func (s *UserService) UploadProfilePictureByID(tx *sqlx.Tx, userID string, pictureData []byte) customerrors.ErrorInterface {

	var txStarted bool
	if tx == nil {
		var txErr error
		tx, txErr = s.storage.StartTransaction()
		if txErr != nil {
			return customerrors.NewDatabaseTransactionBeginError(txErr, "")
		}
		txStarted = true
	}

	if len(pictureData) == 0 {
		pictureData = assets.DefaultProfilePicture
	} else {
		_ = assets.DefaultProfilePicture
	}

	_, _, err := s.fileService.InsertFileProfilePicture(nil, userID, pictureData)
	if err != nil {
		tx.Rollback()
		return err
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
