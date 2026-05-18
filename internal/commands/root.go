package commands

import (
	"context"
	"strings"

	"github.com/spf13/cobra"
	"github.com/twelvedata/twelvedata-cli/internal/output"
	"github.com/twelvedata/twelvedata-cli/internal/update"
)

// outputHelp and errorCodesHelp are appended to every command's usage block
// via the shared UsageTemplate. Output is a single generic line — the schema
// differs per endpoint, so we link to the API reference rather than dumping
// shapes here. The code list comes from output.ErrorCodes (sourced by
// classify() in internal/output/errors.go) so there's no duplication.
const outputHelp = "\n\n" +
	"Output:\n" +
	"  Pretty-printed JSON by default, or streaming CSV with --output csv.\n" +
	"  See https://twelvedata.com/docs for the response schema."

var errorCodesHelp = "\n\n" +
	"Error codes (JSON envelope on stderr in raw mode — --raw, piped stdout, or CI):\n" +
	`  {"error":{"code":"<code>","message":"<message>","status":<http-status>}}` + "\n" +
	"  " + joinChunks(output.ErrorCodes, 4, " | ", "\n  ")

var rootCmd = &cobra.Command{
	Use:           "twelvedata",
	Short:         "Twelve Data CLI",
	Long:          "Twelve Data CLI — REST client for Twelve Data's market data API. Designed for AI agents and humans alike.\n\nResponses render as pretty-printed JSON by default; --output csv switches to the streaming CSV path for endpoints that support it. On an interactive terminal the CLI shows a spinner and colorized errors; pass --raw (or pipe stdout) to force machine-friendly output.",
	SilenceUsage:  true,
	SilenceErrors: true,
	PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
		if _, err := output.ResolveFormat(cmd); err != nil {
			return err
		}
		return promptMissingFlags(cmd)
	},
	// PersistentPostRun fires after a subcommand's RunE returns without error.
	// It runs on the command that defined it (rootCmd), so the notifier hits
	// every successful invocation — including `twelvedata` with no args, which falls
	// through to the Run handler below.
	PersistentPostRun: func(cmd *cobra.Command, _ []string) {
		update.MaybeNotify(cmd)
	},
	Run: func(cmd *cobra.Command, _ []string) {
		if !output.IsRaw(cmd) {
			output.PrintBanner(cmd.OutOrStderr())
		}
		_ = cmd.Help()
	},
}

func init() {
	rootCmd.PersistentFlags().String("api-key", "", "Twelve Data API key (overrides TWELVEDATA_API_KEY)")
	rootCmd.PersistentFlags().StringP("output", "o", "", "Output format: json, csv (default: json)")
	rootCmd.PersistentFlags().Bool("raw", false, "Force machine mode: JSON error envelope, no spinner, no color, no prompts")
	rootCmd.PersistentFlags().StringP("profile", "p", "", "Profile to use (overrides TWELVEDATA_PROFILE)")

	// Rename Cobra's "flag" vocabulary to "option" in usage output:
	//   "Flags:" / "Global Flags:"  headers   -> "Options:" / "Global Options:"
	//   "[flags]" appended by .UseLine()      -> "[options]"
	// The template is inherited by subcommands.
	cobra.AddTemplateFunc("useLineOpts", func(c *cobra.Command) string {
		return strings.Replace(c.UseLine(), "[flags]", "[options]", 1)
	})
	t := rootCmd.UsageTemplate()
	t = strings.ReplaceAll(t, "Flags:", "Options:")
	t = strings.ReplaceAll(t, "{{.UseLine}}", "{{useLineOpts .}}")
	// Append a shared "Error codes" section after Global Options. Inserted
	// outside the {{if .HasAvailableInheritedFlags}} block so it renders for
	// every command — the codes are emitted by the central classifier in
	// internal/output/errors.go and apply uniformly.
	t = strings.Replace(t,
		"{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}",
		"{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}"+outputHelp+errorCodesHelp,
		1)
	rootCmd.SetUsageTemplate(t)
}

// Execute runs the root command and returns a process exit code suitable for
// os.Exit. Exit codes are agent-stable; see internal/output/errors.go.
func Execute() int {
	err := rootCmd.ExecuteContext(context.Background())
	return output.WriteError(rootCmd, err)
}
