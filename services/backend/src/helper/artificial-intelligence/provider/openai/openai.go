package openai

import (
	"github.com/tmc/langchaingo/llms/openai"
)

func CreateOpenAIConnection(opts ...openai.Option) (*openai.LLM, error) {
	openaiArguments := append(opts, openai.WithBaseURL(""))
	openaiArguments = append(openaiArguments, openai.WithToken(""))

	openaiLLM, openaiLLMErr := openai.New(
		openaiArguments...,
	)
	if openaiLLMErr != nil {
		return nil, openaiLLMErr
	}

	return openaiLLM, nil
}
