package commands

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/twelvedata/twelvedata-cli/internal/auth"
	"github.com/twelvedata/twelvedata-cli/internal/output"
)

var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Show current authentication status",
	Long: `Print the profile and API key (masked) that the next command will use.
Local only — no network calls. The source field is one of: flag, env, config,
or secure_storage.`,
	Example: `  td whoami
  td whoami --raw
  td whoami --profile staging`,
	RunE: func(cmd *cobra.Command, args []string) error {
		keyFlag, _ := cmd.Flags().GetString("api-key")
		profileFlag, _ := cmd.Flags().GetString("profile")
		resolved, err := auth.ResolveAPIKey(keyFlag, profileFlag)
		if err != nil {
			return err
		}
		profile := resolved.Profile
		if profile == "" {
			profile = auth.ResolveProfileName(profileFlag)
		}
		configPath := auth.CredentialsPath()

		if output.IsRaw(cmd) {
			enc := json.NewEncoder(cmd.OutOrStdout())
			enc.SetIndent("", "  ")
			return enc.Encode(map[string]any{
				"authenticated": true,
				"profile":       profile,
				"api_key":       auth.MaskKey(resolved.Key),
				"source":        string(resolved.Source),
				"config_path":   configPath,
			})
		}

		fmt.Fprintln(cmd.OutOrStdout())
		fmt.Fprintf(cmd.OutOrStdout(), "  Profile: %s\n", profile)
		fmt.Fprintf(cmd.OutOrStdout(), "  API Key: %s\n", auth.MaskKey(resolved.Key))
		fmt.Fprintf(cmd.OutOrStdout(), "  Source:  %s\n", sourceLabel(resolved.Source))
		fmt.Fprintf(cmd.OutOrStdout(), "  Config:  %s\n", configPath)
		fmt.Fprintln(cmd.OutOrStdout())
		return nil
	},
}

func sourceLabel(s auth.Source) string {
	switch s {
	case auth.SourceFlag:
		return "flag"
	case auth.SourceEnv:
		return "environment variable"
	case auth.SourceConfig:
		return "config file"
	case auth.SourceSecure:
		return "secure storage"
	default:
		return string(s)
	}
}

func init() {
	rootCmd.AddCommand(whoamiCmd)
}
