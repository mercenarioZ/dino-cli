package openai

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestParseResponseTextReturnsOutputText(t *testing.T) {
	body := []byte(`{
		"output_text": "feat: add login"
	}`)

	got, err := parseResponseText(body)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := "feat: add login"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRequestMarshalsStream(t *testing.T) {
	body, err := json.Marshal(Request{
		Model:  "gpt-5.5",
		Input:  "test",
		Stream: true,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(string(body), `"stream":true`) {
		t.Errorf("got body %s, want stream enabled", body)
	}
}
