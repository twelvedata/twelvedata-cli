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

var GetAnalystRatingsLightCmd = &cobra.Command{
	Use:     "analyst-ratings-light",
	Short:   "Analyst ratings snapshot",
	Long:    "The analyst ratings snapshot endpoint provides a streamlined summary of ratings from analyst firms for both US and international markets. It delivers essential data on analyst recommendations, including buy, hold, and sell ratings, allowing users to quickly assess the general sentiment of analysts towards a particular stock.",
	Example: "td analyst-ratings-light --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.AnalysisAPI.GetAnalystRatingsLight(cmd.Context())

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

		if cmd.Flags().Changed("rating-change") {
			req = req.RatingChange(twelvedata.RatingChangeEnum(cmd.Flags().Lookup("rating-change").Value.String()))
		}

		if cmd.Flags().Changed("outputsize") {
			v, _ := cmd.Flags().GetInt64("outputsize")
			req = req.Outputsize(v)
		}

		if v, _ := cmd.Flags().GetString("country"); v != "" {
			req = req.Country(v)
		}

		sp := output.StartSpinner(cmd)
		resp, httpResp, callErr := req.Execute()
		sp.Stop()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetAnalystRatingsUsEquitiesCmd = &cobra.Command{
	Use:     "analyst-ratings-us-equities",
	Short:   "Analyst ratings US equities",
	Long:    "The analyst ratings US equities endpoint provides detailed information on analyst ratings for U.S. stocks. It returns data on the latest ratings issued by various analyst firms, including the rating itself, the firm issuing the rating, and any changes in the rating. This endpoint is useful for users tracking analyst opinions on U.S. equities, allowing them to see how professional analysts view the potential performance of specific stocks.",
	Example: "td analyst-ratings-us-equities --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.AnalysisAPI.GetAnalystRatingsUsEquities(cmd.Context())

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

		if cmd.Flags().Changed("rating-change") {
			req = req.RatingChange(twelvedata.RatingChangeEnum(cmd.Flags().Lookup("rating-change").Value.String()))
		}

		if cmd.Flags().Changed("outputsize") {
			v, _ := cmd.Flags().GetInt64("outputsize")
			req = req.Outputsize(v)
		}

		sp := output.StartSpinner(cmd)
		resp, httpResp, callErr := req.Execute()
		sp.Stop()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetEarningsEstimateCmd = &cobra.Command{
	Use:     "earnings-estimate",
	Short:   "Earnings estimate",
	Long:    "The earnings estimate endpoint provides access to analysts' projected earnings per share (EPS) for a specific company, covering both upcoming quarterly and annual periods. This data is crucial for users who need to track and compare expected financial performance across different timeframes, aiding in the evaluation of a company's future profitability.",
	Example: "td earnings-estimate --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.AnalysisAPI.GetEarningsEstimate(cmd.Context())

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

		if v, _ := cmd.Flags().GetString("country"); v != "" {
			req = req.Country(v)
		}

		if v, _ := cmd.Flags().GetString("exchange"); v != "" {
			req = req.Exchange(v)
		}

		sp := output.StartSpinner(cmd)
		resp, httpResp, callErr := req.Execute()
		sp.Stop()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetEpsRevisionsCmd = &cobra.Command{
	Use:     "eps-revisions",
	Short:   "EPS revisions",
	Long:    "The EPS revisions endpoint provides updated analyst forecasts for a company's earnings per share (EPS) on both a quarterly and annual basis. It delivers data on how these EPS predictions have changed over the past week and month, allowing users to track recent adjustments in analyst expectations. This endpoint is useful for monitoring shifts in market sentiment regarding a company's financial performance.",
	Example: "td eps-revisions --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.AnalysisAPI.GetEpsRevisions(cmd.Context())

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

		if v, _ := cmd.Flags().GetString("country"); v != "" {
			req = req.Country(v)
		}

		if v, _ := cmd.Flags().GetString("exchange"); v != "" {
			req = req.Exchange(v)
		}

		sp := output.StartSpinner(cmd)
		resp, httpResp, callErr := req.Execute()
		sp.Stop()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetEpsTrendCmd = &cobra.Command{
	Use:     "eps-trend",
	Short:   "EPS trend",
	Long:    "The EPS trend endpoint provides detailed historical data on Earnings Per Share (EPS) trends over specified periods. It returns a comprehensive breakdown of estimated EPS changes, allowing users to track and analyze the progression of a company's earnings performance over time. This endpoint is ideal for users seeking to understand historical EPS fluctuations and assess financial growth patterns.",
	Example: "td eps-trend --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.AnalysisAPI.GetEpsTrend(cmd.Context())

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

		if v, _ := cmd.Flags().GetString("country"); v != "" {
			req = req.Country(v)
		}

		if v, _ := cmd.Flags().GetString("exchange"); v != "" {
			req = req.Exchange(v)
		}

		sp := output.StartSpinner(cmd)
		resp, httpResp, callErr := req.Execute()
		sp.Stop()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetGrowthEstimatesCmd = &cobra.Command{
	Use:     "growth-estimates",
	Short:   "Growth estimates",
	Long:    "The growth estimates endpoint provides consensus analyst projections on a company's growth rates over various timeframes. It aggregates and averages estimates from multiple analysts, focusing on key financial metrics such as earnings per share and revenue. This endpoint is useful for obtaining a comprehensive view of expected company performance based on expert analysis.",
	Example: "td growth-estimates --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.AnalysisAPI.GetGrowthEstimates(cmd.Context())

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

		if v, _ := cmd.Flags().GetString("country"); v != "" {
			req = req.Country(v)
		}

		if v, _ := cmd.Flags().GetString("exchange"); v != "" {
			req = req.Exchange(v)
		}

		sp := output.StartSpinner(cmd)
		resp, httpResp, callErr := req.Execute()
		sp.Stop()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetPriceTargetCmd = &cobra.Command{
	Use:     "price-target",
	Short:   "Price target",
	Long:    "The price target endpoint provides detailed projections of a security's future price as estimated by financial analysts. It returns data including the high, low, and average price targets. This endpoint is useful for users seeking to understand potential future valuations of specific securities based on expert analysis.",
	Example: "td price-target --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.AnalysisAPI.GetPriceTarget(cmd.Context())

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

		if v, _ := cmd.Flags().GetString("country"); v != "" {
			req = req.Country(v)
		}

		if v, _ := cmd.Flags().GetString("exchange"); v != "" {
			req = req.Exchange(v)
		}

		sp := output.StartSpinner(cmd)
		resp, httpResp, callErr := req.Execute()
		sp.Stop()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetRecommendationsCmd = &cobra.Command{
	Use:     "recommendations",
	Short:   "Recommendations",
	Long:    "The recommendations endpoint provides a summary of analyst opinions for a specific stock, delivering an average recommendation categorized as Strong Buy, Buy, Hold, or Sell. It also includes a numerical recommendation score, offering a quick overview of market sentiment based on expert analysis.",
	Example: "td recommendations --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.AnalysisAPI.GetRecommendations(cmd.Context())

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

		if v, _ := cmd.Flags().GetString("country"); v != "" {
			req = req.Country(v)
		}

		if v, _ := cmd.Flags().GetString("exchange"); v != "" {
			req = req.Exchange(v)
		}

		sp := output.StartSpinner(cmd)
		resp, httpResp, callErr := req.Execute()
		sp.Stop()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetRevenueEstimateCmd = &cobra.Command{
	Use:     "revenue-estimate",
	Short:   "Revenue estimate",
	Long:    "The revenue estimate endpoint provides a company's projected quarterly and annual revenue figures based on analysts' estimates. This data is useful for users seeking insights into expected company performance, allowing them to compare forecasted sales with historical data or other companies' estimates.",
	Example: "td revenue-estimate --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.AnalysisAPI.GetRevenueEstimate(cmd.Context())

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

		if v, _ := cmd.Flags().GetString("country"); v != "" {
			req = req.Country(v)
		}

		if v, _ := cmd.Flags().GetString("exchange"); v != "" {
			req = req.Exchange(v)
		}

		if cmd.Flags().Changed("dp") {
			v, _ := cmd.Flags().GetInt64("dp")
			req = req.Dp(v)
		}

		sp := output.StartSpinner(cmd)
		resp, httpResp, callErr := req.Execute()
		sp.Stop()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

func init() {

	GetAnalystRatingsLightCmd.Flags().String("symbol", "", "Filter by symbol")

	GetAnalystRatingsLightCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetAnalystRatingsLightCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetAnalystRatingsLightCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetAnalystRatingsLightCmd.Flags().String("exchange", "", "Filter by exchange name")

	flagx.Register(GetAnalystRatingsLightCmd, "rating-change", twelvedata.AllowedRatingChangeEnumEnumValues, "Filter by rating change action")

	GetAnalystRatingsLightCmd.Flags().Int64("outputsize", 0, "Number of records in response")

	GetAnalystRatingsLightCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetAnalystRatingsLightCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetAnalystRatingsLightCmd)

	GetAnalystRatingsUsEquitiesCmd.Flags().String("symbol", "", "Filter by symbol")

	GetAnalystRatingsUsEquitiesCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetAnalystRatingsUsEquitiesCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetAnalystRatingsUsEquitiesCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetAnalystRatingsUsEquitiesCmd.Flags().String("exchange", "", "Filter by exchange name")

	flagx.Register(GetAnalystRatingsUsEquitiesCmd, "rating-change", twelvedata.AllowedRatingChangeEnumEnumValues, "Filter by rating change action")

	GetAnalystRatingsUsEquitiesCmd.Flags().Int64("outputsize", 0, "Number of records in response")

	GetAnalystRatingsUsEquitiesCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetAnalystRatingsUsEquitiesCmd)

	GetEarningsEstimateCmd.Flags().String("symbol", "", "Filter by symbol")

	GetEarningsEstimateCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetEarningsEstimateCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetEarningsEstimateCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetEarningsEstimateCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	GetEarningsEstimateCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetEarningsEstimateCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetEarningsEstimateCmd)

	GetEpsRevisionsCmd.Flags().String("symbol", "", "Filter by symbol")

	GetEpsRevisionsCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetEpsRevisionsCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetEpsRevisionsCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetEpsRevisionsCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetEpsRevisionsCmd.Flags().String("exchange", "", "Filter by exchange name")

	GetEpsRevisionsCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetEpsRevisionsCmd)

	GetEpsTrendCmd.Flags().String("symbol", "", "Filter by symbol")

	GetEpsTrendCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetEpsTrendCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetEpsTrendCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetEpsTrendCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetEpsTrendCmd.Flags().String("exchange", "", "Filter by exchange name")

	GetEpsTrendCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetEpsTrendCmd)

	GetGrowthEstimatesCmd.Flags().String("symbol", "", "Filter by symbol")

	GetGrowthEstimatesCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetGrowthEstimatesCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetGrowthEstimatesCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetGrowthEstimatesCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	GetGrowthEstimatesCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetGrowthEstimatesCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetGrowthEstimatesCmd)

	GetPriceTargetCmd.Flags().String("symbol", "", "Filter by symbol")

	GetPriceTargetCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetPriceTargetCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetPriceTargetCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetPriceTargetCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetPriceTargetCmd.Flags().String("exchange", "", "Filter by exchange name")

	GetPriceTargetCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetPriceTargetCmd)

	GetRecommendationsCmd.Flags().String("symbol", "", "Filter by symbol")

	GetRecommendationsCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetRecommendationsCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetRecommendationsCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetRecommendationsCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	GetRecommendationsCmd.Flags().String("exchange", "", "The exchange name where the instrument is traded, e.g., `Nasdaq`, `NSE`.")

	GetRecommendationsCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetRecommendationsCmd)

	GetRevenueEstimateCmd.Flags().String("symbol", "", "Filter by symbol")

	GetRevenueEstimateCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetRevenueEstimateCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetRevenueEstimateCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetRevenueEstimateCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetRevenueEstimateCmd.Flags().String("exchange", "", "Filter by exchange name")

	GetRevenueEstimateCmd.Flags().Int64("dp", 0, "Number of decimal places for floating values. Should be in range [0,11] inclusive")

	GetRevenueEstimateCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetRevenueEstimateCmd)
}
