package filesize

import (
	"github.com/ccoveille/go-safecast/v2"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
)

func GetFileSite(fileContent []byte) (uint, customerrors.ErrorInterface) {
	fileSizeInt := len(fileContent)

	fileSize, fileSizeErr := safecast.Convert[uint](fileSizeInt)
	if fileSizeErr != nil {
		return 0, customerrors.NewInternalServerError(fileSizeErr, "", "")
	}

	return fileSize, nil
}
