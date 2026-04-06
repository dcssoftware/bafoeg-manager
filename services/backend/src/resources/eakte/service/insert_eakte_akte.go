package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	stateModels "github.com/dcssoftware/bafoeg-manager/src/resources/eakte/states/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *EakteService) InsertEakteAkte(tx *sqlx.Tx, aktenzeichen string, typIdentifier string, vertraulichkeit stateModels.VertraulichkeitEnum) (uuid.UUID, customerrors.ErrorInterface) {
	return s.storage.InsertEakteAkte(tx, aktenzeichen, typIdentifier, vertraulichkeit)
}
