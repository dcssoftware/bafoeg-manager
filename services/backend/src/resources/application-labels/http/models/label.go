package models

import (
	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/application-labels/service/models"
	"github.com/google/uuid"
)

type ApplicationLabelModels struct {
	Labels []ApplicationLabel `json:"labels"`
}

type ApplicationLabel struct {
	ID    uuid.UUID              `json:"id"`
	Name  string                 `json:"name"`
	Style *ApplicationLabelStyle `json:"style"`
}
type ApplicationLabelStyle struct {
	ID                   uuid.UUID `json:"id"`
	Name                 string    `json:"name"`
	ColorLight           string    `json:"colorLight"`
	BackgroundColorLight string    `json:"bgColorLight"`
	ColorDark            string    `json:"colorDark"`
	BackgroundColorDark  string    `json:"bgColorDark"`
}

func ToHttpModel(model *serviceModels.ApplicationLabel) *ApplicationLabel {
	if model == nil {
		return &ApplicationLabel{}
	}

	return &ApplicationLabel{
		ID:   model.ID,
		Name: model.Name,
		Style: &ApplicationLabelStyle{
			ID:                   model.Style.ID,
			Name:                 model.Style.Name,
			ColorDark:            model.Style.ColorDark,
			BackgroundColorDark:  model.Style.BackgroundColorDark,
			ColorLight:           model.Style.ColorLight,
			BackgroundColorLight: model.Style.BackgroundColorLight,
		},
	}
}

func ToHttpModels(models []serviceModels.ApplicationLabel) ApplicationLabelModels {

	var httpModels []ApplicationLabel
	for _, m := range models {
		httpModels = append(httpModels, *ToHttpModel(&m))
	}

	return ApplicationLabelModels{
		Labels: httpModels,
	}
}
