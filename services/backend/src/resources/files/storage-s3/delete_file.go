package storages3

import (
	"context"
	"fmt"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/minio/minio-go/v7"
)

func (s *FilesStorageS3) DeleteApplicationFileByFileID(applicationID, fileID string) customerrors.ErrorInterface {
	err := s.s3.RemoveObject(
		context.Background(),
		configuration.S3StoragePaths.ApplicationDataBucket,
		fmt.Sprintf("/applicant-data/uploaded-documents/%s", fileID),
		minio.RemoveObjectOptions{},
	)

	if err != nil {
		return customerrors.NewS3BucketNoRemoveError(err, "")
	}

	return nil
}
