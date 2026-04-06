package filehash

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
)

func GenerateFileHash(fileContent []byte) (string, customerrors.ErrorInterface) {
	hash := md5.New()
	if _, err := io.Copy(hash, bytes.NewBuffer(fileContent)); err != nil {
		return "", customerrors.NewInternalServerError(err, "", "could not generate file hash")
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
