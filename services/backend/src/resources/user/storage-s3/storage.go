package storages3

import "github.com/minio/minio-go/v7"

type UserStoreS3 struct {
	s3 *minio.Client
}

func NewUserStoreS3(s3 *minio.Client) *UserStoreS3 {
	return &UserStoreS3{
		s3,
	}
}
