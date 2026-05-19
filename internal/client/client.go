package client

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/twelvedata/twelvedata-go/twelvedata"

	"github.com/twelvedata/twelvedata-cli/internal/auth"
)

// apiBaseURL pins the API host so the SDK's TWELVEDATA_API_BASE_URL env-var
// fallback can't redirect requests (and the persisted key) elsewhere.
const apiBaseURL = "https://api.twelvedata.com"

// defaultHTTPTimeout caps every API call so a stalled network never hangs the
// CLI indefinitely. The SDK defaults to http.DefaultClient (no Timeout) when
// none is supplied; we override that. 120s comfortably covers the slowest
// expected endpoint (large time_series) while still bounding indefinite
// stalls. Overridable via TWELVEDATA_HTTP_TIMEOUT.
const defaultHTTPTimeout = 120 * time.Second

// httpTimeoutEnv is the env var that overrides defaultHTTPTimeout. The value
// is a Go duration string (e.g. "30s", "2m", "1m30s"). A bare integer is
// accepted as seconds. "0" disables the timeout (advanced; re-introduces
// indefinite-hang risk). An unparseable value falls back to the default.
const httpTimeoutEnv = "TWELVEDATA_HTTP_TIMEOUT"

// New constructs a Twelve Data API client using the auth resolution chain:
// --api-key flag → TWELVEDATA_API_KEY env → secure storage / config file for
// the active profile. The resolved key is passed to NewConfig as-is so the
// SDK's internal env-var fallback never triggers and the source remains
// inspectable via `twelvedata whoami`.
func New(cmd *cobra.Command) (*twelvedata.APIClient, error) {
	keyFlag, _ := cmd.Flags().GetString("api-key")
	profileFlag, _ := cmd.Flags().GetString("profile")
	resolved, err := auth.ResolveAPIKey(keyFlag, profileFlag)
	if err != nil {
		return nil, err
	}
	cfg, err := twelvedata.NewConfigWithSource("cli", resolved.Key, apiBaseURL)
	if err != nil {
		return nil, err
	}
	cfg.HTTPClient = &http.Client{Timeout: resolveHTTPTimeout()}
	return twelvedata.NewAPIClient(cfg), nil
}

// resolveHTTPTimeout reads TWELVEDATA_HTTP_TIMEOUT and returns the effective
// timeout. Accepts a Go duration string ("30s", "2m") or a bare integer
// treated as seconds. Invalid values fall back to defaultHTTPTimeout silently
// so a typo doesn't break every command — set --raw and run `twelvedata
// doctor` to surface env-config issues during diagnosis.
func resolveHTTPTimeout() time.Duration {
	raw := strings.TrimSpace(os.Getenv(httpTimeoutEnv))
	if raw == "" {
		return defaultHTTPTimeout
	}
	if d, err := time.ParseDuration(raw); err == nil {
		return d
	}
	// Bare integer → seconds, for ergonomic "TWELVEDATA_HTTP_TIMEOUT=60".
	if d, err := time.ParseDuration(raw + "s"); err == nil {
		return d
	}
	return defaultHTTPTimeout
}
