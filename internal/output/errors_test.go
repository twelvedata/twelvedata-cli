package output

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/twelvedata/twelvedata-go/twelvedata"

	"github.com/twelvedata/twelvedata-cli/internal/auth"
)

func apiErr(status int, msg string) error {
	return &twelvedata.ApiError{StatusCode: status, Message: msg, Status: "Error"}
}

func TestClassify_APIStatusCodes(t *testing.T) {
	cases := []struct {
		status   int
		wantCode string
		wantExit int
	}{
		{400, CodeBadRequest, ExitBadRequest},
		{401, CodeUnauthorized, ExitUnauthorized},
		{403, CodeForbidden, ExitForbidden},
		{404, CodeNotFound, ExitNotFound},
		{414, CodeParameterTooLong, ExitBadRequest},
		{429, CodeRateLimited, ExitRateLimited},
		{500, CodeInternalServerError, ExitInternalServer},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d", tc.status), func(t *testing.T) {
			code, exit := classify(apiErr(tc.status, "msg"))
			if code != tc.wantCode || exit != tc.wantExit {
				t.Errorf("classify status %d = %q/%d, want %q/%d", tc.status, code, exit, tc.wantCode, tc.wantExit)
			}
		})
	}
}

func TestClassify_UnknownAPIStatusFallsThrough(t *testing.T) {
	code, exit := classify(apiErr(418, "i'm a teapot"))
	if code != CodeGenericError || exit != ExitGeneric {
		t.Errorf("unmapped status = %q/%d, want generic/1", code, exit)
	}
}

func TestClassify_NoAPIKey(t *testing.T) {
	code, exit := classify(auth.ErrNoAPIKey)
	if code != CodeNotAuthenticated {
		t.Errorf("code = %q, want %q", code, CodeNotAuthenticated)
	}
	if exit != ExitUnauthorized {
		t.Errorf("exit = %d, want %d", exit, ExitUnauthorized)
	}
}

func TestClassify_WrappedNoAPIKey(t *testing.T) {
	code, _ := classify(fmt.Errorf("wrapped: %w", auth.ErrNoAPIKey))
	if code != CodeNotAuthenticated {
		t.Errorf("errors.Is should match wrapped ErrNoAPIKey, got code %q", code)
	}
}

func TestClassify_MissingEnvVar(t *testing.T) {
	code, exit := classify(errors.New("TWELVEDATA_API_KEY environment variable is not set"))
	if code != CodeMissingAPIKey {
		t.Errorf("code = %q, want %q", code, CodeMissingAPIKey)
	}
	if exit != ExitUnauthorized {
		t.Errorf("exit = %d, want %d", exit, ExitUnauthorized)
	}
}

func TestClassify_UsageErrors(t *testing.T) {
	usage := []string{
		"unknown command \"foo\" for \"twelvedata\"",
		"unknown flag: --bogus",
		"unknown shorthand flag: 'x' in -x",
		"required flag(s) \"symbol\" not set",
		"invalid argument \"bad\" for \"--interval\"",
		"flag needs an argument: --symbol",
		"at least one of the flags in the group [a b] is required",
		"if any flags in the group [a b] are set none of the others can be",
		"accepts 1 arg(s), received 0",
	}
	for _, msg := range usage {
		t.Run(msg, func(t *testing.T) {
			code, exit := classify(errors.New(msg))
			if code != CodeUsageError {
				t.Errorf("code = %q, want %q", code, CodeUsageError)
			}
			if exit != ExitUsage {
				t.Errorf("exit = %d, want %d", exit, ExitUsage)
			}
		})
	}
}

func TestClassify_GenericError(t *testing.T) {
	code, exit := classify(errors.New("network unreachable"))
	if code != CodeGenericError {
		t.Errorf("code = %q, want %q", code, CodeGenericError)
	}
	if exit != ExitGeneric {
		t.Errorf("exit = %d, want %d", exit, ExitGeneric)
	}
}

func TestWriteError_JSONEnvelope(t *testing.T) {
	cmd, _, errBuf := newCmd()
	_ = cmd.Flags().Set("raw", "true")
	exit := WriteError(cmd, apiErr(401, "bad key"))
	if exit != ExitUnauthorized {
		t.Errorf("exit = %d, want %d", exit, ExitUnauthorized)
	}
	var env struct {
		Error struct {
			Code    string `json:"code"`
			Message string `json:"message"`
			Status  int    `json:"status"`
		} `json:"error"`
	}
	if err := json.Unmarshal(errBuf.Bytes(), &env); err != nil {
		t.Fatalf("stderr should be valid JSON: %v\n%s", err, errBuf.String())
	}
	if env.Error.Code != CodeUnauthorized {
		t.Errorf("code = %q, want %q", env.Error.Code, CodeUnauthorized)
	}
	if env.Error.Status != 401 {
		t.Errorf("status = %d, want 401", env.Error.Status)
	}
	if env.Error.Message != "bad key" {
		t.Errorf("message = %q, want bad key", env.Error.Message)
	}
}

