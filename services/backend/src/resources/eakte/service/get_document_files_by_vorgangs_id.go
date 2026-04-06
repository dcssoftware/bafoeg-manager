package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/eakte/service/models"
	"github.com/go-sqlx/sqlx"
)

func (s *EakteService) GetFilesByVorgangsID(tx *sqlx.Tx, vorgangsID string) ([]models.DokumentModel, uint, customerrors.ErrorInterface) {
	models, modelsErr := s.storage.GetFilesByVorgangsID(tx, vorgangsID)
	if modelsErr != nil {
		return nil, 0, modelsErr
	}

	modelsCount, modelsCountErr := s.storage.GetFilesByVorgangsIDCount(tx, vorgangsID)
	if modelsCountErr != nil {
		return nil, 0, modelsCountErr
	}

	return models, modelsCount, nil
}
