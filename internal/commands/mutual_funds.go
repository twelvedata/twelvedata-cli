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

var GetMutualFundsFamilyCmd = &cobra.Command{
	Use:   "mutual-funds-family",
	Short: "MFs families",
	Long:  "The mutual funds family endpoint provides a comprehensive list of MF families, which are groups of mutual funds managed by the same investment company. This data is useful for users looking to explore or compare different fund families, understand the range of investment options offered by each, and identify potential investment opportunities within specific fund families.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.MutualFundsAPI.GetMutualFundsFamily(cmd.Context())

		if v, _ := cmd.Flags().GetString("fund-family"); v != "" {
			req = req.FundFamily(v)
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

var GetMutualFundsListCmd = &cobra.Command{
	Use:   "mutual-funds-list",
	Short: "MFs directory",
	Long:  "The mutual funds directory endpoint provides a daily updated list of mutual funds, sorted in descending order by their total assets value. This endpoint is useful for retrieving an organized overview of available mutual funds.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.MutualFundsAPI.GetMutualFundsList(cmd.Context())

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

		if v, _ := cmd.Flags().GetString("country"); v != "" {
			req = req.Country(v)
		}

		if v, _ := cmd.Flags().GetString("fund-family"); v != "" {
			req = req.FundFamily(v)
		}

		if v, _ := cmd.Flags().GetString("fund-type"); v != "" {
			req = req.FundType(v)
		}

		if cmd.Flags().Changed("performance-rating") {
			v, _ := cmd.Flags().GetInt64("performance-rating")
			req = req.PerformanceRating(v)
		}

		if cmd.Flags().Changed("risk-rating") {
			v, _ := cmd.Flags().GetInt64("risk-rating")
			req = req.RiskRating(v)
		}

		if cmd.Flags().Changed("page") {
			v, _ := cmd.Flags().GetInt64("page")
			req = req.Page(v)
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

var GetMutualFundsTypeCmd = &cobra.Command{
	Use:   "mutual-funds-type",
	Short: "MFs types",
	Long:  "This endpoint provides detailed information on various types of mutual funds, such as equity, bond, and balanced funds, allowing users to understand the different investment options available.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.MutualFundsAPI.GetMutualFundsType(cmd.Context())

		if v, _ := cmd.Flags().GetString("fund-type"); v != "" {
			req = req.FundType(v)
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

var GetMutualFundsWorldCmd = &cobra.Command{
	Use:     "mutual-funds-world",
	Short:   "MF full data",
	Long:    "The mutual full data endpoint provides detailed information about global mutual funds. It returns a comprehensive dataset that includes a summary of the fund, its performance metrics, risk assessment, ratings, asset composition, purchase details, and sustainability factors. This endpoint is essential for users seeking in-depth insights into mutual funds on a global scale, allowing them to evaluate various aspects such as investment performance, risk levels, and environmental impact.",
	Example: "td mutual-funds-world --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.MutualFundsAPI.GetMutualFundsWorld(cmd.Context())

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

var GetMutualFundsWorldCompositionCmd = &cobra.Command{
	Use:     "mutual-funds-world-composition",
	Short:   "Composition",
	Long:    "The mutual funds compositions endpoint provides detailed information about the portfolio composition of a specified mutual fund. It returns data on sector allocations, individual holdings, and their respective weighted exposures. This endpoint is useful for users seeking to understand the investment distribution and risk profile of a mutual fund.",
	Example: "td mutual-funds-world-composition --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.MutualFundsAPI.GetMutualFundsWorldComposition(cmd.Context())

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

var GetMutualFundsWorldPerformanceCmd = &cobra.Command{
	Use:     "mutual-funds-world-performance",
	Short:   "Performance",
	Long:    "The mutual funds performances endpoint provides comprehensive performance data for mutual funds globally. It returns metrics such as trailing returns, annual returns, quarterly returns, and load-adjusted returns.",
	Example: "td mutual-funds-world-performance --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.MutualFundsAPI.GetMutualFundsWorldPerformance(cmd.Context())

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

var GetMutualFundsWorldPurchaseInfoCmd = &cobra.Command{
	Use:     "mutual-funds-world-purchase-info",
	Short:   "Purchase info",
	Long:    "The mutual funds purchase information endpoint provides detailed purchasing details for global mutual funds. It returns data on minimum investment requirements, current pricing, and a list of brokerages where the mutual fund can be purchased. This endpoint is useful for users looking to understand the entry requirements and options available for investing in specific mutual funds.",
	Example: "td mutual-funds-world-purchase-info --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.MutualFundsAPI.GetMutualFundsWorldPurchaseInfo(cmd.Context())

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

var GetMutualFundsWorldRatingsCmd = &cobra.Command{
	Use:     "mutual-funds-world-ratings",
	Short:   "Ratings",
	Long:    "The mutual funds ratings endpoint provides detailed ratings for mutual funds across global markets. It returns data on the performance and quality of mutual funds, including ratings calculated in-house by Twelve Data and from various financial institutions.",
	Example: "td mutual-funds-world-ratings --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.MutualFundsAPI.GetMutualFundsWorldRatings(cmd.Context())

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

var GetMutualFundsWorldRiskCmd = &cobra.Command{
	Use:     "mutual-funds-world-risk",
	Short:   "Risk",
	Long:    "The mutual funds risk endpoint provides detailed risk metrics for global mutual funds. It returns data such as standard deviation, beta, and Sharpe ratio, which help assess the volatility and risk profile of mutual funds across different markets.",
	Example: "td mutual-funds-world-risk --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.MutualFundsAPI.GetMutualFundsWorldRisk(cmd.Context())

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

var GetMutualFundsWorldSummaryCmd = &cobra.Command{
	Use:     "mutual-funds-world-summary",
	Short:   "Summary",
	Long:    "The mutual funds summary endpoint provides a concise overview of global mutual funds, including key details such as fund name, symbol, asset class, and region. This endpoint is useful for quickly obtaining essential information about various mutual funds worldwide, aiding in the comparison and selection of funds for investment portfolios.",
	Example: "td mutual-funds-world-summary --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.MutualFundsAPI.GetMutualFundsWorldSummary(cmd.Context())

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

var GetMutualFundsWorldSustainabilityCmd = &cobra.Command{
	Use:     "mutual-funds-world-sustainability",
	Short:   "Sustainability",
	Long:    "The mutual funds sustainability endpoint provides detailed information on the sustainability and Environmental, Social, and Governance (ESG) ratings of global mutual funds. It returns data such as ESG scores, sustainability metrics, and fund identifiers.",
	Example: "td mutual-funds-world-sustainability --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.MutualFundsAPI.GetMutualFundsWorldSustainability(cmd.Context())

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

	GetMutualFundsFamilyCmd.Flags().String("fund-family", "", "Filter by investment company that manages the fund")

	GetMutualFundsFamilyCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	rootCmd.AddCommand(GetMutualFundsFamilyCmd)

	GetMutualFundsListCmd.Flags().String("symbol", "", "Filter by symbol")

	GetMutualFundsListCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetMutualFundsListCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetMutualFundsListCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetMutualFundsListCmd.Flags().String("cik", "", "The CIK of an instrument for which data is requested")

	GetMutualFundsListCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetMutualFundsListCmd.Flags().String("fund-family", "", "Filter by investment company that manages the fund")

	GetMutualFundsListCmd.Flags().String("fund-type", "", "Filter by the type of fund")

	GetMutualFundsListCmd.Flags().Int64("performance-rating", 0, "Filter by performance rating from `0` to `5`")

	GetMutualFundsListCmd.Flags().Int64("risk-rating", 0, "Filter by risk rating from `0` to `5`")

	GetMutualFundsListCmd.Flags().Int64("page", 0, "Page number")

	GetMutualFundsListCmd.Flags().Int64("outputsize", 0, "Number of records in response")

	rootCmd.AddCommand(GetMutualFundsListCmd)

	GetMutualFundsTypeCmd.Flags().String("fund-type", "", "Filter by the type of fund")

	GetMutualFundsTypeCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	rootCmd.AddCommand(GetMutualFundsTypeCmd)

	GetMutualFundsWorldCmd.Flags().String("symbol", "", "Symbol ticker of mutual fund")

	GetMutualFundsWorldCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetMutualFundsWorldCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetMutualFundsWorldCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetMutualFundsWorldCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetMutualFundsWorldCmd.Flags().Int64("dp", 0, "Number of decimal places for floating values. Accepts value in range [0,11]")

	GetMutualFundsWorldCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetMutualFundsWorldCmd)

	GetMutualFundsWorldCompositionCmd.Flags().String("symbol", "", "Symbol ticker of mutual fund")

	GetMutualFundsWorldCompositionCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetMutualFundsWorldCompositionCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetMutualFundsWorldCompositionCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetMutualFundsWorldCompositionCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetMutualFundsWorldCompositionCmd.Flags().Int64("dp", 0, "Number of decimal places for floating values. Accepts value in range [0,11]")

	GetMutualFundsWorldCompositionCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetMutualFundsWorldCompositionCmd)

	GetMutualFundsWorldPerformanceCmd.Flags().String("symbol", "", "Symbol ticker of mutual fund")

	GetMutualFundsWorldPerformanceCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetMutualFundsWorldPerformanceCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetMutualFundsWorldPerformanceCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetMutualFundsWorldPerformanceCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetMutualFundsWorldPerformanceCmd.Flags().Int64("dp", 0, "Number of decimal places for floating values. Accepts value in range [0,11]")

	GetMutualFundsWorldPerformanceCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetMutualFundsWorldPerformanceCmd)

	GetMutualFundsWorldPurchaseInfoCmd.Flags().String("symbol", "", "Symbol ticker of mutual fund")

	GetMutualFundsWorldPurchaseInfoCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetMutualFundsWorldPurchaseInfoCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetMutualFundsWorldPurchaseInfoCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetMutualFundsWorldPurchaseInfoCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetMutualFundsWorldPurchaseInfoCmd.Flags().Int64("dp", 0, "Number of decimal places for floating values. Accepts value in range [0,11]")

	GetMutualFundsWorldPurchaseInfoCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetMutualFundsWorldPurchaseInfoCmd)

	GetMutualFundsWorldRatingsCmd.Flags().String("symbol", "", "Symbol ticker of mutual fund")

	GetMutualFundsWorldRatingsCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetMutualFundsWorldRatingsCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetMutualFundsWorldRatingsCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetMutualFundsWorldRatingsCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetMutualFundsWorldRatingsCmd.Flags().Int64("dp", 0, "Number of decimal places for floating values. Accepts value in range [0,11]")

	GetMutualFundsWorldRatingsCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetMutualFundsWorldRatingsCmd)

	GetMutualFundsWorldRiskCmd.Flags().String("symbol", "", "Symbol ticker of mutual fund")

	GetMutualFundsWorldRiskCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetMutualFundsWorldRiskCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetMutualFundsWorldRiskCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetMutualFundsWorldRiskCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetMutualFundsWorldRiskCmd.Flags().Int64("dp", 0, "Number of decimal places for floating values. Accepts value in range [0,11]")

	GetMutualFundsWorldRiskCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetMutualFundsWorldRiskCmd)

	GetMutualFundsWorldSummaryCmd.Flags().String("symbol", "", "Symbol ticker of mutual fund")

	GetMutualFundsWorldSummaryCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetMutualFundsWorldSummaryCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetMutualFundsWorldSummaryCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetMutualFundsWorldSummaryCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetMutualFundsWorldSummaryCmd.Flags().Int64("dp", 0, "Number of decimal places for floating values. Accepts value in range [0,11]")

	GetMutualFundsWorldSummaryCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetMutualFundsWorldSummaryCmd)

	GetMutualFundsWorldSustainabilityCmd.Flags().String("symbol", "", "Symbol ticker of mutual fund")

	GetMutualFundsWorldSustainabilityCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetMutualFundsWorldSustainabilityCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetMutualFundsWorldSustainabilityCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetMutualFundsWorldSustainabilityCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetMutualFundsWorldSustainabilityCmd.Flags().Int64("dp", 0, "Number of decimal places for floating values. Accepts value in range [0,11]")

	GetMutualFundsWorldSustainabilityCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetMutualFundsWorldSustainabilityCmd)
}
