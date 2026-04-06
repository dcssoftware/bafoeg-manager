package service

import (
	calmavsdk "github.com/dcssoftware/bafoeg-manager/src/helper/calmav-sdk"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	custombadrequestconstraints "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/bad-request-constraints"
	"github.com/dcssoftware/bafoeg-manager/src/helper/files/metadata"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/states"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicationsService) UploadApplicationFile(tx *sqlx.Tx, applicationID string, filename string, fileContent []byte) customerrors.ErrorInterface {

	var commitTx bool
	if tx == nil {
		var txErr error
		tx, txErr = s.storage.StartTx()
		if txErr != nil {
			return customerrors.NewInternalServerError(txErr, "", "cannot start a new database transaction")
		}
		commitTx = true
		defer tx.Rollback()
	}

	application, applicationErr := s.GetApplicationByID(tx, applicationID)
	if applicationErr != nil {
		return applicationErr
	}

	statusIsDone, statusIsDoneErr := states.IsDoneString(application.Status.Identifier)
	if statusIsDoneErr != nil || statusIsDone {
		return customerrors.NewBadRequestError(custombadrequestconstraints.BadRequest_ApplicationStatusChangeInvalid)
	}

	isInfected, isInfectedErr := calmavsdk.ScanFile(fileContent)
	if isInfectedErr != nil {
		return isInfectedErr
	}

	if isInfected {
		return customerrors.NewVirusScannerIncidentError()
	}

	fileSize, mimeType, fileHash, fileMetadataErr := metadata.GetFileMetaData(fileContent)
	if fileMetadataErr != nil {
		return fileMetadataErr
	}

	exists, existsErr := s.storage.ExistsUploadedApplicationFileHash(
		tx,
		application.ID.String(),
		mimeType,
		fileSize,
		fileHash,
	)
	if existsErr != nil {
		return existsErr
	}

	if exists {
		return customerrors.NewDatabaseEntryAlreadyExistsErr()
	}

	fileID, _, fileUploadErr := s.filesService.InsertFileApplication(tx, filename, fileContent)
	if fileUploadErr != nil {
		return fileUploadErr
	}

	_, applicationFileErr := s.storage.UploadApplicationFile(
		tx,
		application.ID.String(),
		fileID,
	)
	if applicationFileErr != nil {
		return applicationFileErr
	}

	if commitTx {
		err := tx.Commit()
		if err != nil {
			return customerrors.NewInternalServerError(err, "", "cannot commit database transaction")
		}
	}

	return nil
}
