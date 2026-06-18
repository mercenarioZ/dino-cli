package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func readDiff(stagedOnly, unstagedOnly bool) (diff string, source string, err error) {
	if stagedOnly {
		diff, err = gitDiff("--cached")
		return diff, "staged", err
	}
	if unstagedOnly {
		diff, err = gitDiff()
		return diff, "unstaged", err
	}

	diff, err = gitDiff("--cached")
	if err != nil {
		return "", "", err
	}
	if strings.TrimSpace(diff) != "" {
		return diff, "staged", nil
	}

	diff, err = gitDiff()
	if err != nil {
		return "", "", err
	}
	return diff, "unstaged", nil
}

func gitDiff(args ...string) (string, error) {
	diffArgs := append([]string{"diff"}, args...)
	cmd := exec.Command("git", diffArgs...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		errText := strings.TrimSpace(stderr.String())
		if errText == "" {
			errText = err.Error()
		}
		return "", fmt.Errorf("git %s failed: %s", strings.Join(diffArgs, " "), errText)
	}

	return stdout.String(), nil
}
