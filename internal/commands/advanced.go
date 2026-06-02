// Twelve Data CLI
//
// NOTE: This code is auto generated, please do not edit it manually.

package commands

import (
	"github.com/spf13/cobra"
	"github.com/twelvedata/twelvedata-go/twelvedata"

	"github.com/twelvedata/twelvedata-cli/internal/client"
	"github.com/twelvedata/twelvedata-cli/internal/flagx"
	"github.com/twelvedata/twelvedata-cli/internal/output"
)

// Sentinel uses for imports that not every per-tag file references in code:
// flagx (only used when an op has enum params), twelvedata (only used when an
// op has enum params or a format setter). Keeping these declared keeps the
// imports "used" in every generated file.
var (
	_ flagx.Enum
	_ twelvedata.APIClient
)

var GetApiUsageCmd = &cobra.Command{
	Use:     "api-usage",
	Short:   "API usage",
	Long:    "The API Usage endpoint provides detailed information on your current API usage statistics. It returns data such as the number of requests made, remaining requests, and the reset time for your usage limits. This endpoint is essential for monitoring and managing your API consumption to ensure you stay within your allocated limits.",
	Example: "twelvedata api-usage",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.AdvancedAPI.GetApiUsage(cmd.Context())

		if v, _ := cmd.Flags().GetString("timezone"); v != "" {
			req = req.Timezone(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		sp := output.StartSpinner(cmd)
		resp, httpResp, callErr := req.Execute()
		sp.Stop()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

func init() {

	GetApiUsageCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports:\n1. UTC for datetime at universal UTC standard\n2. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here.\nTake note that the IANA Timezone name is case-sensitive")

	rootCmd.AddCommand(GetApiUsageCmd)
}
