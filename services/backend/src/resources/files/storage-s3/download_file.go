package storages3

import (
	"context"
	"io"
	"path"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/minio/minio-go/v7"
)

func (s *FilesStorageS3) DownloadFile(bucketname, filepath, filename string) ([]byte, *minio.ObjectInfo, customerrors.ErrorInterface) {

	fileDestination := path.Join(filepath, filename)

	obj, objErr := s.s3.GetObject(
		context.Background(),
		bucketname,
		fileDestination,
		minio.GetObjectOptions{},
	)
	if objErr != nil {
		return nil, nil, customerrors.NewS3BucketDownloadFileNotFoundError(objErr, "cannot read file from "+bucketname+" :: "+fileDestination)
	}

	defer obj.Close()

	// fileInfo, fileInfoErr := obj.Stat()
	// if fileInfoErr != nil {
	// 	return nil, nil, customerrors.NewS3BucketDownloadFailedError(fileInfoErr, "cannot read filedata from "+bucketname+" :: "+fileDestination)
	// }

	content, contentErr := io.ReadAll(obj)
	if contentErr != nil {
		return nil, nil, customerrors.NewS3BucketDownloadFailedError(contentErr, "cannot read file "+bucketname+" :: "+fileDestination)
	}

	return content, nil, nil
}
