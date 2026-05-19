package output

import (
	"bytes"
	"net/http"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

// newCmd builds a cobra.Command with the persistent flags we look at (output,
// raw) plus an out/err buffer pair. Returned buffers expose what the command
// wrote so each test can assert without juggling os.Pipe.
func newCmd() (*cobra.Command, *bytes.Buffer, *bytes.Buffer) {
	cmd := &cobra.Command{Use: "t"}
	cmd.Flags().StringP("output", "o", "", "")
	cmd.Flags().Bool("raw", false, "")
	out, errBuf := &bytes.Buffer{}, &bytes.Buffer{}
	cmd.SetOut(out)
	cmd.SetErr(errBuf)
	return cmd, out, errBuf
}

func TestResolveFormat_Default(t *testing.T) {
	cmd, _, _ := newCmd()
	got, err := ResolveFormat(cmd)
	if err != nil {
		t.Fatal(err)
	}
	if got != FormatJSON {
		t.Errorf("default ResolveFormat = %q, want json", got)
	}
}

func TestResolveFormat_CSV(t *testing.T) {
	cmd, _, _ := newCmd()
	_ = cmd.Flags().Set("output", "csv")
	got, err := ResolveFormat(cmd)
	if err != nil {
		t.Fatal(err)
	}
	if got != FormatCSV {
		t.Errorf("ResolveFormat = %q, want csv", got)
	}
}

func TestResolveFormat_Invalid(t *testing.T) {
	cmd, _, _ := newCmd()
	_ = cmd.Flags().Set("output", "yaml")
	_, err := ResolveFormat(cmd)
	if err == nil {
		t.Fatal("expected error for invalid --output value")
	}
	if !strings.Contains(err.Error(), "yaml") {
		t.Errorf("error should echo bad value, got %q", err)
	}
}

func TestIsRaw(t *testing.T) {
	// Clear env for the parent test so each subtest starts clean.
	t.Setenv("CI", "")
	t.Setenv("GITHUB_ACTIONS", "")
	t.Setenv("TERM", "")

	t.Run("raw flag forces raw", func(t *testing.T) {
		cmd, _, _ := newCmd()
		_ = cmd.Flags().Set("raw", "true")
		if !IsRaw(cmd) {
			t.Error("--raw should force raw mode")
		}
	})

	t.Run("TERM=dumb forces raw", func(t *testing.T) {
		t.Setenv("TERM", "dumb")
		cmd, _, _ := newCmd()
		if !IsRaw(cmd) {
			t.Error("TERM=dumb should force raw mode")
		}
	})

	t.Run("CI env forces raw", func(t *testing.T) {
		t.Setenv("CI", "1")
		cmd, _, _ := newCmd()
		if !IsRaw(cmd) {
			t.Error("CI=1 should force raw mode")
		}
	})

	t.Run("GITHUB_ACTIONS env forces raw", func(t *testing.T) {
		t.Setenv("GITHUB_ACTIONS", "true")
		cmd, _, _ := newCmd()
		if !IsRaw(cmd) {
			t.Error("GITHUB_ACTIONS should force raw mode")
		}
	})

	t.Run("piped stdout (buffer) is raw", func(t *testing.T) {
		cmd, _, _ := newCmd()
		// SetOut is a bytes.Buffer — not a *os.File, so isTerminal == false.
		if !IsRaw(cmd) {
			t.Error("non-TTY stdout should imply raw mode")
		}
	})
}

func TestRender_JSON(t *testing.T) {
	cmd, out, _ := newCmd()
	payload := map[string]any{"a": 1, "b": "two"}
	if err := Render(cmd, payload, nil, nil); err != nil {
		t.Fatal(err)
	}
	body := out.String()
	if !strings.Contains(body, `"a"`) || !strings.Contains(body, `"two"`) {
		t.Errorf("expected pretty JSON, got: %s", body)
	}
}

func TestRender_JSONReturnsCallErr(t *testing.T) {
	cmd, _, _ := newCmd()
	want := errStub("boom")
	if err := Render(cmd, nil, nil, want); err != want {
		t.Errorf("Render must return call error verbatim, got %v", err)
	}
}

func TestRender_CSV_HappyPath(t *testing.T) {
	cmd, out, _ := newCmd()
	_ = cmd.Flags().Set("output", "csv")

	header := http.Header{}
	header.Set("Content-Type", "text/csv")
	resp := &http.Response{
		StatusCode: 200,
		Header:     header,
		Body:       newBody("a,b\n1,2\n"),
	}
	if err := Render(cmd, nil, resp, nil); err != nil {
		t.Fatal(err)
	}
	if got := out.String(); got != "a,b\n1,2\n" {
		t.Errorf("Render(csv) wrote %q", got)
	}
}

func TestRender_CSV_WrongContentType(t *testing.T) {
	cmd, _, _ := newCmd()
	_ = cmd.Flags().Set("output", "csv")
	header := http.Header{}
	header.Set("Content-Type", "application/json")
	resp := &http.Response{StatusCode: 200, Header: header, Body: newBody("{}")}
	err := Render(cmd, nil, resp, nil)
	if err == nil {
		t.Fatal("expected error when --output csv but server returned non-CSV")
	}
	if !strings.Contains(err.Error(), "csv") {
		t.Errorf("error should mention csv: %v", err)
	}
}

func TestRender_CSV_PassesThroughCallErrOnNon2xx(t *testing.T) {
	cmd, _, _ := newCmd()
	_ = cmd.Flags().Set("output", "csv")
	header := http.Header{}
	header.Set("Content-Type", "application/json")
	resp := &http.Response{StatusCode: 401, Header: header, Body: newBody("")}
	want := errStub("unauthorized")
	if err := Render(cmd, nil, resp, want); err != want {
		t.Errorf("non-2xx must propagate the call error, got %v", err)
	}
}

func TestRender_InvalidFormat(t *testing.T) {
	cmd, _, _ := newCmd()
	_ = cmd.Flags().Set("output", "xml")
	if err := Render(cmd, nil, nil, nil); err == nil {
		t.Error("Render should surface invalid --output value")
	}
}

// errStub is a tiny error implementation used to verify Render returns the
// caller's error verbatim rather than re-wrapping it.
type errStub string

func (e errStub) Error() string { return string(e) }

type stubBody struct {
	r *strings.Reader
}

func newBody(s string) *stubBody               { return &stubBody{r: strings.NewReader(s)} }
func (b *stubBody) Read(p []byte) (int, error) { return b.r.Read(p) }
func (b *stubBody) Close() error               { return nil }
