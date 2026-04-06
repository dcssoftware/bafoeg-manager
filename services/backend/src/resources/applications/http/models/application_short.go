package models

import (
	"time"

	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/google/uuid"
)

type ApplicationShortModels struct {
	Count    uint `json:"count"`
	MaxCount uint `json:"maxCount"`

	Application []ApplicationShortModel `json:"application"`
}

type ApplicationShortModel struct {
	ID         uuid.UUID `json:"id"`
	ClassLevel string    `json:"classLevel"`

	Applicant      *ApplicationShortApplicant   `json:"applicant"`
	AssignedUser   *ApplicationAssignedUser     `json:"assignedUser"`
	School         *ApplicationSchoolWithDegree `json:"school"`
	Status         *ApplicationStatus           `json:"status"`
	ProcessingTime *ApplicationProcessingTime   `json:"processingTime"`

	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type ApplicationShortApplicant struct {
	ID        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
}

func ToShortHttpModel(model *serviceModel.ApplicationShortModel) *ApplicationShortModel {
	if model == nil {
		return &ApplicationShortModel{}
	}

	var assignedUser *ApplicationAssignedUser
	if model.AssignedUser != nil {
		assignedUser = &ApplicationAssignedUser{
			ID:          model.AssignedUser.ID,
			Username:    model.AssignedUser.Username,
			DisplayName: model.AssignedUser.DisplayName,
		}
	}

	return &ApplicationShortModel{
		ID:         model.ID,
		ClassLevel: model.ClassLevel,

		Applicant: &ApplicationShortApplicant{
			ID:        model.Applicant.ID,
			Firstname: model.Applicant.Firstname,
			Lastname:  model.Applicant.Lastname,
		},

		AssignedUser: assignedUser,

		School: &ApplicationSchoolWithDegree{
			ID:   model.School.ID,
			Name: model.School.Name,
			Address: &ApplicationSchoolWithDegreeAddress{
				Street:      model.School.Address.Street,
				HouseNumber: model.School.Address.HouseNumber,
				ZipCode:     model.School.Address.ZipCode,
				City:        model.School.Address.City,
				Country:     model.School.Address.Country,
			},
			Type: &ApplicationSchoolWithDegreeType{
				ID:         model.School.Type.ID,
				Name:       model.School.Type.Name,
				Identifier: model.School.Type.Identifier,
			},
			Degree: &ApplicationSchoolWithDegreeDegree{
				ID:                                 model.School.Degree.ID,
				Name:                               model.School.Degree.Name,
				FosBerufsabschlussRequired:         model.School.Degree.FosBerufsabschlussRequired,
				BosBerufsqualifizierenderAbschluss: model.School.Degree.BosBerufsqualifizierenderAbschluss,
				FachschuleBerufsabschlussRequired:  model.School.Degree.FachschuleBerufsabschlussRequired,
			},
		},

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

func ToShortHttpModels(models []serviceModel.ApplicationShortModel, maxCount uint) ApplicationShortModels {

	var httpModels []ApplicationShortModel
	for _, m := range models {
		httpModels = append(httpModels, *ToShortHttpModel(&m))
	}

	return ApplicationShortModels{
		MaxCount:    maxCount,
		Count:       uint(len(httpModels)),
		Application: httpModels,
	}
}
