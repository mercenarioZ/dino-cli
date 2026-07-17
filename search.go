package main

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
