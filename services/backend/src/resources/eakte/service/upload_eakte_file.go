package service

import (
	"io"

	calmavsdk "github.com/dcssoftware/bafoeg-manager/src/helper/calmav-sdk"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *EakteService) UploadEakteFile(tx *sqlx.Tx, filename string, vorgangID uuid.UUID, fileReader io.ReadCloser) customerrors.ErrorInterface {

	content, readErr := io.ReadAll(fileReader)
	fileReader.Close()
	if readErr != nil {
		return customerrors.NewInternalServerError(readErr, "Failed to read file content", "")
	}

	isInfected, isInfectedErr := calmavsdk.ScanFile(content)
	if isInfectedErr != nil {
		return isInfectedErr
	}

	if isInfected {
		return customerrors.NewVirusScannerIncidentError()
	}

	fileID, _, fileErr := s.fileService.InsertFileEakte(tx, filename, content)
	if fileErr != nil {
		return fileErr
	}

	_, documentErr := s.InsertEakteDocument(tx, vorgangID, fileID, "SHAREPOINT", false)
	if documentErr != nil {
		return documentErr
	}

	return nil
}
