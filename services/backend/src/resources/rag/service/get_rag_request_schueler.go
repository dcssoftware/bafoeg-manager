package service

import (
	"context"
	"strings"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/artificial-intelligence/rag"
	ragModels "github.com/dcssoftware/bafoeg-manager/src/helper/artificial-intelligence/rag/models"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
	"github.com/tmc/langchaingo/llms"
)

func (s *RAGService) GetRAGrequestSchüler(tx *sqlx.Tx, conversationID uuid.UUID, userID uuid.UUID, prompt string, streamFunc func(ctx context.Context, reasoningChunk []byte, chunk []byte) error) (response *llms.ContentResponse, currentConversationID uuid.UUID, err customerrors.ErrorInterface) {

	conversation, conversationErr := s.GetRagConversationByID(tx, conversationID)
	if conversationErr != nil {
		return nil, uuid.Nil, conversationErr
	}

	messages, messagesErr := s.GetRagConversationMessagesByConversationID(tx, 1, conversationID)
	if messagesErr != nil && messagesErr.ErrorType() != customerrorconst.ERROR_IDENTIFIER_DATABASE_NOT_FOUND {
		return nil, uuid.Nil, messagesErr
	}

	var ragMessages []ragModels.ConversationMessage
	for _, message := range messages {
		ragMessages = append(ragMessages, ragModels.ConversationMessage{
			Role:    message.Sender,
			Message: message.MessageContent,
		})
	}

	_, insertConversationErr := s.InsertRagConversationMessage(tx, conversation.ID, prompt, true)
	if insertConversationErr != nil {
		return response, uuid.Nil, insertConversationErr
	}

	response, responseErr := rag.RequestRAG(
		s.aiConn,
		prompt,
		configuration.OllamaAPI.DatabaseTablenameRAGSchueler,
		ragMessages,
		streamFunc,
	)
	if responseErr != nil {
		return nil, uuid.Nil, customerrors.NewAIError(responseErr, prompt)
	}

	var responseContent strings.Builder

	for _, resp := range response.Choices {
		responseContent.WriteString(resp.Content)
	}

	_, insertMessageErr := s.InsertRagConversationMessage(tx, conversation.ID, responseContent.String(), false)
	if insertMessageErr != nil {
		return response, conversation.ID, insertMessageErr
	}

	return response, conversation.ID, err
}
