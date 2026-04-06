package storages3

import (
	"github.com/minio/minio-go/v7"
)

type FilesStorageS3 struct {
	s3 *minio.Client
}

func NewFilesStorageS3(s3 *minio.Client) *FilesStorageS3 {
	return &FilesStorageS3{
		s3: s3,
	}
}
