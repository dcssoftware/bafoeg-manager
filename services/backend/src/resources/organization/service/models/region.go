package models

import "github.com/google/uuid"

type Region struct {
	ID         uuid.UUID
	Identifier string
	Name       string
}
