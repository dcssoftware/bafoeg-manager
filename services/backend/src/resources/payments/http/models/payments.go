package models

import (
	"time"

	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/payments/service/models"
	"github.com/google/uuid"
)

type PaymentsModell struct {
	ID            uuid.UUID `json:"id"`
	ApplicantID   uuid.UUID `json:"applicantID"`
	ApplicationID uuid.UUID `json:"applicationID"`

	Amount           float64 `json:"amount"`
	StatusIdentifier string  `json:"statusIdentifier"`
	Description      string  `json:"description"`
	Iban             string  `json:"iban"`
	Bic              string  `json:"bic"`
	Direction        string  `json:"direction"`

	Executed time.Time `json:"executed"`
	Created  time.Time `json:"created"`
}

func ToHttpPaymentModel(model *serviceModels.PaymentsModell) *PaymentsModell {
	if model == nil {
		return nil
	}

	return &PaymentsModell{
		ID: model.ID,

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

func ToHttpPaymentModels(models []serviceModels.PaymentsModell) []PaymentsModell {
	var httpModels []PaymentsModell

	for _, model := range models {
		httpModels = append(httpModels, *ToHttpPaymentModel(&model))
	}

	return httpModels
}
