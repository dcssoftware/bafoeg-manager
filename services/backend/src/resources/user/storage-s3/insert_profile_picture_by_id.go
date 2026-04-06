package storages3

import (
	"bytes"
	"context"
	"fmt"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/minio/minio-go/v7"
)

func (s *UserStoreS3) InsertProfilePictureByID(userID string, pictureJPG []byte) customerrors.ErrorInterface {
	objectName := fmt.Sprintf("/profile-pictures/%s/profilepicture.jpg", userID)

	_, err := s.s3.PutObject(
		context.Background(),
		configuration.S3StoragePaths.ProfilePictureBucket,
		objectName,
		bytes.NewReader(pictureJPG),
		int64(len(pictureJPG)),
		minio.PutObjectOptions{
			ContentType: "image/jpeg",
		},
	)

	if err != nil {
		return customerrors.NewInternalServerError(err, userID, "Failed to upload profile picture to S3")
	}

	return nil
}
