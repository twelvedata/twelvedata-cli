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

var GetCurrencyConversionCmd = &cobra.Command{
	Use:     "currency-conversion",
	Short:   "Currency conversion",
	Long:    "The currency conversion endpoint provides real-time exchange rates and calculates the converted amount for specified currency pairs, including both forex and cryptocurrencies. This endpoint is useful for obtaining up-to-date conversion values between two currencies, facilitating tasks such as financial reporting, e-commerce transactions, and travel budgeting.",
	Example: "twelvedata currency-conversion --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.CurrenciesAPI.GetCurrencyConversion(cmd.Context())

		if v, _ := cmd.Flags().GetString("symbol"); v != "" {
			req = req.Symbol(v)
		}

		if cmd.Flags().Changed("amount") {
			v, _ := cmd.Flags().GetFloat64("amount")
			req = req.Amount(v)
		}

		if v, _ := cmd.Flags().GetString("date"); v != "" {
			req = req.Date(v)
		}

		if cmd.Flags().Changed("dp") {
			v, _ := cmd.Flags().GetInt64("dp")
			req = req.Dp(v)
		}

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

var GetExchangeRateCmd = &cobra.Command{
	Use:     "exchange-rate",
	Short:   "Exchange rate",
	Long:    "The exchange rate endpoint provides real-time exchange rates for specified currency pairs, including both forex and cryptocurrency. It returns the current exchange rate value between two currencies, allowing users to quickly access up-to-date conversion rates for financial transactions or market analysis.",
	Example: "twelvedata exchange-rate --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.CurrenciesAPI.GetExchangeRate(cmd.Context())

		if v, _ := cmd.Flags().GetString("symbol"); v != "" {
			req = req.Symbol(v)
		}

		if v, _ := cmd.Flags().GetString("date"); v != "" {
			req = req.Date(v)
		}

		if cmd.Flags().Changed("dp") {
			v, _ := cmd.Flags().GetInt64("dp")
			req = req.Dp(v)
		}

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

	GetCurrencyConversionCmd.Flags().String("symbol", "", "The currency pair you want to request can be either forex or cryptocurrency. Slash(`/`) delimiter is used. E.g. `EUR/USD` or `BTC/ETH` will be correct")

	GetCurrencyConversionCmd.Flags().Float64("amount", 0, "Amount of base currency to be converted into quote currency. Supports values in the range from `0` and above")

	GetCurrencyConversionCmd.Flags().String("date", "", "If not null, will use exchange rate from a specific date or time. Format `2006-01-02` or `2006-01-02T15:04:05`. Is set in the local exchange time zone, use timezone parameter to specify a specific time zone")

	GetCurrencyConversionCmd.Flags().Int64("dp", 0, "The number of decimal places for the data")

	GetCurrencyConversionCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports:\n1. Exchange for local exchange time\n2. UTC for datetime at universal UTC standard\n3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here.\nTake note that the IANA Timezone name is case-sensitive")

	_ = GetCurrencyConversionCmd.MarkFlagRequired("symbol")

	_ = GetCurrencyConversionCmd.MarkFlagRequired("amount")

	rootCmd.AddCommand(GetCurrencyConversionCmd)

	GetExchangeRateCmd.Flags().String("symbol", "", "The currency pair you want to request can be either forex or cryptocurrency. Slash(`/`) delimiter is used. E.g. `EUR/USD` or `BTC/ETH` will be correct")

	GetExchangeRateCmd.Flags().String("date", "", "If not null, will use exchange rate from a specific date or time. Format `2006-01-02` or `2006-01-02T15:04:05`. Is set in the local exchange time zone, use timezone parameter to specify a specific time zone")

	GetExchangeRateCmd.Flags().Int64("dp", 0, "The number of decimal places for the data")

	GetExchangeRateCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports:\n1. Exchange for local exchange time\n2. UTC for datetime at universal UTC standard\n3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here.\nTake note that the IANA Timezone name is case-sensitive")

	_ = GetExchangeRateCmd.MarkFlagRequired("symbol")

	rootCmd.AddCommand(GetExchangeRateCmd)
}
