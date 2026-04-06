package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *EakteService) InsertEakteDocument(tx *sqlx.Tx, vorgangID uuid.UUID, fileID uuid.UUID, sourceIdentidier string, isXdomeaZipFile bool) (uuid.UUID, customerrors.ErrorInterface) {
	return s.storage.InsertEakteDocument(tx, vorgangID, fileID, sourceIdentidier, isXdomeaZipFile)
}
