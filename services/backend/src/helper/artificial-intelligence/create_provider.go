package artificialintelligence

import (

	// bm = bafög manager

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	bmClaude "github.com/dcssoftware/bafoeg-manager/src/helper/artificial-intelligence/provider/claude"
	bmOllama "github.com/dcssoftware/bafoeg-manager/src/helper/artificial-intelligence/provider/ollama"

	// lc = langchain
	lcClaude "github.com/tmc/langchaingo/llms/anthropic"
	lcOllama "github.com/tmc/langchaingo/llms/ollama"
	lcOpenAI "github.com/tmc/langchaingo/llms/openai"
)

type LlmModel string

const (
	OllamaModel    LlmModel = "ollama"
	OpenAIModel    LlmModel = "openai"
	ClaudeModel    LlmModel = "claude"
	UndefinedModel LlmModel = "unknown"
)

func (m LlmModel) String() string {
	return string(m)
}

type AIProvider struct {
	OllamaEmbedder *lcOllama.LLM
	OllamaRequest  *lcOllama.LLM

	OpenAIRequest *lcOpenAI.LLM
	ClaudeRequest *lcClaude.LLM
}

func CreateAIProvider() (*AIProvider, error) {

	var claude *lcClaude.LLM
	var claudeErr error

	ollamaRequest, ollamaRequestErr := bmOllama.CreateOllamaConnection()
	if ollamaRequestErr != nil {
		return nil, ollamaRequestErr
	}

	ollamaEmbedd, ollamaEmbeddErr := bmOllama.CreateOllamaConnection(
		lcOllama.WithModel(configuration.OllamaAPI.EmbeddingModelname),
	)
	if ollamaEmbeddErr != nil {
		return nil, ollamaEmbeddErr
	}

	// openai, openaiErr := bmOpenai.CreateOpenAIConnection()
	// if openaiErr != nil {
	// 	return nil, openaiErr
	// }

	if configuration.ClaudeAPI.Enabled {
		claude, claudeErr = bmClaude.CreateClaudeConnection()
		if claudeErr != nil {
			return nil, claudeErr
		}
	}

	return &AIProvider{
		OllamaEmbedder: ollamaEmbedd,
		OllamaRequest:  ollamaRequest,
		OpenAIRequest:  nil,
		ClaudeRequest:  claude,
	}, nil
}
