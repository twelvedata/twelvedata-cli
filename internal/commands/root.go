package commands

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/twelvedata/twelvedata-cli/internal/output"
)

var rootCmd = &cobra.Command{
	Use:           "td",
	Short:         "Twelve Data CLI",
	Long:          "Twelve Data CLI — REST client for Twelve Data's market data API. Designed for AI agents and humans alike.\n\nResponses render as pretty-printed JSON by default; --output csv switches to the streaming CSV path for endpoints that support it.",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	rootCmd.PersistentFlags().String("api-key", "", "Twelve Data API key (overrides TWELVEDATA_API_KEY)")
	rootCmd.PersistentFlags().StringP("output", "o", "", "Output format: json, csv (default: json)")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Force JSON error envelopes on stderr even on a TTY")
}

// Execute runs the root command and returns a process exit code suitable for
// os.Exit. Exit codes are agent-stable; see internal/output/errors.go.
func Execute() int {
	err := rootCmd.ExecuteContext(context.Background())
	return output.WriteError(rootCmd, err)
}
