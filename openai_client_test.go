package main

import "testing"

func TestParseResponsesText_ReturnsOutputText(t *testing.T) {
	body := []byte(`{
		"output_text": "feat: add login"
	}`)

	got, err := parseResponsesText(body)

	// assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := "feat: add login"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
