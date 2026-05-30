package documentclassifier

import (
	"context"

	ollamaProvider "github.com/dcssoftware/bafoeg-manager/src/helper/artificial-intelligence/provider/ollama"
)

func ClassifyDocument(content []byte) {
	llm, llmErr := ollamaProvider.CreateOllamaConnection()
	if llmErr != nil {
		return
	}

	llm.Call(context.Background(), string(content))
}
