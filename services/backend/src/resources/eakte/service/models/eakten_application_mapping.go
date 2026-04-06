package models

import "github.com/google/uuid"

type EaktenApplicationMappingModel struct {
	ID               uuid.UUID
	ApplicationID    uuid.UUID
	EakteAkteID      uuid.UUID
	Application      *EaktenApplicationMappingApplicationShortModel
	Applicant        *EaktenApplicationMappingApplicantShortModel
	SchoolWithDegree *EaktenApplicationMappingDegreeShortModel
	School           *EaktenApplicationMappingSchoolShortModel
	Eakte_Akte       *EaktenApplicationMappingAkteShortModel
}

type EaktenApplicationMappingApplicationShortModel struct {
	ID     uuid.UUID
	Status *EaktenApplicationMappingApplicationStatusShortModel
}

type EaktenApplicationMappingApplicationStatusShortModel struct {
	ID         uuid.UUID
	Identifier string
	Name       string
}

type EaktenApplicationMappingApplicantShortModel struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
}

type EaktenApplicationMappingDegreeShortModel struct {
	ID   uuid.UUID
	Name string
}

type EaktenApplicationMappingSchoolShortModel struct {
	ID   uuid.UUID
	Name string
}

type EaktenApplicationMappingAkteShortModel struct {
	ID           uuid.UUID
	Aktenzeichen string
}
