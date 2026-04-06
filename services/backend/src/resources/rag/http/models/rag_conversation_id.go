package models

import "github.com/google/uuid"

type RAGConversationIDResponseModel struct {
	ConversationID string `json:"ID"`
}

func ToHttpRAGConversationIDResponseModel(conversationID uuid.UUID) *RAGConversationIDResponseModel {
	return &RAGConversationIDResponseModel{
		ConversationID: conversationID.String(),
	}
}
