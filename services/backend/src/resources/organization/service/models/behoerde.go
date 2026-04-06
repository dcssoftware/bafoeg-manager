package models

import "github.com/google/uuid"

type Behörde struct {
	ID       uuid.UUID
	Name     string
	RegionID uuid.UUID
}
