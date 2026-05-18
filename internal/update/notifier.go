// Package update implements a lightweight "new release available" notifier.
//
// MaybeNotify reads a cached check from disk and, if the cache is older than
// cacheTTL, fetches the latest GitHub release with a short timeout. When the
// remote version is newer than the embedded build version, it prints a single
// line to stderr suggesting the upgrade command. The check is gated so it
// never runs in machine mode — see shouldSkip.
package update

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/twelvedata/twelvedata-cli/internal/output"
	"github.com/twelvedata/twelvedata-cli/internal/version"
)

const (
	githubReleasesURL = "https://api.github.com/repos/twelvedata/twelvedata-cli/releases/latest"
	upgradeCommand    = "go install github.com/twelvedata/twelvedata-cli/cmd/td@v%s"
	cacheDirName      = "twelvedata-cli"
	cacheFileName     = "update-check.json"
	cacheTTL          = 24 * time.Hour
	fetchTimeout      = 2 * time.Second
	maxResponseBytes  = 1 << 16
	envOptOut         = "TWELVEDATA_NO_UPDATE_NOTIFIER"
)

type cacheState struct {
	LatestVersion string    `json:"latest_version"`
	CheckedAt     time.Time `json:"checked_at"`
}

// MaybeNotify checks for a newer release and, if one exists, prints a one-line
// hint to stderr. Safe to call from rootCmd's PersistentPostRun.
//
// All I/O errors are swallowed — the notifier is a courtesy, never a blocker.
func MaybeNotify(cmd *cobra.Command) {
	if shouldSkip(cmd) {
		return
	}

	state, _ := readCache()
	if isStale(state) {
		if latest, err := fetchLatest(); err == nil {
			state = cacheState{LatestVersion: latest, CheckedAt: time.Now()}
			_ = writeCache(state)
		}
	}
	if state.LatestVersion == "" {
		return
	}
	if compareVersions(state.LatestVersion, version.Version) <= 0 {
		return
	}
	printNotice(cmd, state.LatestVersion)
}

func shouldSkip(cmd *cobra.Command) bool {
	if v := os.Getenv(envOptOut); v != "" && v != "0" {
		return true
	}
	// Machine mode (--raw, piped stdout, CI, TERM=dumb) — never emit cosmetic
	// stderr output, agents and scripts get noisy logs otherwise.
	if output.IsRaw(cmd) {
		return true
	}
	// Skip when the embedded version is empty or a dev placeholder — comparing
	// against an unset version produces spurious "newer" notices.
	v := strings.TrimSpace(version.Version)
	if v == "" || strings.EqualFold(v, "dev") {
		return true
	}
	return false
}

func isStale(s cacheState) bool {
	if s.CheckedAt.IsZero() {
		return true
	}
	return time.Since(s.CheckedAt) > cacheTTL
}

func fetchLatest() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), fetchTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, githubReleasesURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "twelvedata-cli/"+version.Version)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("github releases: %s", resp.Status)
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, maxResponseBytes))
	if err != nil {
		return "", err
	}
	var payload struct {
		TagName string `json:"tag_name"`
	}
	if err := json.Unmarshal(body, &payload); err != nil {
		return "", err
	}
	return strings.TrimPrefix(payload.TagName, "v"), nil
}

func cachePath() (string, error) {
	base, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, cacheDirName, cacheFileName), nil
}

func readCache() (cacheState, error) {
	p, err := cachePath()
	if err != nil {
		return cacheState{}, err
	}
	b, err := os.ReadFile(p)
	if err != nil {
		return cacheState{}, err
	}
	var s cacheState
	if err := json.Unmarshal(b, &s); err != nil {
		return cacheState{}, err
	}
	return s, nil
}

func writeCache(s cacheState) error {
	p, err := cachePath()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		return err
	}
	b, err := json.Marshal(s)
	if err != nil {
		return err
	}
	return os.WriteFile(p, b, 0o644)
}

// compareVersions does a SemVer-ish compare on "MAJOR.MINOR.PATCH" strings.
// Anything after a `-` or `+` (pre-release / build metadata) is ignored. A
// non-numeric segment compares as 0. Returns >0 if a>b, 0 if equal, <0 if a<b.
func compareVersions(a, b string) int {
	ap := splitVersion(a)
	bp := splitVersion(b)
	for i := range 3 {
		if ap[i] != bp[i] {
			return ap[i] - bp[i]
		}
	}
	return 0
}

func splitVersion(v string) [3]int {
	v = strings.TrimPrefix(v, "v")
	if i := strings.IndexAny(v, "-+"); i >= 0 {
		v = v[:i]
	}
	parts := strings.SplitN(v, ".", 3)
	var out [3]int
	for i := 0; i < 3 && i < len(parts); i++ {
		out[i], _ = strconv.Atoi(parts[i])
	}
	return out
}

func printNotice(cmd *cobra.Command, latest string) {
	fmt.Fprintf(cmd.ErrOrStderr(),
		"\nA new version of td is available: v%s → v%s\n  Run: "+upgradeCommand+"\n  Disable: %s=1\n",
		version.Version, latest, latest, envOptOut)
}
