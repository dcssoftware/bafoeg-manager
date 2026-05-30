package rag

import (
	"context"
	"fmt"
	"testing"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	artificialintelligence "github.com/dcssoftware/bafoeg-manager/src/helper/artificial-intelligence"
	ragModels "github.com/dcssoftware/bafoeg-manager/src/helper/artificial-intelligence/rag/models"
	"github.com/stretchr/testify/assert"
)

func TestRAG(t *testing.T) {

	provider, providerErr := artificialintelligence.CreateAIProvider()
	assert.NoError(t, providerErr)

	_, err := RequestRAG(
		provider,
		"Tell me a really good joke. Think hard before you answer!",
		configuration.OllamaAPI.DatabaseTablenameRAGStudierenden,
		[]ragModels.ConversationMessage{},
		func(ctx context.Context, reasoningChunk []byte, chunk []byte) error {
			fmt.Println(string(reasoningChunk))
			fmt.Println(string(chunk))
			return nil
		},
	)

	assert.NoError(t, err)
}
