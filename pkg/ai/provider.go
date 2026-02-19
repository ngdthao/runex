package ai

import (
	"context"
	"os"

	"github.com/runex/runex/internal/detector"
)

type AIProvider interface {
	Name() string
	Analyze(ctx context.Context, err *detector.DetectedError, output string) (string, error)
}

type OpenAIProvider struct {
	apiKey string
	model  string
}

func NewOpenAIProvider() *OpenAIProvider {
	return &OpenAIProvider{
		apiKey: os.Getenv("OPENAI_API_KEY"),
		model:  "gpt-4",
	}
}

func (p *OpenAIProvider) Name() string {
	return "OpenAI"
}

func (p *OpenAIProvider) Analyze(ctx context.Context, err *detector.DetectedError, output string) (string, error) {
	return "", nil
}

type AnthropicProvider struct {
	apiKey string
	model  string
}

func NewAnthropicProvider() *AnthropicProvider {
	return &AnthropicProvider{
		apiKey: os.Getenv("ANTHROPIC_API_KEY"),
		model:  "claude-3-opus-20240229",
	}
}

func (p *AnthropicProvider) Name() string {
	return "Anthropic"
}

func (p *AnthropicProvider) Analyze(ctx context.Context, err *detector.DetectedError, output string) (string, error) {
	return "", nil
}

func GetProvider(name string) AIProvider {
	switch name {
	case "openai":
		return NewOpenAIProvider()
	case "anthropic":
		return NewAnthropicProvider()
	default:
		return NewOpenAIProvider()
	}
}
