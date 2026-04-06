package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/eakte/service/models"
	"github.com/go-sqlx/sqlx"
)

func (s *EakteService) GetFilesByAkteID(tx *sqlx.Tx, akteID string) ([]models.DokumentModel, uint, customerrors.ErrorInterface) {
	models, modelsErr := s.storage.GetFilesByAkteID(tx, akteID)
	if modelsErr != nil {
		return nil, 0, modelsErr
	}

	modelsCount, modelsCountErr := s.storage.GetFilesByAkteIDCount(tx, akteID)
	if modelsCountErr != nil {
		return nil, 0, modelsCountErr
	}

	return models, modelsCount, nil
}
