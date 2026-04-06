package models

import (
	"time"

	serviceModell "github.com/dcssoftware/bafoeg-manager/src/resources/payments/service/models"
	"github.com/google/uuid"
)

type PaymentsModell struct {
	ID            uuid.UUID `db:"id"`
	ApplicantID   uuid.UUID `db:"applicant_id"`
	ApplicationID uuid.UUID `db:"application_id"`

	Amount           float64   `db:"amount"`
	StatusID         uuid.UUID `db:"status_id"`
	StatusIdentifier string    `db:"status_identifier"`
	Description      string    `db:"description"`
	Iban             string    `db:"iban"`
	Bic              string    `db:"bic"`
	Direction        string    `db:"direction"`

	Executed time.Time `db:"executed"`
	Created  time.Time `db:"created"`
}

func ToPaymentsServiceModels(models []PaymentsModell) []serviceModell.PaymentsModell {
	var serviceModel []serviceModell.PaymentsModell
	for _, model := range models {
		serviceModel = append(serviceModel, *model.ToPaymentsServiceModel())
	}
	return serviceModel
}

func (model PaymentsModell) ToPaymentsServiceModel() *serviceModell.PaymentsModell {
	return &serviceModell.PaymentsModell{
		ID:            model.ID,
		ApplicantID:   model.ApplicantID,
		ApplicationID: model.ApplicationID,

		Amount:           model.Amount,
		StatusIdentifier: model.StatusIdentifier,
		Description:      model.Description,
		Iban:             model.Iban,
		Bic:              model.Bic,
		Direction:        model.Direction,

		Executed: model.Executed,
		Created:  model.Created,
	}
}
