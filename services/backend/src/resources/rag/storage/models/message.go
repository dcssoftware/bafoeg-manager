package models

import (
	"time"

	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/rag/service/models"
	"github.com/google/uuid"
)

type ConversationMessage struct {
	ID             uuid.UUID `db:"id"`
	ConversationID uuid.UUID `db:"conversation_id"`
	UserID         uuid.UUID `db:"user_id"`
	BafögType      string    `db:"bafoeg_type"`
	MessageContent string    `db:"message"`
	Sender         string    `db:"sender"`
	Created        time.Time `db:"created"`
}

func (m *ConversationMessage) ToServiceConversationMessageModel() *serviceModel.ConversationMessage {
	return &serviceModel.ConversationMessage{
		ID:             m.ID,
		ConversationID: m.ConversationID,
		UserID:         m.UserID,
		BafögType:      m.BafögType,
		MessageContent: m.MessageContent,
		Sender:         m.Sender,
	}
}
