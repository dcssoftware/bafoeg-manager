package models

import "github.com/google/uuid"

type Conversation struct {
	ID         uuid.UUID
	UserID     string
	BafoegType string
}
