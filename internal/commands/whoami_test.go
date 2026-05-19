package commands

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"

	"github.com/twelvedata/twelvedata-cli/internal/auth"
)

func TestWhoami_FromEnv(t *testing.T) {
	setupTestEnv(t)
	t.Setenv("TWELVEDATA_API_KEY", "env_key_12345")

	stdout, _, err := runRoot(t, nil, "whoami", "--raw")
	if err != nil {
		t.Fatalf("whoami: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(stdout.Bytes(), &payload); err != nil {
		t.Fatalf("stdout not JSON: %v\n%s", err, stdout.String())
	}
	if payload["authenticated"] != true {
		t.Errorf("authenticated = %v, want true", payload["authenticated"])
	}
	if payload["source"] != "env" {
		t.Errorf("source = %v, want env", payload["source"])
	}
	if k, _ := payload["api_key"].(string); k == "env_key_12345" {
		t.Error("raw key leaked into output; should be masked")
	} else if !strings.Contains(k, "...") {
		t.Errorf("api_key %q is not masked", k)
	}
}

func TestWhoami_FromFlag_TakesPriorityOverEnv(t *testing.T) {
	setupTestEnv(t)
	t.Setenv("TWELVEDATA_API_KEY", "env_one")

	stdout, _, err := runRoot(t, nil, "whoami", "--api-key", "flag_two", "--raw")
	if err != nil {
		t.Fatal(err)
	}
	var payload map[string]any
	_ = json.Unmarshal(stdout.Bytes(), &payload)
	if payload["source"] != "flag" {
		t.Errorf("source = %v, want flag (flag must beat env)", payload["source"])
	}
}

func TestWhoami_FromFile(t *testing.T) {
	dir := setupTestEnv(t)
	// Pre-seed credentials so whoami reads from config.
	if _, _, err := runRoot(t, nil, "login", "--key", "stored_key", "--profile", "default", "--raw"); err != nil {
		t.Fatalf("setup login: %v", err)
	}

	stdout, _, err := runRoot(t, nil, "whoami", "--raw")
	if err != nil {
		t.Fatal(err)
	}
	var payload map[string]any
	if err := json.Unmarshal(stdout.Bytes(), &payload); err != nil {
		t.Fatalf("stdout not JSON: %v", err)
	}
	if payload["source"] != "config" {
		t.Errorf("source = %v, want config", payload["source"])
	}
	if payload["profile"] != "default" {
		t.Errorf("profile = %v, want default", payload["profile"])
	}
	wantPath := filepath.Join(dir, "credentials.json")
	if got, _ := payload["config_path"].(string); got != wantPath {
		t.Errorf("config_path = %q, want %q", got, wantPath)
	}
}

func TestWhoami_NoKeyConfigured(t *testing.T) {
	setupTestEnv(t)
	_, _, err := runRoot(t, nil, "whoami", "--raw")
	if err == nil {
		t.Fatal("expected error when no key is resolvable")
	}
	// auth.ErrNoAPIKey should propagate.
	if !strings.Contains(err.Error(), "no Twelve Data API key") {
		t.Errorf("error = %q, want ErrNoAPIKey message", err.Error())
	}
}

func TestSourceLabel_AllVariants(t *testing.T) {
	tests := map[auth.Source]string{
		auth.SourceFlag:   "flag",
		auth.SourceEnv:    "environment variable",
		auth.SourceConfig: "config file",
		auth.SourceSecure: "secure storage",
		auth.Source("unknown"): "unknown",
	}
	for src, want := range tests {
		got := sourceLabel(src)
		if got != want {
			t.Errorf("sourceLabel(%q) = %q, want %q", string(src), got, want)
		}
	}
}
