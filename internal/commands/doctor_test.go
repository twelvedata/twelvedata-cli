package commands

import (
	"bytes"
	"strings"
	"testing"

	"github.com/twelvedata/twelvedata-cli/internal/auth"
)

// NOTE: The full doctor RunE calls os.Exit(1) on any failed check, which is
// fatal in a test process. Tests target the pure check functions and the
// renderer directly. Network-touching checks (checkCLIVersion,
// checkAPIValidation) are deliberately out of scope here — they belong in the
// integration layer where outbound network is acceptable.

func TestCheckAPIKey_NoKeyReportsFail(t *testing.T) {
	setupTestEnv(t)
	c := checkAPIKey("", "")
	if c.Name != "API Key" {
		t.Errorf("Name = %q", c.Name)
	}
	if c.Status != doctorStatusFail {
		t.Errorf("Status = %q, want fail", c.Status)
	}
	if c.Message == "" {
		t.Error("Message should not be empty")
	}
	// Run-`twelvedata login` hint is the actionable detail in the no-key path.
	if !strings.Contains(strings.ToLower(c.Detail), "login") {
		t.Errorf("Detail should hint at `twelvedata login`, got %q", c.Detail)
	}
}

func TestCheckAPIKey_FlagSuppliesKey(t *testing.T) {
	setupTestEnv(t)
	c := checkAPIKey("flag_key_abcdef", "")
	if c.Status != doctorStatusPass {
		t.Errorf("Status = %q, want pass; msg=%q", c.Status, c.Message)
	}
	if !strings.Contains(c.Message, "flag") {
		t.Errorf("Message %q should mention source \"flag\"", c.Message)
	}
	if strings.Contains(c.Message, "flag_key_abcdef") {
		t.Error("raw key must not appear in the message")
	}
}

func TestCheckAPIKey_ConfigFileSourceShownWithProfile(t *testing.T) {
	setupTestEnv(t)
	if _, _, err := runRoot(t, nil, "login", "--key", "stored_key", "--profile", "staging", "--raw"); err != nil {
		t.Fatalf("setup: %v", err)
	}

	c := checkAPIKey("", "staging")
	if c.Status != doctorStatusPass {
		t.Errorf("Status = %q, want pass", c.Status)
	}
	if !strings.Contains(c.Message, "staging") {
		t.Errorf("Message %q should include profile \"staging\"", c.Message)
	}
}

func TestCheckCredentialStorage_FileWarnsAndSurfacesFallback(t *testing.T) {
	setupTestEnv(t)
	// setupTestEnv forces TWELVEDATA_CREDENTIAL_STORE=file, so this is the
	// "user explicitly chose file" path → no fallback Detail.
	c := checkCredentialStorage()
	if c.Status != doctorStatusWarn {
		t.Errorf("Status = %q, want warn (file backend isn't secure)", c.Status)
	}
	if c.Message == "" {
		t.Error("Message should report the backend name")
	}
}

func TestCheckCredentialStorage_SurfacesUnexpectedSecureFallback(t *testing.T) {
	setupTestEnv(t)
	// User did not request file, but credentials.json says secure_storage.
	// On a host without a real keyring the chosen backend is file → Detail
	// should explain the fallback.
	t.Setenv("TWELVEDATA_CREDENTIAL_STORE", "") // clear forced choice
	auth.ResetBackend()
	creds := &auth.CredentialsFile{
		ActiveProfile: "default",
		Storage:       auth.StorageSecure,
		Profiles:      map[string]auth.Profile{"default": {}},
	}
	if _, err := auth.WriteCredentials(creds); err != nil {
		t.Fatalf("seed creds: %v", err)
	}

	c := checkCredentialStorage()
	// On a host where the keyring IS available, the backend is secure and the
	// status is pass — skip the assertion in that case to keep the test
	// portable across CI runners.
	if c.Status == doctorStatusPass {
		t.Skip("keyring backend is available on this host; fallback path not exercised")
	}
	if c.Status != doctorStatusWarn {
		t.Fatalf("Status = %q, want warn", c.Status)
	}
	if !strings.Contains(strings.ToLower(c.Detail), "fall") {
		t.Errorf("Detail = %q, expected mention of fallback", c.Detail)
	}
}

func TestDoctorSourceLabel(t *testing.T) {
	tests := map[auth.Source]string{
		auth.SourceFlag:        "flag",
		auth.SourceEnv:         "env",
		auth.SourceConfig:      "config file",
		auth.SourceSecure:      "secure storage",
		auth.Source("custom"): "custom",
	}
	for src, want := range tests {
		if got := doctorSourceLabel(src); got != want {
			t.Errorf("doctorSourceLabel(%q) = %q, want %q", string(src), got, want)
		}
	}
}

func TestRenderDoctorHuman_GlyphPerStatus(t *testing.T) {
	r := doctorReport{
		OK: false,
		Checks: []doctorCheck{
			{Name: "CLI Version", Status: doctorStatusPass, Message: "v1.0"},
			{Name: "API Key", Status: doctorStatusWarn, Message: "missing"},
			{Name: "Storage", Status: doctorStatusFail, Message: "no", Detail: "broken"},
		},
	}
	var buf bytes.Buffer
	renderDoctorHuman(&buf, r)
	out := buf.String()

	// Order of glyphs matches the slice order.
	pass := strings.Index(out, "✓")
	warn := strings.Index(out, "!")
	fail := strings.Index(out, "✗")
	if pass < 0 || warn < 0 || fail < 0 {
		t.Fatalf("missing glyph in output:\n%s", out)
	}
	if !(pass < warn && warn < fail) {
		t.Errorf("glyph order does not match check order:\n%s", out)
	}
	if !strings.Contains(out, "broken") {
		t.Errorf("Detail not rendered:\n%s", out)
	}
}
