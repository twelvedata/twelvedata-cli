package output

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/twelvedata/twelvedata-go/twelvedata"
	"golang.org/x/term"

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

// Error codes emitted in the JSON error envelope. classify() is the single
// source of truth that returns these; ErrorCodes mirrors that set in the order
// they're surfaced in --help.
const (
	CodeBadRequest          = "bad_request"
	CodeUnauthorized        = "unauthorized"
	CodeForbidden           = "forbidden"
	CodeNotFound            = "not_found"
	CodeParameterTooLong    = "parameter_too_long"
	CodeRateLimited         = "rate_limited"
	CodeInternalServerError = "internal_server_error"
	CodeUsageError          = "usage_error"
	CodeNotAuthenticated    = "not_authenticated"
	CodeMissingAPIKey       = "missing_api_key"
	CodeGenericError        = "error"
)

// ErrorCodes lists every code classify() may emit. Consumed by the Cobra usage
// template in commands/root.go.
var ErrorCodes = []string{
	CodeBadRequest,
	CodeUnauthorized,
	CodeForbidden,
	CodeNotFound,
	CodeParameterTooLong,
	CodeRateLimited,
	CodeInternalServerError,
	CodeUsageError,
	CodeNotAuthenticated,
	CodeMissingAPIKey,
	CodeGenericError,
}

// NotImplementedParam is returned by a generated command when the API spec
// declares a parameter whose type the api.mustache template doesn't handle yet
// (e.g. number/array/file). The classifier maps it to a generic error so the
// user sees a clear "not implemented" message instead of the parameter silently
// being dropped from the request.
func NotImplementedParam(flag, dataType string) error {
	return fmt.Errorf("CLI parameter --%s (type %s) is not implemented", flag, dataType)
}

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
// process exit code. Format resolution mirrors Render: --output flag wins,
// then TTY detection.
func WriteError(cmd *cobra.Command, err error) int {
	if err == nil {
		return ExitOK
	}

	code, exit := classify(err)
	msg := err.Error()
	if code == CodeUsageError {
		msg = cobraWording.Replace(msg)
	}

	w := cmd.ErrOrStderr()
	if IsRaw(cmd) {
		writeJSONError(w, code, msg, err)
	} else {
		writePrettyError(cmd, w, err, msg)
	}
	return exit
}

// writePrettyError renders a human-friendly TTY error: red mark, bold status
// name, dim status code, message indented and word-wrapped at terminal width.
// For non-API errors falls back to a single colorized "Error: msg" line.
func writePrettyError(cmd *cobra.Command, w io.Writer, err error, msg string) {
	color := useColor(cmd, w)

	var apiErr twelvedata.TwelvedataApiError
	if errors.As(err, &apiErr) {
		status := apiErr.GetStatusCode()
		name := statusName(status)
		if color {
			fmt.Fprintf(w, "\x1b[31m✗\x1b[0m \x1b[1m%s\x1b[0m \x1b[2m(%d)\x1b[0m\n", name, status)
		} else {
			fmt.Fprintf(w, "✗ %s (%d)\n", name, status)
		}
		bodyWidth := max(termWidth(w)-2, 20)
		for _, line := range wrapWords(apiErr.GetMessage(), bodyWidth) {
			fmt.Fprintf(w, "  %s\n", line)
		}
		return
	}

	if color {
		fmt.Fprintf(w, "\x1b[31m✗\x1b[0m %s\n", msg)
	} else {
		fmt.Fprintf(w, "Error: %s\n", msg)
	}
}

func statusName(code int) string {
	if name := http.StatusText(code); name != "" {
		return name
	}
	return "Error"
}

func useColor(cmd *cobra.Command, w io.Writer) bool {
	if IsRaw(cmd) {
		return false
	}
	if os.Getenv("NO_COLOR") != "" {
		return false
	}
	return isTerminal(w)
}

// termWidth returns the column width of w, capped at 100 so wide terminals
// don't produce one absurdly long line. Falls back to 80 when w isn't a TTY.
func termWidth(w io.Writer) int {
	f, ok := w.(*os.File)
	if !ok {
		return 80
	}
	width, _, err := term.GetSize(int(f.Fd()))
	if err != nil || width <= 0 {
		return 80
	}
	if width > 100 {
		return 100
	}
	return width
}

// wrapWords splits s on whitespace and greedy-wraps it into lines no wider
// than width. A single word longer than width is kept on its own line rather
// than broken — URLs and identifiers like (figi,figi_composite,...) stay intact.
func wrapWords(s string, width int) []string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return []string{""}
	}
	var lines []string
	var line strings.Builder
	for _, word := range words {
		if line.Len() == 0 {
			line.WriteString(word)
			continue
		}
		if line.Len()+1+len(word) > width {
			lines = append(lines, line.String())
			line.Reset()
			line.WriteString(word)
			continue
		}
		line.WriteByte(' ')
		line.WriteString(word)
	}
	if line.Len() > 0 {
		lines = append(lines, line.String())
	}
	return lines
}

func writeJSONError(w io.Writer, code, msg string, err error) {
	body := envelopeBody{Code: code, Message: msg}
	var apiErr twelvedata.TwelvedataApiError
	if errors.As(err, &apiErr) {
		body.Status = apiErr.GetStatusCode()
		body.Message = apiErr.GetMessage()
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	_ = enc.Encode(envelope{Error: body})
}

// cobraWording rewrites Cobra/pflag's "flag" vocabulary to match the CLI's
// user-facing "option" wording. Applied only to usage errors so it never
// touches API-error messages.
var cobraWording = strings.NewReplacer(
	"at least one of the flags", "at least one of the options",
	"if any flags in the group", "if any options in the group",
	"unknown shorthand flag", "unknown shorthand option",
	"unknown flag", "unknown option",
	"flag needs an argument", "option needs an argument",
	"required flag(s)", "required option(s)",
	"required flag", "required option",
	"flag(s)", "option(s)",
	"\" flag:", "\" option:",
)

// classify maps an error to a stable code string and exit code. The string code
// is what agents branch on; the exit code is what shells branch on.
func classify(err error) (string, int) {
	var apiErr twelvedata.TwelvedataApiError
	if errors.As(err, &apiErr) {
		switch apiErr.GetStatusCode() {
		case 400:
			return CodeBadRequest, ExitBadRequest
		case 401:
			return CodeUnauthorized, ExitUnauthorized
		case 403:
			return CodeForbidden, ExitForbidden
		case 404:
			return CodeNotFound, ExitNotFound
		case 414:
			return CodeParameterTooLong, ExitBadRequest
		case 429:
			return CodeRateLimited, ExitRateLimited
		case 500:
			return CodeInternalServerError, ExitInternalServer
		}
	}
	if isUsageError(err) {
		return CodeUsageError, ExitUsage
	}
	if errors.Is(err, auth.ErrNoAPIKey) {
		return CodeNotAuthenticated, ExitUnauthorized
	}
	if strings.Contains(err.Error(), "TWELVEDATA_API_KEY environment variable is not set") {
		return CodeMissingAPIKey, ExitUnauthorized
	}
	return CodeGenericError, ExitGeneric
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
		"at least one of the flags",
		"if any flags in the group",
		"accepts ", // "accepts N arg(s), received M"
	} {
		if strings.HasPrefix(msg, p) {
			return true
		}
	}
	return false
}
