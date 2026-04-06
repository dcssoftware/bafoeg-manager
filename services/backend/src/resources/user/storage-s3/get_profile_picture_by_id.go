package storages3

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/minio/minio-go/v7"
)

func (s *UserStoreS3) GetProfilePictureByID(id string) ([]byte, customerrors.ErrorInterface) {
	obj, objErr := s.s3.GetObject(
		context.Background(),
		configuration.S3StoragePaths.ProfilePictureBucket,
		fmt.Sprintf("/profile-pictures/%s/profilepicture.jpg", id),
		minio.GetObjectOptions{},
	)
	if objErr != nil {
		log.Fatalln(objErr)
	}

	content, contentErr := io.ReadAll(obj)
	if contentErr != nil {
		log.Fatalln(contentErr)
	}

	return content, nil
}
