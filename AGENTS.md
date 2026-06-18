
# AGENTS.md

## Project: Dino

Please don't use Graphiti for this project since it's not setup. For every change, please explain for me carefully because I'm doing this project for learning purposes.

Dino is a local-first AI-powered CLI that helps developers and creators transform information into useful outputs.

Dino supports two primary workflows:

```text
Content Workflow

source
→ fetch
→ summarize
→ draft
→ export
```

```text
Developer Workflow

git diff
→ analyze
→ summarize
→ commit
→ PR
→ release notes
```

Dino is designed as a collection of practical AI-assisted commands that improve everyday workflows.

## Product Vision

Dino should feel like:

```text
git + jq + curl + AI
```

A lightweight terminal-native assistant.

Not a SaaS.

Not a dashboard.

Not a web application.

The CLI is the product.

## Core Principles

* Local-first
* Human-in-the-loop
* Scriptable
* Fast startup
* Small composable commands
* AI as an assistant, not an autonomous agent

## MVP Scope

Build only a few useful commands.

### Content

```bash
dino fetch <url>
dino summarize <file>
dino draft astro <file>
```

### Git

```bash
dino status
dino suggest
dino commit
```

Everything else comes later.

## Content Workflow

### Fetch

Examples:

```bash
dino fetch https://example.com/article
```

```bash
dino fetch youtube <url>
```

Store raw content locally.

Example:

```text
.dino/
  sources/
```

### Summarize

```bash
dino summarize source.json
```

Generate:

```text
Summary
Key Ideas
Important Takeaways
References
```

### Draft

```bash
dino draft astro summary.md
```

Generate Astro-compatible Markdown.

Example:

```md
---
title: "Generated Title"
description: "Generated description"
pubDate: 2026-06-18
draft: true
tags:
  - tech
source:
  url: "..."
---

## Summary

...

## Notes

...

## References

...
```

The user manually moves the generated draft into their Astro project.

Dino should not publish automatically.

## Git Workflow

### Status

```bash
dino status
```

Displays:

```text
Current branch
Modified files
Staged files
Untracked files
```

### Suggest

```bash
dino suggest
```

Analyze:

```bash
git diff --cached
```

Generate:

```text
feat(auth): add login form validation
```

Optional explanation:

```text
Why:
- validate email format
- validate password length
- improve error feedback
```

### Commit

```bash
dino commit
```

Workflow:

```text
Analyze diff
→ Suggest commit message
→ User reviews
→ Commit
```

Never commit without explicit confirmation.

### Conventional Commits

Default format:

```text
type(scope): description
```

Supported types:

```text
feat
fix
refactor
docs
test
chore
style
build
ci
perf
revert
```

## Future Developer Commands

### PR Generation

```bash
dino pr
```

Generate:

```text
PR title
PR description
Testing notes
Breaking changes
```

### Release Notes

```bash
dino release-notes
```

Generate changelogs from commits.

### Explain Diff

```bash
dino explain
```

Explain:

```text
What changed
Why it changed
Potential risks
```

### Review

```bash
dino review
```

Generate a local review report:

```text
Large functions
Missing tests
Potential bugs
Dead code
```

## AI Usage

AI is used only for:

```text
summarization
classification
rewriting
drafting
explanation
commit generation
PR generation
release notes
```

AI must never:

```text
push code
delete files
rewrite history
force commits
modify repositories automatically
```

## Local Storage

Use:

```text
.dino/
  config.toml
  sources/
  summaries/
  drafts/
  exports/
```

Prefer files over databases in MVP.

Avoid SQLite until a real need appears.

## Configuration

Example:

```toml
[ai]
provider = "openai"
model = "gpt-5.5"
api_key_env = "OPENAI_API_KEY"

[content]
draft_format = "astro"

[git]
commit_style = "conventional"
```

## Architecture

```text
cmd/dino/

internal/config/

internal/source/
  article.go
  youtube.go

internal/summary/

internal/draft/
  astro.go

internal/git/
  status.go
  diff.go

internal/commit/

internal/pr/

internal/ai/
  client.go
```

## Success Criteria

Version 0.1 is successful if all of the following work:

```bash
dino fetch <url>

dino summarize source.json

dino draft astro summary.md

dino suggest

dino commit
```

If these commands provide value in daily usage, Dino is successful.

Everything else is secondary.
