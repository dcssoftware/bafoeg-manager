package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/eakte/service/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *EakteService) GetEakteByID(tx *sqlx.Tx, id uuid.UUID) (*models.EakteModel, customerrors.ErrorInterface) {
	return s.storage.GetEakteByID(tx, id)
}
