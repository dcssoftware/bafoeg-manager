package models

import "github.com/google/uuid"

type ApplicantBySchoolModel struct {
	ID        uuid.UUID
	Firstname string
	Lastname  string

	ClassLevel       string
	StatusIdentifier string
	Degree           *ApplicantBySchoolDegreeModel
	Address          *ApplicantBySchoolAddressModel
}

type ApplicantBySchoolAddressModel struct {
	ZipCode string
	City    string
	Country string
}

type ApplicantBySchoolDegreeModel struct {
	ID   uuid.UUID
	Name string
}
