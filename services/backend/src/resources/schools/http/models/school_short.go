package models

import (
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/schools/service/model"
	"github.com/google/uuid"
)

type SchoolShortResponseModel struct {
	Count    uint `json:"count"`
	MaxCount uint `json:"maxCount"`

	Schools []SchoolShortModel `json:"schools"`
}

type SchoolShortModel struct {
	ID          uuid.UUID        `json:"id"`
	Name        string           `json:"name"`
	Type        *SchoolTypeModel `json:"type"`
	Street      string           `json:"street"`
	HouseNumber string           `json:"houseNumber"`
	City        string           `json:"city"`
	ZipCode     string           `json:"zipCode"`
	Country     string           `json:"country"`
}

func ToHttpSchoolShortModel(model *serviceModel.SchoolShortModel) *SchoolShortModel {
	if model == nil {
		return &SchoolShortModel{}
	}

	return &SchoolShortModel{
		ID:   model.ID,
		Name: model.Name,
		Type: &SchoolTypeModel{
			Name:       model.Type.Name,
			Identifier: model.Type.Identifier,
		},
		Street:      model.Street,
		HouseNumber: model.HouseNumber,
		City:        model.City,
		ZipCode:     model.ZipCode,
		Country:     model.Country,
	}
}

func ToHttpSchoolShortModels(models []serviceModel.SchoolShortModel, maxCount uint) SchoolShortResponseModel {

	var httpModels []SchoolShortModel
	for _, m := range models {
		httpModels = append(httpModels, *ToHttpSchoolShortModel(&m))
	}

	return SchoolShortResponseModel{
		Count:    uint(len(httpModels)),
		MaxCount: uint(maxCount),
		Schools:  httpModels,
	}
}
