package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicationsService) GetApplicationStatusByIdentifier(tx *sqlx.Tx, identifier string) (*models.ApplicationStatus, customerrors.ErrorInterface) {
	return s.storage.GetApplicationStatusByIdentifier(tx, identifier)
}
