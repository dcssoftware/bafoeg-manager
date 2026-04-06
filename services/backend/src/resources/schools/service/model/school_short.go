package model

import "github.com/google/uuid"

type SchoolShortModel struct {
	ID          uuid.UUID
	Name        string
	Type        *SchoolTypeModel
	Street      string
	HouseNumber string
	City        string
	ZipCode     string
	Country     string
}
