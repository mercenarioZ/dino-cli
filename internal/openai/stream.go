package openai

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type streamEvent struct {
	Type  string `json:"type"`
	Delta string `json:"delta,omitempty"`
}

func parseStreamEvent(data []byte) (streamEvent, error) {
	var event streamEvent

	if err := json.Unmarshal(data, &event); err != nil {
		return streamEvent{}, err
	}

	return event, nil
}

const (
	outputTextDeltaEvent   = "response.output_text.delta"
	responseCompletedEvent = "response.completed"
)

func readResponseStream(
	reader io.Reader,
	onDelta func(string),
) (string, error) {
	scanner := bufio.NewScanner(reader)

	// completed events may contain a large response
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)

	var result strings.Builder

	for scanner.Scan() {
		line := scanner.Text()

		if !strings.HasPrefix(line, "data:") {
			continue
		}

		data := strings.TrimSpace(
			strings.TrimPrefix(line, "data:"),
		)

		if data == "" {
			continue
		}

		event, err := parseStreamEvent([]byte(data))
		if err != nil {
			return "", fmt.Errorf("failed stream event: %w", err)
		}

		switch event.Type {
		case outputTextDeltaEvent:
			result.WriteString(event.Delta)
			if onDelta != nil {
				onDelta(event.Delta)
			}
		case responseCompletedEvent:
			return result.String(), nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("read response stream: %w", err)
	}

	return "", fmt.Errorf(
		"response stream ended before %s",
		responseCompletedEvent,
	)
}
