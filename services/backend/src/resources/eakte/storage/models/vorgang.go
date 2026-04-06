package models

import (
	"time"

	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/eakte/service/models"
	"github.com/google/uuid"
)

type VorgangModel struct {
	ID              uuid.UUID `db:"id"`
	EakteID         uuid.UUID `db:"akte_id"`
	Vorgangszeichen string    `db:"vorgangszeichen"`
	Created         time.Time `db:"created"`
}

func (m VorgangModel) ToServiceModel() *serviceModels.VorgangModel {
	return &serviceModels.VorgangModel{
		ID:              m.ID,
		EakteID:         m.EakteID,
		Vorgangszeichen: m.Vorgangszeichen,
		Created:         m.Created,
	}
}
