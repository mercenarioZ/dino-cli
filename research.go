package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
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

func runResearch(args []string) error {
	fs := flag.NewFlagSet("research", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	/*
	* fs.String returns a pointer, so later we need to use *model to get the model value
	 */
	model := fs.String("model", defaultOpenAIModel, "OpenAI model to use")

	fs.Usage = func() {
		fmt.Fprintln(fs.Output(), "Usage: dino research [--model MODEL] TOPIC")

		fs.PrintDefaults()
	}

	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil
		}

		return err
	}

	topic := strings.TrimSpace(strings.Join(fs.Args(), " "))
	if topic == "" {
		return errors.New("research topic is required")
	}

	loading := startSpinner("researching...")
	result, err := researchTopic(context.Background(), *model, topic)

	loading.stop()

	if err != nil {
		return err
	}

	fmt.Printf("%s\n\n", result)

	return nil
}
