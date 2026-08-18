// package service

// import (
// 	"context"
// 	"fmt"

// 	"github.com/dcssoftware/bafoeg-manager/src/configuration"
// 	"github.com/dcssoftware/bafoeg-manager/src/helper/artificial-intelligence/rag"
// 	ragModels "github.com/dcssoftware/bafoeg-manager/src/helper/artificial-intelligence/rag/models"
// )

// func (s *RAGService) GetRAGrequestStudierenden(prompt string) (string, error) {
// 	ctx := context.Background()
// 	response, err := rag.RequestRAG(
// 		ctx,
// 		prompt,
// 		configuration.OllamaAPI.DatabaseTablenameRAGStudierenden,
// 		[]ragModels.ConversationMessage{},
// 		func(ctx context.Context, chunk []byte) error {
// 			fmt.Println(string(chunk))
// 			return nil
// 		},
// 	)
// 	return response, err
// }

package service

import (
	"context"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/artificial-intelligence/rag"
	ragModels "github.com/dcssoftware/bafoeg-manager/src/helper/artificial-intelligence/rag/models"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *RAGService) GetRAGrequestStudierenden(ctx context.Context, tx *sqlx.Tx, conversationID uuid.UUID, userID uuid.UUID, prompt string, streamFunc func(ctx context.Context, chunk []byte) error) (response string, currentConversationID uuid.UUID, err customerrors.ErrorInterface) {

	conversation, conversationErr := s.GetRagConversationByID(tx, conversationID)
	if conversationErr != nil {
		return "", uuid.Nil, conversationErr
	}

	messages, messagesErr := s.GetRagConversationMessagesByConversationID(tx, 1, conversationID)
	if messagesErr != nil && messagesErr.ErrorType() != customerrorconst.ERROR_IDENTIFIER_DATABASE_NOT_FOUND {
		return "", uuid.Nil, messagesErr
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
		ctx,
		prompt,
		configuration.OllamaAPI.DatabaseTablenameRAGStudierenden,
		ragMessages,
		streamFunc,
	)

	if responseErr != nil {
		return "", uuid.Nil, customerrors.NewAIError(responseErr, prompt)
	}

	_, insertMessageErr := s.InsertRagConversationMessage(tx, conversation.ID, response, false)
	if insertMessageErr != nil {
		return response, conversation.ID, insertMessageErr
	}

	return response, conversation.ID, err
}
