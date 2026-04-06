package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/helper/files/metadata"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

func (s *FileService) InsertFileApplication(tx *sqlx.Tx, fileName string, fileContent []byte) (uuid.UUID, *minio.UploadInfo, customerrors.ErrorInterface) {
	fileSize, mimeType, fileHash, fileMetadataErr := metadata.GetFileMetaData(fileContent)
	if fileMetadataErr != nil {
		return uuid.Nil, nil, fileMetadataErr
	}

	fileID, fileErr := s.storage.InsertFile(tx, fileName, mimeType, fileSize, fileHash)
	if fileErr != nil {
		return uuid.Nil, nil, fileErr
	}

	uploadInfo, uploadErr := s.storageS3.UploadFile(
		configuration.S3StoragePaths.ApplicationDataBucket,
		"/application-files/uploaded/",
		fileID.String(),
		fileContent,
	)

	return fileID, uploadInfo, uploadErr
}
