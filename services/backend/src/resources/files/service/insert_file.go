package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *FileService) InsertFile(tx *sqlx.Tx, fileName, fileType string, fileSize uint, fileHash string) (uuid.UUID, customerrors.ErrorInterface) {
	return s.storage.InsertFile(tx, fileName, fileType, fileSize, fileHash)
}
