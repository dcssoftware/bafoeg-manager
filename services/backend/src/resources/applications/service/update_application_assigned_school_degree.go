package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/schools/service/model"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *ApplicationsService) UpdateApplicationAssignedSchoolDegree(tx *sqlx.Tx, applicationID uuid.UUID, newSchoolID uuid.UUID) (*model.SchoolDegreeModel, customerrors.ErrorInterface) {
	isUpdateable, application, applicationErr := s.IsApplicationUpdatableByApplicationID(tx, applicationID)
	if applicationErr != nil {
		return nil, applicationErr
	}

	degree, degreeErr := s.schoolService.GetSchoolDegreeByID(tx, newSchoolID)
	if degreeErr != nil {
		return nil, degreeErr
	}

	if !isUpdateable {
		return nil, customerrors.NewApplicationAlreadyProcessedError()
	}

	updateSchoolDegreeIDErr := s.storage.UpdateApplicationAssignedSchoolDegree(tx, application.ID, degree.ID)
	if updateSchoolDegreeIDErr != nil {
		return nil, updateSchoolDegreeIDErr
	}

	return degree, nil
}
