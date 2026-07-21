package main

import (
	"context"
	"testing"
)

func TestBuildResearchRequest(t *testing.T) {
	request := buildResearchRequest(
		"gpt-5.6",
		"latest Go release",
	)

	if request.Model != "gpt-5.6" {
		t.Errorf("got model %q, want %q", request.Model, "gpt-5.6")
	}

	if request.Input != "latest Go release" {
		t.Errorf("got input %q", request.Input)
	}

	if len(request.Tools) != 1 {
		t.Fatalf("got %d tools, want 1", len(request.Tools))
	}

	if request.Tools[0].Type != "web_search" {
		t.Errorf("got tool %q, want web_search", request.Tools[0].Type)
	}

	if request.ToolChoice != "required" {
		t.Errorf("got tool choice %q, want required", request.ToolChoice)
	}

	if request.Instructions == "" {
		t.Error("research instructions must not be empty")
	}
}

func TestResearchTopicRequiresTopic(t *testing.T) {
	_, err := researchTopic(
		context.Background(),
		"gpt-5.6",
		" ",
	)

	if err == nil {
		t.Fatal("expect an error for empty topic")
	}

	if err.Error() != "research topic is required" {
		t.Errorf("got error %q, want %q", err.Error(), "research topic is required")
	}
}
