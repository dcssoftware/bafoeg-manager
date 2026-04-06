package models

import (
	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service/models"
	"github.com/google/uuid"
)

type ApplicantResponseModels struct {
	Count    uint `json:"count"`
	MaxCount uint `json:"maxCount"`

	Applicants []*ApplicantModel `json:"applicants"`
}

type ApplicantModel struct {
	ID        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`

	BalancePayback  *float64 `json:"balancePayback"`
	BalanceOutgoing *float64 `json:"balanceOutgoing"`

	Address          *ApplicantAddressModel          `json:"address"`
	TrainingsAddress *ApplicantTrainingsAddressModel `json:"trainingsAddress"`
	Contact          *ApplicantContactModel          `json:"contact"`
}

type ApplicantAddressModel struct {
	Street      string `json:"street"`
	HouseNumber string `json:"houseNumber"`
	ZipCode     string `json:"zipCode"`
	City        string `json:"city"`
	Country     string `json:"country"`
}

type ApplicantTrainingsAddressModel struct {
	Street      string `json:"street"`
	HouseNumber string `json:"houseNumber"`
	ZipCode     string `json:"zipCode"`
	City        string `json:"city"`
	Country     string `json:"country"`
}

type ApplicantContactModel struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func ToHttpApplicantModel(model *serviceModels.ApplicantModel) *ApplicantModel {
	if model == nil {
		return nil
	}

	return &ApplicantModel{
		ID:        model.ID,
		Firstname: model.Firstname,
		Lastname:  model.Lastname,

		BalancePayback:  model.BalancePayback,
		BalanceOutgoing: model.BalanceOutgoing,

		Address:          ToHttpApplicantAddressModel(model.Address),
		TrainingsAddress: ToHttpApplicantTrainingsAddressModel(model.TrainingsAddress),
		Contact:          ToHttpApplicantContactModel(model.Contact),
	}
}

func ToHttpApplicantAddressModel(model *serviceModels.ApplicantAddressModel) *ApplicantAddressModel {
	if model == nil {
		return nil
	}

	return &ApplicantAddressModel{
		Street:      model.Street,
		HouseNumber: model.HouseNumber,
		ZipCode:     model.ZipCode,
		City:        model.City,
		Country:     model.Country,
	}
}

func ToHttpApplicantTrainingsAddressModel(model *serviceModels.ApplicantTrainingsAddressModel) *ApplicantTrainingsAddressModel {
	if model == nil {
		return nil
	}

	return &ApplicantTrainingsAddressModel{
		Street:      model.Street,
		HouseNumber: model.HouseNumber,
		ZipCode:     model.ZipCode,
		City:        model.City,
		Country:     model.Country,
	}
}

func ToHttpApplicantContactModel(model *serviceModels.ApplicantContactModel) *ApplicantContactModel {
	if model == nil {
		return nil
	}

	return &ApplicantContactModel{
		Phone: model.Phone,
		Email: model.Email,
	}
}

func ToHttpApplicantModels(models []serviceModels.ApplicantModel, maxCount uint) *ApplicantResponseModels {
	var httpModels []*ApplicantModel

	for _, m := range models {
		httpModels = append(httpModels, ToHttpApplicantModel(&m))
	}

	return &ApplicantResponseModels{
		MaxCount: maxCount,
		Count:    uint(len(httpModels)),

		Applicants: httpModels,
	}
}
