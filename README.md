# Twelve Data CLI

[![Latest release](https://img.shields.io/github/v/release/twelvedata/twelvedata-cli?sort=semver)](https://github.com/twelvedata/twelvedata-cli/releases/latest)
[![Downloads](https://img.shields.io/github/downloads/twelvedata/twelvedata-cli/total)](https://github.com/twelvedata/twelvedata-cli/releases)
[![License](https://img.shields.io/github/license/twelvedata/twelvedata-cli)](https://github.com/twelvedata/twelvedata-cli/blob/master/LICENSE)
[![Issues](https://img.shields.io/github/issues/twelvedata/twelvedata-cli)](https://github.com/twelvedata/twelvedata-cli/issues)

```
  _______       __________ _    ________   ____  ___  _________
 /_  __/ |     / / ____/ /| |  / / ____/  / __ \/   |/_  __/   |
  / /  | | /| / / __/ / / | | / / __/    / / / / /| | / / / /| |
 / /   | |/ |/ / /___/ /__| |/ / /___   / /_/ / ___ |/ / / ___ |
/_/    |__/|__/_____/_____/___/_____/  /_____/_/  |_/_/ /_/  |_|
```

Twelve Data official library. This package supports all main features of the service:
- Real-time and historical market data: time series, quotes, end-of-day prices, exchange rates, and market movers.
- Fundamentals: financial statements, earnings, dividends, splits, company profiles, and key statistics.
- 100+ technical indicators: SMA, EMA, RSI, MACD, Bollinger Bands, Ichimoku, and many more.
- Analysis & estimates: analyst ratings, price targets, EPS trends, revenue and growth estimates.
- ETFs and mutual funds: summaries, composition, performance, risk, and ratings.

🔑 **API key** is required. If you don't have it yet, get it by [signing up here](https://twelvedata.com/pricing).

Every Twelve Data API endpoint is exposed as a subcommand with predictable flags and structured output.

## Documentation
- 📘 [Twelve Data API reference](https://twelvedata.com/docs) — endpoint details and response schemas
- 🛠 `twelvedata --help` — built-in CLI command reference
- 🤖 `twelvedata commands` — full command tree as JSON for agents

## Install

### cURL (macOS / Linux)

```sh
curl -fsSL https://raw.githubusercontent.com/twelvedata/twelvedata-cli/main/install.sh | bash
```

Pin a specific version: append `-s v1.0.10`. Override the install location with `TWELVEDATA_INSTALL=<dir>` (default `~/.twelvedata`).

### Homebrew (macOS / Linux)

```sh
brew install twelvedata/cli/twelvedata
```

Upgrade later with `brew update && brew upgrade twelvedata`. The tap is auto-bumped by the release pipeline for every stable release; pre-releases are skipped.

### PowerShell (Windows)

```sh
irm https://raw.githubusercontent.com/twelvedata/twelvedata-cli/main/install.ps1 | iex
```

Pin a specific version: `$env:TWELVEDATA_VERSION = 'v1.0.10'` before piping. Override the install location with `$env:TWELVEDATA_INSTALL`.

### Go

```sh
go install github.com/twelvedata/twelvedata-cli/cmd/twelvedata@v1.0.10
```

### Prebuilt binaries

Download a tarball / zip for your OS from the [releases page](https://github.com/twelvedata/twelvedata-cli/releases/latest) and put the `twelvedata` binary on your `PATH`.

## Quick start

```sh
twelvedata login
twelvedata time-series --symbol AAPL --interval 1day
```

For scripts, CI, and agent subprocesses, see [Agent & CI/CD usage](#agent--cicd-usage).

`twelvedata docs` and `twelvedata dashboard` open URLs in your default browser. In machine mode (`--raw`, piped stdout, CI) they print the URL to stdout instead of launching a browser, so they're safe to call from scripts.

## Shell completion

`twelvedata` ships completion for bash, zsh, fish, and PowerShell, generated from the live command tree so flags and enum values stay in sync.

```sh
twelvedata completion install     # auto-detects $SHELL; pass a shell name to override
```

Homebrew installs completion automatically (bash users need `brew install bash-completion@2` once). The install command is idempotent — re-running won't append a second copy.

<details>
<summary>Manual setup — print the script and source it yourself</summary>

```sh
source <(twelvedata completion bash)                                        # current shell only
twelvedata completion zsh > "${fpath[1]}/_twelvedata"                       # zsh, system fpath
twelvedata completion fish > ~/.config/fish/completions/twelvedata.fish     # fish
twelvedata completion powershell | Out-String | Invoke-Expression           # PowerShell session
```
</details>

## Output & errors

The CLI auto-switches between two output modes:

| Mode            | When                                         | Stdout                | Stderr                              |
| --------------- | -------------------------------------------- | --------------------- | ----------------------------------- |
| **Interactive** | TTY                                          | Pretty-printed JSON   | Spinner, prompts, colorized errors  |
| **Machine**     | `--raw`, piped stdout, `CI`, or `TERM=dumb`  | Pretty-printed JSON   | JSON error envelope                 |

Machine mode suppresses the spinner and color and skips every interactive prompt (missing-option, masked-key, profile picker, destructive confirmation) — the same arguments that work in CI work from a `--raw` TTY. Errors render as a stable JSON envelope on **stderr**, so **stdout** stays response-only for scripting:

```json
{ "error": { "code": "unauthorized", "message": "Invalid API key", "status": 401 } }
```

`--output` selects the response format (orthogonal to mode):

- `--output json` (default) — pretty-printed JSON.
- `--output csv` — streams the API's CSV response verbatim (sets `format=csv` upstream).

### Exit codes

| Code | Meaning |
| ---: | --- |
| `0` | Success |
| `1` | Generic error |
| `2` | Usage error (bad flag, unknown command, invalid enum value) |
| `3` | Authentication failure (HTTP 401, missing API key) |
| `4` | Forbidden (HTTP 403) |
| `5` | Not found (HTTP 404) |
| `6` | Rate limited (HTTP 429) |
| `7` | Bad request (HTTP 400, 414) |
| `8` | Internal server error (HTTP 500) |

## Authentication

CLI resolves the API key from these sources, in order:

1. `--api-key <key>` flag
2. `TWELVEDATA_API_KEY` environment variable
3. Active profile in `credentials.json` (see `twelvedata whoami`)

> **Avoid putting secrets on the command line.** `--api-key` and `twelvedata login --key` accept the key as a literal argument, which leaks it to shell history, `ps` output, and CI logs. For day-to-day use prefer `TWELVEDATA_API_KEY`, a saved profile, or `twelvedata login --key-stdin` for piped input.

### Profiles

CLI supports named profiles so you can keep separate keys for prototyping, production, or different team accounts.

```sh
twelvedata login                                          # prompts on a TTY (masked input)
printf '%s' "$TWELVEDATA_API_KEY" | twelvedata login --key-stdin
twelvedata login --profile staging --key-stdin <<<"$KEY"  # CI/scripts
twelvedata auth list                                      # list profiles (also: bare `twelvedata auth`)
twelvedata auth switch staging                            # change active profile
twelvedata whoami                                         # show active profile + masked key
```

`twelvedata login --key <value>` still works for ad-hoc use but is discouraged for the leakage reasons above.

Other auth commands:

- `twelvedata logout [--profile <name>]`
- `twelvedata auth rename <old> <new>`
- `twelvedata auth remove <name>`

The `-p` / `--profile` flag (or `TWELVEDATA_PROFILE` env var) overrides the active profile for one invocation.

Interactive behavior — selectors for missing arguments and confirmations for destructive operations — follows the rules in [Output & errors](#output--errors).

### Storage

By default, keys are saved to your OS keyring:

- **macOS** — Keychain
- **Windows** — Credential Manager
- **Linux** — Secret Service

When no keyring is available, the key falls back to a `0600` file:

- **macOS / Linux** — `$XDG_CONFIG_HOME/twelvedata/credentials.json` (or `~/.config/twelvedata/credentials.json`)
- **Windows** — `%APPDATA%\twelvedata\credentials.json`

Override storage with `TWELVEDATA_CREDENTIAL_STORE=file` to force plaintext.

### Environment variables

| Var | Purpose |
| --- | --- |
| `TWELVEDATA_API_KEY` | API key (highest precedence after `--api-key`) |
| `TWELVEDATA_PROFILE` | Profile name override |
| `TWELVEDATA_CREDENTIAL_STORE` | `secure_storage` (default) or `file` |
| `TWELVEDATA_CONFIG_DIR` | Override the credentials directory (mostly for tests) |
| `TWELVEDATA_HTTP_TIMEOUT` | Per-request HTTP timeout. Go duration (`30s`, `2m`, `1m30s`) or a bare integer in seconds. Default `120s`; `0` disables the timeout. |

## Agent & CI/CD usage

### CI/CD

Set `TWELVEDATA_API_KEY` as an environment variable — no `twelvedata login` needed. Machine mode activates automatically in CI (piped stdout, `CI=true`, or `TERM=dumb`), so the spinner is suppressed and errors render as a JSON envelope on stderr. See [Output & errors](#output--errors) for the full contract.

### AI agents

Agents calling the CLI as a subprocess automatically get machine mode (non-TTY detection). The contract:

- **Input** — every required flag must be provided; no interactive prompts in machine mode.
- **Output** — pretty-printed JSON on **stdout** (or CSV with `--output csv`); JSON error envelope on **stderr**, so `stdout` stays response-only for `jq`.
- **Exit code** — stable per category: `0` success, `2` usage error, `3` auth, `4` forbidden, `5` not found, `6` rate limited, `7` bad request, `8` server error. See [Exit codes](#exit-codes).
- **Errors** — always include `code`, `message`, and `status` fields.
- **Discovery** — `twelvedata commands` (alias: `twelvedata schema`) prints the full command tree as JSON (names, flags, types, enum value sets, descriptions) so agents can introspect without scraping `--help`.

### `twelvedata commands`

Dumps the entire CLI command tree as JSON. When stdout is piped, JSON is used automatically; from a TTY pass `--raw` to force machine output.

## Diagnostics

`twelvedata doctor` runs four checks against your local setup and the API, and exits non-zero on any failure — useful as a smoke test in CI and as a triage step when commands start misbehaving.

| Check | What it verifies |
| --- | --- |
| CLI Version        | Installed binary matches the latest GitHub release |
| API Key            | A key is resolvable via flag, env var, or credentials file |
| Credential Storage | Whether keys live in OS secure storage or a plaintext file |
| API Validation     | The resolved key is accepted by `/api_usage` |

Each check reports `pass`, `warn`, or `fail`. Exit code is `1` if any check is `fail`; `warn` does not affect the exit code (so a stale CLI build or a transient network hiccup won't break CI).

```sh
twelvedata doctor                              # human-readable output
twelvedata doctor --raw                        # JSON envelope for scripts and CI
twelvedata doctor --profile staging            # run checks against a specific profile
```

In machine mode the payload shape is:

```json
{
  "ok": true,
  "checks": [
    { "name": "CLI Version",        "status": "pass", "message": "v1.0.10 (latest)" },
    { "name": "API Key",            "status": "pass", "message": "abc...wxyz (source: secure storage, profile: default)" },
    { "name": "Credential Storage", "status": "pass", "message": "Secret Service (libsecret)" },
    { "name": "API Validation",     "status": "pass", "message": "API key accepted" }
  ]
}
```

## Agent skills

This CLI ships with an [agent skill](skills/twelvedata-cli/SKILL.md) that teaches AI coding agents (Claude Code, Cursor, Windsurf, etc.) how to use the Twelve Data CLI effectively, including the machine-mode contract, auth precedence, and common pitfalls.

Drop it into your agent's skills directory or paste the contents into the agent's project instructions. For Claude Code:

```sh
mkdir -p ~/.claude/skills/twelvedata-cli
curl -fsSL https://raw.githubusercontent.com/twelvedata/twelvedata-cli/main/skills/twelvedata-cli/SKILL.md \
  -o ~/.claude/skills/twelvedata-cli/SKILL.md
```

## Local development

Use this when you want to change the CLI and run your build locally.

### Prerequisites

- [Go](https://go.dev/dl/) 1.23+

### Setup

1. **Clone the repo**

   ```sh
   git clone https://github.com/twelvedata/twelvedata-cli.git
   cd twelvedata-cli
   ```

2. **Build the binary**

   ```sh
   go build -o twelvedata ./cmd/twelvedata
   ```

   Output: `./twelvedata`

## Running the CLI locally

Set your API key, then run the built binary:

```sh
export TWELVEDATA_API_KEY=...
./twelvedata quote --symbol AAPL
```

Or run directly from source without producing a binary:

```sh
go run ./cmd/twelvedata quote --symbol AAPL
```

### Making changes

After editing source files, rebuild:

```sh
go build -o twelvedata ./cmd/twelvedata
```

Or skip the build step entirely with `go run`.

## Documentation
Delve deeper with our [official documentation](https://twelvedata.com/docs).

## Support
- 🌐 Visit our [contact page](https://twelvedata.com/contact).
- 🛠 Or our [support center](https://support.twelvedata.com/).

## Stay updated
- Follow us on [𝕏](https://twitter.com/TwelveData).
- Follow us on [Telegram](https://t.me/twelvedata).

## Contributing
1. Fork and clone: `$ git clone https://github.com/twelvedata/twelvedata-cli .`.
2. Branch out: `$ git checkout -b my-feature`.
3. Commit changes and test.
4. Push your branch and open a pull request with a comprehensive description.

For more guidance on contributing, see the [GitHub Docs](https://docs.github.com/en/get-started/quickstart/contributing-to-projects) on GitHub.

## License

This project is licensed under the MIT. See the `LICENSE` file in this repository for more details.
