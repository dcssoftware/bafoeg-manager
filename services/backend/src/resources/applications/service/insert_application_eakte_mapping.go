package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *ApplicationsService) InsertApplicationEakteMapping(tx *sqlx.Tx, applicationID, eakteAkteID uuid.UUID) (uuid.UUID, customerrors.ErrorInterface) {
	return s.storage.InsertApplicationEakteMapping(tx, applicationID, eakteAkteID)
}
