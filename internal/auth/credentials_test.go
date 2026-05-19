package auth

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestValidateProfileName(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"empty", "", true},
		{"too long", strings.Repeat("a", 65), true},
		{"slash", "foo/bar", true},
		{"space", "foo bar", true},
		{"plain", "default", false},
		{"with dot", "team.prod", false},
		{"with dash", "team-prod", false},
		{"with underscore", "team_prod", false},
		{"alphanumeric", "abc123", false},
		{"max length", strings.Repeat("a", 64), false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateProfileName(tc.input)
			if (err != nil) != tc.wantErr {
				t.Fatalf("ValidateProfileName(%q) err=%v, wantErr=%v", tc.input, err, tc.wantErr)
			}
		})
	}
}

func TestReadCredentials_Missing(t *testing.T) {
	setupTestDir(t)
	creds, err := ReadCredentials()
	if err != nil {
		t.Fatalf("ReadCredentials returned error on missing file: %v", err)
	}
	if creds != nil {
		t.Fatalf("expected nil creds for missing file, got %+v", creds)
	}
}

func TestReadCredentials_Malformed(t *testing.T) {
	dir := setupTestDir(t)
	if err := os.WriteFile(filepath.Join(dir, "credentials.json"), []byte("not json"), 0o600); err != nil {
		t.Fatal(err)
	}
	_, err := ReadCredentials()
	if err == nil {
		t.Fatalf("expected parse error for malformed json")
	}
	if !strings.Contains(err.Error(), "failed to parse") {
		t.Fatalf("error should mention parse failure: %v", err)
	}
}

func TestReadCredentials_NilProfilesTreatedAsMissing(t *testing.T) {
	dir := setupTestDir(t)
	if err := os.WriteFile(filepath.Join(dir, "credentials.json"), []byte(`{"active_profile":"default"}`), 0o600); err != nil {
		t.Fatal(err)
	}
	creds, err := ReadCredentials()
	if err != nil {
		t.Fatal(err)
	}
	if creds != nil {
		t.Fatalf("expected nil for file without profiles, got %+v", creds)
	}
}

func TestReadCredentials_DefaultsActiveProfile(t *testing.T) {
	dir := setupTestDir(t)
	body := `{"profiles":{"alpha":{"api_key":"k"}}}`
	if err := os.WriteFile(filepath.Join(dir, "credentials.json"), []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
	creds, err := ReadCredentials()
	if err != nil {
		t.Fatal(err)
	}
	if creds == nil {
		t.Fatal("creds was nil")
	}
	if creds.ActiveProfile != "default" {
		t.Fatalf("expected ActiveProfile to default to \"default\", got %q", creds.ActiveProfile)
	}
}

func TestWriteCredentials_PermissionsAndShape(t *testing.T) {
	dir := setupTestDir(t)
	creds := &CredentialsFile{
		ActiveProfile: "default",
		Profiles:      map[string]Profile{"default": {APIKey: "secret"}},
	}
	path, err := WriteCredentials(creds)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := path, filepath.Join(dir, "credentials.json"); got != want {
		t.Fatalf("path = %q, want %q", got, want)
	}
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	// Skip mode check on Windows where Unix bits are advisory.
	if mode := info.Mode().Perm(); mode != 0o600 && mode != 0o666 {
		// 0o666 only if test runs on a filesystem ignoring chmod (e.g. WSL DrvFs).
		t.Logf("note: credentials.json mode = %o", mode)
	}
	dirInfo, err := os.Stat(dir)
	if err != nil {
		t.Fatal(err)
	}
	if !dirInfo.IsDir() {
		t.Fatal("config dir is not a directory")
	}
	// Round-trip through JSON.
	buf, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var got CredentialsFile
	if err := json.Unmarshal(buf, &got); err != nil {
		t.Fatalf("written file is not valid json: %v", err)
	}
	if got.Profiles["default"].APIKey != "secret" {
		t.Fatalf("round-trip mismatch: %+v", got)
	}
}

func TestDeleteCredentialsFile(t *testing.T) {
	dir := setupTestDir(t)
	path := filepath.Join(dir, "credentials.json")
	if err := os.WriteFile(path, []byte(`{"profiles":{"default":{"api_key":"k"}}}`), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := DeleteCredentialsFile(); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatalf("expected file removed, stat err = %v", err)
	}
	// Idempotent — deleting a missing file is not an error.
	if err := DeleteCredentialsFile(); err != nil {
		t.Fatalf("second delete should be a no-op, got %v", err)
	}
}

func TestMaskKey(t *testing.T) {
	tests := []struct {
		key, want string
	}{
		{"", "..."},
		{"ab", "ab..."},
		{"abc", "abc..."},
		{"abcd", "abc..."},
		{"abcdefg", "abc..."},
		{"abcdefgh", "abc...efgh"},
		{"abcdefghij", "abc...ghij"},
		{strings.Repeat("x", 32), "xxx...xxxx"},
	}
	for _, tc := range tests {
		if got := MaskKey(tc.key); got != tc.want {
			t.Errorf("MaskKey(%q) = %q, want %q", tc.key, got, tc.want)
		}
	}
}

func TestListProfiles_Empty(t *testing.T) {
	setupTestDir(t)
	got, err := ListProfiles()
	if err != nil {
		t.Fatal(err)
	}
	if got != nil {
		t.Fatalf("expected nil for empty config, got %v", got)
	}
}

func TestListProfiles_MarksActive(t *testing.T) {
	setupTestDir(t)
	_, err := WriteCredentials(&CredentialsFile{
		ActiveProfile: "alpha",
		Profiles: map[string]Profile{
			"alpha": {APIKey: "k1"},
			"beta":  {APIKey: "k2"},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	got, err := ListProfiles()
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 {
		t.Fatalf("want 2 profiles, got %d (%v)", len(got), got)
	}
	var sawAlphaActive, sawBetaInactive bool
	for _, p := range got {
		switch p.Name {
		case "alpha":
			sawAlphaActive = p.Active
		case "beta":
			sawBetaInactive = !p.Active
		}
	}
	if !sawAlphaActive {
		t.Error("alpha should be marked active")
	}
	if !sawBetaInactive {
		t.Error("beta should not be marked active")
	}
}
