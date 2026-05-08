package output

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

type Format string

const (
	FormatTable Format = "table"
	FormatJSON  Format = "json"
	FormatCSV   Format = "csv"
)

// ResolveFormat picks an output format using this precedence:
//
//	--output flag > --quiet (forces json) > non-TTY/CI (json) > TTY default (table).
func ResolveFormat(cmd *cobra.Command) (Format, error) {
	out, _ := cmd.Flags().GetString("output")
	if out != "" {
		f := Format(out)
		switch f {
		case FormatJSON, FormatCSV, FormatTable:
			return f, nil
		default:
			return "", fmt.Errorf("invalid --output %q (want json, csv, or table)", out)
		}
	}
	quiet, _ := cmd.Flags().GetBool("quiet")
	if quiet {
		return FormatJSON, nil
	}
	if !isTerminal(os.Stdout) || isCI() {
		return FormatJSON, nil
	}
	return FormatTable, nil
}

// Render writes resp in the resolved format. For --output csv the typed resp is
// ignored and httpResp.Body (the raw CSV bytes returned by the API) is streamed
// instead — the SDK rewinds the body but cannot decode CSV, so it returns a
// benign decode error in that mode which we silently absorb here.
func Render(cmd *cobra.Command, resp any, httpResp *http.Response, callErr error) error {
	f, err := ResolveFormat(cmd)
	if err != nil {
		return err
	}

	if f == FormatCSV {
		if httpResp == nil {
			return callErr
		}
		defer httpResp.Body.Close()
		_, copyErr := io.Copy(cmd.OutOrStdout(), httpResp.Body)
		return copyErr
	}

	if callErr != nil {
		return callErr
	}

	enc := json.NewEncoder(cmd.OutOrStdout())
	enc.SetIndent("", "  ")
	return enc.Encode(resp)
}

func isTerminal(w io.Writer) bool {
	f, ok := w.(*os.File)
	if !ok {
		return false
	}
	return term.IsTerminal(int(f.Fd()))
}

func isCI() bool {
	for _, k := range []string{"CI", "GITHUB_ACTIONS"} {
		if os.Getenv(k) != "" {
			return true
		}
	}
	return false
}
