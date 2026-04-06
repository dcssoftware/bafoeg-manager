package rag

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func ReadPDFWithOCRModel(readingocrOllamaLLM *ollama.LLM, imageInput []byte) (string, error) {

	result, err := readingocrOllamaLLM.GenerateContent(
		context.Background(),
		[]llms.MessageContent{
			{
				Role: "user",
				Parts: []llms.ContentPart{
					llms.BinaryPart("image/png", imageInput),
					llms.TextPart(
						fmt.Sprintf("%v\n<|grounding|>Free OCR.", imageInput),
					),
				},
			},
		},
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate content with OCR model: %w", err)
	}

	// read output text and return it as string
	if len(result.Choices) == 0 {
		return "", fmt.Errorf("no choices returned from OCR model")
	}

	return result.Choices[0].Content, nil
}
