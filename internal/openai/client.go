package openai

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

type Client struct {
	apiKey       string
	responsesURL string
}

func NewClient(apiKey string, responsesURL string) (*Client, error) {
	if strings.TrimSpace(apiKey) == "" {
		return nil, errors.New("set DINO_OPENAI_API_KEY or OPENAI_API_KEY")
	}

	return &Client{
		apiKey:       apiKey,
		responsesURL: responsesURL,
	}, nil
}

func (c *Client) CreateResponse(ctx context.Context, reqBody Request) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	body, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.responsesURL, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
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

	return parseResponseText(respBody)
}
