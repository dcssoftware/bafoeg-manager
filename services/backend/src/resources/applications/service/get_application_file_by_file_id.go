package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicationsService) GetApplicationFileByFileID(tx *sqlx.Tx, applicationFileID string) ([]byte, *models.ApplicationFile, customerrors.ErrorInterface) {

	fileID, fileErr := s.storage.GetFileIDByApplicationFileID(tx, applicationFileID)
	if fileErr != nil {
		return nil, nil, fileErr
	}

	applicationFile, applicationFileErr := s.storage.GetApplicationFileByFileID(tx, applicationFileID)
	if applicationFileErr != nil {
		return []byte{}, nil, applicationFileErr
	}

	fileContent, _, fileContentErr := s.filesService.DownloadFileApplication(tx, fileID.String())
	if fileContentErr != nil {
		return nil, nil, fileContentErr
	}

	return fileContent, applicationFile, nil
}

func (s *ApplicationsService) GetApplicationFileByApplicationIDandFileID(applicationID string, applicationFileID string) ([]byte, customerrors.ErrorInterface) {
	fileContent, _, fileContentErr := s.filesService.DownloadFileApplication(nil, applicationFileID)
	if fileContentErr != nil {
		return nil, fileContentErr
	}

	return fileContent, nil
}
