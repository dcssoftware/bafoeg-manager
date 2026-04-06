package models

import (
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/organization/service/models"
	"github.com/google/uuid"
)

type BehördeResponseModel struct {
	Count    uint `json:"count"`
	MaxCount uint `json:"maxCount"`

	Behörden []Behörde `json:"behoerden"`
}

type Behörde struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	RegionID uuid.UUID `json:"region_id"`
}

func ToBehördeHttpModel(models *serviceModel.Behörde) *Behörde {
	return &Behörde{
		ID:       models.ID,
		Name:     models.Name,
		RegionID: models.RegionID,
	}
}

func toBehördenHttpModels(models []serviceModel.Behörde) []Behörde {
	var result []Behörde
	for _, model := range models {
		result = append(result, *ToBehördeHttpModel(&model))
	}
	return result
}

func ToBehördeResponseModel(maxCount uint, models []serviceModel.Behörde) *BehördeResponseModel {
	return &BehördeResponseModel{
		Count:    uint(len(models)),
		MaxCount: maxCount,
		Behörden: toBehördenHttpModels(models),
	}
}
