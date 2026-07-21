package openai

import (
	"strings"
	"testing"
)

func TestParseStreamEventReturnsTextDelta(t *testing.T) {
	data := []byte(`{
	"type": "response.output_text.delta",
	"delta": "hello"
	}`)

	event, err := parseStreamEvent(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if event.Type != "response.output_text.delta" {
		t.Errorf(
			"got type %q, want %q",
			event.Type,
			"response.output_text.delta",
		)
	}

	if event.Delta != "hello" {
		t.Errorf("got delta %q, want %q", event.Delta, "hello")
	}
}

func TestReadResponseStreamJoinsTextDeltas(t *testing.T) {
	input := strings.Join([]string{
		`event: response.created`,
		`data: {"type":"response.created"}`,
		``,
		`event: response.output_text.delta`,
		`data: {"type":"response.output_text.delta","delta":"Hello "}`,
		``,
		`event: response.output_text.delta`,
		`data: {"type":"response.output_text.delta","delta":"world"}`,
		``,
		`event: response.completed`,
		`data: {"type":"response.completed"}`,
		``,
	}, "\n")

	var received strings.Builder

	got, err := readResponseStream(
		strings.NewReader(input),
		func(delta string) {
			received.WriteString(delta)
		},
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := "Hello world"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if received.String() != want {
		t.Errorf(
			"callback got %q, want %q",
			received.String(),
			want,
		)
	}
}
