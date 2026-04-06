package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *ApplicationsService) IsApplicationUpdatableByApplicationID(tx *sqlx.Tx, applicationID uuid.UUID) (bool, *models.ApplicationModel, customerrors.ErrorInterface) {
	application, applicationErr := s.GetApplicationByID(tx, applicationID.String())
	if applicationErr != nil {
		return false, nil, applicationErr
	}

	return application.Status.IsUpdatable(), application, nil
}
