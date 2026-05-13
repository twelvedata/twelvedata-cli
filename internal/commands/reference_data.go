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

var GetBondsCmd = &cobra.Command{
	Use:   "bonds",
	Short: "Fixed income",
	Long:  "The fixed income endpoint provides a daily updated list of available bonds. It returns an array containing detailed information about each bond, including identifiers, names, and other relevant attributes.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.ReferenceDataAPI.GetBonds(cmd.Context())

		if v, _ := cmd.Flags().GetString("symbol"); v != "" {
			req = req.Symbol(v)
		}

		if v, _ := cmd.Flags().GetString("exchange"); v != "" {
			req = req.Exchange(v)
		}

		if v, _ := cmd.Flags().GetString("country"); v != "" {
			req = req.Country(v)
		}

		if cmd.Flags().Changed("show-plan") {
			v, _ := cmd.Flags().GetBool("show-plan")
			req = req.ShowPlan(v)
		}

		if cmd.Flags().Changed("page") {
			v, _ := cmd.Flags().GetInt64("page")
			req = req.Page(v)
		}

		if cmd.Flags().Changed("outputsize") {
			v, _ := cmd.Flags().GetInt64("outputsize")
			req = req.Outputsize(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetCommoditiesCmd = &cobra.Command{
	Use:   "commodities",
	Short: "Commodities",
	Long:  "The commodities endpoint provides a daily updated list of available commodity pairs, across precious metals, livestock, softs, grains, etc.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.ReferenceDataAPI.GetCommodities(cmd.Context())

		if v, _ := cmd.Flags().GetString("symbol"); v != "" {
			req = req.Symbol(v)
		}

		if v, _ := cmd.Flags().GetString("category"); v != "" {
			req = req.Category(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetCountriesCmd = &cobra.Command{
	Use:   "countries",
	Short: "Countries",
	Long:  "The countries endpoint provides a comprehensive list of countries, including their ISO codes, official names, capitals, and currencies. This data is essential for applications requiring accurate country information for tasks such as localization, currency conversion, or geographic analysis.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.ReferenceDataAPI.GetCountries(cmd.Context())

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetCrossListingsCmd = &cobra.Command{
	Use:   "cross-listings",
	Short: "Cross listings",
	Long:  "The cross_listings endpoint provides a daily updated list of cross-listed symbols for a specified financial instrument. Cross-listed symbols represent the same security available on multiple exchanges. This endpoint is useful for identifying all the exchanges where a particular security is traded, allowing users to access comprehensive trading information across different markets.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.ReferenceDataAPI.GetCrossListings(cmd.Context())

		if v, _ := cmd.Flags().GetString("symbol"); v != "" {
			req = req.Symbol(v)
		}

		if v, _ := cmd.Flags().GetString("exchange"); v != "" {
			req = req.Exchange(v)
		}

		if v, _ := cmd.Flags().GetString("mic-code"); v != "" {
			req = req.MicCode(v)
		}

		if v, _ := cmd.Flags().GetString("country"); v != "" {
			req = req.Country(v)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetCryptocurrenciesCmd = &cobra.Command{
	Use:   "cryptocurrencies",
	Short: "Cryptocurrency pairs",
	Long:  "The cryptocurrencies endpoint provides a daily updated list of all available cryptos. It returns an array containing detailed information about each cryptocurrency, including its symbol, name, and other relevant identifiers. This endpoint is useful for retrieving a comprehensive catalog of cryptocurrencies for applications that require up-to-date market listings or need to display available crypto assets to users.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.ReferenceDataAPI.GetCryptocurrencies(cmd.Context())

		if v, _ := cmd.Flags().GetString("symbol"); v != "" {
			req = req.Symbol(v)
		}

		if v, _ := cmd.Flags().GetString("exchange"); v != "" {
			req = req.Exchange(v)
		}

		if v, _ := cmd.Flags().GetString("currency-base"); v != "" {
			req = req.CurrencyBase(v)
		}

		if v, _ := cmd.Flags().GetString("currency-quote"); v != "" {
			req = req.CurrencyQuote(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetCryptocurrencyExchangesCmd = &cobra.Command{
	Use:   "cryptocurrency-exchanges",
	Short: "Cryptocurrency exchanges",
	Long:  "The cryptocurrency exchanges endpoint provides a daily updated list of available cryptocurrency exchanges. It returns an array containing details about each exchange, such as exchange names and identifiers.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.ReferenceDataAPI.GetCryptocurrencyExchanges(cmd.Context())

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetEarliestTimestampCmd = &cobra.Command{
	Use:   "earliest-timestamp",
	Short: "Earliest timestamp",
	Long:  "The earliest_timestamp endpoint provides the earliest available date and time for a specified financial instrument at a given data interval. This endpoint is useful for determining the starting point of historical data availability for various assets, such as stocks or currencies, allowing users to understand the time range covered by the data.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.ReferenceDataAPI.GetEarliestTimestamp(cmd.Context())

		if cmd.Flags().Changed("interval") {
			req = req.Interval(twelvedata.IntervalEnum(cmd.Flags().Lookup("interval").Value.String()))
		}

		if v, _ := cmd.Flags().GetString("symbol"); v != "" {
			req = req.Symbol(v)
		}

		if v, _ := cmd.Flags().GetString("figi"); v != "" {
			req = req.Figi(v)
		}

		if v, _ := cmd.Flags().GetString("isin"); v != "" {
			req = req.Isin(v)
		}

		if v, _ := cmd.Flags().GetString("cusip"); v != "" {
			req = req.Cusip(v)
		}

		if v, _ := cmd.Flags().GetString("exchange"); v != "" {
			req = req.Exchange(v)
		}

		if v, _ := cmd.Flags().GetString("mic-code"); v != "" {
			req = req.MicCode(v)
		}

		if v, _ := cmd.Flags().GetString("timezone"); v != "" {
			req = req.Timezone(v)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetEtfCmd = &cobra.Command{
	Use:   "etfs",
	Short: "ETFs",
	Long:  "The ETFs endpoint provides a daily updated list of all available Exchange-Traded Funds. It returns an array containing detailed information about each ETF, including its symbol, name, and other relevant identifiers. This endpoint is useful for retrieving a comprehensive catalog of ETFs for portfolio management, investment tracking, or financial analysis.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.ReferenceDataAPI.GetEtf(cmd.Context())

		if v, _ := cmd.Flags().GetString("symbol"); v != "" {
			req = req.Symbol(v)
		}

		if v, _ := cmd.Flags().GetString("figi"); v != "" {
			req = req.Figi(v)
		}

		if v, _ := cmd.Flags().GetString("isin"); v != "" {
			req = req.Isin(v)
		}

		if v, _ := cmd.Flags().GetString("cusip"); v != "" {
			req = req.Cusip(v)
		}

		if v, _ := cmd.Flags().GetString("cik"); v != "" {
			req = req.Cik(v)
		}

		if v, _ := cmd.Flags().GetString("exchange"); v != "" {
			req = req.Exchange(v)
		}

		if v, _ := cmd.Flags().GetString("mic-code"); v != "" {
			req = req.MicCode(v)
		}

		if v, _ := cmd.Flags().GetString("country"); v != "" {
			req = req.Country(v)
		}

		if cmd.Flags().Changed("show-plan") {
			v, _ := cmd.Flags().GetBool("show-plan")
			req = req.ShowPlan(v)
		}

		if cmd.Flags().Changed("include-delisted") {
			v, _ := cmd.Flags().GetBool("include-delisted")
			req = req.IncludeDelisted(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetExchangeScheduleCmd = &cobra.Command{
	Use:   "exchange-schedule",
	Short: "Exchanges schedule",
	Long:  "The exchanges schedule endpoint provides detailed information about various stock exchanges, including their trading hours and operational days. This data is essential for users who need to know when specific exchanges are open for trading, allowing them to plan their activities around the availability of these markets.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.ReferenceDataAPI.GetExchangeSchedule(cmd.Context())

		if v, _ := cmd.Flags().GetString("mic-name"); v != "" {
			req = req.MicName(v)
		}

		if v, _ := cmd.Flags().GetString("mic-code"); v != "" {
			req = req.MicCode(v)
		}

		if v, _ := cmd.Flags().GetString("country"); v != "" {
			req = req.Country(v)
		}

		if v, _ := cmd.Flags().GetString("date"); v != "" {
			req = req.Date(v)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetExchangesCmd = &cobra.Command{
	Use:   "exchanges",
	Short: "Exchanges",
	Long:  "The exchanges endpoint provides a comprehensive list of all available equity exchanges. It returns an array containing detailed information about each exchange, such as exchange code, name, country, and timezone. This data is updated daily.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.ReferenceDataAPI.GetExchanges(cmd.Context())

		if cmd.Flags().Changed("type") {
			req = req.Type_(twelvedata.TypeEnum(cmd.Flags().Lookup("type").Value.String()))
		}

		if v, _ := cmd.Flags().GetString("name"); v != "" {
			req = req.Name(v)
		}

		if v, _ := cmd.Flags().GetString("code"); v != "" {
			req = req.Code(v)
		}

		if v, _ := cmd.Flags().GetString("country"); v != "" {
			req = req.Country(v)
		}

		if cmd.Flags().Changed("show-plan") {
			v, _ := cmd.Flags().GetBool("show-plan")
			req = req.ShowPlan(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetForexPairsCmd = &cobra.Command{
	Use:   "forex-pairs",
	Short: "Forex pairs",
	Long:  "The forex pairs endpoint provides a comprehensive list of all available foreign exchange currency pairs. It returns an array of forex pairs, which is updated daily.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.ReferenceDataAPI.GetForexPairs(cmd.Context())

		if v, _ := cmd.Flags().GetString("symbol"); v != "" {
			req = req.Symbol(v)
		}

		if v, _ := cmd.Flags().GetString("currency-base"); v != "" {
			req = req.CurrencyBase(v)
		}

		if v, _ := cmd.Flags().GetString("currency-quote"); v != "" {
			req = req.CurrencyQuote(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetFundsCmd = &cobra.Command{
	Use:   "funds",
	Short: "Funds",
	Long:  "The funds endpoint provides a daily updated list of available investment funds. It returns an array containing detailed information about each fund, including identifiers, names, and other relevant attributes.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.ReferenceDataAPI.GetFunds(cmd.Context())

		if v, _ := cmd.Flags().GetString("symbol"); v != "" {
			req = req.Symbol(v)
		}

		if v, _ := cmd.Flags().GetString("figi"); v != "" {
			req = req.Figi(v)
		}

		if v, _ := cmd.Flags().GetString("isin"); v != "" {
			req = req.Isin(v)
		}

		if v, _ := cmd.Flags().GetString("cusip"); v != "" {
			req = req.Cusip(v)
		}

		if v, _ := cmd.Flags().GetString("cik"); v != "" {
			req = req.Cik(v)
		}

		if v, _ := cmd.Flags().GetString("exchange"); v != "" {
			req = req.Exchange(v)
		}

		if v, _ := cmd.Flags().GetString("country"); v != "" {
			req = req.Country(v)
		}

		if cmd.Flags().Changed("show-plan") {
			v, _ := cmd.Flags().GetBool("show-plan")
			req = req.ShowPlan(v)
		}

		if cmd.Flags().Changed("page") {
			v, _ := cmd.Flags().GetInt64("page")
			req = req.Page(v)
		}

		if cmd.Flags().Changed("outputsize") {
			v, _ := cmd.Flags().GetInt64("outputsize")
			req = req.Outputsize(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetInstrumentTypeCmd = &cobra.Command{
	Use:   "instrument-type",
	Short: "Instrument type",
	Long:  "The instrument type endpoint lists all available financial instrument types, such as stocks, ETFs, and cryptos. This information is essential for users to identify and categorize different financial instruments when accessing or analyzing market data.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.ReferenceDataAPI.GetInstrumentType(cmd.Context())

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetIntervalsCmd = &cobra.Command{
	Use:   "intervals",
	Short: "Intervals List",
	Long:  "The intervals endpoint provides a list of supported time intervals that can be used for querying financial data.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.ReferenceDataAPI.GetIntervals(cmd.Context())

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetMarketStateCmd = &cobra.Command{
	Use:   "market-state",
	Short: "Market state",
	Long:  "The market state endpoint provides real-time information on the operational status of all available stock exchanges. It returns data on whether each exchange is currently open or closed, along with the time remaining until the next opening or closing. This endpoint is useful for users who need to monitor exchange hours and plan their trading activities accordingly.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.ReferenceDataAPI.GetMarketState(cmd.Context())

		if v, _ := cmd.Flags().GetString("exchange"); v != "" {
			req = req.Exchange(v)
		}

		if v, _ := cmd.Flags().GetString("code"); v != "" {
			req = req.Code(v)
		}

		if v, _ := cmd.Flags().GetString("country"); v != "" {
			req = req.Country(v)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetStocksCmd = &cobra.Command{
	Use:   "stocks",
	Short: "Stocks",
	Long:  "The stocks endpoint provides a daily updated list of all available stock symbols. It returns an array containing the symbols, which can be used to identify and access specific stock data across various services. This endpoint is essential for users needing to retrieve the latest stock symbol information for further data requests or integration into financial applications.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.ReferenceDataAPI.GetStocks(cmd.Context())

		if v, _ := cmd.Flags().GetString("symbol"); v != "" {
			req = req.Symbol(v)
		}

		if v, _ := cmd.Flags().GetString("figi"); v != "" {
			req = req.Figi(v)
		}

		if v, _ := cmd.Flags().GetString("isin"); v != "" {
			req = req.Isin(v)
		}

		if v, _ := cmd.Flags().GetString("cusip"); v != "" {
			req = req.Cusip(v)
		}

		if v, _ := cmd.Flags().GetString("cik"); v != "" {
			req = req.Cik(v)
		}

		if v, _ := cmd.Flags().GetString("exchange"); v != "" {
			req = req.Exchange(v)
		}

		if v, _ := cmd.Flags().GetString("mic-code"); v != "" {
			req = req.MicCode(v)
		}

		if v, _ := cmd.Flags().GetString("country"); v != "" {
			req = req.Country(v)
		}

		if cmd.Flags().Changed("type") {
			req = req.Type_(twelvedata.TypeEnum(cmd.Flags().Lookup("type").Value.String()))
		}

		if cmd.Flags().Changed("show-plan") {
			v, _ := cmd.Flags().GetBool("show-plan")
			req = req.ShowPlan(v)
		}

		if cmd.Flags().Changed("include-delisted") {
			v, _ := cmd.Flags().GetBool("include-delisted")
			req = req.IncludeDelisted(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetSymbolSearchCmd = &cobra.Command{
	Use:   "symbol-search",
	Short: "Symbol search",
	Long:  "The symbol search endpoint allows users to find financial instruments by name or symbol. It returns a list of matching symbols, ordered by relevance, with the most relevant instrument first. This is useful for quickly locating specific stocks, ETFs, or other financial instruments when only partial information is available.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.ReferenceDataAPI.GetSymbolSearch(cmd.Context())

		if v, _ := cmd.Flags().GetString("symbol"); v != "" {
			req = req.Symbol(v)
		}

		if cmd.Flags().Changed("outputsize") {
			v, _ := cmd.Flags().GetInt64("outputsize")
			req = req.Outputsize(v)
		}

		if cmd.Flags().Changed("show-plan") {
			v, _ := cmd.Flags().GetBool("show-plan")
			req = req.ShowPlan(v)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTechnicalIndicatorsCmd = &cobra.Command{
	Use:   "technical-indicators",
	Short: "Technical indicators",
	Long:  "The technical indicators endpoint provides a comprehensive list of available technical indicators, each represented as an object. This endpoint is useful for developers looking to integrate a variety of technical analysis tools into their applications, allowing for streamlined access to indicator data without needing to manually configure each one.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.ReferenceDataAPI.GetTechnicalIndicators(cmd.Context())

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

func init() {

	GetBondsCmd.Flags().String("symbol", "", "The ticker symbol of an instrument for which data is requested")

	GetBondsCmd.Flags().String("exchange", "", "Filter by exchange name")

	GetBondsCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetBondsCmd.Flags().Bool("show-plan", false, "Adds info on which plan symbol is available")

	GetBondsCmd.Flags().Int64("page", 0, "Page number of the results to fetch")

	GetBondsCmd.Flags().Int64("outputsize", 0, "Determines the number of data points returned in the output")

	rootCmd.AddCommand(GetBondsCmd)

	GetCommoditiesCmd.Flags().String("symbol", "", "The ticker symbol of an instrument for which data is requested")

	GetCommoditiesCmd.Flags().String("category", "", "Filter by category of commodity")

	rootCmd.AddCommand(GetCommoditiesCmd)

	rootCmd.AddCommand(GetCountriesCmd)

	GetCrossListingsCmd.Flags().String("symbol", "", "The ticker symbol of an instrument for which data is requested")

	GetCrossListingsCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetCrossListingsCmd.Flags().String("mic-code", "", "Market identifier code (MIC) under ISO 10383 standard")

	GetCrossListingsCmd.Flags().String("country", "", "Country to which stock exchange belongs to")

	rootCmd.AddCommand(GetCrossListingsCmd)

	GetCryptocurrenciesCmd.Flags().String("symbol", "", "The ticker symbol of an instrument for which data is requested")

	GetCryptocurrenciesCmd.Flags().String("exchange", "", "Filter by exchange name. E.g. `Binance`, `Coinbase`, etc.")

	GetCryptocurrenciesCmd.Flags().String("currency-base", "", "Filter by currency base")

	GetCryptocurrenciesCmd.Flags().String("currency-quote", "", "Filter by currency quote")

	rootCmd.AddCommand(GetCryptocurrenciesCmd)

	rootCmd.AddCommand(GetCryptocurrencyExchangesCmd)

	flagx.Register(GetEarliestTimestampCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series.")

	GetEarliestTimestampCmd.Flags().String("symbol", "", "Symbol ticker of the instrument.")

	GetEarliestTimestampCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetEarliestTimestampCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetEarliestTimestampCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetEarliestTimestampCmd.Flags().String("exchange", "", "Exchange where instrument is traded.")

	GetEarliestTimestampCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard.")

	GetEarliestTimestampCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here. Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetEarliestTimestampCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetEarliestTimestampCmd)

	GetEtfCmd.Flags().String("symbol", "", "The ticker symbol of an instrument for which data is requested")

	GetEtfCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetEtfCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetEtfCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetEtfCmd.Flags().String("cik", "", "The CIK of an instrument for which data is requested")

	GetEtfCmd.Flags().String("exchange", "", "Filter by exchange name")

	GetEtfCmd.Flags().String("mic-code", "", "Filter by market identifier code (MIC) under ISO 10383 standard")

	GetEtfCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetEtfCmd.Flags().Bool("show-plan", false, "Adds info on which plan symbol is available")

	GetEtfCmd.Flags().Bool("include-delisted", false, "Include delisted identifiers")

	rootCmd.AddCommand(GetEtfCmd)

	GetExchangeScheduleCmd.Flags().String("mic-name", "", "Filter by exchange name")

	GetExchangeScheduleCmd.Flags().String("mic-code", "", "Filter by market identifier code (MIC) under ISO 10383 standard")

	GetExchangeScheduleCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetExchangeScheduleCmd.Flags().String("date", "", "If a date is provided, the API returns the schedule for the specified date; otherwise, it returns the default (common) schedule. The date can be specified in one of the following formats: An exact date (e.g., 2021-10-27) A human-readable keyword: today or yesterday A full datetime string in UTC (e.g., 2025-04-11T20:00:00) to retrieve the schedule corresponding to the day in the specified time. When using a datetime value, the resulting schedule will correspond to the local calendar day at the specified time. For example, 2025-04-11T20:00:00 UTC corresponds to: 2025-04-11 in the America/New_York timezone 2025-04-12 in the Australia/Sydney timezone")

	rootCmd.AddCommand(GetExchangeScheduleCmd)

	flagx.Register(GetExchangesCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetExchangesCmd.Flags().String("name", "", "Filter by exchange name")

	GetExchangesCmd.Flags().String("code", "", "Filter by market identifier code (MIC) under ISO 10383 standard")

	GetExchangesCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetExchangesCmd.Flags().Bool("show-plan", false, "Adds info on which plan symbol is available")

	rootCmd.AddCommand(GetExchangesCmd)

	GetForexPairsCmd.Flags().String("symbol", "", "The ticker symbol of an instrument for which data is requested")

	GetForexPairsCmd.Flags().String("currency-base", "", "Filter by currency base")

	GetForexPairsCmd.Flags().String("currency-quote", "", "Filter by currency quote")

	rootCmd.AddCommand(GetForexPairsCmd)

	GetFundsCmd.Flags().String("symbol", "", "The ticker symbol of an instrument for which data is requested")

	GetFundsCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetFundsCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetFundsCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetFundsCmd.Flags().String("cik", "", "The CIK of an instrument for which data is requested")

	GetFundsCmd.Flags().String("exchange", "", "Filter by exchange name")

	GetFundsCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetFundsCmd.Flags().Bool("show-plan", false, "Adds info on which plan symbol is available")

	GetFundsCmd.Flags().Int64("page", 0, "Page number of the results to fetch")

	GetFundsCmd.Flags().Int64("outputsize", 0, "Determines the number of data points returned in the output")

	rootCmd.AddCommand(GetFundsCmd)

	rootCmd.AddCommand(GetInstrumentTypeCmd)

	rootCmd.AddCommand(GetIntervalsCmd)

	GetMarketStateCmd.Flags().String("exchange", "", "The exchange name where the instrument is traded.")

	GetMarketStateCmd.Flags().String("code", "", "The Market Identifier Code (MIC) of the exchange where the instrument is traded.")

	GetMarketStateCmd.Flags().String("country", "", "The country where the exchange is located. Takes country name or alpha code.")

	rootCmd.AddCommand(GetMarketStateCmd)

	GetStocksCmd.Flags().String("symbol", "", "The ticker symbol of an instrument for which data is requested")

	GetStocksCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetStocksCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetStocksCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetStocksCmd.Flags().String("cik", "", "The CIK of an instrument for which data is requested")

	GetStocksCmd.Flags().String("exchange", "", "Filter by exchange name")

	GetStocksCmd.Flags().String("mic-code", "", "Filter by market identifier code (MIC) under ISO 10383 standard")

	GetStocksCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	flagx.Register(GetStocksCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetStocksCmd.Flags().Bool("show-plan", false, "Adds info on which plan symbol is available")

	GetStocksCmd.Flags().Bool("include-delisted", false, "Include delisted identifiers")

	rootCmd.AddCommand(GetStocksCmd)

	GetSymbolSearchCmd.Flags().String("symbol", "", "Symbol to search. Supports: Ticker symbol of instrument. International securities identification number (ISIN). ISIN access is activating in the Data add-ons section The FIGI (Financial Instrument Global Identifier) parameter is available on the Ultra plan (individual) and Enterprise plan (business) and above. Composite FIGI parameter is available on the Ultra plan (individual) and Enterprise plan (business) and above. Share Class FIGI parameter is available on the Ultra plan (individual) and Enterprise plan (business) and above.")

	GetSymbolSearchCmd.Flags().Int64("outputsize", 0, "Number of matches in response. Max 120")

	GetSymbolSearchCmd.Flags().Bool("show-plan", false, "Adds info on which plan symbol is available.")

	rootCmd.AddCommand(GetSymbolSearchCmd)

	rootCmd.AddCommand(GetTechnicalIndicatorsCmd)
}
