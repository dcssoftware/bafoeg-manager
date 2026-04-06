package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicationsService) GetApplicationRevisionsByApplicationID(tx *sqlx.Tx, page uint, applicantID string) ([]models.ApplicationRevisionShortModel, customerrors.ErrorInterface) {
	return s.storage.GetApplicationRevisionsByApplicationID(tx, page, applicantID)
}
