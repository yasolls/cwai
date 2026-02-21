<p align="center">
  <img src="assets/logo.png" width="420" alt="cwai — Make commits with AI">
</p>

<p align="center">
  AI-powered conventional commit message generator from staged git changes.
</p>

<p align="center">
  <a href="https://github.com/nikmd1306/cwai/actions/workflows/tests.yml"><img src="https://github.com/nikmd1306/cwai/actions/workflows/tests.yml/badge.svg" alt="CI"></a>
  <a href="https://github.com/nikmd1306/cwai/releases/latest"><img src="https://img.shields.io/github/v/release/nikmd1306/cwai" alt="Release"></a>
  <a href="https://pkg.go.dev/github.com/nikmd1306/cwai"><img src="https://pkg.go.dev/badge/github.com/nikmd1306/cwai.svg" alt="Go Reference"></a>
  <a href="https://goreportcard.com/report/github.com/nikmd1306/cwai"><img src="https://goreportcard.com/badge/github.com/nikmd1306/cwai" alt="Go Report Card"></a>
  <a href="LICENSE"><img src="https://img.shields.io/github/license/nikmd1306/cwai" alt="License"></a>
</p>

---

## What is cwai?

cwai is a command-line tool that writes git commit messages for you. It looks at your staged changes (`git add`), sends them to an AI model, and generates a clear, well-formatted [Conventional Commits](https://www.conventionalcommits.org/) message. You can accept it, edit it, or ask for a new one. All you need is an API key from any OpenAI-compatible provider.

## Features

- Generates [Conventional Commits](https://www.conventionalcommits.org/) messages from `git diff`
- Works with any OpenAI-compatible API (OpenAI, Anthropic, local models)
- Two modes: interactive standalone or silent `prepare-commit-msg` hook
- Structured output support for consistent formatting
- Smart diff truncation to fit model token limits

## Prerequisites

- **Git** installed and available in your terminal
- **API key** from one of these providers (or any OpenAI-compatible API):

  | Provider | Get your key |
  |----------|-------------|
  | OpenAI | https://platform.openai.com/api-keys |
  | OpenRouter | https://openrouter.ai/keys |
  | DeepSeek | https://platform.deepseek.com/api_keys |

  You can also use local models via [Ollama](https://ollama.com/) or [LM Studio](https://lmstudio.ai/) — no API key needed, just point `CWAI_API_URL` to your local server.

## Installation

### Quick install (Linux / macOS)

```bash
curl -fsSL https://raw.githubusercontent.com/nikmd1306/cwai/main/install.sh | bash
```

To install to a custom directory:

```bash
curl -fsSL https://raw.githubusercontent.com/nikmd1306/cwai/main/install.sh | bash -s -- -b ~/.local/bin
```

> **Note:** The default install path `/usr/local/bin` requires `sudo`. If you prefer to install without `sudo`, use `~/.local/bin` and make sure it's in your `PATH`:
> ```bash
> echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
> source ~/.bashrc
> ```
> For zsh, replace `~/.bashrc` with `~/.zshrc`.

### go install

Requires [Go](https://go.dev/dl/) 1.23+.

```bash
go install github.com/nikmd1306/cwai@latest
```

> **Note:** Ensure `$(go env GOPATH)/bin` is in your `PATH`:
> ```bash
> export PATH="$PATH:$(go env GOPATH)/bin"
> ```

### Windows

Download the Windows zip from [Releases](https://github.com/nikmd1306/cwai/releases), extract `cwai.exe`, and add its folder to your `PATH`. Or, if you have Go installed:

```powershell
go install github.com/nikmd1306/cwai@latest
```

### Binary releases

Download pre-built binaries from [Releases](https://github.com/nikmd1306/cwai/releases).

### From source

```bash
git clone https://github.com/nikmd1306/cwai.git
cd cwai
make install
```

## Quick Start

1. **Set up your API key** (one-time):
   ```bash
   cwai setup
   ```
   The wizard will ask for your API key, provider URL, model, and language.
   See [Prerequisites](#prerequisites) for where to get an API key.

2. **Stage your changes**:
   ```bash
   git add .
   ```

3. **Generate a commit message**:
   ```bash
   cwai
   ```
   cwai will show the staged files, generate a message, and let you accept, edit, regenerate, or cancel.

> **Tip:** Use `cwai -y` to auto-accept the generated message (useful in CI/scripts).

### Standalone mode

```console
$ cwai
Staged files (3):
  internal/ai/client.go
  internal/prompt/prompt.go
  cmd/root.go

Generated commit message:

  feat(ai): add structured output support for commit generation

[y]es / [e]dit / [r]egenerate / [n]o:
```

### Hook mode

```bash
cwai hook set    # Install prepare-commit-msg hook
cwai hook unset  # Remove hook
```

## Configuration

Config file: `~/.cwai` (INI format). Use `cwai config set KEY VALUE` or `cwai setup`.

### Essential settings

| Key | Description | Default |
|-----|-------------|---------|
| `CWAI_API_KEY` | API key from your provider (required) | *not set* |
| `CWAI_API_URL` | Base URL of the AI API | `https://api.openai.com/v1` |
| `CWAI_MODEL` | Model to use for generation | `gpt-5-mini` |
| `CWAI_LANGUAGE` | Commit message language (ISO 639-1: `en`, `de`, `fr`, `es`, ...) | `en` |

### Advanced settings

| Key | Description | Default |
|-----|-------------|---------|
| `CWAI_MAX_TOKENS_INPUT` | Max input tokens for diff (higher = more context, more cost) | `4096` |
| `CWAI_MAX_TOKENS_OUTPUT` | Max output tokens for response | `500` |
| `CWAI_TEMPERATURE` | Sampling temperature (`0.0` = deterministic, `1.0` = creative) | *not set* |
| `CWAI_REASONING_EFFORT` | Reasoning effort for reasoning models (`low`, `medium`, `high`) | *not set* |
| `CWAI_VERBOSITY` | Output verbosity level | *not set* |
| `CWAI_STRUCTURED_OUTPUT` | Enable structured JSON output (`true`/`false`) | *not set* |

## Troubleshooting

| Error | Cause | Solution |
|-------|-------|----------|
| `CWAI_API_KEY is not set` | No API key configured | Run `cwai setup` or `cwai config set CWAI_API_KEY <your-key>` |
| `API error (HTTP 401)` | Invalid or expired API key | Regenerate your key at your provider's dashboard |
| `API error (HTTP 429)` | Rate limit exceeded | Wait a moment and try again, or upgrade your API plan |
| `no staged changes` | Nothing added to git staging area | Run `git add <files>` before running `cwai` |
| `not a git repository` | cwai was run outside a git repo | Navigate to a git repository first (`cd your-project`) |
| `cwai: command not found` | Binary not in PATH | See [Installation](#installation) for PATH setup instructions |

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for development setup and guidelines.

## License

[MIT](LICENSE)
