package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicationsService) GetApplicationsMetrics(tx *sqlx.Tx, userID string) (*models.ApplicationsMetrics, customerrors.ErrorInterface) {

	metricTotal, metricTotalErr := s.storage.GetApplicationsMetricsTotal(tx)
	if metricTotalErr != nil {
		return nil, metricTotalErr
	}

	metricToday, metricTodayErr := s.storage.GetApplicationsMetricsNewToday(nil)
	if metricTodayErr != nil {
		return nil, metricTodayErr
	}

	metricInProgress, metricInProgressErr := s.storage.GetApplicationsMetricsInProgress(nil)
	if metricInProgressErr != nil {
		return nil, metricInProgressErr
	}

	var userAssigned uint
	var userAssignedErr customerrors.ErrorInterface

	if userID != "" {
		userAssigned, userAssignedErr = s.storage.GetApplicationsMetricsUserAssigned(nil, userID)
		if userAssignedErr != nil {
			return nil, userAssignedErr
		}
	}

	return &models.ApplicationsMetrics{
		Total:        metricTotal,
		NewToday:     metricToday,
		InProgress:   metricInProgress,
		UserAssigned: userAssigned,
	}, nil
}
