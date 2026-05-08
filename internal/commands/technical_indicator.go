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

var GetTimeSeriesAdCmd = &cobra.Command{
	Use:   "ad",
	Short: "Accumulation/distribution",
	Long:  "The Accumulation/Distribution (AD) endpoint provides data on the cumulative money flow into and out of a financial instrument, using its closing price, price range, and trading volume. This endpoint returns the AD line, which helps users identify potential buying or selling pressure and assess the strength of price movements.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesAd(cmd.Context())

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

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesAdOscCmd = &cobra.Command{
	Use:   "adosc",
	Short: "Accumulation/distribution oscillator",
	Long:  "The Accumulation/Distribution Oscillator endpoint (ADOSC) calculates a momentum indicator that highlights shifts in buying or selling pressure by analyzing price and volume data over different time frames. It returns numerical values that help users identify potential trend reversals in financial markets.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesAdOsc(cmd.Context())

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

		if cmd.Flags().Changed("fast-period") {
			v, _ := cmd.Flags().GetInt64("fast-period")
			req = req.FastPeriod(v)
		}

		if cmd.Flags().Changed("slow-period") {
			v, _ := cmd.Flags().GetInt64("slow-period")
			req = req.SlowPeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Addition",
	Long:  "The Addition (ADD) endpoint calculates the sum of two input data series, such as technical indicators or price data, and returns the combined result. This endpoint is useful for users who need to aggregate data points to create custom indicators or analyze the combined effect of multiple data series in financial analysis.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesAdd(cmd.Context())

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

		if cmd.Flags().Changed("series-type-1") {
			req = req.SeriesType1(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type-1").Value.String()))
		}

		if cmd.Flags().Changed("series-type-2") {
			req = req.SeriesType2(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type-2").Value.String()))
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesAdxCmd = &cobra.Command{
	Use:   "adx",
	Short: "Average directional index",
	Long:  "The Average Directional Index (ADX) endpoint provides data on the strength of a market trend, regardless of its direction. It returns a numerical value that helps users identify whether a market is trending or moving sideways.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesAdx(cmd.Context())

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

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesAdxrCmd = &cobra.Command{
	Use:   "adxr",
	Short: "Average directional movement index rating",
	Long:  "The Average Directional Movement Index Rating (ADXR) endpoint provides a smoothed measure of trend strength for a specified financial instrument. It returns the ADXR values, which help users assess the consistency of a trend over a given period by reducing short-term fluctuations. This endpoint is useful for traders and analysts who need to evaluate the stability of market trends for better timing of entry and exit points in their trading strategies.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesAdxr(cmd.Context())

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

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesApoCmd = &cobra.Command{
	Use:   "apo",
	Short: "Absolute price oscillator",
	Long:  "The Absolute Price Oscillator (APO) endpoint calculates the difference between two specified moving averages of a financial instrument's price, providing data that helps users identify potential price trends and reversals. The response includes the calculated APO values over a specified time period, which can be used to track momentum changes and assess the strength of price movements.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesApo(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("fast-period") {
			v, _ := cmd.Flags().GetInt64("fast-period")
			req = req.FastPeriod(v)
		}

		if cmd.Flags().Changed("slow-period") {
			v, _ := cmd.Flags().GetInt64("slow-period")
			req = req.SlowPeriod(v)
		}

		if cmd.Flags().Changed("ma-type") {
			req = req.MaType(twelvedata.MaTypeEnum(cmd.Flags().Lookup("ma-type").Value.String()))
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesAroonCmd = &cobra.Command{
	Use:   "aroon",
	Short: "Aroon indicator",
	Long:  "The Aroon Indicator endpoint provides data on the time elapsed since the highest high and lowest low within a specified period, helping users identify the presence and strength of market trends. It returns two values: Aroon Up and Aroon Down, which indicate the trend direction and momentum. This endpoint is useful for traders and analysts looking to assess trend patterns and potential reversals in financial markets.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesAroon(cmd.Context())

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

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesAroonOscCmd = &cobra.Command{
	Use:   "aroonosc",
	Short: "Aroon oscillator",
	Long:  "The Aroon Oscillator endpoint provides the calculated difference between the Aroon Up and Aroon Down indicators for a given financial instrument. It returns a time series of values that help users identify the strength and direction of a trend, as well as potential trend reversals. This data is useful for traders and analysts seeking to evaluate market trends over a specified period.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesAroonOsc(cmd.Context())

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

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesAtrCmd = &cobra.Command{
	Use:   "atr",
	Short: "Average true range",
	Long:  "The Average True Range (ATR) endpoint provides data on market volatility by calculating the average range of price movement over a user-defined period. It returns numerical values representing the ATR for each time interval, allowing users to gauge the degree of price fluctuation in a financial instrument. This data is useful for setting stop-loss levels and determining optimal entry and exit points in trading strategies.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesAtr(cmd.Context())

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

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesAvgCmd = &cobra.Command{
	Use:   "avg",
	Short: "Average",
	Long:  "The Average (AVG) endpoint calculates the arithmetic mean of a specified data series over a chosen time period. It returns a smoothed dataset that helps users identify trends by reducing short-term fluctuations. This endpoint is useful for obtaining a clearer view of data trends, particularly in time series analysis.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesAvg(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesAvgPriceCmd = &cobra.Command{
	Use:   "avgprice",
	Short: "Average price",
	Long:  "The Average Price (AVGPRICE) endpoint calculates and returns the mean value of a security's open, high, low, and close prices. This endpoint provides a straightforward metric to assess the overall price level of a security over a specified period.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesAvgPrice(cmd.Context())

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

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesBBandsCmd = &cobra.Command{
	Use:   "bbands",
	Short: "Bollinger bands",
	Long:  "The Bollinger Bands (BBANDS) endpoint calculates and returns three key data points: an upper band, a lower band, and a simple moving average (SMA) for a specified financial instrument. These bands are used to assess market volatility by showing how far prices deviate from the SMA. This information helps users identify potential price reversals and determine whether an asset is overbought or oversold.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesBBands(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("ma-type") {
			req = req.MaType(twelvedata.MaTypeEnum(cmd.Flags().Lookup("ma-type").Value.String()))
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesBetaCmd = &cobra.Command{
	Use:   "beta",
	Short: "Beta indicator",
	Long:  "The Beta Indicator endpoint provides data on a security's sensitivity to market movements by comparing its price changes to a benchmark index. It returns the beta value, which quantifies the systematic risk of the security relative to the market. This information is useful for evaluating how much a security's price is expected to move in relation to market changes.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesBeta(cmd.Context())

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

		if cmd.Flags().Changed("series-type-1") {
			req = req.SeriesType1(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type-1").Value.String()))
		}

		if cmd.Flags().Changed("series-type-2") {
			req = req.SeriesType2(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type-2").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesBopCmd = &cobra.Command{
	Use:   "bop",
	Short: "Balance of power",
	Long:  "The Balance of Power (BOP) endpoint provides data on the buying and selling pressure of a security by analyzing its open, high, low, and close prices. It returns numerical values that help users detect shifts in market sentiment and potential price movements.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesBop(cmd.Context())

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

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesCciCmd = &cobra.Command{
	Use:   "cci",
	Short: "Commodity channel index",
	Long:  "The Commodity Channel Index (CCI) endpoint provides data on the CCI values for a specified security, helping users detect potential price reversals by identifying overbought or oversold conditions. It returns a series of CCI values calculated over a specified time period, allowing users to assess the momentum of a security relative to its average price range.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesCci(cmd.Context())

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

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesCeilCmd = &cobra.Command{
	Use:   "ceil",
	Short: "Ceiling",
	Long:  "The Ceiling (CEIL) endpoint rounds each value in the input data series up to the nearest whole number. It returns a series where each original data point is adjusted to its ceiling value, which can be useful for precise calculations or when integrating with other technical indicators that require integer inputs.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesCeil(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesCmoCmd = &cobra.Command{
	Use:   "cmo",
	Short: "Chande momentum oscillator",
	Long:  "The Chande Momentum Oscillator (CMO) endpoint provides data on the momentum of a security by calculating the relative strength of recent price movements. It returns a numerical value indicating whether a security is potentially overbought or oversold, assisting users in identifying possible trend reversals.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesCmo(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesCoppockCmd = &cobra.Command{
	Use:   "coppock",
	Short: "Coppock curve",
	Long:  "The Coppock Curve is a momentum oscillator used to detect potential long-term trend reversals in financial markets. It returns the calculated values of this indicator over a specified period, allowing users to identify when a security's price may be shifting from a downtrend to an uptrend. This endpoint is particularly useful for analyzing securities in bottoming markets.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesCoppock(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("wma-period") {
			v, _ := cmd.Flags().GetInt64("wma-period")
			req = req.WmaPeriod(v)
		}

		if cmd.Flags().Changed("long-roc-period") {
			v, _ := cmd.Flags().GetInt64("long-roc-period")
			req = req.LongRocPeriod(v)
		}

		if cmd.Flags().Changed("short-roc-period") {
			v, _ := cmd.Flags().GetInt64("short-roc-period")
			req = req.ShortRocPeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesCorrelCmd = &cobra.Command{
	Use:   "correl",
	Short: "Correlation",
	Long:  "The Correlation (CORREL) endpoint calculates the statistical relationship between two securities over a specified time period, returning a correlation coefficient. This coefficient ranges from -1 to 1, indicating the strength and direction of their linear relationship. A value close to 1 suggests a strong positive correlation, while a value near -1 indicates a strong negative correlation. This data is useful for identifying securities that move together or in opposite directions, aiding in strategies like diversification or pairs trading.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesCorrel(cmd.Context())

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

		if cmd.Flags().Changed("series-type-1") {
			req = req.SeriesType1(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type-1").Value.String()))
		}

		if cmd.Flags().Changed("series-type-2") {
			req = req.SeriesType2(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type-2").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesCrsiCmd = &cobra.Command{
	Use:   "crsi",
	Short: "Connors relative strength index",
	Long:  "The Connors Relative Strength Index (CRSI) endpoint provides a detailed analysis of stock momentum by combining three components: the Relative Strength Index, the Rate of Change, and the Up/Down Length. This endpoint returns a numerical value that helps identify potential trend reversals and momentum shifts in a security's price. Ideal for traders seeking to refine entry and exit points, the CRSI offers a nuanced view of market conditions beyond traditional RSI indicators.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesCrsi(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("rsi-period") {
			v, _ := cmd.Flags().GetInt64("rsi-period")
			req = req.RsiPeriod(v)
		}

		if cmd.Flags().Changed("up-down-length") {
			v, _ := cmd.Flags().GetInt64("up-down-length")
			req = req.UpDownLength(v)
		}

		if cmd.Flags().Changed("percent-rank-period") {
			v, _ := cmd.Flags().GetInt64("percent-rank-period")
			req = req.PercentRankPeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesDemaCmd = &cobra.Command{
	Use:   "dema",
	Short: "Double exponential moving average",
	Long:  "The Double Exponential Moving Average (DEMA) endpoint provides a data series that calculates a moving average with reduced lag by emphasizing recent price data. This endpoint returns time-series data that includes the DEMA values for a specified financial instrument, allowing users to track price trends and identify potential trading opportunities.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesDema(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesDivCmd = &cobra.Command{
	Use:   "div",
	Short: "Division",
	Long:  "The Division (DIV) endpoint calculates the result of dividing one data series by another, providing a normalized output. It is commonly used to combine or adjust multiple technical indicators or price data for comparative analysis. This endpoint returns the division results as a time series, allowing users to easily interpret and utilize the normalized data in their financial models or charts.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesDiv(cmd.Context())

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

		if cmd.Flags().Changed("series-type-1") {
			req = req.SeriesType1(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type-1").Value.String()))
		}

		if cmd.Flags().Changed("series-type-2") {
			req = req.SeriesType2(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type-2").Value.String()))
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesDpoCmd = &cobra.Command{
	Use:   "dpo",
	Short: "Detrended price oscillator",
	Long:  "The Detrended Price Oscillator (DPO) endpoint calculates and returns the DPO values for a specified financial instrument over a given time period. This endpoint helps traders by highlighting short-term price cycles and identifying potential overbought or oversold conditions without the influence of long-term trends. The response includes a series of DPO values, which can be used to assess price momentum and cyclical patterns in the market.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesDpo(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("centered") {
			v, _ := cmd.Flags().GetBool("centered")
			req = req.Centered(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesDxCmd = &cobra.Command{
	Use:   "dx",
	Short: "Directional movement index",
	Long:  "Retrieve the Directional Movement Index (DX) values for a given security to assess the strength of its positive and negative price movements. This endpoint provides a time series of DX values, which are useful for evaluating the momentum and trend direction of the security over a specified period.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesDx(cmd.Context())

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

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesEmaCmd = &cobra.Command{
	Use:   "ema",
	Short: "Exponential moving average",
	Long:  "The Exponential Moving Average (EMA) endpoint calculates the EMA for a specified financial instrument over a given time period. It returns a time series of EMA values, which highlight recent price trends by weighting recent data more heavily. This is useful for traders seeking to identify trend directions and potential trade opportunities based on recent price movements.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesEma(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesExpCmd = &cobra.Command{
	Use:   "exp",
	Short: "Exponential",
	Long:  "The Exponential (EXP) Indicator endpoint computes the exponential value of a specified input, providing a numerical result that is commonly applied in complex mathematical and financial computations.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesExp(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesFloorCmd = &cobra.Command{
	Use:   "floor",
	Short: "Floor",
	Long:  "The Floor (FLOOR) endpoint processes numerical input data by rounding each value down to the nearest integer. It returns a series of adjusted data points that can be used for further calculations or combined with other datasets. This endpoint is useful for users needing to simplify data by removing decimal precision, aiding in scenarios where integer values are required.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesFloor(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesHeikinashiCandlesCmd = &cobra.Command{
	Use:   "heikinashicandles",
	Short: "Heikinashi candles",
	Long:  "The heikinashi candles endpoint provides smoothed candlestick data by averaging price information to reduce market noise. It returns a series of Heikin Ashi candles, which include open, high, low, and close values, making it easier to identify trends and potential reversals in asset prices. This endpoint is useful for traders and analysts seeking a clearer view of market trends without the volatility present in traditional candlestick charts.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesHeikinashiCandles(cmd.Context())

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

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesHlc3Cmd = &cobra.Command{
	Use:   "hlc3",
	Short: "High, low, close average",
	Long:  "The High, Low, Close Average (HLC3) endpoint calculates and returns the average of a security's high, low, and close prices for a specified period. This endpoint provides a straightforward metric to assess price trends, helping users quickly identify the average price level of a security over time.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesHlc3(cmd.Context())

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

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesHtDcPeriodCmd = &cobra.Command{
	Use:   "ht-dcperiod",
	Short: "Hilbert transform dominant cycle period",
	Long:  "The Hilbert Transform Dominant Cycle Period (HT_DCPERIOD) endpoint calculates the dominant cycle length of a financial instrument's price data. It returns a numerical value representing the cycle period, which traders can use to identify prevailing market cycles and adjust their trading strategies accordingly.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesHtDcPeriod(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesHtDcPhaseCmd = &cobra.Command{
	Use:   "ht-dcphase",
	Short: "Hilbert transform dominant cycle phase",
	Long:  "The Hilbert Transform Dominant Cycle Phase (HT_DCPHASE) endpoint provides the current phase of the dominant market cycle for a given financial instrument. It returns numerical data indicating the phase angle, which can be used by traders to identify potential market entry and exit points based on cyclical patterns.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesHtDcPhase(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesHtPhasorCmd = &cobra.Command{
	Use:   "ht-phasor",
	Short: "Hilbert transform phasor components",
	Long:  "The Hilbert Transform Phasor Components (HT_PHASOR) endpoint analyzes a price series to return two key components: in-phase and quadrature. These components help identify cyclical patterns and the direction of trends in the data. Use this endpoint to gain precise insights into the timing and strength of market cycles, enhancing your ability to track and predict price movements.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesHtPhasor(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesHtSineCmd = &cobra.Command{
	Use:   "ht-sine",
	Short: "Hilbert transform sine wave",
	Long:  "The Hilbert Transform Sine Wave (HT_SINE) endpoint provides sine and cosine wave components derived from the dominant market cycle. This data helps traders pinpoint potential market turning points and assess trend directions by analyzing cyclical patterns.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesHtSine(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesHtTrendModeCmd = &cobra.Command{
	Use:   "ht-trendmode",
	Short: "Hilbert transform trend vs cycle mode",
	Long:  "The Hilbert Transform Trend vs Cycle Mode (HT_TRENDMODE) endpoint identifies whether a market is in a trending or cyclical phase. It returns data indicating the current market phase, allowing users to adjust their trading strategies based on the prevailing conditions.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesHtTrendMode(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesHtTrendlineCmd = &cobra.Command{
	Use:   "ht-trendline",
	Short: "Hilbert transform instantaneous trendline",
	Long:  "The Hilbert Transform Instantaneous Trendline (HT_TRENDLINE) endpoint provides a smoothed moving average that aligns with the dominant market cycle. It returns data points that help traders identify current market trends and determine potential entry or exit points in trading.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesHtTrendline(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesIchimokuCmd = &cobra.Command{
	Use:   "ichimoku",
	Short: "Ichimoku cloud",
	Long:  "The Ichimoku Cloud endpoint provides data on the Ichimoku Kinko Hyo indicator, offering insights into trend direction, support and resistance levels, and potential entry and exit points. It returns key components such as the Tenkan-sen, Kijun-sen, Senkou Span A, Senkou Span B, and Chikou Span. This data helps users evaluate market trends and identify strategic trading opportunities.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesIchimoku(cmd.Context())

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

		if cmd.Flags().Changed("conversion-line-period") {
			v, _ := cmd.Flags().GetInt64("conversion-line-period")
			req = req.ConversionLinePeriod(v)
		}

		if cmd.Flags().Changed("base-line-period") {
			v, _ := cmd.Flags().GetInt64("base-line-period")
			req = req.BaseLinePeriod(v)
		}

		if cmd.Flags().Changed("leading-span-b-period") {
			v, _ := cmd.Flags().GetInt64("leading-span-b-period")
			req = req.LeadingSpanBPeriod(v)
		}

		if cmd.Flags().Changed("lagging-span-period") {
			v, _ := cmd.Flags().GetInt64("lagging-span-period")
			req = req.LaggingSpanPeriod(v)
		}

		if cmd.Flags().Changed("include-ahead-span-period") {
			v, _ := cmd.Flags().GetBool("include-ahead-span-period")
			req = req.IncludeAheadSpanPeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesKamaCmd = &cobra.Command{
	Use:   "kama",
	Short: "Kaufman adaptive moving average",
	Long:  "The Kaufman Adaptive Moving Average (KAMA) endpoint calculates the KAMA for a specified financial instrument, returning a time series of values that reflect the average price adjusted for market volatility. This endpoint helps users identify trends by smoothing out price fluctuations while remaining sensitive to significant price movements.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesKama(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesKeltnerCmd = &cobra.Command{
	Use:   "keltner",
	Short: "Keltner channel",
	Long:  "The Keltner Channel endpoint provides data for a volatility-based technical indicator that combines the Exponential Moving Average (EMA) and the Average True Range (ATR) to form a channel around a security's price. This endpoint returns the upper, middle, and lower bands of the channel, which can be used to identify potential overbought or oversold conditions, assess trend direction, and detect possible price breakouts.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesKeltner(cmd.Context())

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

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("atr-time-period") {
			v, _ := cmd.Flags().GetInt64("atr-time-period")
			req = req.AtrTimePeriod(v)
		}

		if cmd.Flags().Changed("multiplier") {
			v, _ := cmd.Flags().GetInt64("multiplier")
			req = req.Multiplier(v)
		}

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("ma-type") {
			req = req.MaType(twelvedata.MaTypeEnum(cmd.Flags().Lookup("ma-type").Value.String()))
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesKstCmd = &cobra.Command{
	Use:   "kst",
	Short: "Know sure thing",
	Long:  "The Know Sure Thing (KST) endpoint provides a momentum oscillator that combines four smoothed rates of change into a single trend-following indicator. This endpoint returns data that helps users identify potential trend reversals, as well as overbought or oversold conditions in the market.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesKst(cmd.Context())

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

		if cmd.Flags().Changed("roc-period-1") {
			v, _ := cmd.Flags().GetInt64("roc-period-1")
			req = req.RocPeriod1(v)
		}

		if cmd.Flags().Changed("roc-period-2") {
			v, _ := cmd.Flags().GetInt64("roc-period-2")
			req = req.RocPeriod2(v)
		}

		if cmd.Flags().Changed("roc-period-3") {
			v, _ := cmd.Flags().GetInt64("roc-period-3")
			req = req.RocPeriod3(v)
		}

		if cmd.Flags().Changed("roc-period-4") {
			v, _ := cmd.Flags().GetInt64("roc-period-4")
			req = req.RocPeriod4(v)
		}

		if cmd.Flags().Changed("sma-period-1") {
			v, _ := cmd.Flags().GetInt64("sma-period-1")
			req = req.SmaPeriod1(v)
		}

		if cmd.Flags().Changed("sma-period-2") {
			v, _ := cmd.Flags().GetInt64("sma-period-2")
			req = req.SmaPeriod2(v)
		}

		if cmd.Flags().Changed("sma-period-3") {
			v, _ := cmd.Flags().GetInt64("sma-period-3")
			req = req.SmaPeriod3(v)
		}

		if cmd.Flags().Changed("sma-period-4") {
			v, _ := cmd.Flags().GetInt64("sma-period-4")
			req = req.SmaPeriod4(v)
		}

		if cmd.Flags().Changed("signal-period") {
			v, _ := cmd.Flags().GetInt64("signal-period")
			req = req.SignalPeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesLinearRegCmd = &cobra.Command{
	Use:   "linearreg",
	Short: "Linear regression",
	Long:  "The Linear Regression endpoint (LINEARREG) calculates the best-fit straight line through a series of financial data points. It returns the slope and intercept values of this line, allowing users to determine the overall direction of a market trend and identify potential support or resistance levels.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesLinearReg(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesLinearRegAngleCmd = &cobra.Command{
	Use:   "linearregangle",
	Short: "Linear regression angle",
	Long:  "The Linear Regression Angle endpoint (LINEARREGANGLE) calculates the angle of the linear regression line for a given time series of stock prices. It returns the slope of the trend line, expressed in degrees, which helps users identify the direction and steepness of a trend over a specified period. This data is useful for detecting upward or downward trends in asset prices.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesLinearRegAngle(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesLinearRegInterceptCmd = &cobra.Command{
	Use:   "linearregintercept",
	Short: "Linear regression intercept",
	Long:  "The Linear Regression Intercept endpoint (LINEARREGINTERCEPT) calculates the y-intercept of a linear regression line for a given dataset. It returns the value where the regression line crosses the y-axis, providing a numerical reference point for understanding the starting position of a trend over a specified period. This can be useful for users needing to establish baseline values in their data analysis.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesLinearRegIntercept(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesLinearRegSlopeCmd = &cobra.Command{
	Use:   "linearregslope",
	Short: "Linear regression slope",
	Long:  "The Linear Regression Slope endpoint (LINEARREGSLOPE) calculates the slope of a linear regression line for a given dataset, reflecting the rate of change in the data trend over a specified period. It returns a numerical value representing this slope, which can be used to assess the direction and strength of the trend in the dataset.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesLinearRegSlope(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesLnCmd = &cobra.Command{
	Use:   "ln",
	Short: "Natural logarithm",
	Long:  "The Natural Logarithm (LN) endpoint computes the natural logarithm of a specified input value, returning a numerical result. This endpoint is useful for users needing to perform logarithmic transformations on data, which can be applied in various financial calculations and advanced mathematical analyses.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesLn(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesLog10Cmd = &cobra.Command{
	Use:   "log10",
	Short: "Base-10 logarithm",
	Long:  "The Base-10 Logarithm (LOG10) endpoint computes the base-10 logarithm of a specified input value. It returns a numerical result that represents the power to which the number 10 must be raised to obtain the input value. This endpoint is useful for transforming data into a logarithmic scale, which can simplify the analysis of exponential growth patterns or compress large ranges of data in financial calculations.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesLog10(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesMaCmd = &cobra.Command{
	Use:   "ma",
	Short: "Moving average",
	Long:  "The Moving Average (MA) endpoint provides the average price of a security over a specified time frame, offering a smoothed representation of price data. This endpoint returns the calculated moving average values, which can assist users in identifying price trends and potential support or resistance levels in the market.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesMa(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("ma-type") {
			req = req.MaType(twelvedata.MaTypeEnum(cmd.Flags().Lookup("ma-type").Value.String()))
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesMacdCmd = &cobra.Command{
	Use:   "macd",
	Short: "Moving average convergence divergence",
	Long:  "This endpoint calculates the Moving Average Convergence Divergence (MACD) for a specified financial instrument. It returns the MACD line, signal line, and histogram values, which help users identify potential trend reversals and trading opportunities by analyzing the relationship between two moving averages.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesMacd(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("fast-period") {
			v, _ := cmd.Flags().GetInt64("fast-period")
			req = req.FastPeriod(v)
		}

		if cmd.Flags().Changed("slow-period") {
			v, _ := cmd.Flags().GetInt64("slow-period")
			req = req.SlowPeriod(v)
		}

		if cmd.Flags().Changed("signal-period") {
			v, _ := cmd.Flags().GetInt64("signal-period")
			req = req.SignalPeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesMacdExtCmd = &cobra.Command{
	Use:   "macdext",
	Short: "Moving average convergence divergence extension",
	Long:  "The Moving Average Convergence Divergence Extension (MACDEXT) endpoint provides a customizable version of the MACD indicator, allowing users to specify different moving average types and parameters. It returns data that includes the MACD line, signal line, and histogram values, tailored to the user's chosen settings. This endpoint is useful for traders who require flexibility in analyzing price trends and momentum by adjusting the calculation methods to fit their specific trading strategies.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesMacdExt(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("fast-period") {
			v, _ := cmd.Flags().GetInt64("fast-period")
			req = req.FastPeriod(v)
		}

		if cmd.Flags().Changed("fast-ma-type") {
			req = req.FastMaType(twelvedata.MaTypeEnum(cmd.Flags().Lookup("fast-ma-type").Value.String()))
		}

		if cmd.Flags().Changed("slow-period") {
			v, _ := cmd.Flags().GetInt64("slow-period")
			req = req.SlowPeriod(v)
		}

		if cmd.Flags().Changed("slow-ma-type") {
			req = req.SlowMaType(twelvedata.MaTypeEnum(cmd.Flags().Lookup("slow-ma-type").Value.String()))
		}

		if cmd.Flags().Changed("signal-period") {
			v, _ := cmd.Flags().GetInt64("signal-period")
			req = req.SignalPeriod(v)
		}

		if cmd.Flags().Changed("signal-ma-type") {
			req = req.SignalMaType(twelvedata.MaTypeEnum(cmd.Flags().Lookup("signal-ma-type").Value.String()))
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesMacdSlopeCmd = &cobra.Command{
	Use:   "macd-slope",
	Short: "Moving average convergence divergence slope",
	Long:  "The Moving Average Convergence Divergence (MACD) Slope endpoint provides the rate of change of the MACD line for a given security. It returns data on how quickly the MACD line is rising or falling, offering insights into the momentum shifts in the security's price. This information is useful for traders looking to gauge the speed of price movements and potential trend reversals.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesMacdSlope(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("fast-period") {
			v, _ := cmd.Flags().GetInt64("fast-period")
			req = req.FastPeriod(v)
		}

		if cmd.Flags().Changed("slow-period") {
			v, _ := cmd.Flags().GetInt64("slow-period")
			req = req.SlowPeriod(v)
		}

		if cmd.Flags().Changed("signal-period") {
			v, _ := cmd.Flags().GetInt64("signal-period")
			req = req.SignalPeriod(v)
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesMamaCmd = &cobra.Command{
	Use:   "mama",
	Short: "MESA adaptive moving average",
	Long:  "The MESA Adaptive Moving Average (MAMA) endpoint calculates a moving average that adjusts to the dominant market cycle, offering a balance between quick response to price changes and noise reduction. It returns data that includes the adaptive moving average values, which can be used to identify trends and potential reversal points.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesMama(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesMaxCmd = &cobra.Command{
	Use:   "max",
	Short: "Maximum",
	Long:  "The Maximum (MAX) endpoint calculates and returns the highest value within a specified data series over a given period. This endpoint is useful for identifying potential resistance levels or detecting extreme price movements in financial data.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesMax(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesMaxIndexCmd = &cobra.Command{
	Use:   "maxindex",
	Short: "Maximum Index",
	Long:  "The Maximum Index (MAXINDEX) endpoint identifies the position of the highest value within a specified data series over a given time frame. It returns the index where the peak value occurs, allowing users to pinpoint when the maximum price or value was reached in the series. This is useful for tracking the timing of significant peaks in financial data.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesMaxIndex(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesMcGinleyDynamicCmd = &cobra.Command{
	Use:   "mcginley-dynamic",
	Short: "McGinley dynamic indicator",
	Long:  "This endpoint calculates the McGinley Dynamic (MCGINLEY_DYNAMIC) indicator, which provides a refined moving average that adapts to market volatility. This endpoint returns data that reflects smoother price trends and identifies potential support or resistance levels more accurately than traditional moving averages. It is useful for users seeking to track price movements with reduced lag and enhanced responsiveness to market changes.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesMcGinleyDynamic(cmd.Context())

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

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesMedPriceCmd = &cobra.Command{
	Use:   "medprice",
	Short: "Median price",
	Long:  "The Median Price (MEDPRICE) endpoint calculates and returns the average of the high and low prices of a security for a specified period. This endpoint provides a simplified view of price movements, helping users quickly assess price trends by focusing on the midpoint of price action.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesMedPrice(cmd.Context())

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

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesMfiCmd = &cobra.Command{
	Use:   "mfi",
	Short: "Money flow index",
	Long:  "The Money Flow Index (MFI) endpoint provides a volume-weighted momentum oscillator that quantifies buying and selling pressure by analyzing positive and negative money flow. It returns data indicating potential overbought or oversold conditions in a financial asset, aiding users in understanding market trends and price movements.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesMfi(cmd.Context())

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

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesMidPointCmd = &cobra.Command{
	Use:   "midpoint",
	Short: "Midpoint",
	Long:  "The Midpoint (MIDPOINT) endpoint calculates the average value between the highest and lowest prices of a financial instrument over a specified period. It returns a time series of midpoint values, which can help users identify price trends and smooth out short-term fluctuations in the data.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesMidPoint(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesMidPriceCmd = &cobra.Command{
	Use:   "midprice",
	Short: "Midprice",
	Long:  "The Midprice (MIDPRICE) endpoint calculates and returns the average of a financial instrument's highest and lowest prices over a specified time period. This data provides a smoothed representation of price movements, helping users identify potential support or resistance levels in the market.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesMidPrice(cmd.Context())

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

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesMinCmd = &cobra.Command{
	Use:   "min",
	Short: "Minimum",
	Long:  "The Minimum (MIN) Indicator endpoint provides the lowest value of a specified data series over a chosen time period. This endpoint is useful for identifying potential support levels or detecting extreme price movements in financial data.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesMin(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesMinIndexCmd = &cobra.Command{
	Use:   "minindex",
	Short: "Minimum index",
	Long:  "The Minimum Index (MININDEX) endpoint identifies the position of the lowest value within a specified data series over a given time frame. It returns the index number corresponding to the earliest occurrence of this minimum value. This is useful for pinpointing when the lowest price or value occurred in a dataset, aiding in time-based analysis of data trends.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesMinIndex(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesMinMaxCmd = &cobra.Command{
	Use:   "minmax",
	Short: "Minimum and maximum",
	Long:  "The Minimum and Maximum (MINMAX) endpoint identifies the lowest and highest values within a specified time frame for a given data series. It returns these extreme values, which can be used to detect potential support and resistance levels or significant price fluctuations in the data.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesMinMax(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesMinMaxIndexCmd = &cobra.Command{
	Use:   "minmaxindex",
	Short: "Minimum and maximum index",
	Long:  "The Minimum and Maximum Index (MINMAXINDEX) endpoint identifies the positions of the lowest and highest values within a specified data series period. It returns indices that indicate when these extreme values occur, allowing users to pinpoint significant price changes over time.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesMinMaxIndex(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesMinusDICmd = &cobra.Command{
	Use:   "minus-di",
	Short: "Minus directional indicator",
	Long:  "The Minus Directional Indicator (MINUS_DI) endpoint calculates and returns the strength of a security's downward price movement over a specified period. This data is useful for traders and analysts looking to identify bearish trends and assess the intensity of price declines in financial markets.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesMinusDI(cmd.Context())

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

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesMinusDMCmd = &cobra.Command{
	Use:   "minus-dm",
	Short: "Minus directional movement",
	Long:  "The Minus Directional Movement endpoint (MINUS_DM) calculates the downward price movement of a security over a specified period. It returns a series of values indicating the strength of downward trends, useful for traders to identify potential selling opportunities or confirm bearish market conditions.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesMinusDM(cmd.Context())

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

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesMomCmd = &cobra.Command{
	Use:   "mom",
	Short: "Momentum",
	Long:  "The Momentum (MOM) endpoint provides data on the rate of change in a security's price over a user-defined period. It returns a series of numerical values indicating the speed and direction of the price movement, which can help users detect emerging trends or potential reversals in the market.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesMom(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesMultCmd = &cobra.Command{
	Use:   "mult",
	Short: "Multiplication",
	Long:  "The Multiplication (MULT) endpoint calculates the product of two input data series, returning a new data series that represents the element-wise multiplication of the inputs. This is useful for combining or adjusting technical indicators or price data to create custom metrics or to normalize values across different scales.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesMult(cmd.Context())

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

		if cmd.Flags().Changed("series-type-1") {
			req = req.SeriesType1(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type-1").Value.String()))
		}

		if cmd.Flags().Changed("series-type-2") {
			req = req.SeriesType2(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type-2").Value.String()))
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesNatrCmd = &cobra.Command{
	Use:   "natr",
	Short: "Normalized average true range",
	Long:  "The Normalized Average True Range (NATR) endpoint provides a volatility indicator that calculates the average range of price movement over a specified period, expressed as a percentage of the security's price. This data allows users to compare volatility levels across different securities easily. The endpoint returns a time series of NATR values, which can be used to assess and compare the price volatility of various financial instruments.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesNatr(cmd.Context())

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

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesObvCmd = &cobra.Command{
	Use:   "obv",
	Short: "On balance volume",
	Long:  "The On Balance Volume (OBV) endpoint provides a time series of the OBV indicator, which calculates cumulative volume to reflect buying and selling pressure over time. This endpoint returns data that helps users track volume trends in relation to price movements, aiding in the identification of potential trend continuations or reversals in a security's price.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesObv(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesPercentBCmd = &cobra.Command{
	Use:   "percent-b",
	Short: "Percent B",
	Long:  "The Percent B (%B) endpoint calculates and returns the %B value, which indicates the position of a security's price relative to its Bollinger Bands. This data helps users determine if a security is near the upper or lower band, potentially signaling overbought or oversold conditions.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesPercentB(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("ma-type") {
			req = req.MaType(twelvedata.MaTypeEnum(cmd.Flags().Lookup("ma-type").Value.String()))
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesPivotPointsHLCmd = &cobra.Command{
	Use:   "pivot-points-hl",
	Short: "Pivot points high low",
	Long:  "The Pivot Points High Low (PIVOT_POINTS_HL) endpoint calculates key support and resistance levels for a security by analyzing its highest and lowest prices over a specified period. This endpoint returns data that includes pivot points, support levels, and resistance levels, which can be used to identify potential price reversal zones and optimize trade entry and exit strategies.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesPivotPointsHL(cmd.Context())

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

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesPlusDICmd = &cobra.Command{
	Use:   "plus-di",
	Short: "Plus directional indicator",
	Long:  "The Plus Directional Indicator endpoint (/plus_di) provides data on the strength of a security's upward price movement by calculating the Plus Directional Indicator (PLUS_DI). It returns a time series of PLUS_DI values, which can be used to assess the intensity of upward trends in a security's price over a specified period.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesPlusDI(cmd.Context())

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

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesPlusDMCmd = &cobra.Command{
	Use:   "plus-dm",
	Short: "Plus directional movement",
	Long:  "The Plus Directional Movement (PLUS_DM) endpoint calculates the upward price movement of a financial security over a specified period. It returns numerical values representing the magnitude of upward price changes, which can be used to assess the strength of an uptrend. This data is essential for traders and analysts who need to evaluate the bullish momentum of a security.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesPlusDM(cmd.Context())

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

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesPpoCmd = &cobra.Command{
	Use:   "ppo",
	Short: "Percentage price oscillator",
	Long:  "The Percentage Price Oscillator (PPO) endpoint calculates the percentage difference between two specified moving averages of a financial instrument's price. It returns data that includes the PPO values, which traders can use to identify potential trend reversals and generate trading signals.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesPpo(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("fast-period") {
			v, _ := cmd.Flags().GetInt64("fast-period")
			req = req.FastPeriod(v)
		}

		if cmd.Flags().Changed("slow-period") {
			v, _ := cmd.Flags().GetInt64("slow-period")
			req = req.SlowPeriod(v)
		}

		if cmd.Flags().Changed("ma-type") {
			req = req.MaType(twelvedata.MaTypeEnum(cmd.Flags().Lookup("ma-type").Value.String()))
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesRocCmd = &cobra.Command{
	Use:   "roc",
	Short: "Rate of change",
	Long:  "The Rate of Change (ROC) endpoint calculates the percentage change in a security's price over a defined period, returning a time series of ROC values. This data helps users track momentum by showing how quickly prices are changing, which can be useful for identifying potential price movements.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesRoc(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesRocpCmd = &cobra.Command{
	Use:   "rocp",
	Short: "Rate of change percentage",
	Long:  "The Rate of Change Percentage (ROCP) endpoint calculates and returns the percentage change in the price of a financial security over a user-defined period. This data helps users identify shifts in price momentum and potential trend reversals by providing a clear numerical representation of how much the price has increased or decreased in percentage terms.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesRocp(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesRocrCmd = &cobra.Command{
	Use:   "rocr",
	Short: "Rate of change ratio",
	Long:  "The Rate of Change Ratio (ROCR) endpoint calculates and returns the ratio of a security's current price to its price from a specified number of periods ago. This data helps users track price momentum and identify potential trend reversals by providing a clear numerical value that reflects price changes over time.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesRocr(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesRocr100Cmd = &cobra.Command{
	Use:   "rocr100",
	Short: "Rate of change ratio 100",
	Long:  "The Rate of Change Ratio 100 (ROCR100) endpoint calculates the percentage change in a security's price over a specified period, expressed as a ratio to 100. It returns data that highlights the momentum of the price movement and identifies potential trend reversals. This endpoint is useful for users looking to assess the strength and direction of a security's price trend over time.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesRocr100(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesRsiCmd = &cobra.Command{
	Use:   "rsi",
	Short: "Relative strength index",
	Long:  "The Relative Strength Index (RSI) endpoint provides data on the RSI values for a specified financial instrument over a given period. It returns a series of RSI values, which indicate the momentum of price movements and help identify potential overbought or oversold conditions. This data is useful for traders looking to assess the strength of price trends and anticipate possible trend reversals.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesRsi(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesRvolCmd = &cobra.Command{
	Use:   "rvol",
	Short: "Relative volume",
	Long:  "The Relative Volume endpoint (/rvol) provides a ratio comparing a security's current trading volume to its average volume over a specified period. This data helps users detect unusual trading activity and assess the strength of price movements, offering insights into potential market breakouts.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesRvol(cmd.Context())

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

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesSarCmd = &cobra.Command{
	Use:   "sar",
	Short: "Parabolic stop and reverse",
	Long:  "The Parabolic Stop and Reverse (SAR) endpoint provides data on potential support and resistance levels for a specified security, using its price and time. This endpoint returns numerical values that help traders determine possible entry and exit points in their trading strategies.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesSar(cmd.Context())

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

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesSarExtCmd = &cobra.Command{
	Use:   "sarext",
	Short: "Parabolic stop and reverse extended",
	Long:  "The Parabolic SAR Extended (SAREXT) endpoint provides a customizable version of the Parabolic SAR indicator, which is used to identify potential entry and exit points in trading. Users can adjust parameters such as acceleration factors to tailor the indicator to specific trading strategies. The endpoint returns data points indicating potential trend reversals.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesSarExt(cmd.Context())

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

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesSmaCmd = &cobra.Command{
	Use:   "sma",
	Short: "Simple moving average",
	Long:  "The Simple Moving Average (SMA) endpoint calculates and returns the average price of a security over a user-defined time period. This endpoint provides a series of data points that represent the smoothed price trend, which can help users identify potential price movements and evaluate historical price behavior.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesSma(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesSqrtCmd = &cobra.Command{
	Use:   "sqrt",
	Short: "Square root",
	Long:  "The Square Root (SQRT) endpoint computes the square root of a specified numerical input. It returns a single numerical value representing the square root, which can be used in various mathematical computations or financial models requiring this specific transformation.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesSqrt(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesStdDevCmd = &cobra.Command{
	Use:   "stddev",
	Short: "Standard deviation",
	Long:  "The Standard Deviation (STDDEV) endpoint calculates the dispersion of a financial instrument's price data from its average value. It returns a numerical value representing the volatility of the asset over a specified period. This endpoint is useful for traders and analysts to assess price variability and identify periods of high or low volatility in the market.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesStdDev(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesStochCmd = &cobra.Command{
	Use:   "stoch",
	Short: "Stochastic oscillator",
	Long:  "The Stochastic Oscillator endpoint provides data on a momentum indicator that evaluates a security's closing price relative to its price range over a specified timeframe. It returns values indicating potential overbought or oversold conditions, aiding in identifying possible trend reversals. Users receive the %K and %D values, which are essential for analyzing the momentum and potential turning points in the market.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesStoch(cmd.Context())

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

		if cmd.Flags().Changed("fast-k-period") {
			v, _ := cmd.Flags().GetInt64("fast-k-period")
			req = req.FastKPeriod(v)
		}

		if cmd.Flags().Changed("slow-k-period") {
			v, _ := cmd.Flags().GetInt64("slow-k-period")
			req = req.SlowKPeriod(v)
		}

		if cmd.Flags().Changed("slow-d-period") {
			v, _ := cmd.Flags().GetInt64("slow-d-period")
			req = req.SlowDPeriod(v)
		}

		if cmd.Flags().Changed("slow-kma-type") {
			req = req.SlowKmaType(twelvedata.MaTypeEnum(cmd.Flags().Lookup("slow-kma-type").Value.String()))
		}

		if cmd.Flags().Changed("slow-dma-type") {
			req = req.SlowDmaType(twelvedata.MaTypeEnum(cmd.Flags().Lookup("slow-dma-type").Value.String()))
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesStochFCmd = &cobra.Command{
	Use:   "stochf",
	Short: "Stochastic fast",
	Long:  "The Stochastic Fast (STOCHF) endpoint calculates the fast version of the Stochastic Oscillator, providing data on the momentum of a financial instrument by comparing a particular closing price to a range of its prices over a specified period. This endpoint returns the %K and %D values, which are used to identify potential overbought or oversold conditions in the market. It is useful for traders who need quick, responsive insights into price movements, although it may generate more false signals due to its sensitivity.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesStochF(cmd.Context())

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

		if cmd.Flags().Changed("fast-k-period") {
			v, _ := cmd.Flags().GetInt64("fast-k-period")
			req = req.FastKPeriod(v)
		}

		if cmd.Flags().Changed("fast-d-period") {
			v, _ := cmd.Flags().GetInt64("fast-d-period")
			req = req.FastDPeriod(v)
		}

		if cmd.Flags().Changed("fast-dma-type") {
			req = req.FastDmaType(twelvedata.MaTypeEnum(cmd.Flags().Lookup("fast-dma-type").Value.String()))
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesStochRsiCmd = &cobra.Command{
	Use:   "stochrsi",
	Short: "Stochastic relative strength index",
	Long:  "The Stochastic Relative Strength Index (Stochastic RSI) endpoint calculates the Stochastic RSI values for a given financial instrument, providing data on its momentum and potential price reversals. This endpoint returns time-series data, including the %K and %D lines, which help users identify overbought or oversold conditions. Ideal for traders seeking to refine entry and exit points by analyzing short-term price movements.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesStochRsi(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeStochrsiEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("rsi-length") {
			v, _ := cmd.Flags().GetInt64("rsi-length")
			req = req.RsiLength(v)
		}

		if cmd.Flags().Changed("stoch-length") {
			v, _ := cmd.Flags().GetInt64("stoch-length")
			req = req.StochLength(v)
		}

		if cmd.Flags().Changed("k-period") {
			v, _ := cmd.Flags().GetInt64("k-period")
			req = req.KPeriod(v)
		}

		if v, _ := cmd.Flags().GetString("slow-kma-type"); v != "" {
			req = req.SlowKmaType(v)
		}

		if cmd.Flags().Changed("d-period") {
			v, _ := cmd.Flags().GetInt64("d-period")
			req = req.DPeriod(v)
		}

		if v, _ := cmd.Flags().GetString("slow-dma-type"); v != "" {
			req = req.SlowDmaType(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesSubCmd = &cobra.Command{
	Use:   "sub",
	Short: "Subtraction",
	Long:  "The Subtraction (SUB) endpoint calculates the difference between two input data series, such as technical indicators or price data. It returns a time series of the resulting values, allowing users to compare or normalize data by highlighting the variance between the two series.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesSub(cmd.Context())

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

		if cmd.Flags().Changed("series-type-1") {
			req = req.SeriesType1(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type-1").Value.String()))
		}

		if cmd.Flags().Changed("series-type-2") {
			req = req.SeriesType2(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type-2").Value.String()))
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesSumCmd = &cobra.Command{
	Use:   "sum",
	Short: "Summation",
	Long:  "The Summation (SUM) endpoint calculates the cumulative total of a specified data series over a defined time period. It returns a numerical value representing the sum, which can be used to track the aggregate value of financial data, such as stock prices or trading volumes, over time. This endpoint is useful for users needing to compute the total accumulation of a dataset for further analysis or reporting.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesSum(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesSuperTrendCmd = &cobra.Command{
	Use:   "supertrend",
	Short: "Supertrend",
	Long:  "The Supertrend endpoint provides data on the Supertrend indicator, a tool used to identify potential buy and sell signals in trending markets. It returns values that indicate the current trend direction and potential reversal points based on price, time, and volatility. Users can leverage this data to pinpoint optimal entry and exit points for trades.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesSuperTrend(cmd.Context())

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

		if cmd.Flags().Changed("period") {
			v, _ := cmd.Flags().GetInt64("period")
			req = req.Period(v)
		}

		if cmd.Flags().Changed("multiplier") {
			v, _ := cmd.Flags().GetInt64("multiplier")
			req = req.Multiplier(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesSuperTrendHeikinAshiCandlesCmd = &cobra.Command{
	Use:   "supertrend-heikinashicandles",
	Short: "Supertrend Heikin Ashi candles",
	Long:  "The Supertrend Heikin Ashi candles endpoint provides data combining Supertrend signals with Heikin Ashi candlestick patterns. It returns a series of data points indicating trend direction and smoothed price movements, useful for identifying potential buy or sell opportunities in trading.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesSuperTrendHeikinAshiCandles(cmd.Context())

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

		if cmd.Flags().Changed("period") {
			v, _ := cmd.Flags().GetInt64("period")
			req = req.Period(v)
		}

		if cmd.Flags().Changed("multiplier") {
			v, _ := cmd.Flags().GetInt64("multiplier")
			req = req.Multiplier(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesT3maCmd = &cobra.Command{
	Use:   "t3ma",
	Short: "Triple exponential moving average",
	Long:  "The Triple Exponential Moving Average (T3MA) endpoint calculates a smoothed moving average using three exponential moving averages on price data. It returns a dataset that highlights price trends with reduced lag, offering precise trend analysis. This is useful for identifying trend direction and potential reversal points.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesT3ma(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesTRangeCmd = &cobra.Command{
	Use:   "trange",
	Short: "True range",
	Long:  "The True Range (TRANGE) endpoint calculates the range of price movement for a specified period, providing a measure of market volatility. It returns data that includes the highest and lowest prices over the period, along with the closing price from the previous period. This information is useful for traders to assess market volatility and adjust their trading strategies accordingly.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesTRange(cmd.Context())

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

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesTemaCmd = &cobra.Command{
	Use:   "tema",
	Short: "Triple exponential moving average",
	Long:  "The Triple Exponential Moving Average (TEMA) endpoint calculates and returns the TEMA values for a specified financial instrument over a given time period. This endpoint provides a series of data points that smooth out price fluctuations by applying three layers of exponential moving averages, allowing users to identify and track underlying trends in the instrument's price movement.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesTema(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesTrimaCmd = &cobra.Command{
	Use:   "trima",
	Short: "Triangular moving average",
	Long:  "The Triangular Moving Average (TRIMA) endpoint calculates and returns the smoothed average price of a financial security over a specified period, with a focus on central data points. This endpoint provides a balanced view of price trends by applying a double smoothing process, making it useful for identifying underlying price patterns and reducing short-term fluctuations.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesTrima(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesTsfCmd = &cobra.Command{
	Use:   "tsf",
	Short: "Time series forecast",
	Long:  "The Time Series Forecast (TSF) endpoint provides projected future price levels using linear regression analysis. It returns data that helps users identify potential support and resistance levels, as well as trend direction in a financial market. This endpoint is useful for traders seeking to anticipate price movements and adjust their strategies accordingly.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesTsf(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesTypPriceCmd = &cobra.Command{
	Use:   "typprice",
	Short: "Typical price",
	Long:  "The Typical Price (TYPPRICE) endpoint calculates and returns the average of a financial instrument's high, low, and close prices for a given period. This endpoint provides a simplified metric that reflects the central tendency of price movements, useful for traders and analysts who need a straightforward view of price trends.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesTypPrice(cmd.Context())

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

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesUltOscCmd = &cobra.Command{
	Use:   "ultosc",
	Short: "Ultimate oscillator endpoint",
	Long:  "The Ultimate Oscillator endpoint (/ultosc) calculates a momentum oscillator that integrates short, intermediate, and long-term price movements to detect potential overbought or oversold conditions and possible trend reversals. It returns a time series of oscillator values, which can be used to assess market momentum and identify entry or exit points in trading strategies.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesUltOsc(cmd.Context())

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

		if cmd.Flags().Changed("time-period-1") {
			v, _ := cmd.Flags().GetInt64("time-period-1")
			req = req.TimePeriod1(v)
		}

		if cmd.Flags().Changed("time-period-2") {
			v, _ := cmd.Flags().GetInt64("time-period-2")
			req = req.TimePeriod2(v)
		}

		if cmd.Flags().Changed("time-period-3") {
			v, _ := cmd.Flags().GetInt64("time-period-3")
			req = req.TimePeriod3(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesVarCmd = &cobra.Command{
	Use:   "var",
	Short: "Variance",
	Long:  "The Variance (VAR) endpoint calculates the statistical variance of a financial data series, providing a measure of how much the data points deviate from the average value. It returns a numerical value representing this dispersion, which can be used to assess the volatility of a security over a specified period. This information is crucial for traders and analysts who need to evaluate the risk associated with price fluctuations in the market.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesVar(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesVwapCmd = &cobra.Command{
	Use:   "vwap",
	Short: "Volume weighted average price",
	Long:  "The Volume Weighted Average Price (VWAP) endpoint provides the VWAP value for a specified stock or asset over a given time period. This indicator calculates the average price at which a security has traded throughout the day, based on both volume and price. It is useful for identifying the true average price of an asset, helping traders to assess the current price relative to the day's average.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesVwap(cmd.Context())

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

		if cmd.Flags().Changed("sd-time-period") {
			v, _ := cmd.Flags().GetInt64("sd-time-period")
			req = req.SdTimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesWclPriceCmd = &cobra.Command{
	Use:   "wclprice",
	Short: "Weighted close price",
	Long:  "The Weighted Close Price (WCLPRICE) endpoint calculates a security's average price by giving additional weight to the closing price, using the formula: (High + Low + Close * 2) / 4.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesWclPrice(cmd.Context())

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

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesWillRCmd = &cobra.Command{
	Use:   "willr",
	Short: "Williams %R",
	Long:  "The Williams %R (WILLR) endpoint calculates the Williams Percent Range, a momentum indicator that evaluates a security's closing price relative to its high-low range over a specified period. This endpoint returns data that helps users identify potential overbought or oversold conditions and possible trend reversals in the market.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesWillR(cmd.Context())

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

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTimeSeriesWmaCmd = &cobra.Command{
	Use:   "wma",
	Short: "Weighted moving average",
	Long:  "The Weighted Moving Average (WMA) endpoint calculates and returns the WMA values for a given security over a specified period. This endpoint provides a time series of weighted averages, where recent prices have a higher influence, allowing users to track and analyze short-term price trends effectively.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.TechnicalIndicatorAPI.GetTimeSeriesWma(cmd.Context())

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

		if cmd.Flags().Changed("series-type") {
			req = req.SeriesType(twelvedata.SeriesTypeEnum(cmd.Flags().Lookup("series-type").Value.String()))
		}

		if cmd.Flags().Changed("time-period") {
			v, _ := cmd.Flags().GetInt64("time-period")
			req = req.TimePeriod(v)
		}

		if cmd.Flags().Changed("include-ohlc") {
			v, _ := cmd.Flags().GetBool("include-ohlc")
			req = req.IncludeOhlc(v)
		}

		if f := csvFormat(cmd); f != nil {
			req = req.Format(*f)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

func init() {

	flagx.Register(GetTimeSeriesAdCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesAdCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesAdCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesAdCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesAdCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesAdCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesAdCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesAdCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesAdCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesAdCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesAdCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesAdCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesAdCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesAdCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesAdCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesAdCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesAdCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesAdCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesAdCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesAdCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesAdCmd)

	flagx.Register(GetTimeSeriesAdOscCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesAdOscCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesAdOscCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesAdOscCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesAdOscCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesAdOscCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesAdOscCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesAdOscCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesAdOscCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesAdOscCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesAdOscCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesAdOscCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesAdOscCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesAdOscCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesAdOscCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesAdOscCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesAdOscCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesAdOscCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesAdOscCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesAdOscCmd.Flags().Int64("fast-period", 0, "Number of periods for fast moving average. Takes values in the range from `1` to `800`")

	GetTimeSeriesAdOscCmd.Flags().Int64("slow-period", 0, "Number of periods for slow moving average. Takes values in the range from `1` to `800`")

	GetTimeSeriesAdOscCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesAdOscCmd)

	flagx.Register(GetTimeSeriesAddCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesAddCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesAddCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesAddCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesAddCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesAddCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesAddCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesAddCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesAddCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesAddCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesAddCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesAddCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesAddCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesAddCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesAddCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesAddCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesAddCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesAddCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesAddCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesAddCmd, "series-type-1", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type used as the first part of technical indicator")

	flagx.Register(GetTimeSeriesAddCmd, "series-type-2", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type used as the second part of technical indicator")

	GetTimeSeriesAddCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesAddCmd)

	flagx.Register(GetTimeSeriesAdxCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesAdxCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesAdxCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesAdxCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesAdxCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesAdxCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesAdxCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesAdxCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesAdxCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesAdxCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesAdxCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesAdxCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesAdxCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesAdxCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesAdxCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesAdxCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesAdxCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesAdxCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesAdxCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesAdxCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesAdxCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesAdxCmd)

	flagx.Register(GetTimeSeriesAdxrCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesAdxrCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesAdxrCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesAdxrCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesAdxrCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesAdxrCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesAdxrCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesAdxrCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesAdxrCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesAdxrCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesAdxrCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesAdxrCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesAdxrCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesAdxrCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesAdxrCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesAdxrCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesAdxrCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesAdxrCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesAdxrCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesAdxrCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesAdxrCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesAdxrCmd)

	flagx.Register(GetTimeSeriesApoCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesApoCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesApoCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesApoCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesApoCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesApoCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesApoCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesApoCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesApoCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesApoCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesApoCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesApoCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesApoCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesApoCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesApoCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesApoCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesApoCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesApoCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesApoCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesApoCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesApoCmd.Flags().Int64("fast-period", 0, "Number of periods for fast moving average. Takes values in the range from `1` to `800`")

	GetTimeSeriesApoCmd.Flags().Int64("slow-period", 0, "Number of periods for slow moving average. Takes values in the range from `1` to `800`")

	flagx.Register(GetTimeSeriesApoCmd, "ma-type", twelvedata.AllowedMaTypeEnumEnumValues, "The type of moving average used")

	GetTimeSeriesApoCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesApoCmd)

	flagx.Register(GetTimeSeriesAroonCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesAroonCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesAroonCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesAroonCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesAroonCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesAroonCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesAroonCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesAroonCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesAroonCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesAroonCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesAroonCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesAroonCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesAroonCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesAroonCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesAroonCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesAroonCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesAroonCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesAroonCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesAroonCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesAroonCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesAroonCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesAroonCmd)

	flagx.Register(GetTimeSeriesAroonOscCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesAroonOscCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesAroonOscCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesAroonOscCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesAroonOscCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesAroonOscCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesAroonOscCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesAroonOscCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesAroonOscCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesAroonOscCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesAroonOscCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesAroonOscCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesAroonOscCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesAroonOscCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesAroonOscCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesAroonOscCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesAroonOscCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesAroonOscCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesAroonOscCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesAroonOscCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesAroonOscCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesAroonOscCmd)

	flagx.Register(GetTimeSeriesAtrCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesAtrCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesAtrCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesAtrCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesAtrCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesAtrCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesAtrCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesAtrCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesAtrCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesAtrCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesAtrCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesAtrCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesAtrCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesAtrCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesAtrCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesAtrCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesAtrCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesAtrCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesAtrCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesAtrCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesAtrCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesAtrCmd)

	flagx.Register(GetTimeSeriesAvgCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesAvgCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesAvgCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesAvgCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesAvgCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesAvgCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesAvgCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesAvgCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesAvgCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesAvgCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesAvgCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesAvgCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesAvgCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesAvgCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesAvgCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesAvgCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesAvgCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesAvgCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesAvgCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesAvgCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesAvgCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesAvgCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesAvgCmd)

	flagx.Register(GetTimeSeriesAvgPriceCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesAvgPriceCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesAvgPriceCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesAvgPriceCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesAvgPriceCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesAvgPriceCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesAvgPriceCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesAvgPriceCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesAvgPriceCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesAvgPriceCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesAvgPriceCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesAvgPriceCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesAvgPriceCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesAvgPriceCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesAvgPriceCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesAvgPriceCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesAvgPriceCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesAvgPriceCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesAvgPriceCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesAvgPriceCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesAvgPriceCmd)

	flagx.Register(GetTimeSeriesBBandsCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesBBandsCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesBBandsCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesBBandsCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesBBandsCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesBBandsCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesBBandsCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesBBandsCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesBBandsCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesBBandsCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesBBandsCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesBBandsCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesBBandsCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesBBandsCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesBBandsCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesBBandsCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesBBandsCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesBBandsCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesBBandsCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesBBandsCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesBBandsCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	flagx.Register(GetTimeSeriesBBandsCmd, "ma-type", twelvedata.AllowedMaTypeEnumEnumValues, "The type of moving average used")

	GetTimeSeriesBBandsCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesBBandsCmd)

	flagx.Register(GetTimeSeriesBetaCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesBetaCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesBetaCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesBetaCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesBetaCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesBetaCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesBetaCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesBetaCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesBetaCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesBetaCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesBetaCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesBetaCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesBetaCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesBetaCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesBetaCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesBetaCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesBetaCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesBetaCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesBetaCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesBetaCmd, "series-type-1", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type used as the first part of technical indicator")

	flagx.Register(GetTimeSeriesBetaCmd, "series-type-2", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type used as the second part of technical indicator")

	GetTimeSeriesBetaCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesBetaCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesBetaCmd)

	flagx.Register(GetTimeSeriesBopCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesBopCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesBopCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesBopCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesBopCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesBopCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesBopCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesBopCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesBopCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesBopCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesBopCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesBopCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesBopCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesBopCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesBopCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesBopCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesBopCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesBopCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesBopCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesBopCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesBopCmd)

	flagx.Register(GetTimeSeriesCciCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesCciCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesCciCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesCciCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesCciCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesCciCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesCciCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesCciCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesCciCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesCciCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesCciCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesCciCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesCciCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesCciCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesCciCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesCciCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesCciCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesCciCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesCciCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesCciCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesCciCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesCciCmd)

	flagx.Register(GetTimeSeriesCeilCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesCeilCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesCeilCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesCeilCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesCeilCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesCeilCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesCeilCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesCeilCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesCeilCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesCeilCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesCeilCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesCeilCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesCeilCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesCeilCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesCeilCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesCeilCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesCeilCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesCeilCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesCeilCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesCeilCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesCeilCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesCeilCmd)

	flagx.Register(GetTimeSeriesCmoCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesCmoCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesCmoCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesCmoCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesCmoCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesCmoCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesCmoCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesCmoCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesCmoCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesCmoCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesCmoCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesCmoCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesCmoCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesCmoCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesCmoCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesCmoCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesCmoCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesCmoCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesCmoCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesCmoCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesCmoCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesCmoCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesCmoCmd)

	flagx.Register(GetTimeSeriesCoppockCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesCoppockCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesCoppockCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesCoppockCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesCoppockCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesCoppockCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesCoppockCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesCoppockCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesCoppockCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesCoppockCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesCoppockCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesCoppockCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesCoppockCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesCoppockCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesCoppockCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesCoppockCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesCoppockCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesCoppockCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesCoppockCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesCoppockCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesCoppockCmd.Flags().Int64("wma-period", 0, "Number of periods for weighted moving average. Takes values in the range from `1` to `800`")

	GetTimeSeriesCoppockCmd.Flags().Int64("long-roc-period", 0, "Number of periods for long term rate of change. Takes values in the range from `1` to `800`")

	GetTimeSeriesCoppockCmd.Flags().Int64("short-roc-period", 0, "Number of periods for short term rate of change. Takes values in the range from `1` to `800`")

	GetTimeSeriesCoppockCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesCoppockCmd)

	flagx.Register(GetTimeSeriesCorrelCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesCorrelCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesCorrelCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesCorrelCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesCorrelCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesCorrelCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesCorrelCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesCorrelCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesCorrelCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesCorrelCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesCorrelCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesCorrelCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesCorrelCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesCorrelCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesCorrelCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesCorrelCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesCorrelCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesCorrelCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesCorrelCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesCorrelCmd, "series-type-1", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type used as the first part of technical indicator")

	flagx.Register(GetTimeSeriesCorrelCmd, "series-type-2", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type used as the second part of technical indicator")

	GetTimeSeriesCorrelCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesCorrelCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesCorrelCmd)

	flagx.Register(GetTimeSeriesCrsiCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesCrsiCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesCrsiCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesCrsiCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesCrsiCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesCrsiCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesCrsiCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesCrsiCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesCrsiCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesCrsiCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesCrsiCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesCrsiCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesCrsiCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesCrsiCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesCrsiCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesCrsiCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesCrsiCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesCrsiCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesCrsiCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesCrsiCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesCrsiCmd.Flags().Int64("rsi-period", 0, "Number of periods for RSI used to calculate price momentum. Takes values in the range from `1` to `800`")

	GetTimeSeriesCrsiCmd.Flags().Int64("up-down-length", 0, "Number of periods for RSI used to calculate up/down trend. Takes values in the range from `1` to `800`")

	GetTimeSeriesCrsiCmd.Flags().Int64("percent-rank-period", 0, "Number of periods used to calculate PercentRank. Takes values in the range from `1` to `800`")

	GetTimeSeriesCrsiCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesCrsiCmd)

	flagx.Register(GetTimeSeriesDemaCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesDemaCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesDemaCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesDemaCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesDemaCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesDemaCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesDemaCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesDemaCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesDemaCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesDemaCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesDemaCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesDemaCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesDemaCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesDemaCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesDemaCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesDemaCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesDemaCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesDemaCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesDemaCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesDemaCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesDemaCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesDemaCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesDemaCmd)

	flagx.Register(GetTimeSeriesDivCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesDivCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesDivCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesDivCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesDivCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesDivCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesDivCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesDivCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesDivCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesDivCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesDivCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesDivCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesDivCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesDivCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesDivCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesDivCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesDivCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesDivCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesDivCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesDivCmd, "series-type-1", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type used as the first part of technical indicator")

	flagx.Register(GetTimeSeriesDivCmd, "series-type-2", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type used as the second part of technical indicator")

	GetTimeSeriesDivCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesDivCmd)

	flagx.Register(GetTimeSeriesDpoCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesDpoCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesDpoCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesDpoCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesDpoCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesDpoCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesDpoCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesDpoCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesDpoCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesDpoCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesDpoCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesDpoCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesDpoCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesDpoCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesDpoCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesDpoCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesDpoCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesDpoCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesDpoCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesDpoCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesDpoCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesDpoCmd.Flags().Bool("centered", false, "Specifies if there should be a shift to match the current price")

	GetTimeSeriesDpoCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesDpoCmd)

	flagx.Register(GetTimeSeriesDxCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesDxCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesDxCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesDxCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesDxCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesDxCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesDxCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesDxCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesDxCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesDxCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesDxCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesDxCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesDxCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesDxCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesDxCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesDxCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesDxCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesDxCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesDxCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesDxCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesDxCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesDxCmd)

	flagx.Register(GetTimeSeriesEmaCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesEmaCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesEmaCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesEmaCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesEmaCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesEmaCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesEmaCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesEmaCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesEmaCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesEmaCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesEmaCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesEmaCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesEmaCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesEmaCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesEmaCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesEmaCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesEmaCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesEmaCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesEmaCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesEmaCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesEmaCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesEmaCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesEmaCmd)

	flagx.Register(GetTimeSeriesExpCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesExpCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesExpCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesExpCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesExpCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesExpCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesExpCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesExpCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesExpCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesExpCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesExpCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesExpCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesExpCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesExpCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesExpCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesExpCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesExpCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesExpCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesExpCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesExpCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesExpCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesExpCmd)

	flagx.Register(GetTimeSeriesFloorCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesFloorCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesFloorCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesFloorCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesFloorCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesFloorCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesFloorCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesFloorCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesFloorCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesFloorCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesFloorCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesFloorCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesFloorCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesFloorCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesFloorCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesFloorCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesFloorCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesFloorCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesFloorCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesFloorCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesFloorCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesFloorCmd)

	flagx.Register(GetTimeSeriesHeikinashiCandlesCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesHeikinashiCandlesCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesHeikinashiCandlesCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesHeikinashiCandlesCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesHeikinashiCandlesCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesHeikinashiCandlesCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesHeikinashiCandlesCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesHeikinashiCandlesCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesHeikinashiCandlesCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesHeikinashiCandlesCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesHeikinashiCandlesCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesHeikinashiCandlesCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesHeikinashiCandlesCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesHeikinashiCandlesCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesHeikinashiCandlesCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesHeikinashiCandlesCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesHeikinashiCandlesCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesHeikinashiCandlesCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesHeikinashiCandlesCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesHeikinashiCandlesCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesHeikinashiCandlesCmd)

	flagx.Register(GetTimeSeriesHlc3Cmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesHlc3Cmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesHlc3Cmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesHlc3Cmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesHlc3Cmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesHlc3Cmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesHlc3Cmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesHlc3Cmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesHlc3Cmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesHlc3Cmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesHlc3Cmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesHlc3Cmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesHlc3Cmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesHlc3Cmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesHlc3Cmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesHlc3Cmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesHlc3Cmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesHlc3Cmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesHlc3Cmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesHlc3Cmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesHlc3Cmd)

	flagx.Register(GetTimeSeriesHtDcPeriodCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesHtDcPeriodCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesHtDcPeriodCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesHtDcPeriodCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesHtDcPeriodCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesHtDcPeriodCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesHtDcPeriodCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesHtDcPeriodCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesHtDcPeriodCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesHtDcPeriodCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesHtDcPeriodCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesHtDcPeriodCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesHtDcPeriodCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesHtDcPeriodCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesHtDcPeriodCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesHtDcPeriodCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesHtDcPeriodCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesHtDcPeriodCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesHtDcPeriodCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesHtDcPeriodCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesHtDcPeriodCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesHtDcPeriodCmd)

	flagx.Register(GetTimeSeriesHtDcPhaseCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesHtDcPhaseCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesHtDcPhaseCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesHtDcPhaseCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesHtDcPhaseCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesHtDcPhaseCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesHtDcPhaseCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesHtDcPhaseCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesHtDcPhaseCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesHtDcPhaseCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesHtDcPhaseCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesHtDcPhaseCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesHtDcPhaseCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesHtDcPhaseCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesHtDcPhaseCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesHtDcPhaseCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesHtDcPhaseCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesHtDcPhaseCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesHtDcPhaseCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesHtDcPhaseCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesHtDcPhaseCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesHtDcPhaseCmd)

	flagx.Register(GetTimeSeriesHtPhasorCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesHtPhasorCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesHtPhasorCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesHtPhasorCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesHtPhasorCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesHtPhasorCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesHtPhasorCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesHtPhasorCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesHtPhasorCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesHtPhasorCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesHtPhasorCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesHtPhasorCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesHtPhasorCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesHtPhasorCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesHtPhasorCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesHtPhasorCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesHtPhasorCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesHtPhasorCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesHtPhasorCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesHtPhasorCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesHtPhasorCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesHtPhasorCmd)

	flagx.Register(GetTimeSeriesHtSineCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesHtSineCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesHtSineCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesHtSineCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesHtSineCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesHtSineCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesHtSineCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesHtSineCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesHtSineCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesHtSineCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesHtSineCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesHtSineCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesHtSineCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesHtSineCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesHtSineCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesHtSineCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesHtSineCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesHtSineCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesHtSineCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesHtSineCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesHtSineCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesHtSineCmd)

	flagx.Register(GetTimeSeriesHtTrendModeCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesHtTrendModeCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesHtTrendModeCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesHtTrendModeCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesHtTrendModeCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesHtTrendModeCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesHtTrendModeCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesHtTrendModeCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesHtTrendModeCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesHtTrendModeCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesHtTrendModeCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesHtTrendModeCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesHtTrendModeCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesHtTrendModeCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesHtTrendModeCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesHtTrendModeCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesHtTrendModeCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesHtTrendModeCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesHtTrendModeCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesHtTrendModeCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesHtTrendModeCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesHtTrendModeCmd)

	flagx.Register(GetTimeSeriesHtTrendlineCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesHtTrendlineCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesHtTrendlineCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesHtTrendlineCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesHtTrendlineCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesHtTrendlineCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesHtTrendlineCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesHtTrendlineCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesHtTrendlineCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesHtTrendlineCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesHtTrendlineCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesHtTrendlineCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesHtTrendlineCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesHtTrendlineCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesHtTrendlineCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesHtTrendlineCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesHtTrendlineCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesHtTrendlineCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesHtTrendlineCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesHtTrendlineCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesHtTrendlineCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesHtTrendlineCmd)

	flagx.Register(GetTimeSeriesIchimokuCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesIchimokuCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesIchimokuCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesIchimokuCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesIchimokuCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesIchimokuCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesIchimokuCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesIchimokuCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesIchimokuCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesIchimokuCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesIchimokuCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesIchimokuCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesIchimokuCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesIchimokuCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesIchimokuCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesIchimokuCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesIchimokuCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesIchimokuCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesIchimokuCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesIchimokuCmd.Flags().Int64("conversion-line-period", 0, "The time period used for generating the conversation line. Takes values in the range from `1` to `800`")

	GetTimeSeriesIchimokuCmd.Flags().Int64("base-line-period", 0, "The time period used for generating the base line. Takes values in the range from `1` to `800`")

	GetTimeSeriesIchimokuCmd.Flags().Int64("leading-span-b-period", 0, "The time period used for generating the leading span B line. Takes values in the range from `1` to `800`")

	GetTimeSeriesIchimokuCmd.Flags().Int64("lagging-span-period", 0, "The time period used for generating the lagging span line. Takes values in the range from `1` to `800`")

	GetTimeSeriesIchimokuCmd.Flags().Bool("include-ahead-span-period", false, "Indicates whether to include ahead span period")

	GetTimeSeriesIchimokuCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesIchimokuCmd)

	flagx.Register(GetTimeSeriesKamaCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesKamaCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesKamaCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesKamaCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesKamaCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesKamaCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesKamaCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesKamaCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesKamaCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesKamaCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesKamaCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesKamaCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesKamaCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesKamaCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesKamaCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesKamaCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesKamaCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesKamaCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesKamaCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesKamaCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesKamaCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesKamaCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesKamaCmd)

	flagx.Register(GetTimeSeriesKeltnerCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesKeltnerCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesKeltnerCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesKeltnerCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesKeltnerCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesKeltnerCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesKeltnerCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesKeltnerCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesKeltnerCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesKeltnerCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesKeltnerCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesKeltnerCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesKeltnerCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesKeltnerCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesKeltnerCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesKeltnerCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesKeltnerCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesKeltnerCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesKeltnerCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesKeltnerCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesKeltnerCmd.Flags().Int64("atr-time-period", 0, "The time period used for calculating the Average True Range. Takes values in the range from `1` to `800`")

	GetTimeSeriesKeltnerCmd.Flags().Int64("multiplier", 0, "The factor used to adjust the indicator's sensitivity")

	flagx.Register(GetTimeSeriesKeltnerCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	flagx.Register(GetTimeSeriesKeltnerCmd, "ma-type", twelvedata.AllowedMaTypeEnumEnumValues, "The type of moving average used")

	GetTimeSeriesKeltnerCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesKeltnerCmd)

	flagx.Register(GetTimeSeriesKstCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesKstCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesKstCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesKstCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesKstCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesKstCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesKstCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesKstCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesKstCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesKstCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesKstCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesKstCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesKstCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesKstCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesKstCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesKstCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesKstCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesKstCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesKstCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesKstCmd.Flags().Int64("roc-period-1", 0, "The time period for the first Rate of Change calculation.")

	GetTimeSeriesKstCmd.Flags().Int64("roc-period-2", 0, "The time period for the second Rate of Change calculation.")

	GetTimeSeriesKstCmd.Flags().Int64("roc-period-3", 0, "The time period for the third Rate of Change calculation.")

	GetTimeSeriesKstCmd.Flags().Int64("roc-period-4", 0, "The time period for the forth Rate of Change calculation.")

	GetTimeSeriesKstCmd.Flags().Int64("sma-period-1", 0, "The time period for the first Simple Moving Average.")

	GetTimeSeriesKstCmd.Flags().Int64("sma-period-2", 0, "The time period for the second Simple Moving Average.")

	GetTimeSeriesKstCmd.Flags().Int64("sma-period-3", 0, "The time period for the third Simple Moving Average.")

	GetTimeSeriesKstCmd.Flags().Int64("sma-period-4", 0, "The time period for the forth Simple Moving Average.")

	GetTimeSeriesKstCmd.Flags().Int64("signal-period", 0, "The time period used for generating the signal line.")

	GetTimeSeriesKstCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesKstCmd)

	flagx.Register(GetTimeSeriesLinearRegCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesLinearRegCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesLinearRegCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesLinearRegCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesLinearRegCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesLinearRegCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesLinearRegCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesLinearRegCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesLinearRegCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesLinearRegCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesLinearRegCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesLinearRegCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesLinearRegCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesLinearRegCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesLinearRegCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesLinearRegCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesLinearRegCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesLinearRegCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesLinearRegCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesLinearRegCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesLinearRegCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesLinearRegCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesLinearRegCmd)

	flagx.Register(GetTimeSeriesLinearRegAngleCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesLinearRegAngleCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesLinearRegAngleCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesLinearRegAngleCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesLinearRegAngleCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesLinearRegAngleCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesLinearRegAngleCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesLinearRegAngleCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesLinearRegAngleCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesLinearRegAngleCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesLinearRegAngleCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesLinearRegAngleCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesLinearRegAngleCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesLinearRegAngleCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesLinearRegAngleCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesLinearRegAngleCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesLinearRegAngleCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesLinearRegAngleCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesLinearRegAngleCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesLinearRegAngleCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesLinearRegAngleCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesLinearRegAngleCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesLinearRegAngleCmd)

	flagx.Register(GetTimeSeriesLinearRegInterceptCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesLinearRegInterceptCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesLinearRegInterceptCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesLinearRegInterceptCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesLinearRegInterceptCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesLinearRegInterceptCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesLinearRegInterceptCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesLinearRegInterceptCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesLinearRegInterceptCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesLinearRegInterceptCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesLinearRegInterceptCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesLinearRegInterceptCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesLinearRegInterceptCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesLinearRegInterceptCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesLinearRegInterceptCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesLinearRegInterceptCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesLinearRegInterceptCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesLinearRegInterceptCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesLinearRegInterceptCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesLinearRegInterceptCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesLinearRegInterceptCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesLinearRegInterceptCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesLinearRegInterceptCmd)

	flagx.Register(GetTimeSeriesLinearRegSlopeCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesLinearRegSlopeCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesLinearRegSlopeCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesLinearRegSlopeCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesLinearRegSlopeCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesLinearRegSlopeCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesLinearRegSlopeCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesLinearRegSlopeCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesLinearRegSlopeCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesLinearRegSlopeCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesLinearRegSlopeCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesLinearRegSlopeCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesLinearRegSlopeCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesLinearRegSlopeCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesLinearRegSlopeCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesLinearRegSlopeCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesLinearRegSlopeCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesLinearRegSlopeCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesLinearRegSlopeCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesLinearRegSlopeCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesLinearRegSlopeCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesLinearRegSlopeCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesLinearRegSlopeCmd)

	flagx.Register(GetTimeSeriesLnCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesLnCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesLnCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesLnCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesLnCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesLnCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesLnCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesLnCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesLnCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesLnCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesLnCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesLnCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesLnCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesLnCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesLnCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesLnCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesLnCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesLnCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesLnCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesLnCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesLnCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesLnCmd)

	flagx.Register(GetTimeSeriesLog10Cmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesLog10Cmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesLog10Cmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesLog10Cmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesLog10Cmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesLog10Cmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesLog10Cmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesLog10Cmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesLog10Cmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesLog10Cmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesLog10Cmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesLog10Cmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesLog10Cmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesLog10Cmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesLog10Cmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesLog10Cmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesLog10Cmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesLog10Cmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesLog10Cmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesLog10Cmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesLog10Cmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesLog10Cmd)

	flagx.Register(GetTimeSeriesMaCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesMaCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesMaCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesMaCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesMaCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesMaCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesMaCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesMaCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesMaCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesMaCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesMaCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesMaCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesMaCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesMaCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesMaCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesMaCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesMaCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesMaCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesMaCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesMaCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesMaCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	flagx.Register(GetTimeSeriesMaCmd, "ma-type", twelvedata.AllowedMaTypeEnumEnumValues, "The type of moving average used")

	GetTimeSeriesMaCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesMaCmd)

	flagx.Register(GetTimeSeriesMacdCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesMacdCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesMacdCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesMacdCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesMacdCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesMacdCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesMacdCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesMacdCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesMacdCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesMacdCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesMacdCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesMacdCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesMacdCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesMacdCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesMacdCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesMacdCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesMacdCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesMacdCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesMacdCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesMacdCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesMacdCmd.Flags().Int64("fast-period", 0, "Number of periods for fast moving average. Takes values in the range from `1` to `800`")

	GetTimeSeriesMacdCmd.Flags().Int64("slow-period", 0, "Number of periods for slow moving average. Takes values in the range from `1` to `800`")

	GetTimeSeriesMacdCmd.Flags().Int64("signal-period", 0, "The time period used for generating the signal line.")

	GetTimeSeriesMacdCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesMacdCmd)

	flagx.Register(GetTimeSeriesMacdExtCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesMacdExtCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesMacdExtCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesMacdExtCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesMacdExtCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesMacdExtCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesMacdExtCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesMacdExtCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesMacdExtCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesMacdExtCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesMacdExtCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesMacdExtCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesMacdExtCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesMacdExtCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesMacdExtCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesMacdExtCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesMacdExtCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesMacdExtCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesMacdExtCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesMacdExtCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesMacdExtCmd.Flags().Int64("fast-period", 0, "Number of periods for fast moving average. Takes values in the range from `1` to `800`")

	flagx.Register(GetTimeSeriesMacdExtCmd, "fast-ma-type", twelvedata.AllowedMaTypeEnumEnumValues, "The type of fast moving average used in the calculation.")

	GetTimeSeriesMacdExtCmd.Flags().Int64("slow-period", 0, "Number of periods for slow moving average. Takes values in the range from `1` to `800`")

	flagx.Register(GetTimeSeriesMacdExtCmd, "slow-ma-type", twelvedata.AllowedMaTypeEnumEnumValues, "The type of slow moving average used in the calculation.")

	GetTimeSeriesMacdExtCmd.Flags().Int64("signal-period", 0, "The time period used for generating the signal line.")

	flagx.Register(GetTimeSeriesMacdExtCmd, "signal-ma-type", twelvedata.AllowedMaTypeEnumEnumValues, "The type of fast moving average used for generating the signal line.")

	GetTimeSeriesMacdExtCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesMacdExtCmd)

	flagx.Register(GetTimeSeriesMacdSlopeCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesMacdSlopeCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesMacdSlopeCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesMacdSlopeCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesMacdSlopeCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesMacdSlopeCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesMacdSlopeCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesMacdSlopeCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesMacdSlopeCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesMacdSlopeCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesMacdSlopeCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesMacdSlopeCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesMacdSlopeCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesMacdSlopeCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesMacdSlopeCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesMacdSlopeCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesMacdSlopeCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesMacdSlopeCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesMacdSlopeCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesMacdSlopeCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesMacdSlopeCmd.Flags().Int64("fast-period", 0, "Number of periods for fast moving average. Takes values in the range from `1` to `800`")

	GetTimeSeriesMacdSlopeCmd.Flags().Int64("slow-period", 0, "Number of periods for slow moving average. Takes values in the range from `1` to `800`")

	GetTimeSeriesMacdSlopeCmd.Flags().Int64("signal-period", 0, "The time period used for generating the signal line.")

	GetTimeSeriesMacdSlopeCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesMacdSlopeCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesMacdSlopeCmd)

	flagx.Register(GetTimeSeriesMamaCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesMamaCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesMamaCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesMamaCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesMamaCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesMamaCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesMamaCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesMamaCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesMamaCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesMamaCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesMamaCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesMamaCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesMamaCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesMamaCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesMamaCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesMamaCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesMamaCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesMamaCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesMamaCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesMamaCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesMamaCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesMamaCmd)

	flagx.Register(GetTimeSeriesMaxCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesMaxCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesMaxCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesMaxCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesMaxCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesMaxCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesMaxCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesMaxCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesMaxCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesMaxCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesMaxCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesMaxCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesMaxCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesMaxCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesMaxCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesMaxCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesMaxCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesMaxCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesMaxCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesMaxCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesMaxCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesMaxCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesMaxCmd)

	flagx.Register(GetTimeSeriesMaxIndexCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesMaxIndexCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesMaxIndexCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesMaxIndexCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesMaxIndexCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesMaxIndexCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesMaxIndexCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesMaxIndexCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesMaxIndexCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesMaxIndexCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesMaxIndexCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesMaxIndexCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesMaxIndexCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesMaxIndexCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesMaxIndexCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesMaxIndexCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesMaxIndexCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesMaxIndexCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesMaxIndexCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesMaxIndexCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesMaxIndexCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesMaxIndexCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesMaxIndexCmd)

	flagx.Register(GetTimeSeriesMcGinleyDynamicCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesMcGinleyDynamicCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesMcGinleyDynamicCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesMcGinleyDynamicCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesMcGinleyDynamicCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesMcGinleyDynamicCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesMcGinleyDynamicCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesMcGinleyDynamicCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesMcGinleyDynamicCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesMcGinleyDynamicCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesMcGinleyDynamicCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesMcGinleyDynamicCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesMcGinleyDynamicCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesMcGinleyDynamicCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesMcGinleyDynamicCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesMcGinleyDynamicCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesMcGinleyDynamicCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesMcGinleyDynamicCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesMcGinleyDynamicCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesMcGinleyDynamicCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesMcGinleyDynamicCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesMcGinleyDynamicCmd)

	flagx.Register(GetTimeSeriesMedPriceCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesMedPriceCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesMedPriceCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesMedPriceCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesMedPriceCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesMedPriceCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesMedPriceCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesMedPriceCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesMedPriceCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesMedPriceCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesMedPriceCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesMedPriceCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesMedPriceCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesMedPriceCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesMedPriceCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesMedPriceCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesMedPriceCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesMedPriceCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesMedPriceCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesMedPriceCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesMedPriceCmd)

	flagx.Register(GetTimeSeriesMfiCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesMfiCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesMfiCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesMfiCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesMfiCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesMfiCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesMfiCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesMfiCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesMfiCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesMfiCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesMfiCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesMfiCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesMfiCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesMfiCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesMfiCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesMfiCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesMfiCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesMfiCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesMfiCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesMfiCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesMfiCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesMfiCmd)

	flagx.Register(GetTimeSeriesMidPointCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesMidPointCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesMidPointCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesMidPointCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesMidPointCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesMidPointCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesMidPointCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesMidPointCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesMidPointCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesMidPointCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesMidPointCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesMidPointCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesMidPointCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesMidPointCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesMidPointCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesMidPointCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesMidPointCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesMidPointCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesMidPointCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesMidPointCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesMidPointCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesMidPointCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesMidPointCmd)

	flagx.Register(GetTimeSeriesMidPriceCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesMidPriceCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesMidPriceCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesMidPriceCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesMidPriceCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesMidPriceCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesMidPriceCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesMidPriceCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesMidPriceCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesMidPriceCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesMidPriceCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesMidPriceCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesMidPriceCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesMidPriceCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesMidPriceCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesMidPriceCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesMidPriceCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesMidPriceCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesMidPriceCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesMidPriceCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesMidPriceCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesMidPriceCmd)

	flagx.Register(GetTimeSeriesMinCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesMinCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesMinCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesMinCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesMinCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesMinCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesMinCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesMinCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesMinCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesMinCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesMinCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesMinCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesMinCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesMinCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesMinCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesMinCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesMinCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesMinCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesMinCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesMinCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesMinCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesMinCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesMinCmd)

	flagx.Register(GetTimeSeriesMinIndexCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesMinIndexCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesMinIndexCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesMinIndexCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesMinIndexCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesMinIndexCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesMinIndexCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesMinIndexCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesMinIndexCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesMinIndexCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesMinIndexCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesMinIndexCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesMinIndexCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesMinIndexCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesMinIndexCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesMinIndexCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesMinIndexCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesMinIndexCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesMinIndexCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesMinIndexCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesMinIndexCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesMinIndexCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesMinIndexCmd)

	flagx.Register(GetTimeSeriesMinMaxCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesMinMaxCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesMinMaxCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesMinMaxCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesMinMaxCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesMinMaxCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesMinMaxCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesMinMaxCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesMinMaxCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesMinMaxCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesMinMaxCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesMinMaxCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesMinMaxCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesMinMaxCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesMinMaxCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesMinMaxCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesMinMaxCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesMinMaxCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesMinMaxCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesMinMaxCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesMinMaxCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesMinMaxCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesMinMaxCmd)

	flagx.Register(GetTimeSeriesMinMaxIndexCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesMinMaxIndexCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesMinMaxIndexCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesMinMaxIndexCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesMinMaxIndexCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesMinMaxIndexCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesMinMaxIndexCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesMinMaxIndexCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesMinMaxIndexCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesMinMaxIndexCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesMinMaxIndexCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesMinMaxIndexCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesMinMaxIndexCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesMinMaxIndexCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesMinMaxIndexCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesMinMaxIndexCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesMinMaxIndexCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesMinMaxIndexCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesMinMaxIndexCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesMinMaxIndexCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesMinMaxIndexCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesMinMaxIndexCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesMinMaxIndexCmd)

	flagx.Register(GetTimeSeriesMinusDICmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesMinusDICmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesMinusDICmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesMinusDICmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesMinusDICmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesMinusDICmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesMinusDICmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesMinusDICmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesMinusDICmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesMinusDICmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesMinusDICmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesMinusDICmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesMinusDICmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesMinusDICmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesMinusDICmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesMinusDICmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesMinusDICmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesMinusDICmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesMinusDICmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesMinusDICmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesMinusDICmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesMinusDICmd)

	flagx.Register(GetTimeSeriesMinusDMCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesMinusDMCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesMinusDMCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesMinusDMCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesMinusDMCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesMinusDMCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesMinusDMCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesMinusDMCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesMinusDMCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesMinusDMCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesMinusDMCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesMinusDMCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesMinusDMCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesMinusDMCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesMinusDMCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesMinusDMCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesMinusDMCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesMinusDMCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesMinusDMCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesMinusDMCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesMinusDMCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesMinusDMCmd)

	flagx.Register(GetTimeSeriesMomCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesMomCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesMomCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesMomCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesMomCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesMomCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesMomCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesMomCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesMomCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesMomCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesMomCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesMomCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesMomCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesMomCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesMomCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesMomCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesMomCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesMomCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesMomCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesMomCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesMomCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesMomCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesMomCmd)

	flagx.Register(GetTimeSeriesMultCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesMultCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesMultCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesMultCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesMultCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesMultCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesMultCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesMultCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesMultCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesMultCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesMultCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesMultCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesMultCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesMultCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesMultCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesMultCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesMultCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesMultCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesMultCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesMultCmd, "series-type-1", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type used as the first part of technical indicator")

	flagx.Register(GetTimeSeriesMultCmd, "series-type-2", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type used as the second part of technical indicator")

	GetTimeSeriesMultCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesMultCmd)

	flagx.Register(GetTimeSeriesNatrCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesNatrCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesNatrCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesNatrCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesNatrCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesNatrCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesNatrCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesNatrCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesNatrCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesNatrCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesNatrCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesNatrCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesNatrCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesNatrCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesNatrCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesNatrCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesNatrCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesNatrCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesNatrCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesNatrCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesNatrCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesNatrCmd)

	flagx.Register(GetTimeSeriesObvCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesObvCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesObvCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesObvCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesObvCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesObvCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesObvCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesObvCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesObvCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesObvCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesObvCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesObvCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesObvCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesObvCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesObvCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesObvCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesObvCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesObvCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesObvCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesObvCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesObvCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesObvCmd)

	flagx.Register(GetTimeSeriesPercentBCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesPercentBCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesPercentBCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesPercentBCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesPercentBCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesPercentBCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesPercentBCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesPercentBCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesPercentBCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesPercentBCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesPercentBCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesPercentBCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesPercentBCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesPercentBCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesPercentBCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesPercentBCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesPercentBCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesPercentBCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesPercentBCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesPercentBCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesPercentBCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	flagx.Register(GetTimeSeriesPercentBCmd, "ma-type", twelvedata.AllowedMaTypeEnumEnumValues, "The type of moving average used")

	GetTimeSeriesPercentBCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesPercentBCmd)

	flagx.Register(GetTimeSeriesPivotPointsHLCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesPivotPointsHLCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesPivotPointsHLCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesPivotPointsHLCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesPivotPointsHLCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesPivotPointsHLCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesPivotPointsHLCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesPivotPointsHLCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesPivotPointsHLCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesPivotPointsHLCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesPivotPointsHLCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesPivotPointsHLCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesPivotPointsHLCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesPivotPointsHLCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesPivotPointsHLCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesPivotPointsHLCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesPivotPointsHLCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesPivotPointsHLCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesPivotPointsHLCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesPivotPointsHLCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesPivotPointsHLCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesPivotPointsHLCmd)

	flagx.Register(GetTimeSeriesPlusDICmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesPlusDICmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesPlusDICmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesPlusDICmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesPlusDICmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesPlusDICmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesPlusDICmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesPlusDICmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesPlusDICmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesPlusDICmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesPlusDICmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesPlusDICmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesPlusDICmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesPlusDICmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesPlusDICmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesPlusDICmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesPlusDICmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesPlusDICmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesPlusDICmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesPlusDICmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesPlusDICmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesPlusDICmd)

	flagx.Register(GetTimeSeriesPlusDMCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesPlusDMCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesPlusDMCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesPlusDMCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesPlusDMCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesPlusDMCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesPlusDMCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesPlusDMCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesPlusDMCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesPlusDMCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesPlusDMCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesPlusDMCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesPlusDMCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesPlusDMCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesPlusDMCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesPlusDMCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesPlusDMCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesPlusDMCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesPlusDMCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesPlusDMCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesPlusDMCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesPlusDMCmd)

	flagx.Register(GetTimeSeriesPpoCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesPpoCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesPpoCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesPpoCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesPpoCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesPpoCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesPpoCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesPpoCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesPpoCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesPpoCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesPpoCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesPpoCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesPpoCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesPpoCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesPpoCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesPpoCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesPpoCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesPpoCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesPpoCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesPpoCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesPpoCmd.Flags().Int64("fast-period", 0, "Number of periods for fast moving average. Takes values in the range from `1` to `800`")

	GetTimeSeriesPpoCmd.Flags().Int64("slow-period", 0, "Number of periods for slow moving average. Takes values in the range from `1` to `800`")

	flagx.Register(GetTimeSeriesPpoCmd, "ma-type", twelvedata.AllowedMaTypeEnumEnumValues, "The type of moving average used")

	GetTimeSeriesPpoCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesPpoCmd)

	flagx.Register(GetTimeSeriesRocCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesRocCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesRocCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesRocCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesRocCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesRocCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesRocCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesRocCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesRocCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesRocCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesRocCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesRocCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesRocCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesRocCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesRocCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesRocCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesRocCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesRocCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesRocCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesRocCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesRocCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesRocCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesRocCmd)

	flagx.Register(GetTimeSeriesRocpCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesRocpCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesRocpCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesRocpCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesRocpCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesRocpCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesRocpCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesRocpCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesRocpCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesRocpCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesRocpCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesRocpCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesRocpCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesRocpCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesRocpCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesRocpCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesRocpCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesRocpCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesRocpCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesRocpCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesRocpCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesRocpCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesRocpCmd)

	flagx.Register(GetTimeSeriesRocrCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesRocrCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesRocrCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesRocrCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesRocrCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesRocrCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesRocrCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesRocrCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesRocrCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesRocrCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesRocrCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesRocrCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesRocrCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesRocrCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesRocrCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesRocrCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesRocrCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesRocrCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesRocrCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesRocrCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesRocrCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesRocrCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesRocrCmd)

	flagx.Register(GetTimeSeriesRocr100Cmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesRocr100Cmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesRocr100Cmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesRocr100Cmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesRocr100Cmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesRocr100Cmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesRocr100Cmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesRocr100Cmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesRocr100Cmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesRocr100Cmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesRocr100Cmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesRocr100Cmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesRocr100Cmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesRocr100Cmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesRocr100Cmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesRocr100Cmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesRocr100Cmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesRocr100Cmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesRocr100Cmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesRocr100Cmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesRocr100Cmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesRocr100Cmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesRocr100Cmd)

	flagx.Register(GetTimeSeriesRsiCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesRsiCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesRsiCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesRsiCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesRsiCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesRsiCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesRsiCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesRsiCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesRsiCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesRsiCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesRsiCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesRsiCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesRsiCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesRsiCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesRsiCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesRsiCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesRsiCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesRsiCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesRsiCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesRsiCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesRsiCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesRsiCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesRsiCmd)

	flagx.Register(GetTimeSeriesRvolCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesRvolCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesRvolCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesRvolCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesRvolCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesRvolCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesRvolCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesRvolCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesRvolCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesRvolCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesRvolCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesRvolCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesRvolCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesRvolCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesRvolCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesRvolCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesRvolCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesRvolCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesRvolCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesRvolCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesRvolCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesRvolCmd)

	flagx.Register(GetTimeSeriesSarCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesSarCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesSarCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesSarCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesSarCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesSarCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesSarCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesSarCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesSarCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesSarCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesSarCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesSarCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesSarCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesSarCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesSarCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesSarCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesSarCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesSarCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesSarCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesSarCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesSarCmd)

	flagx.Register(GetTimeSeriesSarExtCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesSarExtCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesSarExtCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesSarExtCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesSarExtCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesSarExtCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesSarExtCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesSarExtCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesSarExtCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesSarExtCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesSarExtCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesSarExtCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesSarExtCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesSarExtCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesSarExtCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesSarExtCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesSarExtCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesSarExtCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesSarExtCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesSarExtCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesSarExtCmd)

	flagx.Register(GetTimeSeriesSmaCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesSmaCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesSmaCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesSmaCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesSmaCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesSmaCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesSmaCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesSmaCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesSmaCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesSmaCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesSmaCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesSmaCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesSmaCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesSmaCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesSmaCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesSmaCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesSmaCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesSmaCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesSmaCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesSmaCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesSmaCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesSmaCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesSmaCmd)

	flagx.Register(GetTimeSeriesSqrtCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesSqrtCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesSqrtCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesSqrtCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesSqrtCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesSqrtCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesSqrtCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesSqrtCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesSqrtCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesSqrtCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesSqrtCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesSqrtCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesSqrtCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesSqrtCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesSqrtCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesSqrtCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesSqrtCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesSqrtCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesSqrtCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesSqrtCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesSqrtCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesSqrtCmd)

	flagx.Register(GetTimeSeriesStdDevCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesStdDevCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesStdDevCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesStdDevCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesStdDevCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesStdDevCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesStdDevCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesStdDevCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesStdDevCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesStdDevCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesStdDevCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesStdDevCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesStdDevCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesStdDevCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesStdDevCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesStdDevCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesStdDevCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesStdDevCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesStdDevCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesStdDevCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesStdDevCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesStdDevCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesStdDevCmd)

	flagx.Register(GetTimeSeriesStochCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesStochCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesStochCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesStochCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesStochCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesStochCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesStochCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesStochCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesStochCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesStochCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesStochCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesStochCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesStochCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesStochCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesStochCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesStochCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesStochCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesStochCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesStochCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesStochCmd.Flags().Int64("fast-k-period", 0, "The time period for the fast %K line in the Stochastic Oscillator. Takes values in the range from `1` to `800`")

	GetTimeSeriesStochCmd.Flags().Int64("slow-k-period", 0, "The time period for the slow %K line in the Stochastic Oscillator. Takes values in the range from `1` to `800`")

	GetTimeSeriesStochCmd.Flags().Int64("slow-d-period", 0, "The time period for the slow %D line in the Stochastic Oscillator. Takes values in the range from `1` to `800`")

	flagx.Register(GetTimeSeriesStochCmd, "slow-kma-type", twelvedata.AllowedMaTypeEnumEnumValues, "The type of slow %K Moving Average used. Default is SMA.")

	flagx.Register(GetTimeSeriesStochCmd, "slow-dma-type", twelvedata.AllowedMaTypeEnumEnumValues, "The type of slow Displaced Moving Average used. Default is SMA.")

	GetTimeSeriesStochCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesStochCmd)

	flagx.Register(GetTimeSeriesStochFCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesStochFCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesStochFCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesStochFCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesStochFCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesStochFCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesStochFCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesStochFCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesStochFCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesStochFCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesStochFCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesStochFCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesStochFCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesStochFCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesStochFCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesStochFCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesStochFCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesStochFCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesStochFCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesStochFCmd.Flags().Int64("fast-k-period", 0, "The time period for the fast %K line in the Stochastic Oscillator. Takes values in the range from `1` to `800`")

	GetTimeSeriesStochFCmd.Flags().Int64("fast-d-period", 0, "The time period for the fast %D line in the Stochastic Oscillator. Takes values in the range from `1` to `800`")

	flagx.Register(GetTimeSeriesStochFCmd, "fast-dma-type", twelvedata.AllowedMaTypeEnumEnumValues, "The type of fast Displaced Moving Average used.")

	GetTimeSeriesStochFCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesStochFCmd)

	flagx.Register(GetTimeSeriesStochRsiCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesStochRsiCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesStochRsiCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesStochRsiCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesStochRsiCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesStochRsiCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesStochRsiCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesStochRsiCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesStochRsiCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesStochRsiCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesStochRsiCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesStochRsiCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesStochRsiCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesStochRsiCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesStochRsiCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesStochRsiCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesStochRsiCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesStochRsiCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesStochRsiCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesStochRsiCmd, "series-type", twelvedata.AllowedSeriesTypeStochrsiEnumEnumValues, "Specifies the price data type: open, high, low, or close.")

	GetTimeSeriesStochRsiCmd.Flags().Int64("rsi-length", 0, "Length of period for calculating the RSI component. Takes values in the range from `1` to `800`")

	GetTimeSeriesStochRsiCmd.Flags().Int64("stoch-length", 0, "Period length for computing the stochastic oscillator of the RSI. Takes values in the range from `1` to `800`")

	GetTimeSeriesStochRsiCmd.Flags().Int64("k-period", 0, "Period for smoothing the %K line. Takes values in the range from `1` to `800`")

	GetTimeSeriesStochRsiCmd.Flags().String("slow-kma-type", "", "")

	GetTimeSeriesStochRsiCmd.Flags().Int64("d-period", 0, "Period for smoothing the %D line, which is a moving average of %K. Takes values in the range from `1` to `800`")

	GetTimeSeriesStochRsiCmd.Flags().String("slow-dma-type", "", "")

	GetTimeSeriesStochRsiCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesStochRsiCmd)

	flagx.Register(GetTimeSeriesSubCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesSubCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesSubCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesSubCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesSubCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesSubCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesSubCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesSubCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesSubCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesSubCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesSubCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesSubCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesSubCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesSubCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesSubCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesSubCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesSubCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesSubCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesSubCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesSubCmd, "series-type-1", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type used as the first part of technical indicator")

	flagx.Register(GetTimeSeriesSubCmd, "series-type-2", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type used as the second part of technical indicator")

	GetTimeSeriesSubCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesSubCmd)

	flagx.Register(GetTimeSeriesSumCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesSumCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesSumCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesSumCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesSumCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesSumCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesSumCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesSumCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesSumCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesSumCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesSumCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesSumCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesSumCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesSumCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesSumCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesSumCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesSumCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesSumCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesSumCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesSumCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesSumCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesSumCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesSumCmd)

	flagx.Register(GetTimeSeriesSuperTrendCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesSuperTrendCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesSuperTrendCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesSuperTrendCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesSuperTrendCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesSuperTrendCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesSuperTrendCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesSuperTrendCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesSuperTrendCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesSuperTrendCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesSuperTrendCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesSuperTrendCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesSuperTrendCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesSuperTrendCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesSuperTrendCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesSuperTrendCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesSuperTrendCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesSuperTrendCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesSuperTrendCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesSuperTrendCmd.Flags().Int64("period", 0, "The period used for calculation in the indicator. Takes values in the range from `1` to `800`")

	GetTimeSeriesSuperTrendCmd.Flags().Int64("multiplier", 0, "The factor used to adjust the indicator's sensitivity.")

	GetTimeSeriesSuperTrendCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesSuperTrendCmd)

	flagx.Register(GetTimeSeriesSuperTrendHeikinAshiCandlesCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesSuperTrendHeikinAshiCandlesCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesSuperTrendHeikinAshiCandlesCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesSuperTrendHeikinAshiCandlesCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesSuperTrendHeikinAshiCandlesCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesSuperTrendHeikinAshiCandlesCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesSuperTrendHeikinAshiCandlesCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesSuperTrendHeikinAshiCandlesCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesSuperTrendHeikinAshiCandlesCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesSuperTrendHeikinAshiCandlesCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesSuperTrendHeikinAshiCandlesCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesSuperTrendHeikinAshiCandlesCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesSuperTrendHeikinAshiCandlesCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesSuperTrendHeikinAshiCandlesCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesSuperTrendHeikinAshiCandlesCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesSuperTrendHeikinAshiCandlesCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesSuperTrendHeikinAshiCandlesCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesSuperTrendHeikinAshiCandlesCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesSuperTrendHeikinAshiCandlesCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesSuperTrendHeikinAshiCandlesCmd.Flags().Int64("period", 0, "The period used for calculation in the indicator. Takes values in the range from `1` to `800`")

	GetTimeSeriesSuperTrendHeikinAshiCandlesCmd.Flags().Int64("multiplier", 0, "The factor used to adjust the indicator's sensitivity.")

	GetTimeSeriesSuperTrendHeikinAshiCandlesCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesSuperTrendHeikinAshiCandlesCmd)

	flagx.Register(GetTimeSeriesT3maCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesT3maCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesT3maCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesT3maCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesT3maCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesT3maCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesT3maCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesT3maCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesT3maCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesT3maCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesT3maCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesT3maCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesT3maCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesT3maCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesT3maCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesT3maCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesT3maCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesT3maCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesT3maCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesT3maCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesT3maCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesT3maCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesT3maCmd)

	flagx.Register(GetTimeSeriesTRangeCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesTRangeCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesTRangeCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesTRangeCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesTRangeCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesTRangeCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesTRangeCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesTRangeCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesTRangeCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesTRangeCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesTRangeCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesTRangeCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesTRangeCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesTRangeCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesTRangeCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesTRangeCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesTRangeCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesTRangeCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesTRangeCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesTRangeCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesTRangeCmd)

	flagx.Register(GetTimeSeriesTemaCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesTemaCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesTemaCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesTemaCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesTemaCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesTemaCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesTemaCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesTemaCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesTemaCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesTemaCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesTemaCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesTemaCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesTemaCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesTemaCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesTemaCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesTemaCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesTemaCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesTemaCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesTemaCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesTemaCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesTemaCmd.Flags().Int64("time-period", 0, "The time period used for calculation in the indicator. Default is 9.")

	GetTimeSeriesTemaCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesTemaCmd)

	flagx.Register(GetTimeSeriesTrimaCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesTrimaCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesTrimaCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesTrimaCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesTrimaCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesTrimaCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesTrimaCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesTrimaCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesTrimaCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesTrimaCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesTrimaCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesTrimaCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesTrimaCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesTrimaCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesTrimaCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesTrimaCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesTrimaCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesTrimaCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesTrimaCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesTrimaCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesTrimaCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesTrimaCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesTrimaCmd)

	flagx.Register(GetTimeSeriesTsfCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesTsfCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesTsfCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesTsfCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesTsfCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesTsfCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesTsfCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesTsfCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesTsfCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesTsfCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesTsfCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesTsfCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesTsfCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesTsfCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesTsfCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesTsfCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesTsfCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesTsfCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesTsfCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesTsfCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesTsfCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesTsfCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesTsfCmd)

	flagx.Register(GetTimeSeriesTypPriceCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesTypPriceCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesTypPriceCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesTypPriceCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesTypPriceCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesTypPriceCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesTypPriceCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesTypPriceCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesTypPriceCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesTypPriceCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesTypPriceCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesTypPriceCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesTypPriceCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesTypPriceCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesTypPriceCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesTypPriceCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesTypPriceCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesTypPriceCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesTypPriceCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesTypPriceCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesTypPriceCmd)

	flagx.Register(GetTimeSeriesUltOscCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesUltOscCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesUltOscCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesUltOscCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesUltOscCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesUltOscCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesUltOscCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesUltOscCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesUltOscCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesUltOscCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesUltOscCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesUltOscCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesUltOscCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesUltOscCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesUltOscCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesUltOscCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesUltOscCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesUltOscCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesUltOscCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesUltOscCmd.Flags().Int64("time-period-1", 0, "The first time period used for calculation in the indicator. Takes values in the range from `1` to `800`")

	GetTimeSeriesUltOscCmd.Flags().Int64("time-period-2", 0, "The second time period used for calculation in the indicator. Takes values in the range from `1` to `800`")

	GetTimeSeriesUltOscCmd.Flags().Int64("time-period-3", 0, "The third time period used for calculation in the indicator. Takes values in the range from `1` to `800`")

	GetTimeSeriesUltOscCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesUltOscCmd)

	flagx.Register(GetTimeSeriesVarCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesVarCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesVarCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesVarCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesVarCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesVarCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesVarCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesVarCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesVarCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesVarCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesVarCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesVarCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesVarCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesVarCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesVarCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesVarCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesVarCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesVarCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesVarCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesVarCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesVarCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesVarCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesVarCmd)

	flagx.Register(GetTimeSeriesVwapCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesVwapCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesVwapCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesVwapCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesVwapCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesVwapCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesVwapCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesVwapCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesVwapCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesVwapCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesVwapCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesVwapCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesVwapCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesVwapCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesVwapCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesVwapCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesVwapCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesVwapCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesVwapCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesVwapCmd.Flags().Int64("sd-time-period", 0, "The time period for the standard deviation calculation. Must be greater than `0`. Recommended value is `9`. This parameter is only used together with `sd`.")

	GetTimeSeriesVwapCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesVwapCmd)

	flagx.Register(GetTimeSeriesWclPriceCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesWclPriceCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesWclPriceCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesWclPriceCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesWclPriceCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesWclPriceCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesWclPriceCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesWclPriceCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesWclPriceCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesWclPriceCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesWclPriceCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesWclPriceCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesWclPriceCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesWclPriceCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesWclPriceCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesWclPriceCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesWclPriceCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesWclPriceCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesWclPriceCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesWclPriceCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesWclPriceCmd)

	flagx.Register(GetTimeSeriesWillRCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesWillRCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesWillRCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesWillRCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesWillRCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesWillRCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesWillRCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesWillRCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesWillRCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesWillRCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesWillRCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesWillRCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesWillRCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesWillRCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesWillRCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesWillRCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesWillRCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesWillRCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesWillRCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	GetTimeSeriesWillRCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesWillRCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesWillRCmd)

	flagx.Register(GetTimeSeriesWmaCmd, "interval", twelvedata.AllowedIntervalEnumEnumValues, "Interval between two consecutive points in time series")

	GetTimeSeriesWmaCmd.Flags().String("symbol", "", "Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ...")

	GetTimeSeriesWmaCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTimeSeriesWmaCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTimeSeriesWmaCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTimeSeriesWmaCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum")

	GetTimeSeriesWmaCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetTimeSeriesWmaCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetTimeSeriesWmaCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetTimeSeriesWmaCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	GetTimeSeriesWmaCmd.Flags().String("timezone", "", "Timezone at which output datetime will be displayed. Supports: 1. Exchange for local exchange time 2. UTC for datetime at universal UTC standard 3. Timezone name according to the IANA Time Zone Database. E.g. America/New_York, Asia/Singapore. Full list of timezones can be found here Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time. Take note that the IANA Timezone name is case-sensitive")

	GetTimeSeriesWmaCmd.Flags().String("start-date", "", "Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05` Default location: Forex and Cryptocurrencies - UTC Stocks - where exchange is located (e.g. for AAPL it will be America/New_York) Both parameters take into account if timezone parameter is provided. If timezone is given then, start_date and end_date will be used in the specified location Examples: 1. &symbol=AAPL&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 New York time up to current date 2. &symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&… Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date 3. &symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&... Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00")

	GetTimeSeriesWmaCmd.Flags().String("end-date", "", "The ending date and time for data selection, see `start_date` description for details.")

	GetTimeSeriesWmaCmd.Flags().String("date", "", "Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday`")

	flagx.Register(GetTimeSeriesWmaCmd, "order", twelvedata.AllowedOrderEnumEnumValues, "Sorting order of the output")

	GetTimeSeriesWmaCmd.Flags().Bool("prepost", false, "Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume")

	GetTimeSeriesWmaCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided")

	GetTimeSeriesWmaCmd.Flags().Bool("previous-close", false, "A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object")

	flagx.Register(GetTimeSeriesWmaCmd, "adjust", twelvedata.AllowedAdjustEnumEnumValues, "Adjusting mode for prices")

	flagx.Register(GetTimeSeriesWmaCmd, "series-type", twelvedata.AllowedSeriesTypeEnumEnumValues, "Price type on which technical indicator is calculated")

	GetTimeSeriesWmaCmd.Flags().Int64("time-period", 0, "Number of periods to average over. Takes values in the range from `1` to `800`")

	GetTimeSeriesWmaCmd.Flags().Bool("include-ohlc", false, "Specify if OHLC values should be added in the output")

	tiCmd.AddCommand(GetTimeSeriesWmaCmd)
}
