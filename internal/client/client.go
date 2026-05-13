package client

import (
	"github.com/spf13/cobra"
	"github.com/twelvedata/twelvedata-go/twelvedata"

	"github.com/twelvedata/twelvedata-cli/internal/auth"
)

// New constructs a Twelve Data API client using the auth resolution chain:
// --api-key flag → TWELVEDATA_API_KEY env → secure storage / config file for
// the active profile. The resolved key is passed to NewConfig as-is so the
// SDK's internal env-var fallback never triggers and the source remains
// inspectable via `td whoami`.
func New(cmd *cobra.Command) (*twelvedata.APIClient, error) {
	keyFlag, _ := cmd.Flags().GetString("api-key")
	profileFlag, _ := cmd.Flags().GetString("profile")
	resolved, err := auth.ResolveAPIKey(keyFlag, profileFlag)
	if err != nil {
		return nil, err
	}
	cfg, err := twelvedata.NewConfigWithSource("cli", resolved.Key)
	if err != nil {
		return nil, err
	}
	return twelvedata.NewAPIClient(cfg), nil
}
