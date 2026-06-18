package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func generateCommitMessage(ctx context.Context, model string, diff string) (string, error) {
	config := loadOpenAIConfig()
	if config.APIKey == "" {
		return "", errors.New("set DINO_OPENAI_API_KEY or OPENAI_API_KEY")
	}

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	reqBody := responsesRequest{
		Model:        model,
		Instructions: commitInstructions(),
		Input:        commitInput(diff),
		MaxTokens:    240,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, config.ResponsesURL, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+config.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("OpenAI API request failed with status %s: %s", resp.Status, strings.TrimSpace(string(respBody)))
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

type responsesRequest struct {
	Model        string `json:"model"`
	Instructions string `json:"instructions"`
	Input        string `json:"input"`
	MaxTokens    int    `json:"max_output_tokens"`
}

type responsesResponse struct {
	OutputText string `json:"output_text"`
	Output     []struct {
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
	} `json:"output"`
}

func parseResponsesText(body []byte) (string, error) {
	var parsed responsesResponse
	if err := json.Unmarshal(body, &parsed); err != nil {
		return "", err
	}

	if parsed.OutputText != "" {
		return parsed.OutputText, nil
	}

	var parts []string
	for _, output := range parsed.Output {
		for _, content := range output.Content {
			if content.Text != "" {
				parts = append(parts, content.Text)
			}
		}
	}

	return strings.Join(parts, "\n"), nil
}
