package models

import (
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/organization/service/models"
	"github.com/google/uuid"
)

type RegionResponseModel struct {
	Count    uint `json:"count"`
	MaxCount uint `json:"maxCount"`

	Regions []Region `json:"regions"`
}
type Region struct {
	ID         uuid.UUID `json:"id"`
	Identifier string    `json:"identifier"`
	Name       string    `json:"name"`
}

func ToRegionsHttpModel(models *serviceModel.Region) *Region {
	return &Region{
		ID:         models.ID,
		Identifier: models.Identifier,
		Name:       models.Name,
	}
}

func toRegionsHttpModels(models []serviceModel.Region) []Region {
	var result []Region
	for _, model := range models {
		result = append(result, *ToRegionsHttpModel(&model))
	}
	return result
}

func ToRegionResponseModel(maxCount uint, models []serviceModel.Region) *RegionResponseModel {
	return &RegionResponseModel{
		Count:    uint(len(models)),
		MaxCount: maxCount,
		Regions:  toRegionsHttpModels(models),
	}
}
