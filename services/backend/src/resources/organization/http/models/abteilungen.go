package models

import (
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/organization/service/models"
	"github.com/google/uuid"
)

type AbteilungenResponseModel struct {
	Count    uint `json:"count"`
	MaxCount uint `json:"maxCount"`

	Abteilungen []Abteilung `json:"abteilungen"`
}

type Abteilung struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	BehördeID uuid.UUID `json:"behoerde_id"`
}

func ToAbteilungHttpModel(models *serviceModel.Abteilung) *Abteilung {
	return &Abteilung{
		ID:        models.ID,
		Name:      models.Name,
		BehördeID: models.BehördeID,
	}
}

func toAbteilungenHttpModels(models []serviceModel.Abteilung) []Abteilung {
	var result []Abteilung
	for _, model := range models {
		result = append(result, *ToAbteilungHttpModel(&model))
	}
	return result
}

func ToAbteilungenResponseModel(maxCount uint, models []serviceModel.Abteilung) *AbteilungenResponseModel {
	return &AbteilungenResponseModel{
		Count:       uint(len(models)),
		MaxCount:    maxCount,
		Abteilungen: toAbteilungenHttpModels(models),
	}
}
