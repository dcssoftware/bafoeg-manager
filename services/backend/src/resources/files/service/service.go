package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

type FileService struct {
	storage   FileStorage
	storageS3 FileStorageS3
}

func NewFileService(storage FileStorage, storageS3 FileStorageS3) *FileService {
	return &FileService{
		storage:   storage,
		storageS3: storageS3,
	}
}

type FileStorage interface {
	StartTx() (*sqlx.Tx, error)
	InsertFile(tx *sqlx.Tx, fileName, fileType string, fileSize uint, fileHash string) (uuid.UUID, customerrors.ErrorInterface)
}

type FileStorageS3 interface {
	UploadFile(bucketname, path, filename string, fileContent []byte) (*minio.UploadInfo, customerrors.ErrorInterface)
	DownloadFile(bucketname, path, filename string) ([]byte, *minio.ObjectInfo, customerrors.ErrorInterface)
}
