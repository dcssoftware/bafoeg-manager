package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicationsService) GetApplicationStatusByID(tx *sqlx.Tx, id string) (*models.ApplicationStatus, customerrors.ErrorInterface) {
	return s.storage.GetApplicationStatusByID(tx, id)
}
