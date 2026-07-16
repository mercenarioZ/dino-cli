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

	diff, _, err := readDiff(*stagedOnly, *unstagedOnly)
	if err != nil {
		return err
	}

	if strings.TrimSpace(diff) == "" {
		switch {
		case *stagedOnly:
			return errors.New("nothing to commit - no staged changes found")
		case *unstagedOnly:
			return errors.New("nothing to commit - no unstaged changes found")
		default:
			return errors.New("nothing to commit - no staged or unstaged changes found")
		}
	}

	loading := startSpinner("generating commit message")
	message, err := generateCommitMessage(context.Background(), *model, diff)
	loading.stop()
	if err != nil {
		return err
	}

	fmt.Printf("%s\n\n", strings.TrimSpace(message))
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
