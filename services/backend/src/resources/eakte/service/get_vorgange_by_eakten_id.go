package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/eakte/service/models"
	"github.com/go-sqlx/sqlx"
)

func (s *EakteService) GetVorgängeByEaktenID(tx *sqlx.Tx, id string) ([]models.VorgangModel, customerrors.ErrorInterface) {
	return s.storage.GetVorgängeByEaktenID(tx, id)
}
