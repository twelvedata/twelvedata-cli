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

var GetETFsFamilyCmd = &cobra.Command{
	Use:     "etfs-family",
	Short:   "ETFs families",
	Long:    "Retrieve a comprehensive list of exchange-traded fund (ETF) families, providing users with detailed information on various ETF groups available in the market. This endpoint is ideal for users looking to explore different ETF categories, compare offerings, or integrate ETF family data into their financial applications.",
	Example: "twelvedata etfs-family",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.EtfsAPI.GetETFsFamily(cmd.Context())

		if v, _ := cmd.Flags().GetString("country"); v != "" {
			req = req.Country(v)
		}

		if v, _ := cmd.Flags().GetString("fund-family"); v != "" {
			req = req.FundFamily(v)
		}

		sp := output.StartSpinner(cmd)
		resp, httpResp, callErr := req.Execute()
		sp.Stop()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetETFsListCmd = &cobra.Command{
	Use:     "etfs-list",
	Short:   "ETFs directory",
	Long:    "The ETFs directory endpoint provides a daily updated list of exchange-traded funds, sorted by total assets in descending order. This endpoint is useful for retrieving comprehensive ETF data, including fund names and asset values, to assist users in quickly identifying the ETFs available.",
	Example: "twelvedata etfs-list",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.EtfsAPI.GetETFsList(cmd.Context())

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

		if cmd.Flags().Changed("dp") {
			v, _ := cmd.Flags().GetInt64("dp")
			req = req.Dp(v)
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

		sp := output.StartSpinner(cmd)
		resp, httpResp, callErr := req.Execute()
		sp.Stop()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetETFsTypeCmd = &cobra.Command{
	Use:     "etfs-type",
	Short:   "ETFs types",
	Long:    "The ETFs Types endpoint provides a concise list of ETF categories by market (e.g., Singapore, United States), including types like 'Equity Precious Metals' and 'Large Blend.' It supports targeted investment research and portfolio diversification.",
	Example: "twelvedata etfs-type",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.EtfsAPI.GetETFsType(cmd.Context())

		if v, _ := cmd.Flags().GetString("country"); v != "" {
			req = req.Country(v)
		}

		if v, _ := cmd.Flags().GetString("fund-type"); v != "" {
			req = req.FundType(v)
		}

		sp := output.StartSpinner(cmd)
		resp, httpResp, callErr := req.Execute()
		sp.Stop()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetETFsWorldCmd = &cobra.Command{
	Use:     "etfs-world",
	Short:   "ETF full data",
	Long:    "The ETF full data endpoint provides detailed information about global Exchange-Traded Funds. It returns comprehensive data, including a summary, performance metrics, risk assessment, and composition details. This endpoint is ideal for users seeking an in-depth analysis of worldwide ETFs, enabling them to access key financial metrics and portfolio breakdowns.",
	Example: "twelvedata etfs-world --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.EtfsAPI.GetETFsWorld(cmd.Context())

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

var GetETFsWorldCompositionCmd = &cobra.Command{
	Use:     "etfs-world-composition",
	Short:   "Composition",
	Long:    "The ETFs composition endpoint provides detailed information about the composition of global Exchange-Traded Funds. It returns data on the sectors included in the ETF, specific holding details, and the weighted exposure of each component. This endpoint is useful for users who need to understand the specific makeup and sector distribution of an ETF portfolio.",
	Example: "twelvedata etfs-world-composition --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.EtfsAPI.GetETFsWorldComposition(cmd.Context())

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

var GetETFsWorldPerformanceCmd = &cobra.Command{
	Use:     "etfs-world-performance",
	Short:   "Performance",
	Long:    "The ETFs performance endpoint provides comprehensive performance data for exchange-traded funds globally. It returns detailed metrics such as trailing returns and annual returns, enabling users to evaluate the historical performance of various ETFs. This endpoint is ideal for users looking to compare ETF performance over different time periods and assess their investment potential.",
	Example: "twelvedata etfs-world-performance --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.EtfsAPI.GetETFsWorldPerformance(cmd.Context())

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

var GetETFsWorldRiskCmd = &cobra.Command{
	Use:     "etfs-world-risk",
	Short:   "Risk",
	Long:    "The ETFs risk endpoint provides essential risk metrics for global Exchange Traded Funds. It returns data such as volatility, beta, and other risk-related indicators, enabling users to assess the potential risk associated with investing in various ETFs worldwide.",
	Example: "twelvedata etfs-world-risk --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.EtfsAPI.GetETFsWorldRisk(cmd.Context())

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

var GetETFsWorldSummaryCmd = &cobra.Command{
	Use:     "etfs-world-summary",
	Short:   "Summary",
	Long:    "The ETFs summary endpoint provides a concise overview of global Exchange-Traded Funds. It returns key data points such as ETF names, symbols, and current market values, enabling users to quickly assess the performance and status of various international ETFs. This summary is ideal for users who need a snapshot of the global ETF landscape without delving into detailed analysis.",
	Example: "twelvedata etfs-world-summary --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.EtfsAPI.GetETFsWorldSummary(cmd.Context())

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

	GetETFsFamilyCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetETFsFamilyCmd.Flags().String("fund-family", "", "Filter by investment company that manages the fund")

	rootCmd.AddCommand(GetETFsFamilyCmd)

	GetETFsListCmd.Flags().String("symbol", "", "Filter by symbol")

	GetETFsListCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetETFsListCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetETFsListCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetETFsListCmd.Flags().String("cik", "", "The CIK of an instrument for which data is requested")

	GetETFsListCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetETFsListCmd.Flags().String("fund-family", "", "Filter by investment company that manages the fund")

	GetETFsListCmd.Flags().String("fund-type", "", "Filter by the type of fund")

	GetETFsListCmd.Flags().Int64("dp", 0, "Number of decimal places for floating values")

	GetETFsListCmd.Flags().Int64("page", 0, "Page number")

	GetETFsListCmd.Flags().Int64("outputsize", 0, "Number of records in response")

	rootCmd.AddCommand(GetETFsListCmd)

	GetETFsTypeCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetETFsTypeCmd.Flags().String("fund-type", "", "Filter by the type of fund")

	rootCmd.AddCommand(GetETFsTypeCmd)

	GetETFsWorldCmd.Flags().String("symbol", "", "Symbol ticker of etf")

	GetETFsWorldCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetETFsWorldCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetETFsWorldCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetETFsWorldCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetETFsWorldCmd.Flags().Int64("dp", 0, "Number of decimal places for floating values. Accepts value in range [0,11]")

	GetETFsWorldCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetETFsWorldCmd)

	GetETFsWorldCompositionCmd.Flags().String("symbol", "", "Symbol ticker of etf")

	GetETFsWorldCompositionCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetETFsWorldCompositionCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetETFsWorldCompositionCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetETFsWorldCompositionCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetETFsWorldCompositionCmd.Flags().Int64("dp", 0, "Number of decimal places for floating values. Accepts value in range [0,11]")

	GetETFsWorldCompositionCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetETFsWorldCompositionCmd)

	GetETFsWorldPerformanceCmd.Flags().String("symbol", "", "Symbol ticker of etf")

	GetETFsWorldPerformanceCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetETFsWorldPerformanceCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetETFsWorldPerformanceCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetETFsWorldPerformanceCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetETFsWorldPerformanceCmd.Flags().Int64("dp", 0, "Number of decimal places for floating values. Accepts value in range [0,11]")

	GetETFsWorldPerformanceCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetETFsWorldPerformanceCmd)

	GetETFsWorldRiskCmd.Flags().String("symbol", "", "Symbol ticker of etf")

	GetETFsWorldRiskCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetETFsWorldRiskCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetETFsWorldRiskCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetETFsWorldRiskCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetETFsWorldRiskCmd.Flags().Int64("dp", 0, "Number of decimal places for floating values. Accepts value in range [0,11]")

	GetETFsWorldRiskCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetETFsWorldRiskCmd)

	GetETFsWorldSummaryCmd.Flags().String("symbol", "", "Symbol ticker of etf")

	GetETFsWorldSummaryCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetETFsWorldSummaryCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetETFsWorldSummaryCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetETFsWorldSummaryCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetETFsWorldSummaryCmd.Flags().Int64("dp", 0, "Number of decimal places for floating values. Accepts value in range [0,11]")

	GetETFsWorldSummaryCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetETFsWorldSummaryCmd)
}
