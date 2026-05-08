package commands

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/twelvedata/twelvedata-cli/internal/output"
)

var rootCmd = &cobra.Command{
	Use:           "td",
	Short:         "Twelve Data CLI",
	Long:          "Twelve Data CLI — REST client for Twelve Data's market data API. Designed for AI agents and humans alike.\n\nGlobal flags --output and --quiet control output rendering. On a TTY the default is a human-friendly format; in pipes and CI the default is JSON.",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	rootCmd.PersistentFlags().String("api-key", "", "Twelve Data API key (overrides TWELVEDATA_API_KEY)")
	rootCmd.PersistentFlags().StringP("output", "o", "", "Output format: json, csv, table (default: table on TTY, json on pipe/CI)")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Suppress decorations; implies --output=json")
}

// Execute runs the root command and returns a process exit code suitable for
// os.Exit. Exit codes are agent-stable; see internal/output/errors.go.
func Execute() int {
	err := rootCmd.ExecuteContext(context.Background())
	return output.WriteError(rootCmd, err)
}
