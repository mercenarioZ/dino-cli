package main

import (
	"context"
	"errors"
	"strings"
)

func researchTopic(ctx context.Context, model, topic string) (string, error) {
	topic = strings.TrimSpace(topic)

	if topic == "" {
		return "", errors.New("research topic is required")
	}

	request := buildResearchRequest(model, topic)

	client, err := newOpenAIClient()
	if err != nil {
		return "", err
	}

	result, err := client.CreateResponse(ctx, request)
	if err != nil {
		return "", err
	}

	result = strings.TrimSpace(result)
	if result == "" {
		return "", errors.New("research result is empty")
	}

	return result, nil
}
