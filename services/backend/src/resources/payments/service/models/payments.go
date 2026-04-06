package models

import (
	"time"

	"github.com/google/uuid"
)

type PaymentsModell struct {
	ID            uuid.UUID
	ApplicantID   uuid.UUID
	ApplicationID uuid.UUID

	Amount           float64
	StatusIdentifier string
	Description      string
	Iban             string
	Bic              string
	Direction        string

	Executed time.Time
	Created  time.Time
}
