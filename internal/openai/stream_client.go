package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const streamResponseTimeout = 5 * time.Minute

func (c *Client) CreateResponseStream(
	ctx context.Context,
	reqBody Request,
	onDelta func(string),
) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, streamResponseTimeout)
	defer cancel()

	reqBody.Stream = true

	body, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.responsesURL,
		bytes.NewReader(body),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBody, readErr := io.ReadAll(resp.Body)

		if readErr != nil {
			return "", readErr
		}

		return "", fmt.Errorf(
			"OpenAI API failed with status %s: %s",
			resp.Status,
			strings.TrimSpace(string(respBody)),
		)
	}

	return readResponseStream(resp.Body, onDelta)
}
