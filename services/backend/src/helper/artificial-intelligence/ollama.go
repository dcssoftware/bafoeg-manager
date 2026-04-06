package rag

import (
	"fmt"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/tmc/langchaingo/llms/ollama"
)

func CreateOllamaConnection(opts ...ollama.Option) (*ollama.LLM, error) {
	var ollamaProtocol string = "http"
	if configuration.OllamaAPI.IsSecure {
		ollamaProtocol = "https"
	}
	ollamaURL := fmt.Sprintf("%s://%s:%d", ollamaProtocol, configuration.OllamaAPI.Address, configuration.OllamaAPI.Port)

	ollamaArguments := append(opts, ollama.WithServerURL(ollamaURL))
	ollamaArguments = append(ollamaArguments, ollama.WithHTTPClient(CreateCustomOllamaHTTPClient()))

	ollamaLLM, ollamaLLMErr := ollama.New(
		ollamaArguments...,
	)
	if ollamaLLMErr != nil {
		return nil, ollamaLLMErr
	}

	return ollamaLLM, nil
}
