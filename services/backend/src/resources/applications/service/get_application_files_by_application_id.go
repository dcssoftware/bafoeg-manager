package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicationsService) GetApplicationFilesByApplicationID(tx *sqlx.Tx, page uint, applicationID string) ([]models.ApplicationFile, customerrors.ErrorInterface) {
	return s.storage.GetApplicationFilesByApplicationID(tx, page, applicationID)
}
