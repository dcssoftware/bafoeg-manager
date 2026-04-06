package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
)

func (s *ApplicationsService) GetApplications(pagination uint, userID string, filterOnlyInProgress bool, filterString string) ([]models.ApplicationShortModel, customerrors.ErrorInterface) {
	var applications []models.ApplicationShortModel
	var applicationsErr customerrors.ErrorInterface

	applications, applicationsErr = s.storage.GetApplications(nil, pagination, userID, filterString, filterOnlyInProgress)
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
