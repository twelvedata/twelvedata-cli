package commands

import (
	"strings"
	"testing"
)

func TestCompletion_BashEmitsNonEmptyScript(t *testing.T) {
	setupTestEnv(t)
	stdout, _, err := runRoot(t, nil, "completion", "bash")
	if err != nil {
		t.Fatalf("completion bash: %v", err)
	}
	if stdout.Len() == 0 {
		t.Fatal("bash completion script is empty")
	}
	// Cobra's bash completion always starts with `# bash completion V2 for ...`.
	if !strings.Contains(stdout.String(), "bash completion") {
		t.Errorf("bash completion output doesn't look like a bash script:\n%s", stdout.String())
	}
}

func TestCompletion_ZshFishPowershellEmitNonEmpty(t *testing.T) {
	for _, shell := range []string{"zsh", "fish", "powershell"} {
		t.Run(shell, func(t *testing.T) {
			setupTestEnv(t)
			stdout, _, err := runRoot(t, nil, "completion", shell)
			if err != nil {
				t.Fatalf("completion %s: %v", shell, err)
			}
			if stdout.Len() == 0 {
				t.Fatalf("%s completion script is empty", shell)
			}
		})
	}
}

func TestResolveCompletionShell_ExplicitArgWins(t *testing.T) {
	for _, s := range []string{"bash", "zsh", "fish", "powershell"} {
		got, err := resolveCompletionShell([]string{s})
		if err != nil {
			t.Fatalf("resolveCompletionShell(%q): %v", s, err)
		}
		if got != s {
			t.Errorf("resolveCompletionShell(%q) = %q", s, got)
		}
	}
}

func TestResolveCompletionShell_ArgCaseInsensitive(t *testing.T) {
	got, err := resolveCompletionShell([]string{"BASH"})
	if err != nil {
		t.Fatalf("resolveCompletionShell(BASH): %v", err)
	}
	if got != "bash" {
		t.Errorf("got %q, want bash", got)
	}
}

func TestResolveCompletionShell_UnknownArgErrors(t *testing.T) {
	_, err := resolveCompletionShell([]string{"tcsh"})
	if err == nil {
		t.Fatal("expected error for unsupported shell")
	}
	if !strings.Contains(err.Error(), "tcsh") {
		t.Errorf("error %q should name the bad shell", err.Error())
	}
}

func TestDetectShell_FromShellEnv(t *testing.T) {
	cases := map[string]string{
		"/bin/bash":      "bash",
		"/usr/bin/zsh":   "zsh",
		"/opt/fish/fish": "fish",
		"/bin/dash":      "", // unsupported
	}
	for in, want := range cases {
		t.Setenv("SHELL", in)
		t.Setenv("PSModulePath", "")
		got := detectShell()
		if got != want {
			t.Errorf("detectShell(SHELL=%q) = %q, want %q", in, got, want)
		}
	}
}

func TestDetectShell_FallsBackToPowerShellEnv(t *testing.T) {
	t.Setenv("SHELL", "")
	t.Setenv("PSModulePath", `C:\Program Files\PowerShell\Modules`)
	if got := detectShell(); got != "powershell" {
		t.Errorf("detectShell with PSModulePath set = %q, want powershell", got)
	}
}

func TestDetectShell_NothingSet(t *testing.T) {
	t.Setenv("SHELL", "")
	t.Setenv("PSModulePath", "")
	if got := detectShell(); got != "" {
		t.Errorf("detectShell with no env clues = %q, want empty", got)
	}
}

func TestCompletionShells_CoversAllCobraTargets(t *testing.T) {
	// The "install" subcommand dispatches on this list; if a new shell is
	// added we want the test to flag the missing branches in installCompletion.
	want := map[string]bool{"bash": true, "zsh": true, "fish": true, "powershell": true}
	for _, s := range completionShells {
		delete(want, s)
	}
	if len(want) != 0 {
		t.Errorf("completionShells missing: %v", want)
	}
}
