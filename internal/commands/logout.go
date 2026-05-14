package commands

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/twelvedata/twelvedata-cli/internal/auth"
	"github.com/twelvedata/twelvedata-cli/internal/output"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove stored API keys",
	Long: `Remove one or all stored API keys. With --profile, only that profile is
removed; without --profile every profile is wiped and credentials.json is
deleted (after a confirmation prompt on a TTY).`,
	Example: `  td logout                       # remove all profiles
  td logout --profile staging     # remove a single profile`,
	RunE: func(cmd *cobra.Command, args []string) error {
		profile, _ := cmd.Flags().GetString("profile")
		if profile != "" {
			if err := auth.RemoveProfile(profile); err != nil {
				return err
			}
			return emitLogout(cmd, profile, false)
		}

		if shouldPrompt(cmd) {
			ok, err := auth.ConfirmDestructive("Remove all stored API keys?")
			if err != nil {
				return err
			}
			if !ok {
				fmt.Fprintln(cmd.OutOrStdout(), "Cancelled.")
				return nil
			}
		}
		if err := auth.RemoveAll(); err != nil {
			return err
		}
		return emitLogout(cmd, "", true)
	},
}

func emitLogout(cmd *cobra.Command, profile string, all bool) error {
	if output.IsRaw(cmd) {
		payload := map[string]any{"success": true}
		if all {
			payload["scope"] = "all"
		} else {
			payload["scope"] = "profile"
			payload["profile"] = profile
		}
		enc := json.NewEncoder(cmd.OutOrStdout())
		enc.SetIndent("", "  ")
		return enc.Encode(payload)
	}
	if all {
		fmt.Fprintln(cmd.OutOrStdout(), "All stored API keys removed.")
	} else {
		fmt.Fprintf(cmd.OutOrStdout(), "Profile %q removed.\n", profile)
	}
	return nil
}

func init() {
	// --profile is inherited from rootCmd's persistent flag; no local flag needed.
	rootCmd.AddCommand(logoutCmd)
}
