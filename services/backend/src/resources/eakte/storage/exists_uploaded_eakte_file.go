package storage

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
)

func (s *EakteStorage) ExistsUploadedEakteFileHash(tx *sqlx.Tx, filetype string, fileSize uint, fileHash string) (bool, customerrors.ErrorInterface) {
	return false, nil
}
