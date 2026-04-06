package models

import (
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/organization/service/models"
	"github.com/google/uuid"
)

type Region struct {
	ID         uuid.UUID `db:"id"`
	Identifier string    `db:"identifier"`
	Name       string    `db:"name"`
}

func (r Region) ToServiceModel() *serviceModel.Region {
	return &serviceModel.Region{
		ID:         r.ID,
		Identifier: r.Identifier,
		Name:       r.Name,
	}
}
