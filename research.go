package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/mercenarioZ/dino/internal/research"
)

func researchTopic(ctx context.Context, model, topic string) (research.Report, error) {
	topic = strings.TrimSpace(topic)

	if topic == "" {
		return research.Report{}, errors.New("research topic is required")
	}

	request := buildResearchRequest(model, topic)

	client, err := newOpenAIClient()
	if err != nil {
		return research.Report{}, err
	}

	raw, err := client.CreateResponseStream(ctx, request, nil)
	if err != nil {
		return research.Report{}, err
	}

	report, err := research.ParseReport(raw)
	if err != nil {
		return research.Report{}, err
	}

	if err := research.ValidateReport(report); err != nil {
		return research.Report{}, fmt.Errorf(
			"validate research report: %w",
			err,
		)
	}

	return report, nil
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

	loading := startSpinner(
		"researching...",
	)
	report, err := researchTopic(context.Background(), *model, topic)
	loading.stop()

	if err != nil {
		return err
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(report); err != nil {
		return err
	}

	fmt.Fprintln(os.Stderr, "Research complete!")

	return nil
}
