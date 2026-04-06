package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/eakte/service/models"
	"github.com/go-sqlx/sqlx"
)

func (s *EakteService) GetEakten(tx *sqlx.Tx, page uint) ([]models.EakteModel, uint, customerrors.ErrorInterface) {
	akten, aktenErr := s.storage.GetEakten(tx, page)
	if aktenErr != nil {
		return nil, 0, aktenErr
	}

	aktenCount, aktenCountErr := s.storage.GetEaktenCount(tx)
	if aktenCountErr != nil {
		return nil, 0, aktenCountErr
	}

	return akten, aktenCount, nil
}
