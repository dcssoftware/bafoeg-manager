package service

import (
	"context"
	"fmt"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/artificial-intelligence/rag"
	ragModels "github.com/dcssoftware/bafoeg-manager/src/helper/artificial-intelligence/rag/models"
)

func (s *RAGService) GetRAGrequestStudierenden(prompt string) (string, error) {
	response, err := rag.RequestRAG(
		prompt,
		configuration.OllamaAPI.DatabaseTablenameRAGStudierenden,
		[]ragModels.ConversationMessage{},
		func(ctx context.Context, chunk []byte) error {
			fmt.Println(string(chunk))
			return nil
		},
	)
	return response, err
}
