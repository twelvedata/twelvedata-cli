package commands

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"

	"github.com/twelvedata/twelvedata-cli/internal/output"
)

const (
	docsURL      = "https://twelvedata.com/docs/"
	dashboardURL = "https://twelvedata.com/account"
)

var docsCmd = &cobra.Command{
	Use:     "docs",
	Short:   "Open the Twelve Data documentation in your browser",
	Long:    "Open the Twelve Data API documentation in your default browser.\n\nIn machine mode (--raw, piped stdout, CI, TERM=dumb) the URL is printed to stdout instead — no browser launch.",
	Example: "  twelvedata docs",
	RunE: func(cmd *cobra.Command, _ []string) error {
		return openOrPrintURL(cmd, docsURL)
	},
}

var dashboardCmd = &cobra.Command{
	Use:     "dashboard",
	Short:   "Open your Twelve Data account dashboard in your browser",
	Long:    "Open your Twelve Data account dashboard in your default browser — manage your API key, plan, and usage.\n\nIn machine mode (--raw, piped stdout, CI, TERM=dumb) the URL is printed to stdout instead — no browser launch.",
	Example: "  twelvedata dashboard",
	RunE: func(cmd *cobra.Command, _ []string) error {
		return openOrPrintURL(cmd, dashboardURL)
	},
}

// openOrPrintURL launches the URL in the user's default browser. In machine
// mode the URL is written to stdout and no browser is started — an agent
// or CI run consumes the URL programmatically. Interactive callers get a
// one-line outcome on stderr so it does not pollute the response stream.
func openOrPrintURL(cmd *cobra.Command, url string) error {
	if output.IsRaw(cmd) {
		_, err := fmt.Fprintln(cmd.OutOrStdout(), url)
		return err
	}
	errOut := cmd.ErrOrStderr()
	if err := openInBrowser(url); err != nil {
		fmt.Fprintf(errOut, "Could not open browser (%s). Visit: %s\n", err, url)
		return nil
	}
	fmt.Fprintf(errOut, "Opened %s\n", url)
	return nil
}

func openInBrowser(url string) error {
	var c *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		c = exec.Command("open", url)
	case "windows":
		// First "" is the window title — without it, cmd treats the URL as
		// the title and never opens it.
		c = exec.Command("cmd", "/c", "start", "", url)
	default:
		c = exec.Command("xdg-open", url)
	}
	if err := c.Start(); err != nil {
		return err
	}
	// Detach: we don't care when the browser exits, and we don't want a
	// zombie subprocess hanging around if twelvedata is invoked in a long-lived
	// shell session.
	return c.Process.Release()
}

func init() {
	rootCmd.AddCommand(docsCmd, dashboardCmd)
}
