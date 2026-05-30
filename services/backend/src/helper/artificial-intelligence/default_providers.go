package artificialintelligence

import (
	"errors"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"

	"github.com/tmc/langchaingo/llms"
)

func (provider *AIProvider) DefaultRequestModelInstance() (llms.Model, LlmModel, error) {
	if provider == nil {
		return nil, UndefinedModel, errors.New("ai provider is nil")
	}

	switch "claude" {
	case "ollama":
		if provider.OllamaRequest == nil {
			return nil, UndefinedModel, errors.New("ollama provider is not configured")
		}
		return provider.OllamaRequest, OllamaModel, nil

	case "claude":
		if provider.ClaudeRequest == nil {
			return nil, UndefinedModel, errors.New("claude provider is not configured")
		}
		return provider.ClaudeRequest, ClaudeModel, nil
	default:
		return nil, UndefinedModel, errors.New("no requesting llm defined")
	}
}

func (provider *AIProvider) DefaultEmbeddingModelInstance() (llms.Model, LlmModel, error) {
	if provider == nil {
		return nil, UndefinedModel, errors.New("ai provider is nil")
	}

	switch "ollama" {
	case "ollama":
		if provider.OllamaEmbedder == nil {
			return nil, UndefinedModel, errors.New("ollama provider is not configured")
		}
		return provider.OllamaEmbedder, OllamaModel, nil
	default:
		return nil, UndefinedModel, errors.New("no requesting llm defined")
	}
}

func (provider *AIProvider) DefaultRequestModelName() string {
	_, name, _ := provider.DefaultRequestModelInstance()

	switch name {
	case OllamaModel:
		return configuration.OllamaAPI.RequestingModelname
	case OpenAIModel:
		return ""
	case ClaudeModel:
		return ""
	default:
		return ""
	}
}

func (provider *AIProvider) DefaultEmbeddingModelName() string {
	_, name, _ := provider.DefaultEmbeddingModelInstance()

	switch name {
	case OllamaModel:
		return configuration.OllamaAPI.EmbeddingModelname
	case OpenAIModel:
		return ""
	case ClaudeModel:
		return ""
	default:
		return ""
	}
}

func (provider *AIProvider) DefaultRequestModelNameByLLM(model LlmModel) string {
	switch model {
	case OllamaModel:
		return configuration.OllamaAPI.RequestingModelname
	case OpenAIModel:
		return ""
	case ClaudeModel:
		return ""
	default:
		return ""
	}
}

func (provider *AIProvider) DefaultEmbeddingModelNameByLLM(model LlmModel) string {
	switch model {
	case OllamaModel:
		return configuration.OllamaAPI.EmbeddingModelname
	case OpenAIModel:
		return ""
	case ClaudeModel:
		return ""
	default:
		return ""
	}
}
