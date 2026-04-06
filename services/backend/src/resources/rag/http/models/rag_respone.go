package models

import "github.com/google/uuid"

type RagResponse struct {
	ConversationID uuid.UUID `json:"conversation_id"`
	Response       string    `json:"response"`
}
