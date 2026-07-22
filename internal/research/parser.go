package research

import (
	"encoding/json"
	"fmt"
	"strings"
)

func ParseReport(raw string) (Report, error) {
	var report Report

	raw = strings.TrimSpace(raw)
	if raw == "" {
		return Report{}, fmt.Errorf(
			"research response is empty",
		)
	}

	if err := json.Unmarshal([]byte(raw), &report); err != nil {
		return Report{}, fmt.Errorf(
			"parse research report: %w",
			err,
		)
	}

	return report, nil
}
