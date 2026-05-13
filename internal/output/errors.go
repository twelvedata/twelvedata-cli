package output

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/twelvedata/twelvedata-go/twelvedata"

	"github.com/twelvedata/twelvedata-cli/internal/auth"
)

// Exit codes communicated to agents and shells.
//
//	0  ok
//	2  CLI usage error (bad flag, missing required, unknown command)
//	3  authentication failure        (HTTP 401)
//	4  authorization / forbidden      (HTTP 403)
//	5  not found                      (HTTP 404)
//	6  rate-limited                   (HTTP 429)
//	7  bad request / param too long   (HTTP 400, 414)
//	8  server-side failure            (HTTP 500)
//	1  any other error
const (
	ExitOK             = 0
	ExitGeneric        = 1
	ExitUsage          = 2
	ExitUnauthorized   = 3
	ExitForbidden      = 4
	ExitNotFound       = 5
	ExitRateLimited    = 6
	ExitBadRequest     = 7
	ExitInternalServer = 8
)

// envelope is the machine-readable error shape written to stderr in JSON mode.
type envelope struct {
	Error envelopeBody `json:"error"`
}

type envelopeBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"status,omitempty"`
}

// WriteError writes the error to stderr in the resolved format and returns the
// process exit code. Format resolution mirrors Render: --output flag wins, then
// --quiet, then TTY detection.
func WriteError(cmd *cobra.Command, err error) int {
	if err == nil {
		return ExitOK
	}

	code, exit := classify(err)

	w := cmd.ErrOrStderr()
	if jsonMode(cmd) {
		writeJSONError(w, code, err)
	} else {
		fmt.Fprintf(w, "Error: %s\n", err.Error())
	}
	return exit
}

func writeJSONError(w io.Writer, code string, err error) {
	body := envelopeBody{Code: code, Message: err.Error()}
	var apiErr twelvedata.TwelvedataApiError
	if errors.As(err, &apiErr) {
		body.Status = apiErr.GetStatusCode()
		body.Message = apiErr.GetMessage()
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	_ = enc.Encode(envelope{Error: body})
}

// classify maps an error to a stable code string and exit code. The string code
// is what agents branch on; the exit code is what shells branch on.
func classify(err error) (string, int) {
	var apiErr twelvedata.TwelvedataApiError
	if errors.As(err, &apiErr) {
		switch apiErr.GetStatusCode() {
		case 400:
			return "bad_request", ExitBadRequest
		case 401:
			return "unauthorized", ExitUnauthorized
		case 403:
			return "forbidden", ExitForbidden
		case 404:
			return "not_found", ExitNotFound
		case 414:
			return "parameter_too_long", ExitBadRequest
		case 429:
			return "rate_limited", ExitRateLimited
		case 500:
			return "internal_server_error", ExitInternalServer
		}
	}
	if isUsageError(err) {
		return "usage_error", ExitUsage
	}
	if errors.Is(err, auth.ErrNoAPIKey) {
		return "not_authenticated", ExitUnauthorized
	}
	if strings.Contains(err.Error(), "TWELVEDATA_API_KEY environment variable is not set") {
		return "missing_api_key", ExitUnauthorized
	}
	return "error", ExitGeneric
}

// isUsageError detects Cobra's flag/argument validation errors. Cobra does not
// expose a sentinel, so we match on the prefixes it uses.
func isUsageError(err error) bool {
	msg := err.Error()
	for _, p := range []string{
		"unknown command",
		"unknown flag",
		"unknown shorthand",
		"required flag",
		"invalid argument",
		"flag needs an argument",
		"accepts ", // "accepts N arg(s), received M"
	} {
		if strings.HasPrefix(msg, p) {
			return true
		}
	}
	return false
}

// jsonMode tells whether stderr should carry the JSON error envelope. Same
// resolution rules as ResolveFormat for stdout.
func jsonMode(cmd *cobra.Command) bool {
	out, _ := cmd.Flags().GetString("output")
	if out != "" {
		return Format(out) == FormatJSON
	}
	quiet, _ := cmd.Flags().GetBool("quiet")
	if quiet {
		return true
	}
	return !isTerminal(os.Stderr) || isCI()
}
