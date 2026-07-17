package main

import (
	"context"
	"errors"
	"strings"
)

func generateCommitMessage(ctx context.Context, model string, diff string) (string, error) {
	reqBody := responsesRequest{
		Model:        model,
		Instructions: commitInstructions(),
		Input:        commitInput(diff),
		MaxTokens:    400,
	}

	respBody, err := sendResponsesRequest(ctx, reqBody)
	if err != nil {
		return "", err
	}

	message, err := parseResponsesText(respBody)
	if err != nil {
		return "", err
	}
	if strings.TrimSpace(message) == "" {
		return "", errors.New("OpenAI returned an empty commit message")
	}

	return message, nil
}
