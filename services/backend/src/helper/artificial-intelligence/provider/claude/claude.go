package claude

import (
	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/tmc/langchaingo/llms/anthropic"
)

func CreateClaudeConnection(opts ...anthropic.Option) (*anthropic.LLM, error) {
	// anthropicArguments := append(opts, anthropic.WithBaseURL(""))
	anthropicArguments := append(opts, anthropic.WithToken(configuration.ClaudeAPI.Token))
	anthropicArguments = append(anthropicArguments, anthropic.WithModel(configuration.ClaudeAPI.RequestingModelname))

	anthropicLLM, anthropicLLMErr := anthropic.New(
		anthropicArguments...,
	)
	if anthropicLLMErr != nil {
		return nil, anthropicLLMErr
	}

	return anthropicLLM, nil
}
