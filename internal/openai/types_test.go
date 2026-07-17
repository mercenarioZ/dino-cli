package openai

import "testing"

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
