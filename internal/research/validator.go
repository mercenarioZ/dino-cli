package research

import (
	"errors"
	"strings"
)

func ValidateReport(report Report) error {
	switch {
	case strings.TrimSpace(report.Topic) == "":
		return errors.New("research report topic is required")

	case strings.TrimSpace(report.Summary) == "":
		return errors.New("research summary cannot be blank")

	case len(report.Findings) == 0:
		return errors.New("research report required findings")
	}
	return nil
}
