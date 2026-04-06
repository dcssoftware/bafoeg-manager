package model

import (
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/schools/service/model"
	"github.com/google/uuid"
)

type SchoolShortModel struct {
	ID                   uuid.UUID `db:"id"`
	Name                 string    `db:"name"`
	Street               string    `db:"street"`
	HouseNumber          string    `db:"house_number"`
	City                 string    `db:"city"`
	ZipCode              string    `db:"zip_code"`
	Country              string    `db:"country"`
	SchoolTypeName       string    `db:"school_type_name"`
	SchoolTypeIdentifier string    `db:"school_type_identifier"`
}

func (m *SchoolShortModel) ToServiceShortModel() *serviceModel.SchoolShortModel {
	if m == nil {
		return nil
	}

	return &serviceModel.SchoolShortModel{
		ID:   m.ID,
		Name: m.Name,
		Type: &serviceModel.SchoolTypeModel{
			Name:       m.SchoolTypeName,
			Identifier: m.SchoolTypeIdentifier,
		},
		Street:      m.Street,
		HouseNumber: m.HouseNumber,
		City:        m.City,
		ZipCode:     m.ZipCode,
		Country:     m.Country,
	}
}
