package service

import (
	"fmt"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
)

func (s *FileService) DownloadFileProfilePicture(userID string) ([]byte, customerrors.ErrorInterface) {

	content, _, err := s.storageS3.DownloadFile(
		configuration.S3StoragePaths.ProfilePictureBucket,
		fmt.Sprintf("/profile-pictures/%s/", userID),
		"profilepicture.webp",
	)
	if err != nil {
		return nil, err
	}

	return content, nil
}
