package service

import (
	"time"

	"github.com/ccoveille/go-safecast/v2"
	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
)

func CalculateRemainingProcessingTime(created time.Time) *models.ApplicationProcessingTime {
	maxLegalProcessingTimeFloat, _ := safecast.Convert[float64](configuration.ApplicationConfiguration.MaxLegalProcessingTime)

	durationInDaysFloat := time.Since(created).Hours() / 24
	remainingDaysFloat := maxLegalProcessingTimeFloat - durationInDaysFloat
	remainingDays, _ := safecast.Convert[int](remainingDaysFloat)

	remainingDaysInPercent := remainingDaysFloat / maxLegalProcessingTimeFloat * 100

	return &models.ApplicationProcessingTime{
		RemainingTimeInDays:    remainingDays,
		RemainingTimeInPercent: remainingDaysInPercent,
		IsStillLegal:           remainingDays >= 0,
	}
}
