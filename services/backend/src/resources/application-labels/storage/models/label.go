package models

import (
	"encoding/json"
	"errors"
	"fmt"

	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/application-labels/service/models"
	"github.com/google/uuid"
)

type ApplicationLabel struct {
	ID    uuid.UUID              `db:"id"`
	Name  string                 `db:"name"`
	Style *ApplicationLabelStyle `db:"style"`
}

type ApplicationLabelStyle struct {
	ID                   uuid.UUID `json:"id"`
	Name                 string    `json:"name"`
	ColorDark            string    `json:"color_dark"`
	BackgroundColorDark  string    `json:"bg_color_dark"`
	ColorLight           string    `json:"color_light"`
	BackgroundColorLight string    `json:"bg_color_light"`
}

func (model *ApplicationLabel) ToApplicationLabelServiceModel() *serviceModels.ApplicationLabel {
	if model == nil {
		return &serviceModels.ApplicationLabel{}
	}

	return &serviceModels.ApplicationLabel{
		ID:   model.ID,
		Name: model.Name,
		Style: &serviceModels.ApplicationLabelStyle{
			ID:                   model.Style.ID,
			Name:                 model.Style.Name,
			ColorDark:            model.Style.ColorDark,
			BackgroundColorDark:  model.Style.BackgroundColorDark,
			ColorLight:           model.Style.ColorLight,
			BackgroundColorLight: model.Style.BackgroundColorLight,
		},
	}
}

func (model *ApplicationLabelStyle) Scan(val any) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, &model)
	case string:
		return json.Unmarshal([]byte(v), &model)
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}
