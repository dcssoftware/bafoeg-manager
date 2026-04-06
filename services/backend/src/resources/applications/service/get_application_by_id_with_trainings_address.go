package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicationsService) GetApplicationByIDWithTrainingsAddress(tx *sqlx.Tx, applicationID string) (*models.ApplicationModel, customerrors.ErrorInterface) {
	application, applicationErr := s.storage.GetApplicationByID(tx, applicationID)
	if applicationErr != nil {
		return application, applicationErr
	}

	trainingsAddress, trainingsAddressErr := s.applicantsService.GetApplicantTrainingsAddressByApplicantID(
		tx,
		application.Applicant.ID.String(),
	)
	if trainingsAddressErr == nil {
		application.Applicant.TrainingsAddress = trainingsAddress
	}

	application.ProcessingTime = CalculateRemainingProcessingTime(application.Created)
	return application, applicationErr
}
