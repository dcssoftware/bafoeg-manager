package storages3

import (
	"github.com/minio/minio-go/v7"
)

type CronjobRagVectorProcessorStorageS3 struct {
	s3 *minio.Client
}

func NewCronjobRagVectorProcessorStorageS3(s3 *minio.Client) *CronjobRagVectorProcessorStorageS3 {
	return &CronjobRagVectorProcessorStorageS3{
		s3: s3,
	}
}
