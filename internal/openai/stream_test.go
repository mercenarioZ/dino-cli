package openai

import "testing"

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

	// delta case
	if event.Delta != "hello" {
		t.Errorf("got delta %q, want %q", event.Delta, "hello")
	}
}
