package commands

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"

	"github.com/twelvedata/twelvedata-cli/internal/auth"
)

func TestLogin_KeyFlag_StoresAndActivates(t *testing.T) {
	dir := setupTestEnv(t)
	stdout, _, err := runRoot(t, nil, "login", "--key", "k_abcdef", "--profile", "staging", "--raw")
	if err != nil {
		t.Fatalf("unexpected error: %v\nstdout=%s", err, stdout.String())
	}

	// Persisted state must match the flag values.
	creds, err := auth.ReadCredentials()
	if err != nil {
		t.Fatalf("ReadCredentials: %v", err)
	}
	if creds == nil {
		t.Fatal("credentials.json not written")
	}
	if creds.ActiveProfile != "staging" {
		t.Errorf("active_profile = %q, want staging", creds.ActiveProfile)
	}
	if creds.Profiles["staging"].APIKey != "k_abcdef" {
		t.Errorf("stored key = %q, want k_abcdef", creds.Profiles["staging"].APIKey)
	}

	// Raw output shape.
	var payload map[string]any
	if err := json.Unmarshal(stdout.Bytes(), &payload); err != nil {
		t.Fatalf("stdout not JSON: %v\n%s", err, stdout.String())
	}
	if payload["success"] != true {
		t.Errorf("success = %v, want true", payload["success"])
	}
	if payload["profile"] != "staging" {
		t.Errorf("profile = %v, want staging", payload["profile"])
	}
	wantPath := filepath.Join(dir, "credentials.json")
	if got, _ := payload["config_path"].(string); got != wantPath {
		t.Errorf("config_path = %q, want %q", got, wantPath)
	}
}

func TestLogin_KeyStdin_ReadsAndTrims(t *testing.T) {
	setupTestEnv(t)
	stdin := strings.NewReader("  k_from_stdin  \n")
	_, _, err := runRoot(t, stdin, "login", "--key-stdin", "--raw")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	creds, _ := auth.ReadCredentials()
	if creds == nil {
		t.Fatal("credentials.json not written")
	}
	if got := creds.Profiles["default"].APIKey; got != "k_from_stdin" {
		t.Errorf("stored key = %q, want trimmed \"k_from_stdin\"", got)
	}
	if creds.ActiveProfile != "default" {
		t.Errorf("active_profile = %q, want default", creds.ActiveProfile)
	}
}

func TestLogin_KeyAndKeyStdin_MutuallyExclusive(t *testing.T) {
	setupTestEnv(t)
	_, _, err := runRoot(t, strings.NewReader("k"), "login", "--key", "literal", "--key-stdin", "--raw")
	if err == nil {
		t.Fatal("expected error when both --key and --key-stdin are set")
	}
	if !strings.Contains(err.Error(), "mutually exclusive") {
		t.Errorf("error message = %q, want mention of mutual exclusion", err.Error())
	}
}

func TestLogin_KeyStdin_EmptyInputErrors(t *testing.T) {
	setupTestEnv(t)
	_, _, err := runRoot(t, strings.NewReader("   \n"), "login", "--key-stdin", "--raw")
	if err == nil {
		t.Fatal("expected error for empty stdin")
	}
	if !strings.Contains(err.Error(), "empty input") {
		t.Errorf("error message = %q, want \"empty input\"", err.Error())
	}
}

func TestLogin_NoKeySource_NonInteractiveErrors(t *testing.T) {
	// TERM=dumb in setupTestEnv forces non-interactive; without a key source,
	// login must error rather than prompt.
	setupTestEnv(t)
	_, _, err := runRoot(t, nil, "login", "--raw")
	if err == nil {
		t.Fatal("expected error when no key source provided in non-interactive mode")
	}
	if !strings.Contains(err.Error(), "key") {
		t.Errorf("error message = %q, want mention of key", err.Error())
	}
}

func TestLogin_InvalidProfileName(t *testing.T) {
	setupTestEnv(t)
	_, _, err := runRoot(t, nil, "login", "--key", "k", "--profile", "bad name!", "--raw")
	if err == nil {
		t.Fatal("expected validation error for invalid profile name")
	}
	if !strings.Contains(err.Error(), "profile name") {
		t.Errorf("error message = %q, want validation mention", err.Error())
	}
}

func TestLogin_OverwritesExistingProfile(t *testing.T) {
	setupTestEnv(t)
	if _, _, err := runRoot(t, nil, "login", "--key", "v1", "--raw"); err != nil {
		t.Fatalf("first login: %v", err)
	}
	if _, _, err := runRoot(t, nil, "login", "--key", "v2", "--raw"); err != nil {
		t.Fatalf("second login: %v", err)
	}
	creds, _ := auth.ReadCredentials()
	if got := creds.Profiles["default"].APIKey; got != "v2" {
		t.Errorf("stored key after re-login = %q, want v2", got)
	}
}
