package models

import "github.com/google/uuid"

type Abteilung struct {
	ID        uuid.UUID
	Name      string
	BehördeID uuid.UUID
}
