package main

import (
	"context"
	"errors"
	"strings"

	"github.com/mercenarioZ/dino/internal/openai"
)

func generateCommitMessage(ctx context.Context, model string, diff string) (string, error) {
	reqBody := openai.Request{
		Model:        model,
		Instructions: commitInstructions(),
		Input:        commitInput(diff),
		MaxTokens:    400,
	}

	client, err := newOpenAIClient()
	if err != nil {
		return "", err
	}

	message, err := client.CreateResponse(ctx, reqBody)
	if err != nil {
		return "", err
	}
	if strings.TrimSpace(message) == "" {
		return "", errors.New("OpenAI returned an empty commit message")
	}

	return message, nil
}
