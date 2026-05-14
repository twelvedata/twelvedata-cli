package commands

import (
	"context"
	"strings"

	"github.com/spf13/cobra"
	"github.com/twelvedata/twelvedata-cli/internal/output"
)

// errorCodesHelp is appended to every command's usage block. The code list is
// pulled from output.ErrorCodes, which classify() in internal/output/errors.go
// returns from — single source of truth, no duplication.
var errorCodesHelp = "\n\n" +
	"Error codes (JSON envelope on stderr in raw mode — --raw, piped stdout, or CI):\n" +
	`  {"error":{"code":"<code>","message":"<message>","status":<http-status>}}` + "\n" +
	"  " + joinChunks(output.ErrorCodes, 4, " | ", "\n  ")

var rootCmd = &cobra.Command{
	Use:           "td",
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
		"{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}"+errorCodesHelp,
		1)
	rootCmd.SetUsageTemplate(t)
}

// Execute runs the root command and returns a process exit code suitable for
// os.Exit. Exit codes are agent-stable; see internal/output/errors.go.
func Execute() int {
	err := rootCmd.ExecuteContext(context.Background())
	return output.WriteError(rootCmd, err)
}
