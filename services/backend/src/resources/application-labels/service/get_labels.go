package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/application-labels/service/models"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicationLabelsService) GetLabels(tx *sqlx.Tx, page uint) ([]models.ApplicationLabel, customerrors.ErrorInterface) {
	return s.storage.GetLabels(tx, page)
}
