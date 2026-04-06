package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/eakte/service/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *EakteService) GetDocumentFileByID(tx *sqlx.Tx, documentFileID uuid.UUID) (*models.EaktenFileModel, []byte, customerrors.ErrorInterface) {

	fileModel, fileModelErr := s.storage.GetFileByFileID(tx, documentFileID)
	if fileModelErr != nil {
		return nil, nil, fileModelErr
	}

	fileContent, _, fileContentErr := s.fileService.DownloadFileEakte(tx, fileModel.FileID)
	if fileContentErr != nil {
		return nil, nil, fileContentErr
	}

	return fileModel, fileContent, nil
}
