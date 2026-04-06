package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/eakte/service/models"
	"github.com/go-sqlx/sqlx"
)

func (s *EakteService) GetEakteApplicationMapping(tx *sqlx.Tx, eakteAkteID string) (*models.EaktenApplicationMappingModel, customerrors.ErrorInterface) {
	return s.storage.GetEakteApplicationMapping(tx, eakteAkteID)
}
