package models

import "github.com/google/uuid"

type ConversationMessage struct {
	ID             uuid.UUID
	ConversationID uuid.UUID
	UserID         uuid.UUID
	BafögType      string
	MessageContent string
	Sender         string
}
