package update

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/spf13/cobra"

	"github.com/twelvedata/twelvedata-cli/internal/version"
)

// newCmd returns a Cobra command with the same persistent flags the root
// command declares, so output.IsRaw can be probed.
func newCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "twelvedata"}
	cmd.Flags().Bool("raw", false, "")
	cmd.Flags().StringP("output", "o", "", "")
	cmd.SetOut(&bytes.Buffer{})
	cmd.SetErr(&bytes.Buffer{})
	return cmd
}

func TestCompareVersions(t *testing.T) {
	cases := []struct {
		a, b string
		want int // sign only matters
	}{
		{"1.0.0", "1.0.0", 0},
		{"v1.0.0", "1.0.0", 0},
		{"1.0.1", "1.0.0", 1},
		{"1.1.0", "1.0.9", 1},
		{"2.0.0", "1.999.999", 1},
		{"1.0.0", "1.0.1", -1},
		// Pre-release / build metadata stripped:
		{"1.2.3-beta", "1.2.3", 0},
		{"1.2.3+meta", "1.2.3", 0},
		// Non-numeric segment compares as 0:
		{"1.x.0", "1.0.0", 0},
		// Missing segments treated as 0:
		{"2", "1.999.999", 1},
	}
	for _, tc := range cases {
		got := compareVersions(tc.a, tc.b)
		gotSign := sign(got)
		if gotSign != tc.want {
			t.Errorf("compareVersions(%q,%q) sign = %d, want %d (raw %d)", tc.a, tc.b, gotSign, tc.want, got)
		}
	}
}

func sign(n int) int {
	switch {
	case n > 0:
		return 1
	case n < 0:
		return -1
	}
	return 0
}

func TestSplitVersion(t *testing.T) {
	cases := []struct {
		in   string
		want [3]int
	}{
		{"1.2.3", [3]int{1, 2, 3}},
		{"v1.2.3", [3]int{1, 2, 3}},
		{"1.2.3-beta1", [3]int{1, 2, 3}},
		{"1.2.3+meta", [3]int{1, 2, 3}},
		{"", [3]int{0, 0, 0}},
		{"abc", [3]int{0, 0, 0}},
	}
	for _, tc := range cases {
		if got := splitVersion(tc.in); got != tc.want {
			t.Errorf("splitVersion(%q) = %v, want %v", tc.in, got, tc.want)
		}
	}
}

func TestIsStale(t *testing.T) {
	if !isStale(cacheState{}) {
		t.Error("zero-time cache must be stale")
	}
	fresh := cacheState{CheckedAt: time.Now()}
	if isStale(fresh) {
		t.Error("just-now cache must be fresh")
	}
	old := cacheState{CheckedAt: time.Now().Add(-25 * time.Hour)}
	if !isStale(old) {
		t.Error("25h-old cache must be stale")
	}
	// Exactly at the boundary is still fresh (>, not >=).
	atTTL := cacheState{CheckedAt: time.Now().Add(-cacheTTL + 1*time.Second)}
	if isStale(atTTL) {
		t.Error("cache within TTL must be fresh")
	}
}

func TestShouldSkip_OptOutEnv(t *testing.T) {
	t.Setenv(envOptOut, "1")
	cmd := newCmd()
	if !shouldSkip(cmd) {
		t.Error("opt-out env must skip")
	}

	// "0" is the documented "still on" sentinel — must not skip on env grounds.
	// (The buffer-writer still trips IsRaw, so we can't assert !shouldSkip here.)
	t.Setenv(envOptOut, "")
	t.Setenv("CI", "")
	t.Setenv("GITHUB_ACTIONS", "")
	t.Setenv("TERM", "xterm")
	// Without opt-out and with a sane version, we still skip when stdout isn't a
	// TTY — the test buffer triggers raw mode.
	if !shouldSkip(cmd) {
		t.Error("non-TTY stdout should imply skip via IsRaw")
	}
}

func TestShouldSkip_DevVersion(t *testing.T) {
	t.Setenv(envOptOut, "")
	t.Setenv("CI", "")
	t.Setenv("GITHUB_ACTIONS", "")

	orig := version.Version
	t.Cleanup(func() { version.Version = orig })

	version.Version = "dev"
	cmd := newCmd()
	if !shouldSkip(cmd) {
		t.Error("dev version must skip")
	}

	version.Version = ""
	if !shouldSkip(cmd) {
		t.Error("empty version must skip")
	}

	version.Version = "DEV"
	if !shouldSkip(cmd) {
		t.Error("case-insensitive dev should still skip")
	}
}

