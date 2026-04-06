package models

import (
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/schools/service/model"
	"github.com/google/uuid"
)

type SchoolResponseModel struct {
	Count    uint `json:"count"`
	MaxCount uint `json:"maxCount"`

	Schools []SchoolModel `json:"schools"`
}

type SchoolModel struct {
	ID          uuid.UUID           `json:"id"`
	Name        string              `json:"name"`
	Type        *SchoolTypeModel    `json:"type"`
	Degree      []SchoolDegreeModel `json:"degree"`
	Street      string              `json:"street"`
	HouseNumber string              `json:"houseNumber"`
	City        string              `json:"city"`
	ZipCode     string              `json:"zipCode"`
	Country     string              `json:"country"`
}

type SchoolTypeModel struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}

func ToHttpSchoolModel(model *serviceModel.SchoolModel) *SchoolModel {
	if model == nil {
		return &SchoolModel{}
	}

	var degrees []SchoolDegreeModel
	for _, degree := range model.Degree {
		degrees = append(degrees, ToHttpSchoolDegreeModel(degree))
	}

	return &SchoolModel{
		ID:   model.ID,
		Name: model.Name,
		Type: &SchoolTypeModel{
			Name:       model.Type.Name,
			Identifier: model.Type.Identifier,
		},
		Degree:      degrees,
		Street:      model.Street,
		HouseNumber: model.HouseNumber,
		City:        model.City,
		ZipCode:     model.ZipCode,
		Country:     model.Country,
	}
}

func ToHttpResponseSchoolModels(models []serviceModel.SchoolModel, maxCount uint) []SchoolModel {

	var httpModels []SchoolModel
	for _, m := range models {
		httpModels = append(httpModels, *ToHttpSchoolModel(&m))
	}

	return httpModels
}
