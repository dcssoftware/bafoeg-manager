package storages3

import (
	"bytes"
	"context"
	"path"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/minio/minio-go/v7"
)

func (s *FilesStorageS3) UploadFile(bucketname, filepath, filename string, fileContent []byte) (*minio.UploadInfo, customerrors.ErrorInterface) {

	reader := bytes.NewReader(fileContent)

	fileDestination := path.Join(filepath, filename)

	obj, objErr := s.s3.PutObject(
		context.Background(),
		bucketname,
		fileDestination,
		reader,
		int64(len(fileContent)),
		minio.PutObjectOptions{},
	)
	if objErr != nil {
		return nil, customerrors.NewS3BucketUploadFailedError(objErr, "cannot upload file")
	}

	return &obj, nil
}
