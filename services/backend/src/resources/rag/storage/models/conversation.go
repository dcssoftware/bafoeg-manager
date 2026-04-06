package models

import (
	"time"

	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/rag/service/models"
	"github.com/google/uuid"
)

type Conversation struct {
	ID         uuid.UUID `db:"id"`
	UserID     string    `db:"user_id"`
	BafoegType string    `db:"bafoeg_type"`
	Created    time.Time `db:"created"`
}

func (c Conversation) ToServiceConversationModel() *serviceModel.Conversation {
	return &serviceModel.Conversation{
		ID:         c.ID,
		UserID:     c.UserID,
		BafoegType: c.BafoegType,
	}
}
