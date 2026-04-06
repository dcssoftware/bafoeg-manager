package models

import (
	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/eakte/service/models"
	"github.com/google/uuid"
)

type EaktenApplicationMappingModel struct {
	ID               uuid.UUID                                      `json:"id"`
	ApplicationID    uuid.UUID                                      `json:"application_id"`
	EakteAkteID      uuid.UUID                                      `json:"eakte_akte_id"`
	Application      *EaktenApplicationMappingApplicationShortModel `json:"application"`
	Applicant        *EaktenApplicationMappingApplicantShortModel   `json:"applicant"`
	SchoolWithDegree *EaktenApplicationMappingDegreeShortModel      `json:"school_degree"`
	School           *EaktenApplicationMappingSchoolShortModel      `json:"school"`
	Eakte_Akte       *EaktenApplicationMappingAkteShortModel        `json:"eakte_akte"`
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

func EaktenApplicationMappingToHttpResponse(model serviceModels.EaktenApplicationMappingModel) EaktenApplicationMappingModel {
	return EaktenApplicationMappingModel{
		ID:               model.ID,
		ApplicationID:    model.ApplicationID,
		EakteAkteID:      model.EakteAkteID,
		Application:      EaktenApplicationMappingApplicationShortModelToHttpResponse(model.Application),
		Applicant:        EaktenApplicationMappingApplicantShortModelToHttpResponse(model.Applicant),
		SchoolWithDegree: EaktenApplicationMappingDegreeShortModelToHttpResponse(model.SchoolWithDegree),
		School:           EaktenApplicationMappingSchoolShortModelToHttpResponse(model.School),
		Eakte_Akte:       EaktenApplicationMappingAkteShortModelToHttpResponse(model.Eakte_Akte),
	}
}

func EaktenApplicationMappingApplicationShortModelToHttpResponse(model *serviceModels.EaktenApplicationMappingApplicationShortModel) *EaktenApplicationMappingApplicationShortModel {
	return &EaktenApplicationMappingApplicationShortModel{
		ID:     model.ID,
		Status: EaktenApplicationMappingApplicationStatusShortModelToHttpResponse(model.Status),
	}
}

func EaktenApplicationMappingApplicationStatusShortModelToHttpResponse(model *serviceModels.EaktenApplicationMappingApplicationStatusShortModel) *EaktenApplicationMappingApplicationStatusShortModel {
	return &EaktenApplicationMappingApplicationStatusShortModel{
		ID:         model.ID,
		Identifier: model.Identifier,
		Name:       model.Name,
	}
}

func EaktenApplicationMappingApplicantShortModelToHttpResponse(model *serviceModels.EaktenApplicationMappingApplicantShortModel) *EaktenApplicationMappingApplicantShortModel {
	return &EaktenApplicationMappingApplicantShortModel{
		ID:        model.ID,
		FirstName: model.FirstName,
		LastName:  model.LastName,
	}
}

func EaktenApplicationMappingDegreeShortModelToHttpResponse(model *serviceModels.EaktenApplicationMappingDegreeShortModel) *EaktenApplicationMappingDegreeShortModel {
	return &EaktenApplicationMappingDegreeShortModel{
		ID:   model.ID,
		Name: model.Name,
	}
}
func EaktenApplicationMappingSchoolShortModelToHttpResponse(model *serviceModels.EaktenApplicationMappingSchoolShortModel) *EaktenApplicationMappingSchoolShortModel {
	return &EaktenApplicationMappingSchoolShortModel{
		ID:   model.ID,
		Name: model.Name,
	}
}
func EaktenApplicationMappingAkteShortModelToHttpResponse(model *serviceModels.EaktenApplicationMappingAkteShortModel) *EaktenApplicationMappingAkteShortModel {
	return &EaktenApplicationMappingAkteShortModel{
		ID:           model.ID,
		Aktenzeichen: model.Aktenzeichen,
	}
}
