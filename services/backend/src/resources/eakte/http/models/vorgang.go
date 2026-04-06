package models

import (
	"time"

	"github.com/google/uuid"

	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/eakte/service/models"
)

type VorgangModel struct {
	ID              uuid.UUID `json:"id"`
	EakteID         uuid.UUID `json:"eakte_id"`
	Vorgangszeichen string    `json:"vorgangszeichen"`
	Created         time.Time `json:"created"`
}

func ToVorgangHttpModel(model *serviceModels.VorgangModel) *VorgangModel {
	return &VorgangModel{
		ID:              model.ID,
		EakteID:         model.EakteID,
		Vorgangszeichen: model.Vorgangszeichen,
		Created:         model.Created,
	}
}

func ToVorgängeHttpModel(models []serviceModels.VorgangModel) []VorgangModel {
	var httpModels []VorgangModel

	for _, model := range models {
		httpModels = append(httpModels, *ToVorgangHttpModel(&model))
	}

	return httpModels
}
