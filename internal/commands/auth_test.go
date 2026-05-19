package commands

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/twelvedata/twelvedata-cli/internal/auth"
)

// seedProfiles populates credentials.json with the given (name, key) pairs by
// calling the storage backend directly. We avoid going through `twelvedata
// login` because login activates the most-recently-stored profile (it calls
// SetActiveProfile), which makes "first one wins" assertions impossible.
// Using auth.StoreAPIKey preserves the first-stored-is-active behavior of
// FileBackend, which is what auth list's "active" marker actually surfaces.
func seedProfiles(t *testing.T, pairs ...string) {
	t.Helper()
	if len(pairs)%2 != 0 {
		t.Fatal("seedProfiles needs name/key pairs")
	}
	for i := 0; i < len(pairs); i += 2 {
		if _, _, err := auth.StoreAPIKey(pairs[i], pairs[i+1]); err != nil {
			t.Fatalf("seed %q: %v", pairs[i], err)
		}
	}
}

func TestAuthList_Empty(t *testing.T) {
	setupTestEnv(t)
	stdout, _, err := runRoot(t, nil, "auth", "list", "--raw")
	if err != nil {
		t.Fatal(err)
	}
	var payload map[string]any
	if err := json.Unmarshal(stdout.Bytes(), &payload); err != nil {
		t.Fatalf("not JSON: %v\n%s", err, stdout.String())
	}
	items, ok := payload["profiles"].([]any)
	if !ok {
		t.Fatalf("profiles is not an array: %T (%v)", payload["profiles"], payload["profiles"])
	}
	if len(items) != 0 {
		t.Errorf("profiles length = %d, want 0", len(items))
	}
}

func TestAuthList_TwoProfiles_MarksActive(t *testing.T) {
	setupTestEnv(t)
	seedProfiles(t, "alpha", "k1", "beta", "k2")

	stdout, _, err := runRoot(t, nil, "auth", "list", "--raw")
	if err != nil {
		t.Fatal(err)
	}
	var payload struct {
		Profiles []struct {
			Name   string `json:"name"`
			Active bool   `json:"active"`
		} `json:"profiles"`
	}
	if err := json.Unmarshal(stdout.Bytes(), &payload); err != nil {
		t.Fatalf("not JSON: %v\n%s", err, stdout.String())
	}
	if len(payload.Profiles) != 2 {
		t.Fatalf("profiles len = %d, want 2", len(payload.Profiles))
	}
	activeCount := 0
	for _, p := range payload.Profiles {
		if p.Active {
			activeCount++
			if p.Name != "alpha" {
				t.Errorf("active profile = %q, want alpha (first stored)", p.Name)
			}
		}
	}
	if activeCount != 1 {
		t.Errorf("expected exactly one active profile, got %d", activeCount)
	}
}

func TestAuthList_DefaultsToListWhenNoSubcommand(t *testing.T) {
	setupTestEnv(t)
	seedProfiles(t, "only", "k")

	stdout, _, err := runRoot(t, nil, "auth", "--raw")
	if err != nil {
		t.Fatal(err)
	}
	// `auth` with no subcommand should produce the same JSON shape as `auth list`.
	var payload map[string]any
	if err := json.Unmarshal(stdout.Bytes(), &payload); err != nil {
		t.Fatalf("auth (no subcommand) didn't emit list JSON: %v\n%s", err, stdout.String())
	}
	if _, ok := payload["profiles"]; !ok {
		t.Error(`expected "profiles" key in JSON, got: ` + stdout.String())
	}
}

func TestAuthSwitch_ChangesActive(t *testing.T) {
	setupTestEnv(t)
	seedProfiles(t, "alpha", "k1", "beta", "k2")

	if _, _, err := runRoot(t, nil, "auth", "switch", "beta", "--raw"); err != nil {
		t.Fatal(err)
	}
	creds, _ := auth.ReadCredentials()
	if creds.ActiveProfile != "beta" {
		t.Errorf("active = %q, want beta", creds.ActiveProfile)
	}
}

func TestAuthSwitch_UnknownProfileErrors(t *testing.T) {
	setupTestEnv(t)
	seedProfiles(t, "alpha", "k1")

	_, _, err := runRoot(t, nil, "auth", "switch", "ghost", "--raw")
	if err == nil {
		t.Fatal("expected ProfileNotFoundError")
	}
	if !strings.Contains(err.Error(), "ghost") {
		t.Errorf("error %q should name the missing profile", err.Error())
	}
}

func TestAuthSwitch_MissingArgInNonInteractive(t *testing.T) {
	setupTestEnv(t)
	seedProfiles(t, "alpha", "k1")

	_, _, err := runRoot(t, nil, "auth", "switch", "--raw")
	if err == nil {
		t.Fatal("expected error when arg missing in non-interactive mode")
	}
}

func TestAuthRename_Succeeds(t *testing.T) {
	setupTestEnv(t)
	seedProfiles(t, "old", "k")

	if _, _, err := runRoot(t, nil, "auth", "rename", "old", "fresh", "--raw"); err != nil {
		t.Fatal(err)
	}
	creds, _ := auth.ReadCredentials()
	if _, ok := creds.Profiles["old"]; ok {
		t.Error("old profile should be gone")
	}
	if got := creds.Profiles["fresh"].APIKey; got != "k" {
		t.Errorf("renamed profile key = %q, want preserved \"k\"", got)
	}
	if creds.ActiveProfile != "fresh" {
		t.Errorf("active should follow rename: got %q, want fresh", creds.ActiveProfile)
	}
}

func TestAuthRename_InvalidNewName(t *testing.T) {
	setupTestEnv(t)
	seedProfiles(t, "old", "k")

	_, _, err := runRoot(t, nil, "auth", "rename", "old", "bad name!", "--raw")
	if err == nil {
		t.Fatal("expected validation error")
	}
}

func TestAuthRemove_RemovesProfile(t *testing.T) {
	setupTestEnv(t)
	seedProfiles(t, "alpha", "k1", "beta", "k2")

	// --raw disables the destructive confirmation prompt (shouldPrompt returns
	// false in raw mode).
	if _, _, err := runRoot(t, nil, "auth", "remove", "beta", "--raw"); err != nil {
		t.Fatal(err)
	}
	creds, _ := auth.ReadCredentials()
	if _, ok := creds.Profiles["beta"]; ok {
		t.Error("beta should be removed")
	}
	if _, ok := creds.Profiles["alpha"]; !ok {
		t.Error("alpha should remain")
	}
}

func TestAuthRemove_LastProfileDeletesFile(t *testing.T) {
	dir := setupTestEnv(t)
	seedProfiles(t, "only", "k")

	if _, _, err := runRoot(t, nil, "auth", "remove", "only", "--raw"); err != nil {
		t.Fatal(err)
	}
	if creds, _ := auth.ReadCredentials(); creds != nil {
		t.Errorf("credentials.json should be deleted; got %+v", creds)
	}
	_ = dir // path checked indirectly via ReadCredentials
}

func TestSanitizePlaceholder(t *testing.T) {
	cases := map[string]string{
		"good_name-1.2":    "good_name-1.2",
		"bad name!":        "bad-name-",
		"with/slash":       "with-slash",
		"":                 "",
		"weird ✨ unicode": "weird---unicode",
	}
	for in, want := range cases {
		if got := sanitizePlaceholder(in); got != want {
			t.Errorf("sanitizePlaceholder(%q) = %q, want %q", in, got, want)
		}
	}
}
