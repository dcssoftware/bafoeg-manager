package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

func (s *FileService) DownloadFileEakte(tx *sqlx.Tx, fileID uuid.UUID) ([]byte, *minio.ObjectInfo, customerrors.ErrorInterface) {
	return s.storageS3.DownloadFile(
		configuration.S3StoragePaths.EAktenDataBucket,
		"/eakte-files/uploaded/",
		fileID.String(),
	)
}
