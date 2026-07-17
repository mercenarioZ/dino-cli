package openai

import (
	"encoding/json"
	"strings"
)

type Tool struct {
	Type string `json:"type"`
}

type Request struct {
	Model        string `json:"model"`
	Instructions string `json:"instructions"`
	Input        string `json:"input"`
	MaxTokens    int    `json:"max_output_tokens"`
	Tools        []Tool `json:"tools,omitempty"`
	ToolChoice   string `json:"tool_choice,omitempty"`
}

type response struct {
	OutputText string `json:"output_text"`
	Output     []struct {
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
	} `json:"output"`
}

func parseResponseText(body []byte) (string, error) {
	var parsed response

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
