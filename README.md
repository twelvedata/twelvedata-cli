# Twelve Data CLI

`twelvedata-cli` is the official command-line client for the [Twelve Data API](https://twelvedata.com/), built for AI agents and humans alike.

Every API endpoint is available as a subcommand, with predictable flags and structured output.

```
  _______       __________ _    ________   ____  ___  _________
 /_  __/ |     / / ____/ /| |  / / ____/  / __ \/   |/_  __/   |
  / /  | | /| / / __/ / / | | / / __/    / / / / /| | / / / /| |
 / /   | |/ |/ / /___/ /__| |/ / /___   / /_/ / ___ |/ / / ___ |
/_/    |__/|__/_____/_____/___/_____/  /_____/_/  |_/_/ /_/  |_|
```

## Install

### cURL (macOS / Linux)

```sh
curl -fsSL https://raw.githubusercontent.com/twelvedata/twelvedata-cli/main/install.sh | bash
```

Pin a specific version: append `-s v1.0.0`. Override the install location with `TWELVEDATA_INSTALL=<dir>` (default `~/.twelvedata`).

### PowerShell (Windows)

```sh
irm https://raw.githubusercontent.com/twelvedata/twelvedata-cli/main/install.ps1 | iex
```

Pin a specific version: `$env:TWELVEDATA_VERSION = 'v1.0.0'` before piping. Override the install location with `$env:TWELVEDATA_INSTALL`.

### Go

```sh
go install github.com/twelvedata/twelvedata-cli/cmd/twelvedata@v1.0.0
```

### Prebuilt binaries

Download a tarball / zip for your OS from the [releases page](https://github.com/twelvedata/twelvedata-cli/releases/latest) and put the `twelvedata` binary on your `PATH`.

## Quick start

```sh
twelvedata login
twelvedata time-series --symbol AAPL --interval 1day
```

For scripts and CI, skip `twelvedata login` and provide the key inline — either set `TWELVEDATA_API_KEY=...` in the environment or pass `--api-key <key>` on each invocation. See [Authentication](#authentication).

`twelvedata docs` and `twelvedata dashboard` open URLs in your default browser. In machine mode (`--raw`, piped stdout, CI) they print the URL to stdout instead of launching a browser, so they're safe to call from scripts.

## Output behavior

The CLI has two output modes:

| Mode            | When                                         | Stdout                | Stderr                              |
| --------------- | -------------------------------------------- | --------------------- | ----------------------------------- |
| **Interactive** | TTY                                          | Pretty-printed JSON   | Spinner, prompts, colorized errors  |
| **Machine**     | `--raw`, piped stdout, `CI`, or `TERM=dumb`  | Pretty-printed JSON   | JSON error envelope                 |

Switching is automatic — pipe stdout and machine mode activates:

```sh
twelvedata quote --symbol AAPL | jq .price
```

Use `--raw` to force machine mode from a TTY (e.g. when an agent captures both streams):

```sh
twelvedata quote --symbol AAPL --raw
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
    { "name": "CLI Version",        "status": "pass", "message": "v1.0.0 (latest)" },
    { "name": "API Key",            "status": "pass", "message": "abc...wxyz (source: secure storage, profile: default)" },
    { "name": "Credential Storage", "status": "pass", "message": "Secret Service (libsecret)" },
    { "name": "API Validation",     "status": "pass", "message": "API key accepted" }
  ]
}
```

## Update notifications

After a successful command, `twelvedata` checks GitHub releases at most once every 24 hours and prints a one-line hint to **stderr** when a newer version is available:

```
A new version of twelvedata is available: v1.0.0 → v<latest>
  Run: <upgrade command for your install method>
  Disable: TWELVEDATA_NO_UPDATE_NOTIFIER=1
```

The upgrade hint is inferred from the running binary's path so it matches how you installed `twelvedata`:

| Detected path | Hint |
| --- | --- |
| `~/.twelvedata/bin/twelvedata` (or `.exe`) | `curl … install.sh \| bash` or `irm … install.ps1 \| iex` |
| `…/Cellar/twelvedata/…` or `…/homebrew/…/twelvedata` | `brew update && brew upgrade twelvedata` |
| `$GOPATH/bin/twelvedata` or `~/go/bin/twelvedata` | `go install github.com/twelvedata/twelvedata-cli/cmd/twelvedata@v<latest>` |
| anything else | `https://github.com/twelvedata/twelvedata-cli/releases/latest` (shown as "Visit:") |

The check is skipped automatically in machine mode (`--raw`, piped stdout, `CI`, or `TERM=dumb`), so scripts and agents never see the notice. The result is cached at `$XDG_CACHE_HOME/twelvedata-cli/update-check.json` (or your OS cache dir). Set `TWELVEDATA_NO_UPDATE_NOTIFIER=1` to opt out entirely.

## Agent discovery

`twelvedata commands` dumps the entire command tree as JSON — names, flags, types, enum value sets, descriptions — so an LLM can introspect what commands and arguments are available without scraping `--help` text. `twelvedata schema` is kept as an alias.

## Shell completion

`twelvedata` ships completion for bash, zsh, fish, and PowerShell. Each script is generated from the live command tree, so flags and enum values stay in sync with the binary.

One-line install (auto-detects `$SHELL`):

```sh
twelvedata completion install
```

Then restart your shell. Pass a shell name to override detection (`twelvedata completion install zsh`).

Manual setup — print the script and source it yourself:

```sh
source <(twelvedata completion bash)                                # current shell only
twelvedata completion zsh > "${fpath[1]}/_twelvedata"                       # zsh, system fpath
twelvedata completion fish > ~/.config/fish/completions/twelvedata.fish     # fish
twelvedata completion powershell | Out-String | Invoke-Expression   # PowerShell session
```

`twelvedata completion install` is idempotent — re-running won't append a second copy.

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

## Verifying releases

Each GitHub release ships:

- `twelvedata-<os>-<arch>.tar.gz` (or `.zip` on Windows) — the binary archive.
- `checksums.txt` — SHA256 of every archive.
- `checksums.txt.sig` and `checksums.txt.pem` — keyless [cosign](https://docs.sigstore.dev/cosign/overview/) signature of `checksums.txt`, produced by this repo's release workflow via Sigstore. The signature is recorded in the public [Rekor](https://docs.sigstore.dev/logging/overview/) transparency log.

The install scripts already verify the archive's SHA256 against `checksums.txt`. To additionally verify that `checksums.txt` was produced by the expected release workflow — a full supply-chain check that catches a compromised GitHub release — use cosign:

```sh
VERSION=v1.0.0
BASE=https://github.com/twelvedata/twelvedata-cli/releases/download/$VERSION

curl -fLO $BASE/checksums.txt
curl -fLO $BASE/checksums.txt.sig
curl -fLO $BASE/checksums.txt.pem

cosign verify-blob \
  --certificate checksums.txt.pem \
  --signature checksums.txt.sig \
  --certificate-identity-regexp '^https://github\.com/twelvedata/twelvedata-cli/\.github/workflows/release\.yml@refs/tags/v.+$' \
  --certificate-oidc-issuer 'https://token.actions.githubusercontent.com' \
  checksums.txt
```

A successful run prints `Verified OK`. If the command fails, do not install the release — the archive was not signed by the expected workflow identity.

## License

MIT — see [LICENSE](LICENSE).
