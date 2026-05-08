# Twelve Data CLI

`twelvedata-cli` is the official command-line client for the [Twelve Data API](https://twelvedata.com/), built for AI agents and humans alike. Every endpoint exposed by the [`twelvedata-go`](https://github.com/twelvedata/twelvedata-cli) SDK is reachable as a Cobra subcommand, with predictable flags, structured output, and stable exit codes.

## Install

```sh
go install github.com/twelvedata/twelvedata-cli/cmd/td@v1.0.0
```

## Quick start

```sh
export TWELVEDATA_API_KEY=...
td quote --symbol AAPL
td time-series --symbol AAPL --interval 1day --outputsize 5
td ti kst --symbol AAPL --interval 1day
```

## Output formats

`td` renders responses based on `--output` and TTY detection:

- `--output json` (or piped / `CI=true` / `--quiet`): pretty-printed JSON.
- `--output csv`: streams the API's CSV response verbatim. Sets `format=csv` upstream.
- `--output table` (default on a TTY): pretty-printed JSON for now; richer table rendering is a future enhancement.

## Authentication

`td` reads the API key from these sources in order:

1. `--api-key <key>` flag
2. `TWELVEDATA_API_KEY` environment variable

## Agent discovery

`td schema --json` dumps the entire command tree — names, flags, types, enum value sets, descriptions — so an LLM can introspect what commands and arguments are available without scraping `--help` text.

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
