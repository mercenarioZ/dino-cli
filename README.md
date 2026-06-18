# dino

`dino` is a small Go CLI for developer workflows.

## Commit Message Generation

Generate a Conventional Commit message from the current git diff:

```sh
dino commit
```

By default, `dino commit` reads staged changes first with `git diff --cached`.
If there are no staged changes, it falls back to unstaged changes with `git diff`.

You can force either source:

```sh
dino commit --staged
dino commit --unstaged
```

## OpenAI Configuration

Set an API key before running AI-powered commands:

```sh
export DINO_OPENAI_API_KEY="your-api-key"
export OPENAI_API_KEY="your-api-key"
```

`DINO_OPENAI_API_KEY` takes priority when both are set.

By default, Dino calls the OpenAI Responses API:

```text
https://api.openai.com/v1/responses
```

If you use a custom OpenAI-compatible endpoint, set:

```sh
export DINO_OPENAI_RESPONSES_URL="https://your-endpoint.example/v1/responses"
```

For local development, you can copy `.env.example` to `.env`, edit it, then load it:

```sh
cp .env.example .env
source .env
```

The `.env` file is ignored by Git so secrets do not get pushed to GitHub.

## Build

```sh
go build -o dino .
```
