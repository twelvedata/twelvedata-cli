//go:build integration

package integration

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
)

// Auth-lifecycle tests exercise login / whoami / auth switch / logout against a
// throw-away credentials.json. They do not burn API quota — only the local
// credential store is touched. Fake keys are used since whoami performs no
// server-side validation.

// runCLIStdin is like runCLI but pipes stdin into the subprocess. Used to feed
// a key into `twelvedata login --key-stdin`.
func runCLIStdin(t *testing.T, stdin string, args ...string) []byte {
	t.Helper()
	args = append(args, "--raw")
	cmd := exec.Command(binaryPath, args...)
	cmd.Env = append(os.Environ(), "TWELVEDATA_NO_UPDATE_NOTIFIER=1", "TERM=dumb")
	cmd.Stdin = bytes.NewBufferString(stdin)
	out, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			t.Fatalf("twelvedata %v: exit=%d\nstdout:\n%s\nstderr:\n%s",
				args, ee.ExitCode(), out, ee.Stderr)
		}
		t.Fatalf("twelvedata %v: %v\nstdout:\n%s", args, err, out)
	}
	return out
}

// isolateAuth points every credential lookup at a throw-away dir and forces
// file storage, so the OS keyring is not touched.
func isolateAuth(t *testing.T) {
	t.Helper()
	t.Setenv("TWELVEDATA_CONFIG_DIR", t.TempDir())
	t.Setenv("TWELVEDATA_CREDENTIAL_STORE", "file")
	t.Setenv("TWELVEDATA_API_KEY", "")
	t.Setenv("TWELVEDATA_PROFILE", "")
}

func TestCLIAuthLoginWhoami(t *testing.T) {
	isolateAuth(t)

	loginOut := runCLIStdin(t, "fake-key-1", "login", "--key-stdin")
	var login struct {
		Success bool   `json:"success"`
		Profile string `json:"profile"`
	}
	mustJSON(t, loginOut, &login)
	if !login.Success {
		t.Fatalf("login.success=%v body=%s", login.Success, loginOut)
	}
	if login.Profile != "default" {
		t.Fatalf("login.profile=%q want default", login.Profile)
	}

	whoamiOut := runCLI(t, "whoami")
	var who struct {
		Authenticated bool   `json:"authenticated"`
		Profile       string `json:"profile"`
		APIKey        string `json:"api_key"`
		Source        string `json:"source"`
	}
	mustJSON(t, whoamiOut, &who)
	if !who.Authenticated {
		t.Fatalf("whoami.authenticated=false body=%s", whoamiOut)
	}
	if who.Profile != "default" {
		t.Fatalf("whoami.profile=%q want default", who.Profile)
	}
	if who.APIKey == "" {
		t.Fatal("whoami.api_key empty")
	}
	if who.Source != "config" {
		t.Fatalf("whoami.source=%q want %q", who.Source, "config")
	}
}

func TestCLIAuthSwitch(t *testing.T) {
	isolateAuth(t)

	runCLIStdin(t, "fake-key-a", "login", "--key-stdin", "--profile", "alpha")
	runCLIStdin(t, "fake-key-b", "login", "--key-stdin", "--profile", "beta")

	// Most recent login wins as active profile.
	whoami1 := runCLI(t, "whoami")
	var who1 struct {
		Profile string `json:"profile"`
	}
	mustJSON(t, whoami1, &who1)
	if who1.Profile != "beta" {
		t.Fatalf("after login: profile=%q want beta", who1.Profile)
	}

	switchOut := runCLI(t, "auth", "switch", "alpha")
	var sw struct {
		Success       bool   `json:"success"`
		ActiveProfile string `json:"active_profile"`
	}
	mustJSON(t, switchOut, &sw)
	if !sw.Success || sw.ActiveProfile != "alpha" {
		t.Fatalf("auth switch: %+v body=%s", sw, switchOut)
	}

	whoami2 := runCLI(t, "whoami")
	var who2 struct {
		Profile string `json:"profile"`
	}
	mustJSON(t, whoami2, &who2)
	if who2.Profile != "alpha" {
		t.Fatalf("after switch: profile=%q want alpha", who2.Profile)
	}

	listOut := runCLI(t, "auth", "list")
	var list struct {
		Profiles []struct {
			Name   string `json:"name"`
			Active bool   `json:"active"`
		} `json:"profiles"`
	}
	mustJSON(t, listOut, &list)
	if len(list.Profiles) != 2 {
		t.Fatalf("expected 2 profiles, got %d: %s", len(list.Profiles), listOut)
	}
	var activeFound bool
	for _, p := range list.Profiles {
		if p.Name == "alpha" && p.Active {
			activeFound = true
		}
	}
	if !activeFound {
		t.Fatalf("alpha not marked active in list: %s", listOut)
	}
}

func TestCLIAuthLogout(t *testing.T) {
	isolateAuth(t)

	runCLIStdin(t, "fake-key-1", "login", "--key-stdin")
	logoutOut := runCLI(t, "logout")
	var lo struct {
		Success bool   `json:"success"`
		Scope   string `json:"scope"`
	}
	mustJSON(t, logoutOut, &lo)
	if !lo.Success {
		t.Fatalf("logout.success=false body=%s", logoutOut)
	}
	if lo.Scope != "all" {
		t.Fatalf("logout.scope=%q want all", lo.Scope)
	}

	// After logout, whoami must exit non-zero (ErrNoAPIKey → exit 3).
	cmd := exec.Command(binaryPath, "whoami", "--raw")
	cmd.Env = append(os.Environ(), "TWELVEDATA_NO_UPDATE_NOTIFIER=1", "TERM=dumb")
	out, err := cmd.Output()
	if err == nil {
		t.Fatalf("expected whoami to fail after logout; stdout=%s", out)
	}
	ee, ok := err.(*exec.ExitError)
	if !ok {
		t.Fatalf("whoami: %v", err)
	}
	if ee.ExitCode() != 3 {
		t.Fatalf("whoami exit code = %d, want 3; stderr=%s", ee.ExitCode(), ee.Stderr)
	}
}
