package client

import (
	"github.com/spf13/cobra"
	"github.com/twelvedata/twelvedata-go/twelvedata"
)

// New constructs a Twelve Data API client, honoring --api-key over the
// TWELVEDATA_API_KEY env var. The SDK's NewConfig handles the env-var fallback,
// the X-API-Version header, and the source=client-go attribution transport.
func New(cmd *cobra.Command) (*twelvedata.APIClient, error) {
	key, _ := cmd.Flags().GetString("api-key")
	cfg, err := twelvedata.NewConfig(key)
	if err != nil {
		return nil, err
	}
	return twelvedata.NewAPIClient(cfg), nil
}
