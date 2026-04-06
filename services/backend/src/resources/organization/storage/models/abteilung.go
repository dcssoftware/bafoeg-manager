package models

import (
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/organization/service/models"
	"github.com/google/uuid"
)

type Abteilung struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	BehördeID uuid.UUID `db:"behoerde_id"`
}

func (r Abteilung) ToServiceModel() *serviceModel.Abteilung {
	return &serviceModel.Abteilung{
		ID:        r.ID,
		Name:      r.Name,
		BehördeID: r.BehördeID,
	}
}
