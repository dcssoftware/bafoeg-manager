package models

type CreateApplicantModel struct {
	Firstname string `json:"firstname" xml:"firstname" form:"firstname"`
	Lastname  string `json:"lastname" xml:"lastname" form:"lastname"`

	Street      string `json:"street" xml:"street" form:"street"`
	HouseNumber string `json:"houseNumber" xml:"houseNumber" form:"houseNumber"`
	ZipCode     string `json:"zipCode" xml:"zipCode" form:"zipCode"`
	City        string `json:"city" xml:"city" form:"city"`
	Country     string `json:"country" xml:"country" form:"country"`
}
