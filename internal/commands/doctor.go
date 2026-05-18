package commands

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/twelvedata/twelvedata-go/twelvedata"

	"github.com/twelvedata/twelvedata-cli/internal/auth"
	"github.com/twelvedata/twelvedata-cli/internal/client"
	"github.com/twelvedata/twelvedata-cli/internal/output"
	"github.com/twelvedata/twelvedata-cli/internal/version"
)

const (
	doctorAPITimeout = 5 * time.Second
	githubLatestURL  = "https://api.github.com/repos/twelvedata/twelvedata-cli/releases/latest"

	doctorStatusPass = "pass"
	doctorStatusWarn = "warn"
	doctorStatusFail = "fail"
)

type doctorCheck struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`
}

type doctorReport struct {
	OK     bool          `json:"ok"`
	Checks []doctorCheck `json:"checks"`
}

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Diagnose the CLI setup (version, key, storage, API access)",
	Long: `Run four sequential checks against the local installation and the Twelve Data API:

  CLI Version          Compare the installed version against the latest GitHub release.
  API Key              Confirm a key is resolvable via flag, env var, or credentials file.
  Credential Storage   Whether keys live in OS secure storage or in a plaintext file.
  API Validation       Call /api_usage to confirm the resolved key is accepted.

Each check reports pass | warn | fail; the command exits 1 if any check is fail.
Use --raw (or pipe stdout) to get the result as a JSON envelope suitable for CI.`,
	Example: `  td doctor
  td doctor --raw
  td doctor --profile staging`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		keyFlag, _ := cmd.Flags().GetString("api-key")
		profileFlag, _ := cmd.Flags().GetString("profile")

		sp := output.StartSpinner(cmd)
		checks := []doctorCheck{
			checkCLIVersion(ctx),
			checkAPIKey(keyFlag, profileFlag),
			checkCredentialStorage(),
			checkAPIValidation(cmd, keyFlag, profileFlag),
		}
		sp.Stop()

		hasFails := false
		for _, c := range checks {
			if c.Status == doctorStatusFail {
				hasFails = true
				break
			}
		}
		report := doctorReport{OK: !hasFails, Checks: checks}

		if output.IsRaw(cmd) {
			enc := json.NewEncoder(cmd.OutOrStdout())
			enc.SetIndent("", "  ")
			if err := enc.Encode(report); err != nil {
				return err
			}
		} else {
			renderDoctorHuman(cmd.OutOrStdout(), report)
		}

		if hasFails {
			os.Exit(1)
		}
		return nil
	},
}

func checkCLIVersion(parent context.Context) doctorCheck {
	c := doctorCheck{Name: "CLI Version"}
	unknown := func() doctorCheck {
		c.Status = doctorStatusWarn
		c.Message = fmt.Sprintf("v%s (could not check for updates)", version.Version)
		return c
	}

	ctx, cancel := context.WithTimeout(parent, doctorAPITimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, githubLatestURL, nil)
	if err != nil {
		return unknown()
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return unknown()
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return unknown()
	}
	var body struct {
		TagName    string `json:"tag_name"`
		Prerelease bool   `json:"prerelease"`
		Draft      bool   `json:"draft"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil || body.TagName == "" {
		return unknown()
	}
	if body.Prerelease || body.Draft {
		return unknown()
	}
	latest := strings.TrimPrefix(body.TagName, "v")
	if latest == version.Version {
		c.Status = doctorStatusPass
		c.Message = fmt.Sprintf("v%s (latest)", version.Version)
		return c
	}
	c.Status = doctorStatusWarn
	c.Message = fmt.Sprintf("v%s (latest: v%s)", version.Version, latest)
	c.Detail = "Update available"
	return c
}

