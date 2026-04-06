package models

import (
	"time"

	applicantHandlerModel "github.com/dcssoftware/bafoeg-manager/src/resources/applicants/http/models"
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/google/uuid"
)

type ApplicationModel struct {
	ID         uuid.UUID `json:"id"`
	ClassLevel string    `json:"classLevel"`

	Labels           *[]ApplicationLabel          `json:"labels"`
	Applicant        *ApplicationApplicant        `json:"applicant"`
	SchoolWithDegree *ApplicationSchoolWithDegree `json:"school"`
	AssignedUser     *ApplicationAssignedUser     `json:"assignedUser"`
	Status           *ApplicationStatus           `json:"status"`
	ProcessingTime   *ApplicationProcessingTime   `json:"processingTime"`

	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
type ApplicationLabel struct {
	ID    uuid.UUID              `json:"id"`
	Name  string                 `json:"name"`
	Color *ApplicationLabelColor `json:"color"`
}
type ApplicationLabelColor struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`

	ColorDark       string `json:"colorDark"`
	BgColorDark     string `json:"bgColorDark"`
	BorderColorDark string `json:"borderColorDark"`

	ColorLight       string `json:"colorLight"`
	BgColorLight     string `json:"bgColorLight"`
	BorderColorLight string `json:"borderColorLight"`
}

type ApplicationApplicant struct {
	ID               uuid.UUID                                             `json:"id"`
	Firstname        string                                                `json:"firstname"`
	Lastname         string                                                `json:"lastname"`
	Address          *ApplicationApplicantAddress                          `json:"address"`
	TrainingsAddress *applicantHandlerModel.ApplicantTrainingsAddressModel `json:"trainingsAddress"`
	ContactData      *ApplicationApplicantContactData                      `json:"contactData"`
}

type ApplicationApplicantAddress struct {
	Street      string `json:"street"`
	HouseNumber string `json:"houseNumber"`
	ZipCode     string `json:"zipCode"`
	City        string `json:"city"`
	Country     string `json:"country"`
}

type ApplicationApplicantContactData struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type ApplicationSchoolWithDegree struct {
	ID      uuid.UUID                           `json:"id"`
	Name    string                              `json:"name"`
	Address *ApplicationSchoolWithDegreeAddress `json:"address"`
	Type    *ApplicationSchoolWithDegreeType    `json:"type"`
	Degree  *ApplicationSchoolWithDegreeDegree  `json:"degree"`
}

type ApplicationSchoolWithDegreeAddress struct {
	Street      string `json:"street"`
	HouseNumber string `json:"houseNumber"`
	ZipCode     string `json:"zipCode"`
	City        string `json:"city"`
	Country     string `json:"country"`
}

type ApplicationSchoolWithDegreeType struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Identifier string    `json:"identifier"`
}

type ApplicationSchoolWithDegreeDegree struct {
	ID                                 uuid.UUID `json:"id"`
	Name                               string    `json:"name"`
	FosBerufsabschlussRequired         bool      `json:"fosBerufsabschlussRequired"`
	BosBerufsqualifizierenderAbschluss bool      `json:"bosBerufsqualifizierenderAbschluss"`
	FachschuleBerufsabschlussRequired  bool      `json:"fachschuleBerufsschule"`
}

type ApplicationAssignedUser struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	DisplayName string    `json:"displayName"`
}

type ApplicationStatus struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Identifier string    `json:"identifier"`
}

type ApplicationProcessingTime struct {
	MaxValidity            uint    `json:"maxValidity"`
	RemainingTimeInDays    int     `json:"remainingTimeInDays"`
	RemainingTimeInPercent float64 `json:"remainingTimeInPercent"`
	IsStillLegal           bool    `json:"isStillLegal"`
}

func ToHttpModel(model *serviceModel.ApplicationModel) *ApplicationModel {
	if model == nil {
		return nil
	}

	var newLabelModel []ApplicationLabel
	for _, label := range *model.Labels {
		newLabelModel = append(newLabelModel, ApplicationLabel{
			ID:   label.ID,
			Name: label.Name,
			Color: &ApplicationLabelColor{
				ID:   label.Color.ID,
				Name: label.Color.Name,

				ColorDark:       label.Color.ColorDark,
				BgColorDark:     label.Color.BgColorDark,
				BorderColorDark: label.Color.BorderColorDark,

				ColorLight:       label.Color.ColorLight,
				BgColorLight:     label.Color.BgColorLight,
				BorderColorLight: label.Color.BorderColorLight,
			},
		})
	}

	var trainingsAddress *applicantHandlerModel.ApplicantTrainingsAddressModel
	if model.Applicant.TrainingsAddress != nil {
		trainingsAddress = &applicantHandlerModel.ApplicantTrainingsAddressModel{
			Street:      model.Applicant.TrainingsAddress.Street,
			HouseNumber: model.Applicant.TrainingsAddress.HouseNumber,
			ZipCode:     model.Applicant.TrainingsAddress.ZipCode,
			City:        model.Applicant.TrainingsAddress.City,
			Country:     model.Applicant.TrainingsAddress.Country,
		}
	}

	var assignedUser *ApplicationAssignedUser
	if model.AssignedUser != nil {
		assignedUser = &ApplicationAssignedUser{
			ID:          model.AssignedUser.ID,
			Username:    model.AssignedUser.Username,
			DisplayName: model.AssignedUser.DisplayName,
		}
	}

	return &ApplicationModel{
		ID:         model.ID,
		ClassLevel: model.ClassLevel,

		Labels: &newLabelModel,
		Applicant: &ApplicationApplicant{
			ID:        model.Applicant.ID,
			Firstname: model.Applicant.Firstname,
			Lastname:  model.Applicant.Lastname,
			Address: &ApplicationApplicantAddress{
				Street:      model.Applicant.Address.Street,
				HouseNumber: model.Applicant.Address.HouseNumber,
				ZipCode:     model.Applicant.Address.ZipCode,
				City:        model.Applicant.Address.City,
				Country:     model.Applicant.Address.Country,
			},
			TrainingsAddress: trainingsAddress,
			ContactData: &ApplicationApplicantContactData{
				Email: model.Applicant.ContactData.Email,
				Phone: model.Applicant.ContactData.Phone,
			},
		},
		SchoolWithDegree: &ApplicationSchoolWithDegree{
			ID:   model.SchoolWithDegree.ID,
			Name: model.SchoolWithDegree.Name,
			Address: &ApplicationSchoolWithDegreeAddress{
				Street:      model.SchoolWithDegree.Address.Street,
				HouseNumber: model.SchoolWithDegree.Address.HouseNumber,
				ZipCode:     model.SchoolWithDegree.Address.ZipCode,
				City:        model.SchoolWithDegree.Address.City,
				Country:     model.SchoolWithDegree.Address.Country,
			},
			Type: &ApplicationSchoolWithDegreeType{
				ID:         model.SchoolWithDegree.Type.ID,
				Name:       model.SchoolWithDegree.Type.Name,
				Identifier: model.SchoolWithDegree.Type.Identifier,
			},
			Degree: &ApplicationSchoolWithDegreeDegree{
				ID:                                 model.SchoolWithDegree.Degree.ID,
				Name:                               model.SchoolWithDegree.Degree.Name,
				FosBerufsabschlussRequired:         model.SchoolWithDegree.Degree.FosBerufsabschlussRequired,
				BosBerufsqualifizierenderAbschluss: model.SchoolWithDegree.Degree.BosBerufsqualifizierenderAbschluss,
				FachschuleBerufsabschlussRequired:  model.SchoolWithDegree.Degree.FachschuleBerufsabschlussRequired,
			},
		},
		AssignedUser: assignedUser,
		Status: &ApplicationStatus{
			ID:         model.Status.ID,
			Name:       model.Status.Name,
			Identifier: model.Status.Identifier,
		},
		ProcessingTime: &ApplicationProcessingTime{
			MaxValidity:            model.ProcessingTime.MaxValidity,
			RemainingTimeInDays:    model.ProcessingTime.RemainingTimeInDays,
			RemainingTimeInPercent: model.ProcessingTime.RemainingTimeInPercent,
			IsStillLegal:           model.ProcessingTime.IsStillLegal,
		},

		Created: model.Created,
		Updated: model.Updated,
	}
}
