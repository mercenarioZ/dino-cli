package main

import (
	"os"
	"strings"

	"github.com/mercenarioZ/dino/internal/openai"
)

const (
	defaultOpenAIModel        = "gpt-5.5"
	defaultOpenAIResponsesURL = "https://api.openai.com/v1/responses"
)

type openAIConfig struct {
	APIKey       string
	ResponsesURL string
}

func loadOpenAIConfig() openAIConfig {
	return openAIConfig{
		APIKey:       firstEnv("DINO_OPENAI_API_KEY", "OPENAI_API_KEY"),
		ResponsesURL: envOrDefault("DINO_OPENAI_RESPONSES_URL", defaultOpenAIResponsesURL),
	}
}

func newOpenAIClient() (*openai.Client, error) {
	config := loadOpenAIConfig()
	return openai.NewClient(config.APIKey, config.ResponsesURL)
}

func firstEnv(names ...string) string {
	for _, name := range names {
		if value := strings.TrimSpace(os.Getenv(name)); value != "" {
			return value
		}
	}
	return ""
}

func envOrDefault(name string, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(name)); value != "" {
		return value
	}
	return fallback
}
