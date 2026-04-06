package models

import (
	"time"

	// serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"

	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/google/uuid"
)

// ApplicationLabelModels is a slice of ApplicationLabelModel with custom scanner implementation
type ApplicationLabelModels []ApplicationLabelModel

type ApplicationModel struct {
	ApplicationID         uuid.UUID `db:"id"`
	ApplicationClassLevel string    `db:"class_level"`

	Labels           ApplicationLabelModels           `db:"labels"`
	Applicant        ApplicationApplicantModel        `db:"applicant"`
	SchoolWithDegree ApplicationSchoolWithDegreeModel `db:"school"`
	AssignedUser     *ApplicationAssignedUserModel    `db:"assigned_user"`
	Status           ApplicationStatusModel           `db:"status"`

	Created time.Time `db:"application_created"`
	Updated time.Time `db:"application_updated"`
}
type ApplicationLabelModel struct {
	ID    uuid.UUID                  `json:"id"`
	Name  string                     `json:"name"`
	Color ApplicationLabelColorModel `json:"color"`
}
type ApplicationLabelColorModel struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`

	ColorDark       string `json:"color_dark"`
	BgColorDark     string `json:"bg_color_dark"`
	BorderColorDark string `json:"border_color_dark"`

	ColorLight       string `json:"color_light"`
	BgColorLight     string `json:"bg_color_light"`
	BorderColorLight string `json:"border_color_light"`
}

type ApplicationApplicantModel struct {
	ID          uuid.UUID                            `json:"id"`
	Firstname   string                               `json:"firstname"`
	Lastname    string                               `json:"lastname"`
	Address     ApplicationApplicantAddressModel     `json:"address"`
	ContactData ApplicationApplicantContactDataModel `json:"contact_data"`
}

type ApplicationApplicantAddressModel struct {
	Street      string `json:"street"`
	HouseNumber string `json:"house_number"`
	ZipCode     string `json:"zip_code"`
	City        string `json:"city"`
	Country     string `json:"country"`
}

type ApplicationApplicantContactDataModel struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type ApplicationSchoolWithDegreeModel struct {
	ID      uuid.UUID                               `json:"id"`
	Name    string                                  `json:"name"`
	Address ApplicationSchoolWithDegreeAddressModel `json:"address"`
	Type    ApplicationSchoolWithDegreeTypeModel    `json:"type"`
	Degree  ApplicationSchoolWithDegreeDegreeModel  `json:"degree"`
}
type ApplicationSchoolWithDegreeAddressModel struct {
	Street      string `json:"street"`
	HouseNumber string `json:"house_number"`
	ZipCode     string `json:"zip_code"`
	City        string `json:"city"`
	Country     string `json:"country"`
}

type ApplicationSchoolWithDegreeTypeModel struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Identifier string    `json:"identifier"`
}

type ApplicationSchoolWithDegreeDegreeModel struct {
	ID                                 uuid.UUID `json:"id"`
	Name                               string    `json:"name"`
	FosBerufsabschlussRequired         bool      `json:"fos_berufsabschluss_required"`
	BosBerufsqualifizierenderAbschluss bool      `json:"bos_berufsqualifizierender_abschluss"`
	FachschuleBerufsabschlussRequired  bool      `json:"fachschule_berufsabschluss_required"`
}

type ApplicationAssignedUserModel struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	DisplayName string    `json:"display_name"`
}

type ApplicationStatusModel struct {
	ID         uuid.UUID `db:"id" json:"id"`
	Name       string    `db:"name" json:"name"`
	Identifier string    `db:"identifier" json:"identifier"`
}

func (m *ApplicationModel) ToServiceModel() *serviceModel.ApplicationModel {
	model := &serviceModel.ApplicationModel{
		ID:         m.ApplicationID,
		ClassLevel: m.ApplicationClassLevel,

		Labels:           m.Labels.ToServiceModel(),
		Applicant:        m.Applicant.ToServiceModel(),
		SchoolWithDegree: m.SchoolWithDegree.ToServiceModel(),
		AssignedUser:     m.AssignedUser.ToServiceModel(),
		Status:           m.Status.ToServiceModel(),
		ProcessingTime:   nil,

		Created: m.Created,
		Updated: m.Updated,
	}

	return model
}

func (models ApplicationLabelModels) ToServiceModel() *[]serviceModel.ApplicationLabel {
	var serviceModels []serviceModel.ApplicationLabel

	for _, model := range models {
		serviceModels = append(serviceModels, serviceModel.ApplicationLabel{
			ID:   model.ID,
			Name: model.Name,
			Color: &serviceModel.ApplicationLabelColor{
				ID:   model.Color.ID,
				Name: model.Color.Name,

				ColorDark:       model.Color.ColorDark,
				BgColorDark:     model.Color.BgColorDark,
				BorderColorDark: model.Color.BorderColorDark,

				ColorLight:       model.Color.ColorLight,
				BgColorLight:     model.Color.BgColorLight,
				BorderColorLight: model.Color.BorderColorLight,
			},
		})
	}
	return &serviceModels
}

func (model ApplicationApplicantModel) ToServiceModel() *serviceModel.ApplicationApplicant {
	return &serviceModel.ApplicationApplicant{
		ID:          model.ID,
		Firstname:   model.Firstname,
		Lastname:    model.Lastname,
		Address:     model.Address.ToServiceModel(),
		ContactData: model.ContactData.ToServiceModel(),
	}
}
func (model ApplicationApplicantAddressModel) ToServiceModel() *serviceModel.ApplicationApplicantAddress {
	return &serviceModel.ApplicationApplicantAddress{
		Street:      model.Street,
		HouseNumber: model.HouseNumber,
		ZipCode:     model.ZipCode,
		City:        model.City,
		Country:     model.Country,
	}
}
func (model ApplicationApplicantContactDataModel) ToServiceModel() *serviceModel.ApplicationApplicantContactData {
	return &serviceModel.ApplicationApplicantContactData{
		Email: model.Email,
		Phone: model.Phone,
	}
}

func (model ApplicationSchoolWithDegreeModel) ToServiceModel() *serviceModel.ApplicationSchoolWithDegree {
	return &serviceModel.ApplicationSchoolWithDegree{
		ID:      model.ID,
		Name:    model.Name,
		Address: model.Address.ToServiceModel(),
		Type:    model.Type.ToServiceModel(),
		Degree:  model.Degree.ToServiceModel(),
	}
}

func (model ApplicationSchoolWithDegreeAddressModel) ToServiceModel() *serviceModel.ApplicationSchoolWithDegreeAddress {
	return &serviceModel.ApplicationSchoolWithDegreeAddress{
		Street:      model.Street,
		HouseNumber: model.HouseNumber,
		ZipCode:     model.ZipCode,
		City:        model.City,
		Country:     model.Country,
	}
}

func (model ApplicationSchoolWithDegreeTypeModel) ToServiceModel() *serviceModel.ApplicationSchoolWithDegreeType {
	return &serviceModel.ApplicationSchoolWithDegreeType{
		ID:         model.ID,
		Name:       model.Name,
		Identifier: model.Identifier,
	}
}

func (model ApplicationSchoolWithDegreeDegreeModel) ToServiceModel() *serviceModel.ApplicationSchoolWithDegreeDegree {
	return &serviceModel.ApplicationSchoolWithDegreeDegree{
		ID:                                 model.ID,
		Name:                               model.Name,
		FosBerufsabschlussRequired:         model.FosBerufsabschlussRequired,
		BosBerufsqualifizierenderAbschluss: model.BosBerufsqualifizierenderAbschluss,
		FachschuleBerufsabschlussRequired:  model.FachschuleBerufsabschlussRequired,
	}
}

func (model *ApplicationAssignedUserModel) ToServiceModel() *serviceModel.ApplicationAssignedUser {
	if model == nil {
		return nil
	}

	return &serviceModel.ApplicationAssignedUser{
		ID:          model.ID,
		Username:    model.Username,
		DisplayName: model.DisplayName,
	}
}

func (model ApplicationStatusModel) ToServiceModel() *serviceModel.ApplicationStatus {
	return &serviceModel.ApplicationStatus{
		ID:         model.ID,
		Name:       model.Name,
		Identifier: model.Identifier,
	}
}
