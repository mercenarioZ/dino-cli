package main

import "testing"

func TestWebSearchReq_RequiresWebSearch(t *testing.T) {
	request := buildWebSearchRequest(
		"gpt-5.6",
		"latest Go release",
	)

	if request.Model != "gpt-5.6" {
		t.Errorf("got model %q", request.Model)
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
}