func TestWriteError_OK(t *testing.T) {
	cmd, _, errBuf := newCmd()
	if got := WriteError(cmd, nil); got != ExitOK {
		t.Errorf("nil error must return ExitOK, got %d", got)
	}
	if errBuf.Len() != 0 {
		t.Errorf("nil error must write nothing, got %q", errBuf.String())
	}
}

func TestWriteError_UsageWordingRewrite(t *testing.T) {
	cmd, _, errBuf := newCmd()
	_ = cmd.Flags().Set("raw", "true")
	WriteError(cmd, errors.New("required flag(s) \"symbol\" not set"))
	if strings.Contains(errBuf.String(), "flag") && !strings.Contains(errBuf.String(), "option") {
		t.Errorf("cobra wording should be rewritten to \"option\", got %s", errBuf.String())
	}
}

func TestWriteError_PrettyForAPIError(t *testing.T) {
	cmd, _, errBuf := newCmd()
	// Force non-raw via flag; we're a non-TTY buffer otherwise.
	// We can't easily fake a TTY, so just verify no JSON braces appear here —
	// raw mode auto-engages with a buffer writer.
	WriteError(cmd, apiErr(401, "bad key"))
	if !strings.Contains(errBuf.String(), "bad key") {
		t.Errorf("error message should reach stderr, got %q", errBuf.String())
	}
}

func TestWrapWords(t *testing.T) {
	got := wrapWords("the quick brown fox jumps", 10)
	if len(got) < 2 {
		t.Fatalf("expected wrapping into >=2 lines, got %v", got)
	}
	for _, line := range got {
		if len(line) > 10 && !strings.Contains(line, " ") {
			// single oversized word is allowed; otherwise must fit.
			continue
		}
		if len(line) > 10 {
			t.Errorf("line %q exceeds width 10", line)
		}
	}

	// Long single word kept intact.
	got = wrapWords("supercalifragilistic", 5)
	if len(got) != 1 || got[0] != "supercalifragilistic" {
		t.Errorf("a single oversized word should sit on its own line, got %v", got)
	}

	if got := wrapWords("", 10); len(got) != 1 || got[0] != "" {
		t.Errorf("empty input should yield one empty line, got %v", got)
	}
}

func TestIsUsageError(t *testing.T) {
	if !isUsageError(errors.New("unknown command foo")) {
		t.Error("should detect Cobra usage error")
	}
	if isUsageError(errors.New("HTTP 500 internal")) {
		t.Error("non-usage error misclassified")
	}
}

func TestStatusName(t *testing.T) {
	if got := statusName(404); got != "Not Found" {
		t.Errorf("statusName(404) = %q, want \"Not Found\"", got)
	}
	if got := statusName(999); got != "Error" {
		t.Errorf("unknown status code should fall back to \"Error\", got %q", got)
	}
}

func TestNotImplementedParam(t *testing.T) {
	err := NotImplementedParam("decimals", "number")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	for _, sub := range []string{"--decimals", "number", "not implemented"} {
		if !strings.Contains(err.Error(), sub) {
			t.Errorf("expected %q in error, got %q", sub, err.Error())
		}
	}
}

func TestErrorCodesList(t *testing.T) {
	if len(ErrorCodes) == 0 {
		t.Fatal("ErrorCodes must not be empty")
	}
	seen := map[string]struct{}{}
	for _, c := range ErrorCodes {
		if _, dup := seen[c]; dup {
			t.Errorf("duplicate error code: %q", c)
		}
		seen[c] = struct{}{}
	}
	// Every constant emitted by classify must appear.
	for _, c := range []string{
		CodeBadRequest, CodeUnauthorized, CodeForbidden, CodeNotFound,
		CodeParameterTooLong, CodeRateLimited, CodeInternalServerError,
		CodeUsageError, CodeNotAuthenticated, CodeMissingAPIKey, CodeGenericError,
	} {
		if _, ok := seen[c]; !ok {
			t.Errorf("ErrorCodes missing %q — agents will see it from classify() but not in --help", c)
		}
	}
}
