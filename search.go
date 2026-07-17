package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/mercenarioZ/dino/internal/openai"
)

func buildWebSearchRequest(model string, query string) openai.Request {
	return openai.Request{
		Model:        model,
		Instructions: "Search the live web before answering, include sources!",
		Input:        query,
		MaxTokens:    2048,
		Tools: []openai.Tool{
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

	client, err := newOpenAIClient()
	if err != nil {
		return "", err
	}

	ans, err := client.CreateResponse(ctx, request)
	if err != nil {
		return "", err
	}

	if strings.TrimSpace(ans) == "" {
		return "", errors.New("OpenAI returned an empty search result")
	}

	return strings.TrimSpace(ans), nil
}

func runSearch(args []string) error {
	fs := flag.NewFlagSet("search", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	model := fs.String("model", defaultOpenAIModel, "OpenAi model to use")

	fs.Usage = func() {
		fmt.Fprintln(fs.Output(), "Usage: dino search [--model MODEL] QUERY")

		fs.PrintDefaults()
	}

	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil
		}

		return err
	}

	query := strings.TrimSpace(strings.Join(fs.Args(), " "))

	if query == "" {
		return errors.New("search query is required")
	}

	loading := startSpinner("searching the web")
	answer, err := searchWeb(context.Background(), *model, query)
	loading.stop()

	if err != nil {
		return err
	}

	fmt.Printf("%s\n\n", answer)

	return nil
}
