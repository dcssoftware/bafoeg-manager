package mimetype

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/h2non/filetype"
)

func DetectMineType(fileContent []byte) (string, customerrors.ErrorInterface) {
	kind, kindErr := filetype.Match(fileContent)
	if kindErr != nil {
		return "", customerrors.NewInternalServerError(kindErr, "", "cannot detect mimetype of file")
	}

	return kind.MIME.Value, nil
}
