package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
	"github.com/minio/minio-go/v7"
)

func (s *FileService) DownloadFileApplication(tx *sqlx.Tx, fileID string) ([]byte, *minio.ObjectInfo, customerrors.ErrorInterface) {
	return s.storageS3.DownloadFile(
		configuration.S3StoragePaths.ApplicationDataBucket,
		"application-files/uploaded",
		fileID,
	)
}
