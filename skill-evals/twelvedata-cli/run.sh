#!/usr/bin/env bash
#
# run.sh — fire SKILL.md evals against claude and/or codex, then (optionally)
#          grade the responses via an LLM-as-judge pass.
#
# Usage:
#   ./run.sh                       # run all evals on claude + codex
#   ./run.sh --runtime claude      # claude only
#   ./run.sh --runtime codex       # codex only
#   ./run.sh 9 10 11               # run only evals with these ids
#   ./run.sh --grade <run-dir>     # grade a prior run via LLM-as-judge
#
# Setup (do this once before running):
#   Claude: cp SKILL.md ~/.claude/skills/twelvedata-cli/SKILL.md
#   Codex:  cp SKILL.md AGENTS.md   # in the dir you run this script from
#   API:    export TWELVEDATA_API_KEY=... (needed if agents will actually run
#           twelvedata commands; required for evals that fetch live data)
#   CLI:    twelvedata installed and on PATH (the agent will shell out to it)
#
# Note: both agents are invoked with their fullest permission tier so they can
# actually execute `twelvedata` (Claude: --permission-mode bypassPermissions,
# Codex: --sandbox danger-full-access). Run from a throwaway working dir.
#
# Env vars:
#   JUDGE_CMD   command used for grading (default: "claude -p")
#               set to "codex exec" to use Codex as judge, or a different model
#               (e.g. "claude -p --model claude-sonnet-4-6") to reduce same-
#               model bias when grading a Claude run.

set -euo pipefail

SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
EVALS_FILE="$SCRIPT_DIR/evals.json"
RUNS_DIR="$SCRIPT_DIR/runs"
JUDGE_CMD=${JUDGE_CMD:-"claude -p"}

RUNTIMES=()
FILTER_IDS=()
MODE=run
GRADE_DIR=""

die()  { echo "error: $*" >&2; exit 1; }
warn() { echo "warn:  $*" >&2; }

usage() { sed -n '2,/^set -euo pipefail/p' "$0" | sed -n 's/^# \{0,1\}//p'; }

while [[ $# -gt 0 ]]; do
  case $1 in
    --runtime) RUNTIMES+=("$2"); shift 2;;
    --grade)   MODE=grade; GRADE_DIR=$2; shift 2;;
    -h|--help) usage; exit 0;;
    [0-9]*)    FILTER_IDS+=("$1"); shift;;
    *)         die "unknown arg: $1 (try --help)";;
  esac
done

