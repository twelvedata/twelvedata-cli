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

var GetEodCmd = &cobra.Command{
	Use:   "eod",
	Short: "End of day price",
	Long:  "The End of Day (EOD) Prices endpoint provides the closing price and other relevant metadata for a financial instrument at the end of a trading day. This endpoint is useful for retrieving daily historical data for stocks, ETFs, or other securities, allowing users to track performance over time and compare daily market movements.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.MarketDataAPI.GetEod(cmd.Context())

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

		if v, _ := cmd.Flags().GetString("country"); v != "" {
			req = req.Country(v)
		}

		if cmd.Flags().Changed("type") {
			req = req.Type_(twelvedata.TypeEnum(cmd.Flags().Lookup("type").Value.String()))
		}

		if v, _ := cmd.Flags().GetString("date"); v != "" {
			req = req.Date(v)
		}

		if cmd.Flags().Changed("prepost") {
			v, _ := cmd.Flags().GetBool("prepost")
			req = req.Prepost(v)
		}

		if cmd.Flags().Changed("dp") {
			v, _ := cmd.Flags().GetInt64("dp")
			req = req.Dp(v)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetPriceCmd = &cobra.Command{
	Use:   "price",
	Short: "Latest price",
	Long:  "The latest price endpoint provides the latest market price for a specified financial instrument. It returns a single data point representing the current (or the most recently available) trading price.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.MarketDataAPI.GetPrice(cmd.Context())

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

		if v, _ := cmd.Flags().GetString("country"); v != "" {
			req = req.Country(v)
		}

		if cmd.Flags().Changed("type") {
			req = req.Type_(twelvedata.TypeEnum(cmd.Flags().Lookup("type").Value.String()))
		}

		if cmd.Flags().Changed("prepost") {
			v, _ := cmd.Flags().GetBool("prepost")
			req = req.Prepost(v)
		}

		if cmd.Flags().Changed("dp") {
			v, _ := cmd.Flags().GetInt64("dp")
			req = req.Dp(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetQuoteCmd = &cobra.Command{
	Use:   "quote",
	Short: "Quote",
	Long:  "The quote endpoint provides real-time data for a selected financial instrument, returning essential information such as the latest price, open, high, low, close, volume, and price change. This endpoint is ideal for users needing up-to-date market data to track price movements and trading activity for specific stocks, ETFs, or other securities.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.MarketDataAPI.GetQuote(cmd.Context())

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

		if cmd.Flags().Changed("interval") {
			req = req.Interval(twelvedata.IntervalEnum(cmd.Flags().Lookup("interval").Value.String()))
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

		if cmd.Flags().Changed("volume-time-period") {
			v, _ := cmd.Flags().GetInt64("volume-time-period")
			req = req.VolumeTimePeriod(v)
		}

		if cmd.Flags().Changed("type") {
			req = req.Type_(twelvedata.TypeEnum(cmd.Flags().Lookup("type").Value.String()))
		}

		if cmd.Flags().Changed("prepost") {
			v, _ := cmd.Flags().GetBool("prepost")
			req = req.Prepost(v)
		}

		if cmd.Flags().Changed("eod") {
			v, _ := cmd.Flags().GetBool("eod")
			req = req.Eod(v)
		}

		if cmd.Flags().Changed("rolling-period") {
			v, _ := cmd.Flags().GetInt64("rolling-period")
			req = req.RollingPeriod(v)
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

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesCmd = &cobra.Command{
	Use:   "time-series",
	Short: "Time series",
	Long:  "The time series endpoint provides detailed historical data for a specified financial instrument. It returns two main components: metadata, which includes essential information about the instrument, and a time series dataset. The time series consists of chronological entries with Open, High, Low, and Close prices, and for applicable instruments, it also includes trading volume. This endpoint is ideal for retrieving comprehensive historical price data for analysis or visualization purposes.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.MarketDataAPI.GetTimeSeries(cmd.Context())

		if cmd.Flags().Changed("interval") {
			req = req.Interval(twelvedata.IntervalEnum(cmd.Flags().Lookup("interval").Value.String()))
		}

		if v, _ := cmd.Flags().GetString("symbol"); v != "" {
			req = req.Symbol(v)
		}

		if v, _ := cmd.Flags().GetString("isin"); v != "" {
			req = req.Isin(v)
		}

		if v, _ := cmd.Flags().GetString("figi"); v != "" {
			req = req.Figi(v)
		}

		if v, _ := cmd.Flags().GetString("cusip"); v != "" {
			req = req.Cusip(v)
		}

		if cmd.Flags().Changed("outputsize") {
			v, _ := cmd.Flags().GetInt64("outputsize")
			req = req.Outputsize(v)
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

		if v, _ := cmd.Flags().GetString("timezone"); v != "" {
			req = req.Timezone(v)
		}

		if v, _ := cmd.Flags().GetString("start-date"); v != "" {
			req = req.StartDate(v)
		}

		if v, _ := cmd.Flags().GetString("end-date"); v != "" {
			req = req.EndDate(v)
		}

		if v, _ := cmd.Flags().GetString("date"); v != "" {
			req = req.Date(v)
		}

		if cmd.Flags().Changed("order") {
			req = req.Order(twelvedata.OrderEnum(cmd.Flags().Lookup("order").Value.String()))
		}

		if cmd.Flags().Changed("prepost") {
			v, _ := cmd.Flags().GetBool("prepost")
			req = req.Prepost(v)
		}

		if cmd.Flags().Changed("dp") {
			v, _ := cmd.Flags().GetInt64("dp")
			req = req.Dp(v)
		}

		if cmd.Flags().Changed("previous-close") {
			v, _ := cmd.Flags().GetBool("previous-close")
			req = req.PreviousClose(v)
		}

		if cmd.Flags().Changed("adjust") {
			req = req.Adjust(twelvedata.AdjustEnum(cmd.Flags().Lookup("adjust").Value.String()))
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesCrossCmd = &cobra.Command{
	Use:   "time-series-cross",
	Short: "Time series cross",
	Long:  "The Time Series Cross endpoint calculates and returns historical cross-rate data for exotic forex pairs, cryptocurrencies, or stocks (e.g., Apple Inc. price in Indian Rupees) on the fly. It provides metadata about the requested symbol and a time series array with Open, High, Low, and Close prices, sorted descending by time, enabling analysis of price history and market trends.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.MarketDataAPI.GetTimeSeriesCross(cmd.Context())

		if v, _ := cmd.Flags().GetString("base"); v != "" {
			req = req.Base(v)
		}

		if v, _ := cmd.Flags().GetString("quote"); v != "" {
			req = req.Quote(v)
		}

		if cmd.Flags().Changed("interval") {
			req = req.Interval(twelvedata.IntervalEnum(cmd.Flags().Lookup("interval").Value.String()))
		}

		if v, _ := cmd.Flags().GetString("base-type"); v != "" {
			req = req.BaseType(v)
		}

		if v, _ := cmd.Flags().GetString("base-exchange"); v != "" {
			req = req.BaseExchange(v)
		}

		if v, _ := cmd.Flags().GetString("base-mic-code"); v != "" {
			req = req.BaseMicCode(v)
		}

		if v, _ := cmd.Flags().GetString("quote-type"); v != "" {
			req = req.QuoteType(v)
		}

		if v, _ := cmd.Flags().GetString("quote-exchange"); v != "" {
			req = req.QuoteExchange(v)
		}

		if v, _ := cmd.Flags().GetString("quote-mic-code"); v != "" {
			req = req.QuoteMicCode(v)
		}

		if cmd.Flags().Changed("outputsize") {
			v, _ := cmd.Flags().GetInt64("outputsize")
			req = req.Outputsize(v)
		}

		if cmd.Flags().Changed("prepost") {
			v, _ := cmd.Flags().GetBool("prepost")
			req = req.Prepost(v)
		}

		if v, _ := cmd.Flags().GetString("start-date"); v != "" {
			req = req.StartDate(v)
		}

		if v, _ := cmd.Flags().GetString("end-date"); v != "" {
			req = req.EndDate(v)
		}

		if cmd.Flags().Changed("adjust") {
			v, _ := cmd.Flags().GetBool("adjust")
			req = req.Adjust(v)
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

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

func init() {

	GetEodCmd.Flags().String("symbol", "", "Symbol ticker of the instrument")

	GetEodCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetEodCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetEodCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetEodCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetEodCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetEodCmd.Flags().String("country", "", "Country where instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetEodCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetEodCmd.Flags().String("date", "", "If not null, then return data from a specific date")

	GetEodCmd.Flags().Bool("prepost", false, "Parameter is optional. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetEodCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values Should be in range [0,11] inclusive")

	rootCmd.AddCommand(GetEodCmd)

	GetPriceCmd.Flags().String("symbol", "", "Symbol ticker of the instrument")

	GetPriceCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetPriceCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetPriceCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetPriceCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetPriceCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetPriceCmd.Flags().String("country", "", "Country where instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetPriceCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetPriceCmd.Flags().Bool("prepost", false, "Parameter is optional. Only for Pro or Venture, and above plans. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume.")

	GetPriceCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0,11] inclusive")

	rootCmd.AddCommand(GetPriceCmd)

	GetQuoteCmd.Flags().String("symbol", "", "Symbol ticker of the instrument")

	GetQuoteCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetQuoteCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetQuoteCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	flagx.Register(GetQuoteCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval of the quote")

	GetQuoteCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetQuoteCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetQuoteCmd.Flags().String("country", "", "Country where instrument is traded, e.g., `United States` or `US`")

	GetQuoteCmd.Flags().Int64("volume-time-period", 0, "Number of periods for Average Volume")

	flagx.Register(GetQuoteCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetQuoteCmd.Flags().Bool("prepost", false, "Parameter is optional. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume.")

	GetQuoteCmd.Flags().Bool("eod", false, "If true, then return data for closed day")

	GetQuoteCmd.Flags().Int64("rolling-period", 0, "Number of hours for calculate rolling change at period. By default set to 24, it can be in range [1, 168].")

	GetQuoteCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values Should be in range [0,11] inclusive")

	GetQuoteCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here. Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	rootCmd.AddCommand(GetQuoteCmd)

	flagx.Register(GetTimeSeriesCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	rootCmd.AddCommand(GetTimeSeriesCmd)

	GetTimeSeriesCrossCmd.Flags().String("base", "", "Base currency symbol")

	GetTimeSeriesCrossCmd.Flags().String("quote", "", "Quote currency symbol")

	flagx.Register(GetTimeSeriesCrossCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesCrossCmd.Flags().String("base-type", "", "Base instrument type according to the `/instrument_type` endpoint")

	GetTimeSeriesCrossCmd.Flags().String("base-exchange", "", "Base exchange")

	GetTimeSeriesCrossCmd.Flags().String("base-mic-code", "", "Base MIC code")

	GetTimeSeriesCrossCmd.Flags().String("quote-type", "", "Quote instrument type according to the `/instrument_type` endpoint")

	GetTimeSeriesCrossCmd.Flags().String("quote-exchange", "", "Quote exchange")

	GetTimeSeriesCrossCmd.Flags().String("quote-mic-code", "", "Quote MIC code")

	GetTimeSeriesCrossCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesCrossCmd.Flags().Bool("prepost", false, "Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume.")

	GetTimeSeriesCrossCmd.Flags().String("start-date", "", "Start date for the time series data")

	GetTimeSeriesCrossCmd.Flags().String("end-date", "", "End date for the time series data")

	GetTimeSeriesCrossCmd.Flags().Bool("adjust", false, "Specifies if there should be an adjustment")

	GetTimeSeriesCrossCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive.")

	GetTimeSeriesCrossCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here. Take note that the IANA Timezone name is case-sensitive")

	rootCmd.AddCommand(GetTimeSeriesCrossCmd)
}
