package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

const maxDiffChars = 120_000

func runCommit(args []string) error {
	fs := flag.NewFlagSet("commit", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	stagedOnly := fs.Bool("staged", false, "use only staged changes")
	unstagedOnly := fs.Bool("unstaged", false, "use only unstaged changes")
	model := fs.String("model", defaultOpenAIModel, "OpenAI model to use")
	fs.Usage = func() {
		fmt.Fprintln(fs.Output(), "Usage: dino commit [--staged | --unstaged] [--model MODEL]")
		fs.PrintDefaults()
	}

	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil
		}
		return err
	}
	if fs.NArg() > 0 {
		return fmt.Errorf("unexpected argument %q", fs.Arg(0))
	}
	if *stagedOnly && *unstagedOnly {
		return errors.New("choose only one of --staged or --unstaged")
	}

	statusf("checking the diff before making any claims")
	diff, source, err := readDiff(*stagedOnly, *unstagedOnly)
	if err != nil {
		return err
	}
	if strings.TrimSpace(diff) == "" {
		return errors.New("no git diff found")
	}
	statusf("using %s changes", source)

	if len(diff) > maxDiffChars {
		statusf("the diff is chunky, so only the first %d characters will be sent", maxDiffChars)
	}
	statusf("asking the model for a conventional commit message")

	loading := startSpinner("generating commit message")
	message, err := generateCommitMessage(context.Background(), *model, diff)
	loading.stop()
	if err != nil {
		return err
	}
	statusf("message ready")

	fmt.Printf("%s\n\n", strings.TrimSpace(message))
	fmt.Fprintf(os.Stderr, "Generated from %s diff.\n", source)
	return nil
}

func commitInstructions() string {
	return strings.Join([]string{
		"You generate Conventional Commit messages from git diffs.",
		"Return only the commit message with no markdown and no explanation.",
		"Use this format: <type>(optional scope): <description>.",
		"Use one of these types: feat, fix, docs, style, refactor, perf, test, build, ci, chore, revert.",
		"Prefer a single-line commit message unless a short body is genuinely useful.",
		"Keep the subject under 72 characters when possible.",
	}, "\n")
}

func commitInput(diff string) string {
	return fmt.Sprintf("Generate a Conventional Commit message for this diff:\n\n```diff\n%s\n```", truncateDiff(diff))
}

func truncateDiff(diff string) string {
	if len(diff) <= maxDiffChars {
		return diff
	}

	return diff[:maxDiffChars] + "\n\n[Diff truncated because it exceeded the CLI limit.]"
}