[[ ${#RUNTIMES[@]} -eq 0 ]] && RUNTIMES=(claude codex)

[[ -f "$EVALS_FILE" ]] || die "evals.json not found: $EVALS_FILE"
command -v jq >/dev/null || die "jq is required"

# ---- grade mode -------------------------------------------------------------

if [[ $MODE == grade ]]; then
  [[ -d "$GRADE_DIR" ]] || die "run dir not found: $GRADE_DIR"
  out="$GRADE_DIR/grades.jsonl"
  : > "$out"
  echo "Grading with: $JUDGE_CMD"

  while read -r e; do
    id=$(jq -r '.id' <<<"$e")
    prompt=$(jq -r '.prompt' <<<"$e")
    expectations=$(jq -r '.expectations | "- " + join("\n- ")' <<<"$e")

    for rt in claude codex; do
      resp_file="$GRADE_DIR/$rt-$id.txt"
      [[ -f "$resp_file" ]] || continue
      response=$(cat "$resp_file")

      verdict=$($JUDGE_CMD "You are grading an LLM response against a checklist.
Output ONLY one line of JSON: {\"pass\": <bool>, \"failed\": [<short reasons>]}.
Pass = true ONLY if every expectation is satisfied.

PROMPT:
$prompt

RESPONSE (verbatim):
$response

EXPECTATIONS:
$expectations" </dev/null 2>/dev/null | tr -d '\r' | grep -m1 '^{' || echo '{"pass":null,"failed":["judge produced no JSON"]}')

      jq -c --argjson id "$id" --arg rt "$rt" --argjson v "$verdict" \
         '{id:$id, runtime:$rt} + $v' <<<'{}' >> "$out"
      echo "[$id $rt] $verdict"
    done
  done < <(jq -c '.evals[]' "$EVALS_FILE")

  echo
  echo "Summary ($out):"
  jq -s 'group_by(.runtime) | map({
    runtime: .[0].runtime,
    pass:    map(select(.pass == true))  | length,
    fail:    map(select(.pass == false)) | length,
    errored: map(select(.pass == null))  | length
  })' "$out"
  exit 0
fi

# ---- run mode ---------------------------------------------------------------

for rt in "${RUNTIMES[@]}"; do
  command -v "$rt" >/dev/null || die "$rt not on PATH (skip with --runtime <other>)"
done

if [[ ${#FILTER_IDS[@]} -gt 0 ]]; then
  filter=$(printf ',%s' "${FILTER_IDS[@]}" | sed 's/^,//')
  query=".evals[] | select(.id as \$i | [$filter] | index(\$i))"
else
  query=".evals[]"
fi

[[ " ${RUNTIMES[*]} " == *" claude "* ]] && [[ ! -f "$HOME/.claude/skills/twelvedata-cli/SKILL.md" ]] && \
  warn "skill not at ~/.claude/skills/twelvedata-cli/SKILL.md — claude may answer without it"
[[ " ${RUNTIMES[*]} " == *" codex "* ]] && [[ ! -f "AGENTS.md" ]] && \
  warn "no AGENTS.md in cwd ($PWD) — codex may answer without the skill"
[[ -z "${TWELVEDATA_API_KEY:-}" ]] && \
  warn "TWELVEDATA_API_KEY not set — agents granted Bash can't actually run twelvedata commands"

OUT="$RUNS_DIR/$(date +%Y%m%d-%H%M%S)"
mkdir -p "$OUT"
echo "Writing to $OUT"
echo "Runtimes: ${RUNTIMES[*]}"

while read -r e; do
  id=$(jq -r '.id' <<<"$e")
  prompt=$(jq -r '.prompt' <<<"$e")
  echo
  echo "===== eval $id ====="
  echo "$prompt"

  for rt in "${RUNTIMES[@]}"; do
    echo "  -> $rt"
    # `</dev/null` is critical: both runtimes read stdin and would otherwise
    # drain the eval JSON stream, terminating the loop after one iteration.
    case $rt in
      # --strict-mcp-config with no --mcp-config disables all MCP servers,
      # so the agent can only reach the API via the CLI (the thing under test).
      # Without this, an installed twelvedata MCP server shadows the CLI for
      # endpoints with a direct MCP equivalent (price, currency, indicators).
      claude) claude -p --permission-mode bypassPermissions --strict-mcp-config "$prompt" \
                </dev/null >"$OUT/claude-$id.txt" 2>"$OUT/claude-$id.err" || warn "claude exit $? on eval $id";;
      # --sandbox danger-full-access skips bubblewrap entirely, which is
      # required on Linux distros where AppArmor restricts unprivileged user
      # namespaces (Ubuntu 24+, recent Debian) — the default workspace-write
      # sandbox fails with `bwrap: loopback: Failed RTM_NEWADDR` there.
      # Symmetrical to Claude's --permission-mode bypassPermissions: same
      # risk envelope, agent runs as your user with full env access.
      codex)  codex exec --sandbox danger-full-access --skip-git-repo-check "$prompt" \
                </dev/null >"$OUT/codex-$id.txt"  2>"$OUT/codex-$id.err"  || warn "codex exit $? on eval $id";;
      *)      die "unknown runtime: $rt";;
    esac
  done
done < <(jq -c "$query" "$EVALS_FILE")

echo
echo "Done."
echo "Eyeball: less $OUT/*.txt"
echo "Grade:   $0 --grade $OUT"
