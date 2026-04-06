package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *EakteService) InsertEakteVorgang(tx *sqlx.Tx, akteID uuid.UUID, vorgangszeichen string) (uuid.UUID, customerrors.ErrorInterface) {
	return s.storage.InsertEakteVorgang(tx, akteID, vorgangszeichen)
}
