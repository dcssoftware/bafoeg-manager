package models

import (
	"time"

	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/google/uuid"
)

type ApplicationShortModel struct {
	ID                    uuid.UUID `db:"id"`
	ApplicationClassLevel string    `db:"class_level"`

	// for filter purposes onlym not in use, use struct instead
	AssignedUserID   *string `db:"assigned_user_id"`
	ApplicantID      string  `db:"applicant_id"`
	StatusIdentifier string  `db:"status_identifier"`

	Applicant    ApplicationShortApplicantModel   `db:"applicant"`
	AssignedUser *ApplicationAssignedUserModel    `db:"assigned_user"`
	School       ApplicationSchoolWithDegreeModel `db:"school"`
	Status       ApplicationStatusModel           `db:"status"`

	Created time.Time `db:"application_created"`
	Updated time.Time `db:"application_updated"`
}

type ApplicationShortApplicantModel struct {
	ID        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
}

func (m *ApplicationShortModel) ToServiceShortModel() *serviceModel.ApplicationShortModel {

	model := &serviceModel.ApplicationShortModel{
		ID:         m.ID,
		ClassLevel: m.ApplicationClassLevel,

		Applicant:    m.Applicant.ToServiceModel(),
		AssignedUser: m.AssignedUser.ToServiceModel(),
		School:       m.School.ToServiceModel(),
		Status:       m.Status.ToServiceModel(),

		Created: m.Created,
		Updated: m.Updated,
	}

	return model
}

func (model ApplicationShortApplicantModel) ToServiceModel() *serviceModel.ApplicationShortApplicant {
	return &serviceModel.ApplicationShortApplicant{
		ID:        model.ID,
		Firstname: model.Firstname,
		Lastname:  model.Lastname,
	}
}
