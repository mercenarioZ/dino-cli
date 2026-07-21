package main

import "github.com/mercenarioZ/dino/internal/openai"

func buildResearchRequest(model, topic string) openai.Request {
	return openai.Request{
		Model:        model,
		Instructions: researchInstructions(),
		Input:        topic,
		MaxTokens:    8192,
		Tools: []openai.Tool{
			{Type: "web_search"},
		},
		ToolChoice: "required",
	}
}

