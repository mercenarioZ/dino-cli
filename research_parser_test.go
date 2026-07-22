package main

import "testing"

func TestParseResearchReport(t *testing.T) {
	raw := `{
		"topic": "Server-Sent Events",
		"summary": "SSE streams server updates over HTTP.",
		"findings": [
			{
				"claim": "SSE uses a long-lived HTTP response.",
				"source_ids": ["S1"]
			}
		],
		"conflicts": [],
		"angles": ["Understanding streaming as incremental HTTP"],
		"sources": [
			{
				"id": "S1",
				"title": "HTML Standard",
				"url": "https://html.spec.whatwg.org/"
			}
		]
	}`

	report, err := parseResearchReport(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if report.Topic != "Server-Sent Events" {
		t.Errorf(
			"got topic %q, want %q",
			report.Topic,
			"Server-Sent Events",
		)
	}
	if len(report.Findings) != 1 {
		t.Fatalf("got %d findings, want 1", len(report.Findings))
	}

	if report.Findings[0].SourceIDs[0] != "S1" {
		t.Errorf(
			"got source ID %q, want %q",
			report.Findings[0].SourceIDs[0],
			"S1",
		)
	}
}
