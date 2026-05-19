package commands

import (
	"strings"
	"testing"
)

// In setupTestEnv we set TERM=dumb, which makes output.IsRaw() return true
// even without --raw. The browser commands take the raw-mode branch (print
// URL to stdout, never spawn a process) so these tests never launch a real
// browser.

func TestDocsCmd_RawPrintsURL(t *testing.T) {
	setupTestEnv(t)
	stdout, stderr, err := runRoot(t, nil, "docs")
	if err != nil {
		t.Fatalf("docs: %v", err)
	}
	if got := strings.TrimSpace(stdout.String()); got != docsURL {
		t.Errorf("stdout = %q, want %q", got, docsURL)
	}
	if stderr.Len() != 0 {
		t.Errorf("expected empty stderr in raw mode, got %q", stderr.String())
	}
}

func TestDashboardCmd_RawPrintsURL(t *testing.T) {
	setupTestEnv(t)
	stdout, stderr, err := runRoot(t, nil, "dashboard")
	if err != nil {
		t.Fatalf("dashboard: %v", err)
	}
	if got := strings.TrimSpace(stdout.String()); got != dashboardURL {
		t.Errorf("stdout = %q, want %q", got, dashboardURL)
	}
	if stderr.Len() != 0 {
		t.Errorf("expected empty stderr in raw mode, got %q", stderr.String())
	}
}

func TestBrowserURLs_AreHTTPS(t *testing.T) {
	// Cheap guard against accidentally regressing the URL constants — both
	// must be https://twelvedata.com/... to keep the install hint trustworthy.
	for _, url := range []string{docsURL, dashboardURL} {
		if !strings.HasPrefix(url, "https://twelvedata.com/") {
			t.Errorf("browser URL %q not under https://twelvedata.com/", url)
		}
	}
}
