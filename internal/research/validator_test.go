package research

import "testing"

func TestValidateReport(t *testing.T) {
	tests := []struct {
		name    string
		modify  func(*Report)
		wantErr string
	}{
		{
			name:   "valid report",
			modify: func(*Report) {},
		},
		{
			name: "empty topic",
			modify: func(report *Report) {
				report.Topic = " "
			},
			wantErr: "research report topic is required",
		},
		{
			name: "empty summary",
			modify: func(report *Report) {
				report.Summary = ""
			},
			wantErr: "research summary cannot be blank",
		},
		{
			name: "no findings",
			modify: func(report *Report) {
				report.Findings = nil
			},
			wantErr: "research report required findings",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			report := validReport()
			test.modify(&report)

			err := ValidateReport(report)

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

func validReport() Report {
	return Report{
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
