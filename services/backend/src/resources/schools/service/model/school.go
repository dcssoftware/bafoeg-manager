package model

import "github.com/google/uuid"

type SchoolModel struct {
	ID          uuid.UUID
	Name        string
	Type        *SchoolTypeModel
	Degree      []SchoolDegreeModel
	Street      string
	HouseNumber string
	City        string
	ZipCode     string
	Country     string
}

type SchoolTypeModel struct {
	Identifier string
	Name       string
}

type SchoolDegreeModel struct {
	ID                                 uuid.UUID
	Name                               string
	SchoolID                           uuid.UUID
	FosBerufsabschlussRequired         bool
	BosBerufsqualifizierenderAbschluss bool
	FachschuleBerufsschuleRequired     bool
}
