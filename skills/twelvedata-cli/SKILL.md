---
name: twelvedata-cli
description: >
  Query the Twelve Data API from the terminal — real-time and historical prices, 60+
  technical indicators, fundamentals, market intelligence, regulatory filings — via
  the `twelvedata` CLI. Use when running Twelve Data commands in shells, scripts, or
  CI/CD pipelines, or when an agent invokes the CLI as a subprocess. Load this skill
  before running `twelvedata` commands — it contains the machine-mode contract,
  auth precedence, and gotchas that cause silent breakage otherwise.
license: MIT
metadata:
  author: twelvedata
  homepage: https://twelvedata.com/docs
  source: https://github.com/twelvedata/twelvedata-cli
  openclaw:
    primaryEnv: TWELVEDATA_API_KEY
    requires:
      env:
        - TWELVEDATA_API_KEY
      bins:
        - twelvedata
    envVars:
      - name: TWELVEDATA_API_KEY
        required: true
        description: Twelve Data API key for authenticating CLI commands
      - name: TWELVEDATA_PROFILE
        required: false
        description: Named auth profile for multi-account setups
    install:
      - kind: shell
        command: "curl -fsSL https://raw.githubusercontent.com/twelvedata/twelvedata-cli/main/install.sh | bash"
        bins: [twelvedata]
        label: Twelve Data CLI
    links:
      repository: https://github.com/twelvedata/twelvedata-cli
      documentation: https://twelvedata.com/docs
inputs:
  - name: TWELVEDATA_API_KEY
    description: Twelve Data API key for authenticating CLI commands. Get yours at https://twelvedata.com/account/api-keys
    required: true
  - name: TWELVEDATA_PROFILE
    description: Named auth profile for multi-account setups. Selects which stored API key to use (see `twelvedata auth`).
    required: false
---

# Twelve Data CLI

## Installation

Before running any `twelvedata` commands, check whether the CLI is installed:

```bash
twelvedata --version
```

If the command is not found, install it using one of the methods below:

**cURL (macOS / Linux):**
```bash
curl -fsSL https://raw.githubusercontent.com/twelvedata/twelvedata-cli/main/install.sh | bash
```

**Homebrew (macOS / Linux):**
```bash
brew install twelvedata/cli/twelvedata
```

**PowerShell (Windows):**
```powershell
irm https://raw.githubusercontent.com/twelvedata/twelvedata-cli/main/install.ps1 | iex
```

**Go:**
```bash
go install github.com/twelvedata/twelvedata-cli/cmd/twelvedata@latest
```

After installing, verify:
```bash
twelvedata --version
```

## Agent Protocol

The CLI auto-detects non-TTY environments and switches to **machine mode** automatically — no flag needed. Machine mode also activates when stdout is piped, when `CI=true` is set, when `TERM=dumb`, or when `--raw` is passed.

**Rules for agents:**
- Supply ALL required flags. The CLI does NOT prompt when stdin is not a TTY.
- Exit `0` = success; non-zero categorized per the table below.
- Response (JSON or CSV) goes to **stdout**; error envelope goes to **stderr**:
  ```json
  { "error": { "code": "unauthorized", "message": "Invalid API key", "status": 401 } }
  ```
- Use `--api-key` flag or `TWELVEDATA_API_KEY` env var. Never rely on interactive login.
- Stdout is response-only — safe to pipe directly into `jq`.

**Exit codes** (stable per category):

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

The CLI resolves the API key from these sources, in order:

1. `--api-key <key>` flag
2. `TWELVEDATA_API_KEY` environment variable
3. Active profile in `credentials.json` (see `twelvedata whoami`)

Keys are saved to the OS keyring (macOS Keychain, Windows Credential Manager, Linux Secret Service) by default, with a `0600` plaintext-file fallback when no keyring is available.

## Global Flags

| Flag | Description |
|------|-------------|
| `--api-key <key>` | Override API key for this invocation |
| `-p, --profile <name>` | Select stored profile (or set `TWELVEDATA_PROFILE`) |
| `-o, --output <fmt>` | Response format: `json` (default) or `csv` |
| `--raw` | Force machine mode from a TTY: JSON error envelope on stderr, no spinner, no color, no prompts |

