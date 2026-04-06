package storages3

import "github.com/minio/minio-go/v7"

type RAGStorageS3 struct {
	s3 *minio.Client
}

func NewRAGStorageS3(s3 *minio.Client) *RAGStorageS3 {
	return &RAGStorageS3{
		s3: s3,
	}
}
