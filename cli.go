package main

import "fmt"

func run(args []string) error {
	if len(args) == 0 {
		printUsage()
		return nil
	}

	switch args[0] {
	case "commit":
		return runCommit(args[1:])
	case "help", "-h", "--help":
		printUsage()
		return nil
	case "search":
		return runSearch(args[1:])
	default:
		return fmt.Errorf("unknown command %q", args[0])
	}
}

func printUsage() {
	fmt.Print(`dino is a small AI helper for developer workflows.

Usage:
  dino commit [--staged | --unstaged] [--model MODEL]
	dino search [--model MODEL] QUERY

Commands:
  commit    Generate a Conventional Commit message from the current git diff
	search    Search the live web and answer a question

Configuration:
  Set DINO_OPENAI_API_KEY or OPENAI_API_KEY.
  Optional: set DINO_OPENAI_RESPONSES_URL for a custom endpoint.
`)
}
