package auth

import (
	"errors"
	"testing"
)

func TestResolveProfileName_Precedence(t *testing.T) {
	setupTestDir(t)

	// Empty everywhere -> "default".
	if got := ResolveProfileName(""); got != "default" {
		t.Errorf("default fallback: got %q, want default", got)
	}

	// Env var beats default.
	t.Setenv("TWELVEDATA_PROFILE", "from-env")
	if got := ResolveProfileName(""); got != "from-env" {
		t.Errorf("env should win over default: got %q", got)
	}

	// Flag beats env.
	if got := ResolveProfileName("from-flag"); got != "from-flag" {
		t.Errorf("flag should win over env: got %q", got)
	}

	// active_profile in creds beats default but not env or flag.
	t.Setenv("TWELVEDATA_PROFILE", "")
	if _, err := WriteCredentials(&CredentialsFile{
		ActiveProfile: "from-file",
		Profiles:      map[string]Profile{"from-file": {APIKey: "k"}},
	}); err != nil {
		t.Fatal(err)
	}
	if got := ResolveProfileName(""); got != "from-file" {
		t.Errorf("creds active_profile should be used: got %q", got)
	}
}

func TestResolveAPIKey_FlagWins(t *testing.T) {
	setupTestDir(t)
	t.Setenv("TWELVEDATA_API_KEY", "from-env")
	r, err := ResolveAPIKey("from-flag", "")
	if err != nil {
		t.Fatal(err)
	}
	if r.Key != "from-flag" || r.Source != SourceFlag {
		t.Errorf("got key=%q source=%q, want from-flag/flag", r.Key, r.Source)
	}
}

func TestResolveAPIKey_EnvWinsOverFile(t *testing.T) {
	setupTestDir(t)
	if _, err := WriteCredentials(&CredentialsFile{
		ActiveProfile: "default",
		Profiles:      map[string]Profile{"default": {APIKey: "from-file"}},
	}); err != nil {
		t.Fatal(err)
	}
	t.Setenv("TWELVEDATA_API_KEY", "from-env")
	r, err := ResolveAPIKey("", "")
	if err != nil {
		t.Fatal(err)
	}
	if r.Key != "from-env" || r.Source != SourceEnv {
		t.Errorf("got key=%q source=%q, want from-env/env", r.Key, r.Source)
	}
}

func TestResolveAPIKey_FromConfigFile(t *testing.T) {
	setupTestDir(t)
	if _, err := WriteCredentials(&CredentialsFile{
		ActiveProfile: "alpha",
		Profiles:      map[string]Profile{"alpha": {APIKey: "k-alpha"}},
	}); err != nil {
		t.Fatal(err)
	}
	r, err := ResolveAPIKey("", "")
	if err != nil {
		t.Fatal(err)
	}
	if r.Key != "k-alpha" || r.Source != SourceConfig || r.Profile != "alpha" {
		t.Errorf("got %+v, want key=k-alpha source=config profile=alpha", r)
	}
}

func TestResolveAPIKey_ProfileFlagOverridesActive(t *testing.T) {
	setupTestDir(t)
	if _, err := WriteCredentials(&CredentialsFile{
		ActiveProfile: "alpha",
		Profiles: map[string]Profile{
			"alpha": {APIKey: "k-alpha"},
			"beta":  {APIKey: "k-beta"},
		},
	}); err != nil {
		t.Fatal(err)
	}
	r, err := ResolveAPIKey("", "beta")
	if err != nil {
		t.Fatal(err)
	}
	if r.Key != "k-beta" || r.Profile != "beta" {
		t.Errorf("expected beta, got %+v", r)
	}
}

func TestResolveAPIKey_NothingConfigured(t *testing.T) {
	setupTestDir(t)
	_, err := ResolveAPIKey("", "")
	if !errors.Is(err, ErrNoAPIKey) {
		t.Errorf("expected ErrNoAPIKey, got %v", err)
	}
}

