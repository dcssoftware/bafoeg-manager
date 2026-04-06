package models

import (
	"time"

	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/google/uuid"
)

type ApplicationRevisionShortModel struct {
	ID          uuid.UUID `json:"id"`
	Header      string    `json:"header"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
}

type ApplicationRevisionShortModels struct {
	Count     uint                            `json:"count"`
	MaxCount  uint                            `json:"maxCount"`
	Revisions []ApplicationRevisionShortModel `json:"revisions"`
}

func ToRevisionShortHttpModel(model *serviceModel.ApplicationRevisionShortModel) *ApplicationRevisionShortModel {
	if model == nil {
		return &ApplicationRevisionShortModel{}
	}

	return &ApplicationRevisionShortModel{
		ID:          model.ID,
		Header:      model.Header,
		Description: model.Description,
		Created:     model.Created,
	}
}

func ToRevisionShortHttpModels(models []serviceModel.ApplicationRevisionShortModel, maxCount uint) ApplicationRevisionShortModels {

	var httpModels []ApplicationRevisionShortModel
	for _, m := range models {
		httpModels = append(httpModels, *ToRevisionShortHttpModel(&m))
	}

	return ApplicationRevisionShortModels{
		MaxCount:  maxCount,
		Count:     uint(len(httpModels)),
		Revisions: httpModels,
	}
}
