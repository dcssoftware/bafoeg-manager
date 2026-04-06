package models

import (
	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/eakte/service/models"
	"github.com/google/uuid"
)

type EaktenApplicationMappingModel struct {
	ID               uuid.UUID                                      `db:"id"`
	ApplicationID    uuid.UUID                                      `db:"application_id"`
	EakteAkteID      uuid.UUID                                      `db:"eakte_akte_id"`
	Application      *EaktenApplicationMappingApplicationShortModel `db:"application"`
	Applicant        *EaktenApplicationMappingApplicantShortModel   `db:"applicant"`
	SchoolWithDegree *EaktenApplicationMappingDegreeShortModel      `db:"school_degree"`
	School           *EaktenApplicationMappingSchoolShortModel      `db:"school"`
	Eakte_Akte       *EaktenApplicationMappingAkteShortModel        `db:"eakte_akte"`
}

type EaktenApplicationMappingApplicationShortModel struct {
	ID     uuid.UUID                                            `json:"id"`
	Status *EaktenApplicationMappingApplicationStatusShortModel `json:"status"`
}

type EaktenApplicationMappingApplicationStatusShortModel struct {
	ID         uuid.UUID `json:"id"`
	Identifier string    `json:"identifier"`
	Name       string    `json:"name"`
}

type EaktenApplicationMappingApplicantShortModel struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
}

type EaktenApplicationMappingDegreeShortModel struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type EaktenApplicationMappingSchoolShortModel struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type EaktenApplicationMappingAkteShortModel struct {
	ID           uuid.UUID `json:"id"`
	Aktenzeichen string    `json:"aktenzeichen"`
}

func (model EaktenApplicationMappingModel) ToServiceModel() *serviceModels.EaktenApplicationMappingModel {
	return &serviceModels.EaktenApplicationMappingModel{
		ID:               model.ID,
		ApplicationID:    model.ApplicationID,
		EakteAkteID:      model.EakteAkteID,
		Application:      model.Application.ToServiceModel(),
		Applicant:        model.Applicant.ToServiceModel(),
		SchoolWithDegree: model.SchoolWithDegree.ToServiceModel(),
		School:           model.School.ToServiceModel(),
		Eakte_Akte:       model.Eakte_Akte.ToServiceModel(),
	}
}

func (e *EaktenApplicationMappingApplicantShortModel) ToServiceModel() *serviceModels.EaktenApplicationMappingApplicantShortModel {
	return &serviceModels.EaktenApplicationMappingApplicantShortModel{
		ID:        e.ID,
		FirstName: e.FirstName,
		LastName:  e.LastName,
	}
}
func (e *EaktenApplicationMappingApplicationShortModel) ToServiceModel() *serviceModels.EaktenApplicationMappingApplicationShortModel {
	return &serviceModels.EaktenApplicationMappingApplicationShortModel{
		ID:     e.ID,
		Status: e.Status.ToServiceModel(),
	}
}
func (e *EaktenApplicationMappingApplicationStatusShortModel) ToServiceModel() *serviceModels.EaktenApplicationMappingApplicationStatusShortModel {
	return &serviceModels.EaktenApplicationMappingApplicationStatusShortModel{
		ID:         e.ID,
		Identifier: e.Identifier,
		Name:       e.Name,
	}
}
func (e *EaktenApplicationMappingDegreeShortModel) ToServiceModel() *serviceModels.EaktenApplicationMappingDegreeShortModel {
	return &serviceModels.EaktenApplicationMappingDegreeShortModel{
		ID:   e.ID,
		Name: e.Name,
	}
}
func (e *EaktenApplicationMappingSchoolShortModel) ToServiceModel() *serviceModels.EaktenApplicationMappingSchoolShortModel {
	return &serviceModels.EaktenApplicationMappingSchoolShortModel{
		ID:   e.ID,
		Name: e.Name,
	}
}
func (e *EaktenApplicationMappingAkteShortModel) ToServiceModel() *serviceModels.EaktenApplicationMappingAkteShortModel {
	return &serviceModels.EaktenApplicationMappingAkteShortModel{
		ID:           e.ID,
		Aktenzeichen: e.Aktenzeichen,
	}
}
