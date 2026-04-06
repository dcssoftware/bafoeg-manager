package models

import (
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
)

type ApplicationsMetrics struct {
	Total        uint `json:"total"`
	NewToday     uint `json:"newToday"`
	InProgress   uint `json:"inProgress"`
	UserAssigned uint `json:"userAssigned"`
}

func ToApplicationsMetricsHttpModel(model *serviceModel.ApplicationsMetrics) *ApplicationsMetrics {
	if model == nil {
		return &ApplicationsMetrics{}
	}

	return &ApplicationsMetrics{
		Total:        model.Total,
		NewToday:     model.NewToday,
		InProgress:   model.InProgress,
		UserAssigned: model.UserAssigned,
	}
}
