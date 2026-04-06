package service

import (
	"archive/zip"
	"bytes"
	"io"
	"strings"

	calmavsdk "github.com/dcssoftware/bafoeg-manager/src/helper/calmav-sdk"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/helper/eakte/filename"
	"github.com/dcssoftware/bafoeg-manager/src/helper/files/metadata"
	stateModels "github.com/dcssoftware/bafoeg-manager/src/resources/eakte/states/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *EakteService) UploadEakte(tx *sqlx.Tx, fileContent []byte, fileName string) (uuid.UUID, customerrors.ErrorInterface) {

	var txStarted bool
	if tx == nil {
		var txErr error
		tx, txErr = s.storage.StartTx()
		if txErr != nil {
			return uuid.Nil, customerrors.NewDatabaseTransactionBeginError(txErr, "")
		}
		txStarted = true
	}

	isInfected, isInfectedErr := calmavsdk.ScanFile(fileContent)
	if isInfectedErr != nil {
		return uuid.Nil, isInfectedErr
	}

	if isInfected {
		return uuid.Nil, customerrors.NewVirusScannerIncidentError()
	}

	fileSize, mimeType, fileHash, fileMetadataErr := metadata.GetFileMetaData(fileContent)
	if fileMetadataErr != nil {
		return uuid.Nil, fileMetadataErr
	}

	exists, existsErr := s.storage.ExistsUploadedEakteFileHash(
		tx,
		mimeType,
		fileSize,
		fileHash,
	)
	if existsErr != nil {
		return uuid.Nil, existsErr
	}

	if exists {
		return uuid.Nil, customerrors.NewDatabaseEntryAlreadyExistsErr()
	}

	reader := bytes.NewReader(fileContent)
	zipReader, zipErr := zip.NewReader(reader, int64(len(fileContent)))
	if zipErr != nil {
		tx.Rollback()
		return uuid.Nil, customerrors.NewInternalServerError(zipErr, "Failed to read ZIP file", "")
	}

	_, transportFileXMLErr := GetTransportFile(zipReader)
	if transportFileXMLErr != nil {
		tx.Rollback()
		return uuid.Nil, transportFileXMLErr
	}

	fileID, _, fileErr := s.fileService.InsertFileEakte(tx, fileName, fileContent)
	if fileErr != nil {
		tx.Rollback()
		return uuid.Nil, fileErr
	}

	// mapping to eakte / vorgang / or create new
	akteID, akteErr := s.InsertEakteAkte(tx, "AKTENZEICHEN-12345", "ANTRAG", stateModels.VertraulichkeitVertraulich)
	if akteErr != nil {
		tx.Rollback()
		return uuid.Nil, akteErr
	}

	vorgangsID, vorgangsErr := s.InsertEakteVorgang(tx, akteID, "VORGANGSZEICHEN-12345")
	if vorgangsErr != nil {
		tx.Rollback()
		return uuid.Nil, vorgangsErr
	}

	_, documentErr := s.InsertEakteDocument(tx, vorgangsID, fileID, "SHAREPOINT", true)
	if documentErr != nil {
		tx.Rollback()
		return uuid.Nil, documentErr
	}

	// loop for every file in zip
	for _, zipFile := range zipReader.File {

		if zipFile.FileInfo().IsDir() || strings.HasPrefix(zipFile.FileInfo().Name(), ".") || zipFile.FileInfo().Size() == 0 {
			continue
		}

		fileReader, openErr := zipFile.Open()
		if openErr != nil {
			tx.Rollback()
			return uuid.Nil, customerrors.NewInternalServerError(openErr, "Failed to open file in ZIP", "")
		}

		uploadErr := s.UploadEakteFile(tx, zipFile.FileInfo().Name(), vorgangsID, fileReader)
		if uploadErr != nil {
			tx.Rollback()
			return uuid.Nil, uploadErr
		}
	}

	if txStarted {
		commitErr := tx.Commit()
		if commitErr != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				return uuid.Nil, customerrors.NewInternalServerError(rollbackErr, "", "")
			}
			return uuid.Nil, customerrors.NewDatabaseTransactionCommitError(commitErr, "")
		}
	}

	return akteID, nil
}

func GetTransportFile(zipReader *zip.Reader) ([]byte, customerrors.ErrorInterface) {
	for _, zipFile := range zipReader.File {

		if zipFile.FileInfo().IsDir() || strings.HasPrefix(zipFile.FileInfo().Name(), ".") || zipFile.FileInfo().Size() == 0 {
			continue
		}

		fileReader, openErr := zipFile.Open()
		if openErr != nil {
			return nil, customerrors.NewInternalServerError(openErr, "Failed to open file in ZIP", "")
		}

		_, transportFileStructureErr := filename.ParseTransportDateiname(zipFile.FileInfo().Name())
		if transportFileStructureErr == nil {
			transportFileXMLContent, transportFileXMLErr := io.ReadAll(fileReader)
			if transportFileXMLErr != nil {
				return nil, customerrors.NewInternalServerError(transportFileXMLErr, "", "")
			}

			return transportFileXMLContent, nil
		}
	}
	return nil, nil
}
