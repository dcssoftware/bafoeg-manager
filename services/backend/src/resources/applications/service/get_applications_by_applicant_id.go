package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicationsService) GetApplicationsByApplicantID(tx *sqlx.Tx, pagination uint, applicantID string) ([]models.ApplicationShortModel, customerrors.ErrorInterface) {
	applications, applicationsErr := s.storage.GetApplicationsByApplicantID(tx, pagination, applicantID)
	if applicationsErr != nil {
		return applications, applicationsErr
	}

	var newApplications []models.ApplicationShortModel
	for _, application := range applications {
		application.ProcessingTime = CalculateRemainingProcessingTime(application.Created)
		newApplications = append(newApplications, application)
	}

	return newApplications, nil
}
