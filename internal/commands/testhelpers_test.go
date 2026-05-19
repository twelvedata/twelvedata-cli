package commands

import (
	"bytes"
	"io"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/twelvedata/twelvedata-cli/internal/auth"
)

// setupTestEnv isolates a test from the host's real config, keyring, and shell.
// It returns the temp config dir so callers can assert paths against it.
//
// Why we set each one:
//   TWELVEDATA_CONFIG_DIR           — redirect credentials.json into a temp dir.
//   TWELVEDATA_CREDENTIAL_STORE=file — never touch the OS keyring.
//   TWELVEDATA_API_KEY=""           — kill resolver fast-path via env.
//   TWELVEDATA_PROFILE=""           — same for profile env.
//   XDG_CONFIG_HOME=""              — defeat XDG fallback inside GetConfigDir.
//   TERM=dumb                       — force output.IsRaw() to return true,
//                                     suppressing prompts/spinner/color and
//                                     short-circuiting the update notifier.
//   TWELVEDATA_NO_UPDATE_NOTIFIER=1 — belt-and-braces; never hit GitHub.
//   NO_COLOR=1                      — keep error/banner output ANSI-clean so
//                                     stdout/stderr buffers are easy to assert.
func setupTestEnv(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	t.Setenv("TWELVEDATA_CONFIG_DIR", dir)
	t.Setenv("TWELVEDATA_CREDENTIAL_STORE", "file")
	t.Setenv("TWELVEDATA_API_KEY", "")
	t.Setenv("TWELVEDATA_PROFILE", "")
	t.Setenv("XDG_CONFIG_HOME", "")
	t.Setenv("TERM", "dumb")
	t.Setenv("TWELVEDATA_NO_UPDATE_NOTIFIER", "1")
	t.Setenv("NO_COLOR", "1")
	auth.ResetBackend()
	resetRootFlags()
	t.Cleanup(func() {
		auth.ResetBackend()
		resetRootFlags()
	})
	return dir
}

// resetRootFlags walks rootCmd and every subcommand, restoring each flag to
// its default value. Cobra mutates flag state in place during Execute(); without
// this, a flag set by an earlier test leaks into the next one.
func resetRootFlags() {
	var visit func(c *cobra.Command)
	visit = func(c *cobra.Command) {
		reset := func(f *pflag.Flag) {
			_ = f.Value.Set(f.DefValue)
			f.Changed = false
		}
		c.Flags().VisitAll(reset)
		c.PersistentFlags().VisitAll(reset)
		for _, sub := range c.Commands() {
			visit(sub)
		}
	}
	visit(rootCmd)
}

// runRoot drives rootCmd end-to-end with isolated I/O. It returns stdout,
// stderr, and the error returned by Execute() so a single call can assert all
// three (Cobra writes some errors via SilenceErrors=false, but rootCmd silences
// them — they only surface via the return value).
func runRoot(t *testing.T, stdin io.Reader, args ...string) (stdout, stderr *bytes.Buffer, err error) {
	t.Helper()
	stdout = new(bytes.Buffer)
	stderr = new(bytes.Buffer)
	rootCmd.SetArgs(args)
	if stdin != nil {
		rootCmd.SetIn(stdin)
	} else {
		rootCmd.SetIn(bytes.NewReader(nil))
	}
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)
	err = rootCmd.Execute()
	return stdout, stderr, err
}

// newTestCmd builds a standalone cobra.Command wired with the four persistent
// flags rootCmd defines. Use it when a test wants to invoke a command's RunE
// directly without driving Cobra's parse/dispatch over the global rootCmd
// (handy when the command-under-test has its own subtree, or to keep test
// state hermetic).
func newTestCmd() *cobra.Command {
	c := &cobra.Command{Use: "test"}
	c.PersistentFlags().String("api-key", "", "")
	c.PersistentFlags().StringP("output", "o", "", "")
	c.PersistentFlags().Bool("raw", false, "")
	c.PersistentFlags().StringP("profile", "p", "", "")
	return c
}
