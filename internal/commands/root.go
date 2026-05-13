package commands

import (
	"context"
	"strings"

	"github.com/spf13/cobra"
	"github.com/twelvedata/twelvedata-cli/internal/output"
)

var rootCmd = &cobra.Command{
	Use:           "td",
	Short:         "Twelve Data CLI",
	Long:          "Twelve Data CLI — REST client for Twelve Data's market data API. Designed for AI agents and humans alike.\n\nResponses render as pretty-printed JSON by default; --output csv switches to the streaming CSV path for endpoints that support it.",
	SilenceUsage:  true,
	SilenceErrors: true,
	PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
		_, err := output.ResolveFormat(cmd)
		return err
	},
}

func init() {
	rootCmd.PersistentFlags().String("api-key", "", "Twelve Data API key (overrides TWELVEDATA_API_KEY)")
	rootCmd.PersistentFlags().StringP("output", "o", "", "Output format: json, csv (default: json)")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Force JSON error envelopes on stderr even on a TTY")
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
	rootCmd.SetUsageTemplate(t)
}

// Execute runs the root command and returns a process exit code suitable for
// os.Exit. Exit codes are agent-stable; see internal/output/errors.go.
func Execute() int {
	err := rootCmd.ExecuteContext(context.Background())
	return output.WriteError(rootCmd, err)
}