func TestResolveAPIKey_UnknownProfile(t *testing.T) {
	setupTestDir(t)
	if _, err := WriteCredentials(&CredentialsFile{
		ActiveProfile: "alpha",
		Profiles:      map[string]Profile{"alpha": {APIKey: "k"}},
	}); err != nil {
		t.Fatal(err)
	}
	_, err := ResolveAPIKey("", "nonexistent")
	if !errors.Is(err, ErrNoAPIKey) {
		t.Errorf("expected ErrNoAPIKey for unknown profile, got %v", err)
	}
}

func TestResolveAPIKey_EmptyAPIKeyTreatedAsMissing(t *testing.T) {
	setupTestDir(t)
	if _, err := WriteCredentials(&CredentialsFile{
		ActiveProfile: "default",
		Profiles:      map[string]Profile{"default": {APIKey: ""}},
	}); err != nil {
		t.Fatal(err)
	}
	_, err := ResolveAPIKey("", "")
	if !errors.Is(err, ErrNoAPIKey) {
		t.Errorf("empty stored key should resolve to ErrNoAPIKey, got %v", err)
	}
}

func TestSetActiveProfile(t *testing.T) {
	setupTestDir(t)
	if _, err := WriteCredentials(&CredentialsFile{
		ActiveProfile: "alpha",
		Profiles: map[string]Profile{
			"alpha": {APIKey: "k1"},
			"beta":  {APIKey: "k2"},
		},
	}); err != nil {
		t.Fatal(err)
	}
	if err := SetActiveProfile("beta"); err != nil {
		t.Fatal(err)
	}
	creds, _ := ReadCredentials()
	if creds.ActiveProfile != "beta" {
		t.Errorf("ActiveProfile = %q, want beta", creds.ActiveProfile)
	}
}

func TestSetActiveProfile_InvalidName(t *testing.T) {
	setupTestDir(t)
	if err := SetActiveProfile(""); err == nil {
		t.Error("expected error for empty name")
	}
	if err := SetActiveProfile("bad name"); err == nil {
		t.Error("expected error for invalid name")
	}
}

func TestSetActiveProfile_NoCredentialsFile(t *testing.T) {
	setupTestDir(t)
	err := SetActiveProfile("default")
	if err == nil {
		t.Fatal("expected error when no credentials file exists")
	}
}

func TestSetActiveProfile_UnknownProfile(t *testing.T) {
	setupTestDir(t)
	if _, err := WriteCredentials(&CredentialsFile{
		ActiveProfile: "alpha",
		Profiles:      map[string]Profile{"alpha": {APIKey: "k"}},
	}); err != nil {
		t.Fatal(err)
	}
	err := SetActiveProfile("ghost")
	var pne *ProfileNotFoundError
	if !errors.As(err, &pne) {
		t.Fatalf("expected ProfileNotFoundError, got %v", err)
	}
	if pne.Name != "ghost" {
		t.Errorf("ProfileNotFoundError.Name = %q, want ghost", pne.Name)
	}
	if len(pne.Available) != 1 || pne.Available[0] != "alpha" {
		t.Errorf("ProfileNotFoundError.Available = %v, want [alpha]", pne.Available)
	}
}

func TestProfileNotFoundError_Message(t *testing.T) {
	e := &ProfileNotFoundError{Name: "missing", Available: []string{"a", "b"}}
	got := e.Error()
	for _, sub := range []string{"missing", "a", "b"} {
		if !contains2(got, sub) {
			t.Errorf("error %q should mention %q", got, sub)
		}
	}

	e = &ProfileNotFoundError{Name: "x"}
	if got := e.Error(); !contains2(got, "No profiles configured") {
		t.Errorf("empty-available message should mention \"No profiles configured\", got %q", got)
	}
}

// Tiny local helper so this file doesn't pull in strings just for one substring check.
func contains2(haystack, needle string) bool {
	if needle == "" {
		return true
	}
	for i := 0; i+len(needle) <= len(haystack); i++ {
		if haystack[i:i+len(needle)] == needle {
			return true
		}
	}
	return false
}
