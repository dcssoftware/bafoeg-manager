package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *ApplicationsService) UpdateApplicationAssignedApplicant(tx *sqlx.Tx, applicationID uuid.UUID, applicantID uuid.UUID) (*models.ApplicantModel, customerrors.ErrorInterface) {
	isUpdateable, application, applicationErr := s.IsApplicationUpdatableByApplicationID(tx, applicationID)
	if applicationErr != nil {
		return nil, applicationErr
	}

	applicant, applicantErr := s.applicantsService.GetApplicantByID(tx, applicantID.String())
	if applicantErr != nil {
		return nil, applicantErr
	}

	if !isUpdateable {
		return nil, customerrors.NewApplicationAlreadyProcessedError()
	}

	updateApplicantIDErr := s.storage.UpdateApplicationAssignedApplicant(tx, application.ID, applicant.ID)
	if updateApplicantIDErr != nil {
		return nil, updateApplicantIDErr
	}

	return applicant, nil
}
