package models

import (
	"time"

	"github.com/google/uuid"
)

type ApplicationRevisionModel struct {
	ID            uuid.UUID
	Header        string
	Description   string
	ApplicationID uuid.UUID

	TrainingsAddressID uuid.UUID

	Created time.Time
}

type ApplicationRevisionShortModel struct {
	ID            uuid.UUID
	Header        string
	Description   string
	ApplicationID uuid.UUID
	Created       time.Time
}

func (model *ApplicationRevisionModel) ToDataModel() *ApplicationRevisionDataModel {
	if model == nil {
		return &ApplicationRevisionDataModel{}
	}

	return &ApplicationRevisionDataModel{
		TrainingsAddressID: model.TrainingsAddressID.String(),
	}
}

type ApplicationRevisionDataModel struct {
	TrainingsAddressID string
}

func GetTrainingsAddressID(baseModel *ApplicationRevisionDataModel, newModel *ApplicationRevisionDataModel) *string {
	if baseModel == nil && newModel == nil {
		var returnValue *string
		return returnValue
	}

	if newModel != nil {
		return &newModel.TrainingsAddressID
	}

	return &baseModel.TrainingsAddressID

}