## Command Discovery

Do NOT scrape `--help`. The CLI ships a stable JSON schema dump for agents:

```bash
twelvedata commands           # alias: twelvedata schema
```

The output is a tree of every command and subcommand, with each node containing `name`, `path`, `short`, `long`, `aliases`, `flags[]` (with `type`, `default`, `required`, `enum[]` if applicable), and `subcommands[]`. This is the source of truth for what flags exist and which values are valid.

## Available Commands

| Command Group | What it does |
|--------------|-------------|
| `time-series`, `time-series-cross`, `quote`, `price`, `eod` | Market data — poll `quote` (or `price`) for the latest tick; `time-series` for history |
| `exchange-rate`, `currency-conversion` | Forex/crypto pair rates (rate alone, or rate × amount) |
| `ti <indicator>` | 60+ technical indicators (RSI, MACD, SMA, BBANDS, ATR, ...) — see `ti --help` |
| `earnings`, `earnings-calendar`, `dividends`, `dividends-calendar`, `splits`, `splits-calendar`, `ipo-calendar` | Earnings, dividends, splits, IPO calendars |
| `profile`, `statistics`, `market-cap`, `key-executives`, `logo`, `press-releases` | Company profile and stats |
| `balance-sheet`, `balance-sheet-consolidated`, `cash-flow`, `cash-flow-consolidated`, `income-statement`, `income-statement-consolidated` | Financial statements (single + consolidated) |
| `analyst-ratings-light`, `analyst-ratings-us-equities`, `recommendations`, `price-target`, `earnings-estimate`, `revenue-estimate`, `growth-estimates`, `eps-trend`, `eps-revisions` | Analyst ratings, price targets, forecasts |
| `etfs-list`, `etfs-family`, `etfs-type`, `etfs-world`, `etfs-world-composition`, `etfs-world-performance`, `etfs-world-risk`, `etfs-world-summary` | ETF lists, composition, performance, risk |
| `mutual-funds-list`, `mutual-funds-family`, `mutual-funds-type`, `mutual-funds-world`, `mutual-funds-world-composition`, `mutual-funds-world-performance`, `mutual-funds-world-purchase-info`, `mutual-funds-world-ratings`, `mutual-funds-world-risk`, `mutual-funds-world-summary`, `mutual-funds-world-sustainability` | Mutual fund catalogs and details |
| `insider-transactions`, `institutional-holders`, `direct-holders`, `fund-holders`, `edgar-filings-archive`, `tax-info` | SEC/EDGAR filings, holders, tax data |
| `stocks`, `forex-pairs`, `cryptocurrencies`, `cryptocurrency-exchanges`, `etfs`, `funds`, `bonds`, `commodities`, `exchanges`, `exchange-schedule`, `countries`, `cross-listings`, `instrument-type`, `intervals`, `technical-indicators`, `market-state`, `symbol-search`, `earliest-timestamp` | Symbol catalogs, exchange/calendar metadata, lookups |
| `api-usage` | Plan limits and current usage |
| `auth` (`list`, `switch`, `rename`, `remove`), `login`, `logout`, `whoami` | Multi-profile credential management |
| `doctor` | Setup + API smoke checks |
| `commands` / `schema` | Dump the full command tree as JSON |
| `completion` | Shell completion for bash/zsh/fish/powershell |
| `docs`, `dashboard` | Open URLs (prints URL in machine mode instead) |

Run `twelvedata commands` for the authoritative, machine-readable list with all flags.

## Common Mistakes

