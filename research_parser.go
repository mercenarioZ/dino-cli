package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func parseResearchReport(raw string) (ResearchReport, error) {
	var report ResearchReport

	raw = strings.TrimSpace(raw)
	if raw == "" {
		return ResearchReport{}, fmt.Errorf(
			"research response is empty",
		)
	}

	if err := json.Unmarshal([]byte(raw), &report); err != nil {
		return ResearchReport{}, fmt.Errorf(
			"parse research report: %w",
			err,
		)
	}

	return report, nil
}
