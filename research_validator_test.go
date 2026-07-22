package main

import "testing"

func TestValidateResearchReport(t *testing.T) {
	tests := []struct {
		name    string
		modify  func(*ResearchReport)
		wantErr string
	}{
		{
			name:   "valid report",
			modify: func(*ResearchReport) {},
		},
		{
			name: "empty topic",
			modify: func(report *ResearchReport) {
				report.Topic = " "
			},
			wantErr: "research report topic is required",
		},
		{
			name: "empty summary",
			modify: func(report *ResearchReport) {
				report.Summary = ""
			},
			wantErr: "research summary cannot be blank",
		},
		{
			name: "no findings",
			modify: func(report *ResearchReport) {
				report.Findings = nil
			},
			wantErr: "research report required findings",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			report := validResearchReport()
			test.modify(&report)

			err := validateResearchReport(report)

			if test.wantErr == "" {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}

			if err == nil {
				t.Fatalf("expected error %q", test.wantErr)
			}

			if err.Error() != test.wantErr {
				t.Fatalf("got %q, want %q", err.Error(), test.wantErr)
			}
		})
	}
}

func validResearchReport() ResearchReport {
	return ResearchReport{
		Topic:   "Go runtime",
		Summary: "How the Go runtime works",
		Findings: []Finding{
			{
				Claim:     "Go uses the G-M-P scheduler model",
				SourceIDs: []string{"S1"},
			},
		},
		Sources: []Source{
			{
				ID:    "S1",
				Title: "Go runtime source",
				URL:   "https://go.dev/src/runtime",
			},
		},
	}
}
