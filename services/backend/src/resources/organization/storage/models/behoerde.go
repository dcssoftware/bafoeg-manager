package models

import (
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/organization/service/models"
	"github.com/google/uuid"
)

type Behörde struct {
	ID       uuid.UUID `db:"id"`
	Name     string    `db:"name"`
	RegionID uuid.UUID `db:"region_id"`
}

func (r Behörde) ToServiceModel() *serviceModel.Behörde {
	return &serviceModel.Behörde{
		ID:       r.ID,
		Name:     r.Name,
		RegionID: r.RegionID,
	}
}
