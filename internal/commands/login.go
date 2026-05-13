package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"

	"github.com/twelvedata/twelvedata-cli/internal/auth"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Save a Twelve Data API key for a profile",
	Long: `Store a Twelve Data API key in the active credential backend (OS keyring when
available, plaintext credentials.json otherwise) under the chosen profile.

On a TTY, the key is prompted for in masked form. In non-interactive shells
either --key-stdin (read the key from stdin) or --key <value> is required;
prefer --key-stdin so the secret never appears in shell history, process
listings, or CI logs.`,
	Example: `  td login                                          # prompts on a TTY
  printf '%s' "$TWELVEDATA_API_KEY" | td login --key-stdin
  td login --profile staging --key-stdin <<<"$KEY"
  td login --key abc123                             # discouraged: leaks to shell history`,
	RunE: func(cmd *cobra.Command, args []string) error {
		key, _ := cmd.Flags().GetString("key")
		key = strings.TrimSpace(key)
		keyStdin, _ := cmd.Flags().GetBool("key-stdin")
		profileFlag, _ := cmd.Flags().GetString("profile")
		profileFlag = strings.TrimSpace(profileFlag)

		if keyStdin && key != "" {
			return errors.New("--key and --key-stdin are mutually exclusive")
		}
		if keyStdin {
			b, err := io.ReadAll(cmd.InOrStdin())
			if err != nil {
				return fmt.Errorf("failed to read key from stdin: %w", err)
			}
			key = strings.TrimSpace(string(b))
			if key == "" {
				return errors.New("--key-stdin: empty input")
			}
		}
		if key == "" {
			v, err := auth.PromptAPIKey()
			if err != nil {
				if errors.Is(err, auth.ErrNotInteractive) {
					return errors.New(`required flag "key" or "key-stdin" not set: not running in an interactive terminal`)
				}
				return err
			}
			key = strings.TrimSpace(v)
		}
		if key == "" {
			return errors.New("API key is required")
		}

		if profileFlag != "" {
			if err := auth.ValidateProfileName(profileFlag); err != nil {
				return err
			}
		}

		profile := profileFlag
		if profile == "" && auth.IsInteractive() {
			chosen, err := chooseLoginProfile()
			if err != nil {
				return err
			}
			profile = chosen
		}
		if profile == "" {
			profile = "default"
		}

		path, backend, err := auth.StoreAPIKey(profile, key)
		if err != nil {
			return err
		}
		if profileFlag != "" {
			if err := auth.SetActiveProfile(profile); err != nil {
				return err
			}
		}

		if isJSON(cmd) {
			payload := map[string]any{
				"success":     true,
				"profile":     profile,
				"storage":     backend.Name(),
				"config_path": path,
			}
			enc := json.NewEncoder(cmd.OutOrStdout())
			enc.SetIndent("", "  ")
			return enc.Encode(payload)
		}
		where := fmt.Sprintf("in %s", backend.Name())
		if !backend.IsSecure() {
			where = fmt.Sprintf("at %s", path)
		}
		fmt.Fprintf(cmd.OutOrStdout(), "API key stored for profile %q %s\n", profile, where)
		return nil
	},
}

func chooseLoginProfile() (string, error) {
	profiles, _ := auth.ListProfiles()
	if len(profiles) == 0 {
		return "default", nil
	}
	name, err := auth.SelectProfile("Save API key to which profile?", append(profiles, auth.ProfileInfo{Name: "+ Create new profile"}))
	if err != nil {
		return "", err
	}
	if name == "+ Create new profile" {
		return auth.PromptText("Enter a name for the new profile", "", func(v string) error {
			return auth.ValidateProfileName(strings.TrimSpace(v))
		})
	}
	return name, nil
}

func init() {
	loginCmd.Flags().String("key", "", "API key as a literal value (leaks to shell history/process listings — prefer --key-stdin)")
	loginCmd.Flags().Bool("key-stdin", false, "Read the API key from stdin (preferred for CI/scripts)")
	rootCmd.AddCommand(loginCmd)
}

// isJSON tells whether a non-error subcommand should emit JSON to stdout.
// --quiet implies JSON; --output json explicitly opts in. Otherwise human text.
func isJSON(cmd *cobra.Command) bool {
	if q, _ := cmd.Flags().GetBool("quiet"); q {
		return true
	}
	if out, _ := cmd.Flags().GetString("output"); out == "json" {
		return true
	}
	return false
}
