package models

import (
	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service/models"
	"github.com/google/uuid"
)

type ApplicantBySchoolModels struct {
	Count      uint                      `json:"count"`
	MaxCount   uint                      `json:"maxCount"`
	Applicants []*ApplicantBySchoolModel `json:"applicants"`
}

type ApplicantBySchoolModel struct {
	ID               uuid.UUID `json:"id"`
	Firstname        string    `json:"firstname"`
	Lastname         string    `json:"lastname"`
	ClassLevel       string    `json:"classLevel"`
	StatusIdentifier string    `json:"statusIdentifier"`

	Address *ApplicantBySchoolAddressModel `json:"address"`
	Degree  *ApplicantBySchoolDegreeModel  `json:"degree"`
}

type ApplicantBySchoolAddressModel struct {
	ZipCode string `json:"zipCode"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type ApplicantBySchoolDegreeModel struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func ToHttpApplicantBySchoolModel(model *serviceModels.ApplicantBySchoolModel) *ApplicantBySchoolModel {
	if model == nil {
		return nil
	}

	return &ApplicantBySchoolModel{
		ID:               model.ID,
		Firstname:        model.Firstname,
		Lastname:         model.Lastname,
		ClassLevel:       model.ClassLevel,
		StatusIdentifier: model.StatusIdentifier,

		Address: ToHttpApplicantBySchoolAddressModel(model.Address),
		Degree:  ToHttpApplicantBySchoolDegreeModel(model.Degree),
	}
}

func ToHttpApplicantBySchoolAddressModel(model *serviceModels.ApplicantBySchoolAddressModel) *ApplicantBySchoolAddressModel {
	if model == nil {
		return nil
	}

	return &ApplicantBySchoolAddressModel{
		ZipCode: model.ZipCode,
		City:    model.City,
		Country: model.Country,
	}
}

func ToHttpApplicantBySchoolDegreeModel(model *serviceModels.ApplicantBySchoolDegreeModel) *ApplicantBySchoolDegreeModel {
	if model == nil {
		return nil
	}

	return &ApplicantBySchoolDegreeModel{
		ID:   model.ID,
		Name: model.Name,
	}
}

func ToHttpApplicantsBySchoolModels(models []serviceModels.ApplicantBySchoolModel, maxCount uint) *ApplicantBySchoolModels {
	var httpModels []*ApplicantBySchoolModel

	for _, m := range models {
		httpModels = append(httpModels, ToHttpApplicantBySchoolModel(&m))
	}

	return &ApplicantBySchoolModels{
		Count:      uint(len(httpModels)),
		MaxCount:   maxCount,
		Applicants: httpModels,
	}
}