func checkAPIKey(keyFlag, profileFlag string) doctorCheck {
	c := doctorCheck{Name: "API Key"}
	resolved, err := auth.ResolveAPIKey(keyFlag, profileFlag)
	if err != nil {
		c.Status = doctorStatusFail
		if errors.Is(err, auth.ErrNoAPIKey) {
			c.Message = "No API key found"
			c.Detail = "Run: td login"
		} else {
			c.Message = err.Error()
		}
		return c
	}
	c.Status = doctorStatusPass
	src := doctorSourceLabel(resolved.Source)
	if resolved.Profile != "" {
		c.Message = fmt.Sprintf("%s (source: %s, profile: %s)", auth.MaskKey(resolved.Key), src, resolved.Profile)
	} else {
		c.Message = fmt.Sprintf("%s (source: %s)", auth.MaskKey(resolved.Key), src)
	}
	return c
}

func checkCredentialStorage() doctorCheck {
	c := doctorCheck{Name: "Credential Storage"}
	backend := auth.GetBackend()
	if backend.IsSecure() {
		c.Status = doctorStatusPass
		c.Message = backend.Name()
		return c
	}
	c.Status = doctorStatusWarn
	c.Message = backend.Name()
	// Surface unexpected fallback: secure was requested but file is in use.
	requested := os.Getenv("TWELVEDATA_CREDENTIAL_STORE") == string(auth.StorageSecure)
	if !requested {
		creds, _ := auth.ReadCredentials()
		if creds != nil && creds.Storage == auth.StorageSecure {
			requested = true
		}
	}
	if requested {
		c.Detail = "Secure backend unavailable — falling back to plaintext"
	}
	return c
}

func checkAPIValidation(cmd *cobra.Command, keyFlag, profileFlag string) doctorCheck {
	c := doctorCheck{Name: "API Validation"}
	if _, err := auth.ResolveAPIKey(keyFlag, profileFlag); err != nil {
		c.Status = doctorStatusFail
		c.Message = "Skipped — no API key"
		return c
	}
	api, err := client.New(cmd)
	if err != nil {
		c.Status = doctorStatusFail
		c.Message = err.Error()
		return c
	}
	ctx, cancel := context.WithTimeout(cmd.Context(), doctorAPITimeout)
	defer cancel()
	_, _, callErr := api.AdvancedAPI.GetApiUsage(ctx).Execute()
	if callErr == nil {
		c.Status = doctorStatusPass
		c.Message = "API key accepted"
		return c
	}
	if errors.Is(callErr, context.DeadlineExceeded) {
		c.Status = doctorStatusWarn
		c.Message = "Request timed out"
		return c
	}
	// The API returns HTTP 200 with an error body for invalid keys, so the
	// SDK's typed error is the source of truth for the status — not the raw
	// http.Response.
	var apiErr twelvedata.TwelvedataApiError
	if errors.As(callErr, &apiErr) {
		status := apiErr.GetStatusCode()
		switch status {
		case http.StatusUnauthorized, http.StatusForbidden:
			c.Status = doctorStatusFail
			c.Message = "API key rejected"
			c.Detail = strings.TrimSpace(apiErr.GetMessage())
			return c
		}
		c.Status = doctorStatusWarn
		c.Message = fmt.Sprintf("API error [%d]", status)
		c.Detail = strings.TrimSpace(apiErr.GetMessage())
		return c
	}
	c.Status = doctorStatusWarn
	c.Message = "Network error"
	if msg := strings.TrimSpace(callErr.Error()); msg != "" {
		c.Detail = msg
	}
	return c
}

func doctorSourceLabel(s auth.Source) string {
	switch s {
	case auth.SourceFlag:
		return "flag"
	case auth.SourceEnv:
		return "env"
	case auth.SourceConfig:
		return "config file"
	case auth.SourceSecure:
		return "secure storage"
	default:
		return string(s)
	}
}

func renderDoctorHuman(w io.Writer, r doctorReport) {
	fmt.Fprintln(w)
	fmt.Fprintln(w, "  Twelve Data Doctor")
	fmt.Fprintln(w)
	for _, c := range r.Checks {
		glyph := "✓"
		switch c.Status {
		case doctorStatusWarn:
			glyph = "!"
		case doctorStatusFail:
			glyph = "✗"
		}
		fmt.Fprintf(w, "  %s %-20s %s\n", glyph, c.Name, c.Message)
		if c.Detail != "" {
			fmt.Fprintf(w, "    %s\n", c.Detail)
		}
	}
	fmt.Fprintln(w)
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
