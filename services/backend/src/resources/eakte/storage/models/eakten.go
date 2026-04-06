package models

import (
	"time"

	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/eakte/service/models"
	stateModels "github.com/dcssoftware/bafoeg-manager/src/resources/eakte/states/models"
	"github.com/google/uuid"
)

type EakteModel struct {
	ID              uuid.UUID                       `db:"id"`
	Aktenzeichen    string                          `db:"aktenzeichen"`
	Typ             EakteModelType                  `db:"typ"`
	Vertraulichkeit stateModels.VertraulichkeitEnum `db:"vertraulichkeit"`
	Created         *time.Time                      `db:"created"`
}

type EakteModelType struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Identifier string    `json:"identifier"`
}

func (m EakteModel) ToServiceModel() *serviceModels.EakteModel {
	return &serviceModels.EakteModel{
		ID:           m.ID,
		Aktenzeichen: m.Aktenzeichen,
		Typ: serviceModels.EakteModelType{
			ID:         m.Typ.ID,
			Identifier: m.Typ.Identifier,
			Name:       m.Typ.Name,
		},
		Vertraulichkeit: m.Vertraulichkeit,
		Created:         m.Created,
	}
}
