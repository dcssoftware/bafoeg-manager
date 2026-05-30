package service

import (
	"context"
	"fmt"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/artificial-intelligence/rag"
	ragModels "github.com/dcssoftware/bafoeg-manager/src/helper/artificial-intelligence/rag/models"
	"github.com/tmc/langchaingo/llms"
)

func (s *RAGService) GetRAGrequestStudierenden(prompt string) (*llms.ContentResponse, error) {
	response, err := rag.RequestRAG(
		s.aiConn,
		prompt,
		configuration.OllamaAPI.DatabaseTablenameRAGStudierenden,
		[]ragModels.ConversationMessage{},
		func(ctx context.Context, reasoningChunk []byte, chunk []byte) error {
			fmt.Println(string(chunk))
			return nil
		},
	)
	return response, err
}
