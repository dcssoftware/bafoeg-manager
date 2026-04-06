package models

import (
	"time"

	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/eakte/service/models"
	"github.com/google/uuid"
)

type EaktenResponseModels struct {
	Count    uint `json:"count"`
	MaxCount uint `json:"maxCount"`

	Eakten []EakteModel `json:"eakten"`
}

type EakteModel struct {
	ID              uuid.UUID      `json:"id"`
	Aktenzeichen    string         `json:"aktenzeichen"`
	Typ             EakteModelType `json:"type"`
	Vertraulichkeit string         `json:"vertraulichkeit"`
	Created         *time.Time     `json:"created"`
}

type EakteModelType struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Identifier string    `json:"identifier"`
}

func ToEakteHttpModel(model *serviceModels.EakteModel) *EakteModel {

	typ := EakteModelType{
		ID:         model.Typ.ID,
		Identifier: model.Typ.Identifier,
		Name:       model.Typ.Name,
	}

	return &EakteModel{
		ID:              model.ID,
		Aktenzeichen:    model.Aktenzeichen,
		Typ:             typ,
		Vertraulichkeit: model.Vertraulichkeit.String(),
		Created:         model.Created,
	}
}

func ToEaktenHttpModel(models []serviceModels.EakteModel, maxCount uint) *EaktenResponseModels {

	var httpModels []EakteModel
	for _, model := range models {
		httpModels = append(httpModels, *ToEakteHttpModel(&model))
	}

	return &EaktenResponseModels{
		MaxCount: maxCount,
		Count:    uint(len(httpModels)),

		Eakten: httpModels,
	}

}
