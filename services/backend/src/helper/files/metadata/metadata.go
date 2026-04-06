package metadata

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/helper/files/filehash"
	"github.com/dcssoftware/bafoeg-manager/src/helper/files/filesize"
	"github.com/dcssoftware/bafoeg-manager/src/helper/files/mimetype"
)

func GetFileMetaData(fileContent []byte) (fileSize uint, mimeType string, fileHash string, err customerrors.ErrorInterface) {
	fileSize, err = filesize.GetFileSite(fileContent)
	if err != nil {
		return
	}

	fileHash, err = filehash.GenerateFileHash(fileContent)
	if err != nil {
		return
	}

	mimeType, err = mimetype.DetectMineType(fileContent)
	if err != nil {
		return
	}

	return
}
