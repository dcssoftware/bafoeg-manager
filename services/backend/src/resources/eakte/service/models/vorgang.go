package models

import (
	"time"

	"github.com/google/uuid"
)

type VorgangModel struct {
	ID              uuid.UUID
	EakteID         uuid.UUID
	Vorgangszeichen string
	Created         time.Time
}
