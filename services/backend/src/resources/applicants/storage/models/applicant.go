package models

import (
	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service/models"
	"github.com/google/uuid"
)

type ApplicantModel struct {
	ID        uuid.UUID `db:"id"`
	Applicant *ApplicantApplicantModel
}

type ApplicantTrainingsAddressModel struct {
	ID          uuid.UUID `db:"id"`
	Street      string    `db:"street"`
	HouseNumber string    `db:"house_number"`
	ZipCode     string    `db:"zip_code"`
	City        string    `db:"city"`
	Country     string    `db:"country"`
}

type ApplicantApplicantModel struct {
	ID        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`

	Address     *ApplicantApplicantAddressModel     `json:"address"`
	ContactData *ApplicantApplicantContactDataModel `json:"contact_data"`
}

type ApplicantApplicantAddressModel struct {
	Street      string `json:"street"`
	HouseNumber string `json:"house_number"`
	ZipCode     string `json:"zip_code"`
	City        string `json:"city"`
	Country     string `json:"country"`
}

type ApplicantApplicantContactDataModel struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func (u *ApplicantTrainingsAddressModel) ToServiceModel() *serviceModels.ApplicantTrainingsAddressModel {
	if u == nil {
		return nil
	}

	return &serviceModels.ApplicantTrainingsAddressModel{
		Street:      u.Street,
		HouseNumber: u.HouseNumber,
		ZipCode:     u.ZipCode,
		City:        u.City,
		Country:     u.Country,
	}
}

func (u *ApplicantModel) ToServiceModel() *serviceModels.ApplicantModel {
	if u == nil {
		return nil
	}

	return &serviceModels.ApplicantModel{
		ID:        u.ID,
		Firstname: u.Applicant.Firstname,
		Lastname:  u.Applicant.Lastname,
		Address:   u.Applicant.Address.ToServiceModel(),
		Contact:   u.Applicant.ContactData.ToServiceModel(),
	}
}

func (model *ApplicantApplicantAddressModel) ToServiceModel() *serviceModels.ApplicantAddressModel {
	return &serviceModels.ApplicantAddressModel{
		Street:      model.Street,
		HouseNumber: model.HouseNumber,
		ZipCode:     model.ZipCode,
		City:        model.City,
		Country:     model.Country,
	}
}

func (model *ApplicantApplicantContactDataModel) ToServiceModel() *serviceModels.ApplicantContactModel {
	return &serviceModels.ApplicantContactModel{
		Phone: model.Phone,
		Email: model.Email,
	}
}

func ToApplicationServiceModels(models []ApplicantModel) []serviceModels.ApplicantModel {
	var newModels []serviceModels.ApplicantModel
	for _, model := range models {
		newModels = append(newModels, *model.ToServiceModel())
	}
	return newModels
}

// func (u *ApplicantModel) toApplicantAddressServiceModel() *serviceModels.ApplicantAddressModel {
// 	if u == nil {
// 		return nil
// 	}

// 	return &serviceModels.ApplicantAddressModel{
// 		Street:      u.Street,
// 		HouseNumber: u.HouseNumber,
// 		ZipCode:     u.ZipCode,
// 		City:        u.City,
// 		Country:     u.Country,
// 	}
// }

// func (u *ApplicantModel) toApplicantContactServiceModel() *serviceModels.ApplicantContactModel {
// 	if u == nil {
// 		return nil
// 	}

// 	return &serviceModels.ApplicantContactModel{
// 		Phone: u.Phone,
// 		Email: u.Email,
// 	}
// }

// func ToApplicantServiceModels(user []ApplicantModel) []serviceModels.ApplicantModel {
// 	var serviceModels []serviceModels.ApplicantModel

// 	for _, u := range user {
// 		serviceModels = append(serviceModels, *u.ToApplicantServiceModel())
// 	}

// 	return serviceModels
// }