func TestUpgradeHint(t *testing.T) {
	cases := []struct {
		name    string
		path    string
		latest  string
		hint    string
		isURL   bool
	}{
		{"install script unix", "/home/u/.twelvedata/bin/twelvedata", "1.0.1", "curl -fsSL https://raw.githubusercontent.com/twelvedata/twelvedata-cli/main/install.sh | bash", false},
		{"homebrew cellar", "/opt/homebrew/Cellar/twelvedata/1.0.0/bin/twelvedata", "1.0.1", "brew update && brew upgrade twelvedata", false},
		{"go install", "/home/u/go/bin/twelvedata", "1.0.1", "go install github.com/twelvedata/twelvedata-cli/cmd/twelvedata@v1.0.1", false},
		{"fallback github", "/usr/local/bin/twelvedata", "1.0.1", "https://github.com/twelvedata/twelvedata-cli/releases/latest", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			hint, isURL := upgradeHint(tc.path, tc.latest)
			if hint != tc.hint || isURL != tc.isURL {
				t.Errorf("upgradeHint(%q) = %q,%v; want %q,%v", tc.path, hint, isURL, tc.hint, tc.isURL)
			}
		})
	}
}

func TestCacheRoundTrip(t *testing.T) {
	// os.UserCacheDir reads XDG_CACHE_HOME on Linux, ~/Library/Caches on macOS,
	// %LocalAppData% on Windows. We set both XDG_CACHE_HOME and HOME so the test
	// is hermetic on Linux without depending on macOS-specific env knobs.
	dir := t.TempDir()
	t.Setenv("XDG_CACHE_HOME", dir)
	t.Setenv("HOME", dir)

	want := cacheState{LatestVersion: "9.9.9", CheckedAt: time.Now().UTC().Truncate(time.Second)}
	if err := writeCache(want); err != nil {
		t.Fatalf("writeCache: %v", err)
	}
	got, err := readCache()
	if err != nil {
		t.Fatalf("readCache: %v", err)
	}
	if got.LatestVersion != want.LatestVersion {
		t.Errorf("LatestVersion = %q, want %q", got.LatestVersion, want.LatestVersion)
	}
	if !got.CheckedAt.Equal(want.CheckedAt) {
		t.Errorf("CheckedAt = %v, want %v", got.CheckedAt, want.CheckedAt)
	}
}

func TestReadCache_Missing(t *testing.T) {
	t.Setenv("XDG_CACHE_HOME", t.TempDir())
	t.Setenv("HOME", t.TempDir())
	s, err := readCache()
	if err == nil {
		t.Error("missing cache should surface an error from os.ReadFile")
	}
	if s.LatestVersion != "" {
		t.Errorf("missing cache should yield zero state, got %+v", s)
	}
}

func TestReadCache_Malformed(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("XDG_CACHE_HOME", dir)
	t.Setenv("HOME", dir)
	p, err := cachePath()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(p, []byte("not json"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := readCache(); err == nil {
		t.Error("readCache should return parse error on malformed json")
	}
}

func TestMaybeNotify_SkipsInRawMode(t *testing.T) {
	cmd := newCmd()
	_ = cmd.Flags().Set("raw", "true")
	// Even with a recently-cached newer release, raw mode must suppress the notice.
	t.Setenv("XDG_CACHE_HOME", t.TempDir())
	t.Setenv("HOME", t.TempDir())
	p, _ := cachePath()
	_ = os.MkdirAll(filepath.Dir(p), 0o755)
	enc, _ := json.Marshal(cacheState{LatestVersion: "99.0.0", CheckedAt: time.Now()})
	_ = os.WriteFile(p, enc, 0o644)

	out := &bytes.Buffer{}
	cmd.SetErr(out)
	MaybeNotify(cmd)
	if out.Len() != 0 {
		t.Errorf("raw mode must produce no notice, got %q", out.String())
	}
}

func TestMaybeNotify_QuietWhenUpToDate(t *testing.T) {
	t.Setenv(envOptOut, "")
	t.Setenv("CI", "")
	t.Setenv("GITHUB_ACTIONS", "")
	t.Setenv("TERM", "xterm")
	t.Setenv("XDG_CACHE_HOME", t.TempDir())
	t.Setenv("HOME", t.TempDir())

	orig := version.Version
	t.Cleanup(func() { version.Version = orig })
	version.Version = "1.0.0"

	// Seed a cache entry showing we're already on the latest.
	p, _ := cachePath()
	_ = os.MkdirAll(filepath.Dir(p), 0o755)
	enc, _ := json.Marshal(cacheState{LatestVersion: "1.0.0", CheckedAt: time.Now()})
	_ = os.WriteFile(p, enc, 0o644)

	cmd := newCmd()
	out := &bytes.Buffer{}
	cmd.SetErr(out)
	MaybeNotify(cmd)
	if out.Len() != 0 {
		t.Errorf("up-to-date version must produce no notice, got %q", out.String())
	}
}
