# Twelve Data CLI

`twelvedata-cli` is the official command-line client for the [Twelve Data API](https://twelvedata.com/), built for AI agents and humans alike.

Every API endpoint is reachable as a Cobra subcommand, with predictable flags, structured output, and stable exit codes.

## Install

```sh
go install github.com/twelvedata/twelvedata-cli/cmd/td@v1.0.0
```

## Quick start

```sh
td login                          # save your API key (prompts on a TTY)
td quote --symbol AAPL
td time-series --symbol AAPL --interval 1day --outputsize 5
td ti rsi --symbol AAPL --interval 1day
```

For scripts and CI, skip `td login` and provide the key inline — either set `TWELVEDATA_API_KEY=...` in the environment or pass `--api-key <key>` on each invocation. See [Authentication](#authentication).

## Output behavior

The CLI has two output modes:

| Mode            | When                                         | Stdout                | Stderr                              |
| --------------- | -------------------------------------------- | --------------------- | ----------------------------------- |
| **Interactive** | TTY                                          | Pretty-printed JSON   | Spinner, prompts, colorized errors  |
| **Machine**     | `--raw`, piped stdout, `CI`, or `TERM=dumb`  | Pretty-printed JSON   | JSON error envelope                 |

Switching is automatic — pipe stdout and machine mode activates:

```sh
td quote --symbol AAPL | jq .price
```

Use `--raw` to force machine mode from a TTY (e.g. when an agent captures both streams):

```sh
td quote --symbol AAPL --raw
```

In machine mode the spinner and color are suppressed, errors render as a JSON envelope on stderr, and every interactive helper (missing-option prompt, masked-key prompt, profile picker, destructive confirmation) is skipped — the same arguments that work in CI work from a `--raw` TTY.

`--output` selects the response format (orthogonal to mode):

- `--output json` (default): pretty-printed JSON.
- `--output csv`: streams the API's CSV response verbatim. Sets `format=csv` upstream.

### Error output

Errors exit with a stable code (see [Exit codes](#exit-codes)). On **stderr** the format depends on the mode:

- **Machine**: a JSON envelope, so **stdout** stays response-only for scripting:

  ```json
  { "error": { "code": "unauthorized", "message": "Invalid API key", "status": 401 } }
  ```

- **Interactive**: a human line with a red mark, the status name, and the word-wrapped body.

## Authentication

CLI resolves the API key from these sources, in order:

1. `--api-key <key>` flag
2. `TWELVEDATA_API_KEY` environment variable
3. Active profile in `credentials.json` (see `td whoami`)

> **Avoid putting secrets on the command line.** `--api-key` and `td login --key` accept the key as a literal argument, which leaks it to shell history, `ps` output, and CI logs. For day-to-day use prefer `TWELVEDATA_API_KEY`, a saved profile, or `td login --key-stdin` for piped input.

### Profiles

CLI supports named profiles so you can keep separate keys for prototyping, production, or different team accounts.

```sh
td login                                          # prompts on a TTY (masked input)
printf '%s' "$TWELVEDATA_API_KEY" | td login --key-stdin
td login --profile staging --key-stdin <<<"$KEY"  # CI/scripts
td auth list                                      # list profiles (also: bare `td auth`)
td auth switch staging                            # change active profile
td whoami                                         # show active profile + masked key
```

`td login --key <value>` still works for ad-hoc use but is discouraged for the leakage reasons above.

Other auth commands:

- `td logout [--profile <name>]`
- `td auth rename <old> <new>`
- `td auth remove <name>`

The `-p` / `--profile` flag (or `TWELVEDATA_PROFILE` env var) overrides the active profile for one invocation.

Interactive behavior — selectors for missing arguments and confirmations for destructive operations — follows the rules in [Output behavior](#output-behavior).

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

## Agent discovery

`td commands` dumps the entire command tree as JSON — names, flags, types, enum value sets, descriptions — so an LLM can introspect what commands and arguments are available without scraping `--help` text. `td schema` is kept as an alias.

## Shell completion

`td` ships completion for bash, zsh, fish, and PowerShell. Each script is generated from the live command tree, so flags and enum values stay in sync with the binary.

One-line install (auto-detects `$SHELL`):

```sh
td completion install
```

Then restart your shell. Pass a shell name to override detection (`td completion install zsh`).

Manual setup — print the script and source it yourself:

```sh
source <(td completion bash)                                # current shell only
td completion zsh > "${fpath[1]}/_td"                       # zsh, system fpath
td completion fish > ~/.config/fish/completions/td.fish     # fish
td completion powershell | Out-String | Invoke-Expression   # PowerShell session
```

`td completion install` is idempotent — re-running won't append a second copy.

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
   go build -o td ./cmd/td
   ```

   Output: `./td`

## Running the CLI locally

Set your API key, then run the built binary:

```sh
export TWELVEDATA_API_KEY=...
./td quote --symbol AAPL
```

Or run directly from source without producing a binary:

```sh
go run ./cmd/td quote --symbol AAPL
```

### Making changes

After editing source files, rebuild:

```sh
go build -o td ./cmd/td
```

Or skip the build step entirely with `go run`.

## Exit codes

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

## License

MIT — see [LICENSE](LICENSE).