| # | Mistake | Fix |
|---|---------|-----|
| 1 | **Passing the API key as a literal `--api-key` argument** | Leaks the key to shell history, `ps` output, and CI logs. Use `TWELVEDATA_API_KEY` env var, or `twelvedata login --key-stdin` piped from a secret |
| 2 | **Looking for technical indicators as top-level commands** (`twelvedata rsi`) | They live under `twelvedata ti <indicator>` — e.g. `twelvedata ti rsi --symbol AAPL --interval 1day` |
| 3 | **Scraping `--help` to discover flags** | Use `twelvedata commands` — stable JSON, includes enum value sets per flag |
| 4 | **Expecting errors on stdout** | Errors always go to **stderr** as a JSON envelope; stdout stays response-only |
| 5 | **Passing `--output csv` and parsing as JSON** | `--output csv` streams the API's CSV response verbatim (sets `format=csv` upstream); only `--output json` (default) returns JSON |
| 6 | **Calling `twelvedata docs` / `dashboard` in CI expecting a browser** | In machine mode (piped/`CI=true`/`--raw`) these print the URL to stdout instead of launching a browser — safe to call from scripts, but no browser opens |
| 7 | **Running with no API key and getting exit code `1`** | Auth failures exit `3`, not `1`. If you get `1`, it's a generic error — check stderr for the envelope |
| 8 | **HTTP request hanging in CI** | Set `TWELVEDATA_HTTP_TIMEOUT` (Go duration: `30s`, `2m`; or bare integer seconds). Default is `120s`; `0` disables the timeout |
| 9 | **Forgetting that `doctor` exits `1` on any `fail` check** | `warn` does NOT affect exit code, but a single `fail` flips exit to `1`. Useful as a CI smoke test |
| 10 | **Using `quote`/`price` for an FX pair rate** | They accept pair symbols (`USD/EUR`) but return market-style payloads. For just the rate use `exchange-rate --symbol USD/EUR`; to convert an amount use `currency-conversion --symbol USD/EUR --amount 100` |
| 11 | **Calling `auth switch` to use a different profile for one command** | `auth switch` changes the *persisted* active profile. For a single invocation, use `--profile <name>` (or `TWELVEDATA_PROFILE=<name>`) — they only override resolution for that call and leave the active profile untouched |

## Environment Variables

| Var | Purpose |
| --- | --- |
| `TWELVEDATA_API_KEY` | API key (highest precedence after `--api-key`) |
| `TWELVEDATA_PROFILE` | Profile name override |
| `TWELVEDATA_CREDENTIAL_STORE` | `secure_storage` (default) or `file` to force plaintext |
| `TWELVEDATA_CONFIG_DIR` | Override the credentials directory (mostly for tests) |
| `TWELVEDATA_HTTP_TIMEOUT` | Per-request HTTP timeout (Go duration or seconds). Default `120s`; `0` disables |

## Common Patterns

**Quick price lookup:**
```bash
twelvedata price --symbol AAPL
twelvedata quote --symbol AAPL
```

**Historical time series:**
```bash
twelvedata time-series --symbol AAPL --interval 1day --outputsize 30
```

**Technical indicator:**
```bash
twelvedata ti rsi --symbol AAPL --interval 1day --time-period 14
twelvedata ti macd --symbol BTC/USD --interval 1h
```

**CSV streaming for spreadsheets / pipelines:**
```bash
twelvedata time-series --symbol AAPL --interval 1day --output csv > aapl.csv
```

**Crypto / forex / indices:**
```bash
twelvedata price --symbol BTC/USD
twelvedata quote --symbol EUR/USD
twelvedata quote --symbol SPX
```

**CI smoke check (no login needed):**
```bash
export TWELVEDATA_API_KEY=...
twelvedata doctor --raw         # JSON output; exits 1 on any fail check
```

**Login workflow (interactive then scripted):**
```bash
twelvedata login                                         # masked prompt on a TTY
printf '%s' "$KEY" | twelvedata login --key-stdin        # safe for scripts
twelvedata login --profile staging --key-stdin <<<"$KEY" # named profile
twelvedata auth list                                     # list profiles
twelvedata auth switch staging                           # change active profile
twelvedata whoami                                        # masked key + active profile
```

**Per-invocation profile override:**
```bash
twelvedata quote --symbol AAPL --profile staging
TWELVEDATA_PROFILE=staging twelvedata quote --symbol AAPL
```

**Discover all commands + flags as JSON:**
```bash
twelvedata commands | jq '.subcommands[] | {name, short}'
twelvedata commands | jq '.. | objects | select(.name=="time-series") | .flags'
```

**Agent subprocess pattern (Bash):**
```bash
if out=$(twelvedata quote --symbol "$SYM" 2>err.json); then
  echo "$out" | jq -r '.close'
else
  code=$?
  jq -r '.error.message' < err.json
  exit "$code"
fi
```
