package main

import (
	"context"
	"errors"
	"strings"
)

func buildWebSearchRequest(model string, query string) responsesRequest {
	return responsesRequest{
		Model:        model,
		Instructions: "Search the live web before answering, include sources!",
		Input:        query,
		MaxTokens:    2048,
		Tools: []responseTool{
			{Type: "web_search"},
		},
		ToolChoice: "required",
	}
}

func searchWeb(ctx context.Context, model string, query string) (string, error) {
	query = strings.TrimSpace(query)

	if query == "" {
		return "", errors.New("search query is required")
	}

	request := buildWebSearchRequest(model, query)

	body, err := sendResponsesRequest(ctx, request)

	if err != nil {
		return "", err
	}

	ans, err := parseResponsesText(body)

	if err != nil {
		return "", err
	}

	if strings.TrimSpace(ans) == "" {
		return "", errors.New("OpenAI returned an empty search result")
	}

	return strings.TrimSpace(ans), nil
}
