package documentclassifier

import (
	"context"

	ai "github.com/dcssoftware/bafoeg-manager/src/helper/artificial-intelligence"
)

func ClassifyDocument(content []byte) {
	llm, llmErr := ai.CreateOllamaConnection()
	if llmErr != nil {
		return
	}

	llm.Call(context.Background(), string(content))
}
