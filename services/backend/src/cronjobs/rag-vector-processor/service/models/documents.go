package models

import (
	"time"

	"github.com/google/uuid"
)

type DocumentModel struct {
	ID          uuid.UUID
	Filename    string
	FileType    string
	FileSize    float64
	Status      DocumentStatusModel
	Created     time.Time
	CreatedFrom DocumentCreatedFromModel
}

type DocumentCreatedFromModel struct {
	ID          uuid.UUID
	DisplayName string
	Username    string
}

type DocumentStatusModel struct {
	ID         uuid.UUID `json:"id"`
	Identifier string    `json:"identifier"`
	Name       string    `json:"name"`
}
