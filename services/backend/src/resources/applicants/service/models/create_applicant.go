package models

type CreateApplicantModel struct {
	Firstname string
	Lastname  string

	Street      string
	HouseNumber string
	ZipCode     string
	City        string
	Country     string
}
