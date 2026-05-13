package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/twelvedata/twelvedata-cli/internal/auth"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage authentication profiles",
	Long: `Manage named profiles. Each profile holds one Twelve Data API key.

Environment variables:
  TWELVEDATA_API_KEY           API key (overrides stored credentials)
  TWELVEDATA_PROFILE           Profile name (overrides config default)
  TWELVEDATA_CREDENTIAL_STORE  Storage method: "secure_storage" or "file"`,
	Example: `  td auth list
  td auth switch staging
  td auth rename staging production
  td auth remove staging`,
	// Default behavior of `td auth` with no subcommand is `td auth list`.
	RunE: func(cmd *cobra.Command, args []string) error {
		return authListCmd.RunE(cmd, args)
	},
}

var authListCmd = &cobra.Command{
	Use:   "list",
	Short: "List profiles",
	RunE: func(cmd *cobra.Command, args []string) error {
		profiles, err := auth.ListProfiles()
		if err != nil {
			return err
		}
		if isJSON(cmd) {
			items := make([]map[string]any, 0, len(profiles))
			for _, p := range profiles {
				items = append(items, map[string]any{"name": p.Name, "active": p.Active})
			}
			enc := json.NewEncoder(cmd.OutOrStdout())
			enc.SetIndent("", "  ")
			return enc.Encode(map[string]any{"profiles": items})
		}
		if len(profiles) == 0 {
			fmt.Fprintln(cmd.OutOrStdout(), "No profiles configured. Run: td login")
			return nil
		}
		fmt.Fprintln(cmd.OutOrStdout())
		fmt.Fprintln(cmd.OutOrStdout(), "  Profiles")
		fmt.Fprintln(cmd.OutOrStdout())
		hasInvalid := false
		for _, p := range profiles {
			marker := "  "
			suffix := ""
			if p.Active {
				marker = "▸ "
				suffix = " (active)"
			}
			invalid := ""
			if auth.ValidateProfileName(p.Name) != nil {
				invalid = " (invalid name)"
				hasInvalid = true
			}
			fmt.Fprintf(cmd.OutOrStdout(), "  %s%s%s%s\n", marker, p.Name, suffix, invalid)
		}
		if hasInvalid {
			fmt.Fprintln(cmd.OutOrStdout())
			fmt.Fprintln(cmd.OutOrStdout(), "  Rename profiles with invalid names via `td auth rename`.")
		}
		fmt.Fprintln(cmd.OutOrStdout())
		return nil
	},
}

var authSwitchCmd = &cobra.Command{
	Use:     "switch [name]",
	Short:   "Switch the active profile",
	Args:    cobra.MaximumNArgs(1),
	Example: "  td auth switch staging",
	RunE: func(cmd *cobra.Command, args []string) error {
		name := ""
		if len(args) == 1 {
			name = strings.TrimSpace(args[0])
		}
		if name == "" {
			if !auth.IsInteractive() {
				return errors.New("profile name is required in non-interactive mode")
			}
			profiles, err := auth.ListProfiles()
			if err != nil {
				return err
			}
			if len(profiles) == 0 {
				return errors.New("no profiles configured. Run `td login` first")
			}
			chosen, err := auth.SelectProfile("Switch to which profile?", profiles)
			if err != nil {
				return err
			}
			name = chosen
		}
		if err := auth.SetActiveProfile(name); err != nil {
			return err
		}
		if isJSON(cmd) {
			enc := json.NewEncoder(cmd.OutOrStdout())
			enc.SetIndent("", "  ")
			return enc.Encode(map[string]any{"success": true, "active_profile": name})
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Switched to profile %q.\n", name)
		return nil
	},
}

var authRenameCmd = &cobra.Command{
	Use:     "rename <old> <new>",
	Short:   "Rename a profile",
	Args:    cobra.MaximumNArgs(2),
	Example: "  td auth rename staging production",
	RunE: func(cmd *cobra.Command, args []string) error {
		var oldName, newName string
		if len(args) >= 1 {
			oldName = strings.TrimSpace(args[0])
		}
		if len(args) == 2 {
			newName = strings.TrimSpace(args[1])
		}

		if oldName == "" {
			if !auth.IsInteractive() {
				return errors.New("old and new profile names are required in non-interactive mode")
			}
			profiles, err := auth.ListProfiles()
			if err != nil {
				return err
			}
			if len(profiles) == 0 {
				return errors.New("no profiles configured")
			}
			chosen, err := auth.SelectProfile("Rename which profile?", profiles)
			if err != nil {
				return err
			}
			oldName = chosen
		}
		if newName == "" {
			if !auth.IsInteractive() {
				return errors.New("new profile name is required in non-interactive mode")
			}
			placeholder := sanitizePlaceholder(oldName)
			v, err := auth.PromptText(fmt.Sprintf("New name for %q", oldName), placeholder, func(v string) error {
				return auth.ValidateProfileName(strings.TrimSpace(v))
			})
			if err != nil {
				return err
			}
			newName = strings.TrimSpace(v)
		}

		if err := auth.RenameProfile(oldName, newName); err != nil {
			return err
		}
		if isJSON(cmd) {
			enc := json.NewEncoder(cmd.OutOrStdout())
			enc.SetIndent("", "  ")
			return enc.Encode(map[string]any{"success": true, "from": oldName, "to": newName})
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Renamed %q to %q.\n", oldName, newName)
		return nil
	},
}

var authRemoveCmd = &cobra.Command{
	Use:     "remove [name]",
	Short:   "Remove a profile",
	Args:    cobra.MaximumNArgs(1),
	Example: "  td auth remove staging",
	RunE: func(cmd *cobra.Command, args []string) error {
		name := ""
		if len(args) == 1 {
			name = strings.TrimSpace(args[0])
		}
		if name == "" {
			if !auth.IsInteractive() {
				return errors.New("profile name is required in non-interactive mode")
			}
			profiles, err := auth.ListProfiles()
			if err != nil {
				return err
			}
			if len(profiles) == 0 {
				return errors.New("no profiles configured")
			}
			chosen, err := auth.SelectProfile("Remove which profile?", profiles)
			if err != nil {
				return err
			}
			name = chosen
		}
		if auth.IsInteractive() && !isJSON(cmd) {
			ok, err := auth.ConfirmDestructive(fmt.Sprintf("Remove profile %q?", name))
			if err != nil {
				return err
			}
			if !ok {
				return errors.New("remove cancelled")
			}
		}
		if err := auth.RemoveProfile(name); err != nil {
			return err
		}
		if isJSON(cmd) {
			enc := json.NewEncoder(cmd.OutOrStdout())
			enc.SetIndent("", "  ")
			return enc.Encode(map[string]any{"success": true, "profile": name})
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Removed profile %q.\n", name)
		return nil
	},
}

// sanitizePlaceholder replaces any character outside the valid profile-name set
// with '-', so users renaming a legacy invalid profile get a sensible default.
func sanitizePlaceholder(s string) string {
	var b strings.Builder
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z',
			r >= 'A' && r <= 'Z',
			r >= '0' && r <= '9',
			r == '.', r == '_', r == '-':
			b.WriteRune(r)
		default:
			b.WriteRune('-')
		}
	}
	return b.String()
}

func init() {
	authCmd.AddCommand(authListCmd)
	authCmd.AddCommand(authSwitchCmd)
	authCmd.AddCommand(authRenameCmd)
	authCmd.AddCommand(authRemoveCmd)
	rootCmd.AddCommand(authCmd)
}
