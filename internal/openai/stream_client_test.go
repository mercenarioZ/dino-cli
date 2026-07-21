package openai

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateResponseStream(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var request Request
			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				t.Errorf("decode request: %v", err)
				return
			}

			if !request.Stream {
				t.Error("expected stream to be enabled")
			}

			if r.Header.Get("Accept") != "text/event-stream" {
				t.Errorf(
					"got Accept %q",
					r.Header.Get("Accept"),
				)
			}

			w.Header().Set("Content-Type", "text/event-stream")

			fmt.Fprintln(
				w,
				`data: {"type":"response.output_text.delta","delta":"Hello "}`,
			)
			fmt.Fprintln(w)

			fmt.Fprintln(
				w,
				`data: {"type":"response.output_text.delta","delta":"world"}`,
			)
			fmt.Fprintln(w)

			fmt.Fprintln(
				w,
				`data: {"type":"response.completed"}`,
			)
			fmt.Fprintln(w)
		},
	))
	defer server.Close()

	client, err := NewClient("test-key", server.URL)
	if err != nil {
		t.Fatalf("create client: %v", err)
	}

	got, err := client.CreateResponseStream(
		context.Background(),
		Request{
			Model: "gpt-5.5",
			Input: "test",
		},
		nil,
	)
	if err != nil {
		t.Fatalf("create streaming response: %v", err)
	}

	if got != "Hello world" {
		t.Errorf("got %q, want %q", got, "Hello world")
	}
}
