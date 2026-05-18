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

var GetBalanceSheetCmd = &cobra.Command{
	Use:     "balance-sheet",
	Short:   "Balance sheet",
	Long:    "The balance sheet endpoint provides a detailed financial statement for a company, outlining its assets, liabilities, and shareholders' equity. This endpoint returns structured data that includes current and non-current assets, total liabilities, and equity figures, enabling users to assess a company's financial health and stability.",
	Example: "twelvedata balance-sheet --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.FundamentalsAPI.GetBalanceSheet(cmd.Context())

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

		if cmd.Flags().Changed("period") {
			req = req.Period(twelvedata.PeriodEnum(cmd.Flags().Lookup("period").Value.String()))
		}

		if v, _ := cmd.Flags().GetString("start-date"); v != "" {
			req = req.StartDate(v)
		}

		if v, _ := cmd.Flags().GetString("end-date"); v != "" {
			req = req.EndDate(v)
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

var GetBalanceSheetConsolidatedCmd = &cobra.Command{
	Use:     "balance-sheet-consolidated",
	Short:   "Balance sheet consolidated",
	Long:    "The balance sheet consolidated endpoint provides a detailed overview of a company's raw balance sheet, including a comprehensive summary of its assets, liabilities, and shareholders' equity. This endpoint is useful for retrieving financial data that reflects the overall financial position of a company, allowing users to access critical information about its financial health and structure.",
	Example: "twelvedata balance-sheet-consolidated --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.FundamentalsAPI.GetBalanceSheetConsolidated(cmd.Context())

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

		if cmd.Flags().Changed("period") {
			req = req.Period(twelvedata.PeriodEnum(cmd.Flags().Lookup("period").Value.String()))
		}

		if v, _ := cmd.Flags().GetString("start-date"); v != "" {
			req = req.StartDate(v)
		}

		if v, _ := cmd.Flags().GetString("end-date"); v != "" {
			req = req.EndDate(v)
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

var GetCashFlowCmd = &cobra.Command{
	Use:     "cash-flow",
	Short:   "Cash flow",
	Long:    "The cash flow endpoint provides detailed information on a company's cash flow activities, including the net cash and cash equivalents moving in and out of the business. This data includes operating, investing, and financing cash flows, offering a comprehensive view of the company's liquidity and financial health.",
	Example: "twelvedata cash-flow --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.FundamentalsAPI.GetCashFlow(cmd.Context())

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

		if cmd.Flags().Changed("period") {
			req = req.Period(twelvedata.PeriodEnum(cmd.Flags().Lookup("period").Value.String()))
		}

		if v, _ := cmd.Flags().GetString("start-date"); v != "" {
			req = req.StartDate(v)
		}

		if v, _ := cmd.Flags().GetString("end-date"); v != "" {
			req = req.EndDate(v)
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

var GetCashFlowConsolidatedCmd = &cobra.Command{
	Use:     "cash-flow-consolidated",
	Short:   "Cash flow consolidated",
	Long:    "The cash flow consolidated endpoint provides raw data on a company's consolidated cash flow, including the net cash and cash equivalents moving in and out of the business. It returns information on operating, investing, and financing activities, helping users track liquidity and financial health over a specified period.",
	Example: "twelvedata cash-flow-consolidated --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.FundamentalsAPI.GetCashFlowConsolidated(cmd.Context())

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

		if cmd.Flags().Changed("period") {
			req = req.Period(twelvedata.PeriodEnum(cmd.Flags().Lookup("period").Value.String()))
		}

		if v, _ := cmd.Flags().GetString("start-date"); v != "" {
			req = req.StartDate(v)
		}

		if v, _ := cmd.Flags().GetString("end-date"); v != "" {
			req = req.EndDate(v)
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

var GetDividendsCmd = &cobra.Command{
	Use:     "dividends",
	Short:   "Dividends",
	Long:    "The dividends endpoint provides historical dividend data for a specified stock, in many cases covering over a decade. It returns information on dividend payouts, including the ex-date, amount, and frequency. This endpoint is ideal for users tracking dividend histories or evaluating the income potential of stocks.",
	Example: "twelvedata dividends --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.FundamentalsAPI.GetDividends(cmd.Context())

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

		if cmd.Flags().Changed("range") {
			req = req.Range_(twelvedata.RangeEnum(cmd.Flags().Lookup("range").Value.String()))
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

		sp := output.StartSpinner(cmd)
		resp, httpResp, callErr := req.Execute()
		sp.Stop()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetDividendsCalendarCmd = &cobra.Command{
	Use:     "dividends-calendar",
	Short:   "Dividends calendar",
	Long:    "The dividends calendar endpoint provides a detailed schedule of upcoming and past dividend events for specified date ranges. By using the `start_date` and `end_date` parameters, users can retrieve a list of companies issuing dividends, including the ex-dividend date and dividend amount. This endpoint is ideal for tracking dividend payouts and planning investment strategies based on dividend schedules.",
	Example: "twelvedata dividends-calendar --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.FundamentalsAPI.GetDividendsCalendar(cmd.Context())

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

		if v, _ := cmd.Flags().GetString("start-date"); v != "" {
			req = req.StartDate(v)
		}

		if v, _ := cmd.Flags().GetString("end-date"); v != "" {
			req = req.EndDate(v)
		}

		if cmd.Flags().Changed("outputsize") {
			v, _ := cmd.Flags().GetInt64("outputsize")
			req = req.Outputsize(v)
		}

		if cmd.Flags().Changed("page") {
			v, _ := cmd.Flags().GetInt64("page")
			req = req.Page(v)
		}

		sp := output.StartSpinner(cmd)
		resp, httpResp, callErr := req.Execute()
		sp.Stop()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetEarningsCmd = &cobra.Command{
	Use:     "earnings",
	Short:   "Earnings",
	Long:    "The earnings endpoint provides comprehensive earnings data for a specified company, including both the estimated and actual Earnings Per Share (EPS) figures. This endpoint delivers historical earnings information, allowing users to track a company's financial performance over time.",
	Example: "twelvedata earnings --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.FundamentalsAPI.GetEarnings(cmd.Context())

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

		if cmd.Flags().Changed("period") {
			req = req.Period(twelvedata.PeriodEarningsEnum(cmd.Flags().Lookup("period").Value.String()))
		}

		if cmd.Flags().Changed("outputsize") {
			v, _ := cmd.Flags().GetInt64("outputsize")
			req = req.Outputsize(v)
		}

		if v, _ := cmd.Flags().GetString("start-date"); v != "" {
			req = req.StartDate(v)
		}

		if v, _ := cmd.Flags().GetString("end-date"); v != "" {
			req = req.EndDate(v)
		}

		if cmd.Flags().Changed("dp") {
			v, _ := cmd.Flags().GetInt64("dp")
			req = req.Dp(v)
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

var GetEarningsCalendarCmd = &cobra.Command{
	Use:   "earnings-calendar",
	Short: "Earnings calendar",
	Long:  "The earnings calendar endpoint provides a schedule of company earnings announcements for a specified date range. By default, it returns earnings data for the current day. Users can customize the date range using the `start_date` and `end_date` parameters to retrieve earnings information for specific periods. This endpoint is useful for tracking upcoming earnings reports and planning around key financial announcements.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.FundamentalsAPI.GetEarningsCalendar(cmd.Context())

		if v, _ := cmd.Flags().GetString("exchange"); v != "" {
			req = req.Exchange(v)
		}

		if v, _ := cmd.Flags().GetString("mic-code"); v != "" {
			req = req.MicCode(v)
		}

		if v, _ := cmd.Flags().GetString("country"); v != "" {
			req = req.Country(v)
		}

		if v, _ := cmd.Flags().GetString("start-date"); v != "" {
			req = req.StartDate(v)
		}

		if v, _ := cmd.Flags().GetString("end-date"); v != "" {
			req = req.EndDate(v)
		}

		if cmd.Flags().Changed("dp") {
			v, _ := cmd.Flags().GetInt64("dp")
			req = req.Dp(v)
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

var GetIncomeStatementCmd = &cobra.Command{
	Use:     "income-statement",
	Short:   "Income statement",
	Long:    "The income statement endpoint provides detailed financial data on a company's income statement, including revenues, expenses, and net income for specified periods, either annually or quarterly. This endpoint is essential for retrieving comprehensive financial performance metrics of a company, allowing users to access historical and current financial results.",
	Example: "twelvedata income-statement --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.FundamentalsAPI.GetIncomeStatement(cmd.Context())

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

		if cmd.Flags().Changed("period") {
			req = req.Period(twelvedata.PeriodEnum(cmd.Flags().Lookup("period").Value.String()))
		}

		if v, _ := cmd.Flags().GetString("start-date"); v != "" {
			req = req.StartDate(v)
		}

		if v, _ := cmd.Flags().GetString("end-date"); v != "" {
			req = req.EndDate(v)
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

var GetIncomeStatementConsolidatedCmd = &cobra.Command{
	Use:     "income-statement-consolidated",
	Short:   "Income statement consolidated",
	Long:    "The income statement consolidated endpoint provides a company's raw income statement, detailing revenue, expenses, and net income for specified periods, either annually or quarterly. This data is essential for evaluating a company's financial performance over time, allowing users to access comprehensive financial results in a structured format.",
	Example: "twelvedata income-statement-consolidated --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.FundamentalsAPI.GetIncomeStatementConsolidated(cmd.Context())

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

		if cmd.Flags().Changed("period") {
			req = req.Period(twelvedata.PeriodEnum(cmd.Flags().Lookup("period").Value.String()))
		}

		if v, _ := cmd.Flags().GetString("start-date"); v != "" {
			req = req.StartDate(v)
		}

		if v, _ := cmd.Flags().GetString("end-date"); v != "" {
			req = req.EndDate(v)
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

var GetIpoCalendarCmd = &cobra.Command{
	Use:   "ipo-calendar",
	Short: "IPO calendar",
	Long:  "The IPO Calendar endpoint provides detailed information on initial public offerings (IPOs), including those that have occurred in the past, are happening today, or are scheduled for the future. Users can access data such as company names, IPO dates, and offering details, allowing them to track and monitor IPO activity efficiently.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.FundamentalsAPI.GetIpoCalendar(cmd.Context())

		if v, _ := cmd.Flags().GetString("exchange"); v != "" {
			req = req.Exchange(v)
		}

		if v, _ := cmd.Flags().GetString("mic-code"); v != "" {
			req = req.MicCode(v)
		}

		if v, _ := cmd.Flags().GetString("country"); v != "" {
			req = req.Country(v)
		}

		if v, _ := cmd.Flags().GetString("start-date"); v != "" {
			req = req.StartDate(v)
		}

		if v, _ := cmd.Flags().GetString("end-date"); v != "" {
			req = req.EndDate(v)
		}

		sp := output.StartSpinner(cmd)
		resp, httpResp, callErr := req.Execute()
		sp.Stop()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetKeyExecutivesCmd = &cobra.Command{
	Use:     "key-executives",
	Short:   "Key executives",
	Long:    "The key executives endpoint provides detailed information about a company's key executives identified by a specific stock symbol. It returns data such as names, titles, and roles of the executives, which can be useful for understanding the leadership structure of the company.",
	Example: "twelvedata key-executives --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.FundamentalsAPI.GetKeyExecutives(cmd.Context())

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

		sp := output.StartSpinner(cmd)
		resp, httpResp, callErr := req.Execute()
		sp.Stop()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetLogoCmd = &cobra.Command{
	Use:     "logo",
	Short:   "Logo",
	Long:    "The logo endpoint provides the official logo image for a specified company, cryptocurrency, or forex pair. This endpoint is useful for integrating visual branding elements into financial applications, websites, or reports, ensuring that users can easily identify and associate the correct logo with the respective financial asset.",
	Example: "twelvedata logo --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.FundamentalsAPI.GetLogo(cmd.Context())

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

		sp := output.StartSpinner(cmd)
		resp, httpResp, callErr := req.Execute()
		sp.Stop()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetMarketCapCmd = &cobra.Command{
	Use:     "market-cap",
	Short:   "Market capitalization",
	Long:    "The Market Capitalization History endpoint provides historical data on a company's market capitalization over a specified time period. It returns a time series of market cap values, allowing users to track changes in a company's market value.",
	Example: "twelvedata market-cap --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.FundamentalsAPI.GetMarketCap(cmd.Context())

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

		if v, _ := cmd.Flags().GetString("start-date"); v != "" {
			req = req.StartDate(v)
		}

		if v, _ := cmd.Flags().GetString("end-date"); v != "" {
			req = req.EndDate(v)
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

var GetProfileCmd = &cobra.Command{
	Use:     "profile",
	Short:   "Profile",
	Long:    "The profile endpoint provides detailed company information, including the company's name, industry, sector, CEO, headquarters location, and market capitalization. This data is useful for obtaining a comprehensive overview of a company's business and financial standing.",
	Example: "twelvedata profile --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.FundamentalsAPI.GetProfile(cmd.Context())

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

		sp := output.StartSpinner(cmd)
		resp, httpResp, callErr := req.Execute()
		sp.Stop()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetSplitsCmd = &cobra.Command{
	Use:     "splits",
	Short:   "Splits",
	Long:    "The splits endpoint provides historical data on stock split events for a specified company. It returns details including the date of each split and the corresponding split factor.",
	Example: "twelvedata splits --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.FundamentalsAPI.GetSplits(cmd.Context())

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

		if cmd.Flags().Changed("range") {
			req = req.Range_(twelvedata.RangeSplitsEnum(cmd.Flags().Lookup("range").Value.String()))
		}

		if v, _ := cmd.Flags().GetString("start-date"); v != "" {
			req = req.StartDate(v)
		}

		if v, _ := cmd.Flags().GetString("end-date"); v != "" {
			req = req.EndDate(v)
		}

		sp := output.StartSpinner(cmd)
		resp, httpResp, callErr := req.Execute()
		sp.Stop()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetSplitsCalendarCmd = &cobra.Command{
	Use:     "splits-calendar",
	Short:   "Splits calendar",
	Long:    "The splits calendar endpoint provides a detailed calendar of stock split events within a specified date range. By setting the `start_date` and `end_date` parameters, users can retrieve a list of upcoming or past stock splits, including the company name, split ratio, and effective date. This endpoint is useful for tracking changes in stock structure and planning investment strategies around these events.",
	Example: "twelvedata splits-calendar --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.FundamentalsAPI.GetSplitsCalendar(cmd.Context())

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

		if v, _ := cmd.Flags().GetString("start-date"); v != "" {
			req = req.StartDate(v)
		}

		if v, _ := cmd.Flags().GetString("end-date"); v != "" {
			req = req.EndDate(v)
		}

		if cmd.Flags().Changed("outputsize") {
			v, _ := cmd.Flags().GetInt64("outputsize")
			req = req.Outputsize(v)
		}

		if v, _ := cmd.Flags().GetString("page"); v != "" {
			req = req.Page(v)
		}

		sp := output.StartSpinner(cmd)
		resp, httpResp, callErr := req.Execute()
		sp.Stop()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetStatisticsCmd = &cobra.Command{
	Use:     "statistics",
	Short:   "Statistics",
	Long:    "The statistics endpoint provides a comprehensive snapshot of a company's key financial statistics, including valuation metrics, revenue figures, profit margins, and other essential financial data. This endpoint is ideal for users seeking detailed insights into a company's financial health and performance metrics.",
	Example: "twelvedata statistics --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.FundamentalsAPI.GetStatistics(cmd.Context())

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

		sp := output.StartSpinner(cmd)
		resp, httpResp, callErr := req.Execute()
		sp.Stop()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var PressReleasesListParametersCmd = &cobra.Command{
	Use:     "press-releases",
	Short:   "Press releases",
	Long:    "The press releases endpoint offers structured, real-time access to official company press releases and corporate announcements from public entities across global markets.",
	Example: "twelvedata press-releases --symbol AAPL",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.FundamentalsAPI.PressReleasesListParameters(cmd.Context())

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

		if v, _ := cmd.Flags().GetString("start-date"); v != "" {
			req = req.StartDate(v)
		}

		if v, _ := cmd.Flags().GetString("end-date"); v != "" {
			req = req.EndDate(v)
		}

		if v, _ := cmd.Flags().GetString("timezone"); v != "" {
			req = req.Timezone(v)
		}

		if v, _ := cmd.Flags().GetString("language"); v != "" {
			req = req.Language(v)
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

func init() {

	GetBalanceSheetCmd.Flags().String("symbol", "", "Symbol ticker of instrument. For preferred stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct")

	GetBalanceSheetCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetBalanceSheetCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetBalanceSheetCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetBalanceSheetCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetBalanceSheetCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetBalanceSheetCmd.Flags().String("country", "", "Country where instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetBalanceSheetCmd, "period", twelvedata.AllowedPeriodEnumEnumValues, "The reporting period for the balance sheet data")

	GetBalanceSheetCmd.Flags().String("start-date", "", "Begin date for filtering items by fiscal date. Returns income statements with fiscal dates on or after this date. Format `2006-01-02`")

	GetBalanceSheetCmd.Flags().String("end-date", "", "End date for filtering items by fiscal date. Returns income statements with fiscal dates on or before this date. Format `2006-01-02`")

	GetBalanceSheetCmd.Flags().Int64("outputsize", 0, "Number of records in response")

	GetBalanceSheetCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetBalanceSheetCmd)

	GetBalanceSheetConsolidatedCmd.Flags().String("symbol", "", "Symbol ticker of instrument. For preferred stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct")

	GetBalanceSheetConsolidatedCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetBalanceSheetConsolidatedCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetBalanceSheetConsolidatedCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetBalanceSheetConsolidatedCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetBalanceSheetConsolidatedCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetBalanceSheetConsolidatedCmd.Flags().String("country", "", "Country where instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetBalanceSheetConsolidatedCmd, "period", twelvedata.AllowedPeriodEnumEnumValues, "The reporting period for the balance sheet data.")

	GetBalanceSheetConsolidatedCmd.Flags().String("start-date", "", "Begin date for filtering items by fiscal date. Returns income statements with fiscal dates on or after this date. Format `2006-01-02`")

	GetBalanceSheetConsolidatedCmd.Flags().String("end-date", "", "End date for filtering items by fiscal date. Returns income statements with fiscal dates on or before this date. Format `2006-01-02`")

	GetBalanceSheetConsolidatedCmd.Flags().Int64("outputsize", 0, "Number of records in response")

	GetBalanceSheetConsolidatedCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetBalanceSheetConsolidatedCmd)

	GetCashFlowCmd.Flags().String("symbol", "", "Symbol ticker of instrument. For preferred stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct")

	GetCashFlowCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetCashFlowCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetCashFlowCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetCashFlowCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetCashFlowCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetCashFlowCmd.Flags().String("country", "", "Country where instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetCashFlowCmd, "period", twelvedata.AllowedPeriodEnumEnumValues, "The reporting period for the cash flow statements")

	GetCashFlowCmd.Flags().String("start-date", "", "Start date for filtering cash flow statements. Only cash flow statements with fiscal dates on or after this date will be included. Format `2006-01-02`")

	GetCashFlowCmd.Flags().String("end-date", "", "End date for filtering cash flow statements. Only cash flow statements with fiscal dates on or before this date will be included. Format `2006-01-02`")

	GetCashFlowCmd.Flags().Int64("outputsize", 0, "Number of records in response")

	GetCashFlowCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetCashFlowCmd)

	GetCashFlowConsolidatedCmd.Flags().String("symbol", "", "Symbol ticker of instrument. For preferred stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct")

	GetCashFlowConsolidatedCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetCashFlowConsolidatedCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetCashFlowConsolidatedCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetCashFlowConsolidatedCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetCashFlowConsolidatedCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetCashFlowConsolidatedCmd.Flags().String("country", "", "Country where instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetCashFlowConsolidatedCmd, "period", twelvedata.AllowedPeriodEnumEnumValues, "The reporting period for the cash flow statements")

	GetCashFlowConsolidatedCmd.Flags().String("start-date", "", "Start date for filtering cash flow statements. Only cash flow statements with fiscal dates on or after this date will be included. Format `2006-01-02`")

	GetCashFlowConsolidatedCmd.Flags().String("end-date", "", "End date for filtering cash flow statements. Only cash flow statements with fiscal dates on or before this date will be included. Format `2006-01-02`")

	GetCashFlowConsolidatedCmd.Flags().Int64("outputsize", 0, "Number of records in response")

	GetCashFlowConsolidatedCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetCashFlowConsolidatedCmd)

	GetDividendsCmd.Flags().String("symbol", "", "Symbol ticker of instrument. For preferred stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct")

	GetDividendsCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetDividendsCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetDividendsCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetDividendsCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetDividendsCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetDividendsCmd.Flags().String("country", "", "Country where instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetDividendsCmd, "range", twelvedata.AllowedRangeEnumEnumValues, "Specifies the time range for which to retrieve dividend data. Accepts values such as `last` (most recent dividend), `next` (upcoming dividend), `1m` - `5y` for respective periods, or `full` for all available data. If provided together with `start-date` and/or `end-date`, this parameter takes precedence.")

	GetDividendsCmd.Flags().String("start-date", "", "Start date for the dividend data query. Only dividends with dates on or after this date will be returned. Format `2006-01-02`. If provided together with `range` parameter, `range` will take precedence.")

	GetDividendsCmd.Flags().String("end-date", "", "End date for the dividend data query. Only dividends with dates on or before this date will be returned. Format `2006-01-02`. If provided together with `range` parameter, `range` will take precedence.")

	GetDividendsCmd.Flags().Bool("adjust", false, "Specifies if there should be an adjustment")

	GetDividendsCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetDividendsCmd)

	GetDividendsCalendarCmd.Flags().String("symbol", "", "Symbol ticker of instrument. For preferred stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct")

	GetDividendsCalendarCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetDividendsCalendarCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetDividendsCalendarCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetDividendsCalendarCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetDividendsCalendarCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetDividendsCalendarCmd.Flags().String("country", "", "Country where instrument is traded, e.g., `United States` or `US`")

	GetDividendsCalendarCmd.Flags().String("start-date", "", "Start date for the dividends calendar query. Only dividends with ex-dates on or after this date will be returned. Format `2006-01-02`")

	GetDividendsCalendarCmd.Flags().String("end-date", "", "End date for the dividends calendar query. Only dividends with ex-dates on or before this date will be returned. Format `2006-01-02`")

	GetDividendsCalendarCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `500`. Default `100` when no date parameters are set, otherwise set to maximum")

	GetDividendsCalendarCmd.Flags().Int64("page", 0, "Page number")

	GetDividendsCalendarCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetDividendsCalendarCmd)

	GetEarningsCmd.Flags().String("symbol", "", "Symbol ticker of instrument. For preferred stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct")

	GetEarningsCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetEarningsCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetEarningsCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetEarningsCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetEarningsCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetEarningsCmd.Flags().String("country", "", "Country where instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetEarningsCmd, "type", twelvedata.AllowedTypeEnumEnumValues, "The asset class to which the instrument belongs")

	flagx.Register(GetEarningsCmd, "period", twelvedata.AllowedPeriodEarningsEnumEnumValues, "Type of earning, returns only 1 record. When is not empty, dates and outputsize parameters are ignored")

	GetEarningsCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `1000`. Default `10` when no date parameters are set, otherwise set to maximum")

	GetEarningsCmd.Flags().String("start-date", "", "The date from which the data is requested. The date format is `YYYY-MM-DD`.")

	GetEarningsCmd.Flags().String("end-date", "", "The date to which the data is requested. The date format is `YYYY-MM-DD`.")

	GetEarningsCmd.Flags().Int64("dp", 0, "The number of decimal places in the response data. Should be in range [0,11] inclusive")

	GetEarningsCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetEarningsCmd)

	GetEarningsCalendarCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetEarningsCalendarCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetEarningsCalendarCmd.Flags().String("country", "", "Country where instrument is traded, e.g., `United States` or `US`")

	GetEarningsCalendarCmd.Flags().String("start-date", "", "Can be used separately and together with end_date. Format `2006-01-02` or `2006-01-02T15:04:05`")

	GetEarningsCalendarCmd.Flags().String("end-date", "", "Can be used separately and together with start_date. Format `2006-01-02` or `2006-01-02T15:04:05`")

	GetEarningsCalendarCmd.Flags().Int64("dp", 0, "Specifies the number of decimal places for floating values. Should be in range [0,11] inclusive")

	rootCmd.AddCommand(GetEarningsCalendarCmd)

	GetIncomeStatementCmd.Flags().String("symbol", "", "Symbol ticker of instrument. For preferred stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct")

	GetIncomeStatementCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetIncomeStatementCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetIncomeStatementCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetIncomeStatementCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetIncomeStatementCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetIncomeStatementCmd.Flags().String("country", "", "Country where instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetIncomeStatementCmd, "period", twelvedata.AllowedPeriodEnumEnumValues, "The reporting period for the income statement data")

	GetIncomeStatementCmd.Flags().String("start-date", "", "Begin date for filtering income statements by fiscal date. Returns income statements with fiscal dates on or after this date. Format `2006-01-02`")

	GetIncomeStatementCmd.Flags().String("end-date", "", "End date for filtering income statements by fiscal date. Returns income statements with fiscal dates on or before this date. Format `2006-01-02`")

	GetIncomeStatementCmd.Flags().Int64("outputsize", 0, "Number of records in response")

	GetIncomeStatementCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetIncomeStatementCmd)

	GetIncomeStatementConsolidatedCmd.Flags().String("symbol", "", "Symbol ticker of instrument. For preferred stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct")

	GetIncomeStatementConsolidatedCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetIncomeStatementConsolidatedCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetIncomeStatementConsolidatedCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetIncomeStatementConsolidatedCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetIncomeStatementConsolidatedCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetIncomeStatementConsolidatedCmd.Flags().String("country", "", "Country where instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetIncomeStatementConsolidatedCmd, "period", twelvedata.AllowedPeriodEnumEnumValues, "The reporting period for the income statement data")

	GetIncomeStatementConsolidatedCmd.Flags().String("start-date", "", "Begin date for filtering income statements by fiscal date. Returns income statements with fiscal dates on or after this date. Format `2006-01-02`")

	GetIncomeStatementConsolidatedCmd.Flags().String("end-date", "", "End date for filtering income statements by fiscal date. Returns income statements with fiscal dates on or before this date. Format `2006-01-02`")

	GetIncomeStatementConsolidatedCmd.Flags().Int64("outputsize", 0, "Number of records in response")

	GetIncomeStatementConsolidatedCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetIncomeStatementConsolidatedCmd)

	GetIpoCalendarCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetIpoCalendarCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetIpoCalendarCmd.Flags().String("country", "", "Country where instrument is traded, e.g., `United States` or `US`")

	GetIpoCalendarCmd.Flags().String("start-date", "", "The earliest IPO date to include in the results. Format: `2006-01-02`")

	GetIpoCalendarCmd.Flags().String("end-date", "", "The latest IPO date to include in the results. Format: `2006-01-02`")

	rootCmd.AddCommand(GetIpoCalendarCmd)

	GetKeyExecutivesCmd.Flags().String("symbol", "", "Symbol ticker of instrument. For preferred stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct")

	GetKeyExecutivesCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetKeyExecutivesCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetKeyExecutivesCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetKeyExecutivesCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetKeyExecutivesCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetKeyExecutivesCmd.Flags().String("country", "", "Country where instrument is traded, e.g., `United States` or `US`")

	GetKeyExecutivesCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetKeyExecutivesCmd)

	GetLogoCmd.Flags().String("symbol", "", "The ticker symbol of an instrument for which data is requested, e.g., `AAPL`, `BTC/USD`, `EUR/USD`.")

	GetLogoCmd.Flags().String("exchange", "", "The exchange name where the instrument is traded, e.g., `NASDAQ`, `NSE`")

	GetLogoCmd.Flags().String("mic-code", "", "The Market Identifier Code (MIC) of the exchange where the instrument is traded, e.g., `XNAS`, `XLON`")

	GetLogoCmd.Flags().String("country", "", "The country where the instrument is traded, e.g., `United States` or `US`")

	_ = GetLogoCmd.MarkFlagRequired("symbol")

	rootCmd.AddCommand(GetLogoCmd)

	GetMarketCapCmd.Flags().String("symbol", "", "Filter by symbol")

	GetMarketCapCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetMarketCapCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetMarketCapCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetMarketCapCmd.Flags().String("exchange", "", "Filter by exchange name")

	GetMarketCapCmd.Flags().String("mic-code", "", "Filter by market identifier code (MIC) under ISO 10383 standard")

	GetMarketCapCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetMarketCapCmd.Flags().String("start-date", "", "Start date for market capitalization data retrieval. Data will be returned from this date onwards. Format `2006-01-02`")

	GetMarketCapCmd.Flags().String("end-date", "", "End date for market capitalization data retrieval. Data will be returned up to and including this date. Format `2006-01-02`")

	GetMarketCapCmd.Flags().Int64("page", 0, "Page number")

	GetMarketCapCmd.Flags().Int64("outputsize", 0, "Number of records in response")

	GetMarketCapCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetMarketCapCmd)

	GetProfileCmd.Flags().String("symbol", "", "Symbol ticker of instrument. For preferred stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct")

	GetProfileCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetProfileCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetProfileCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetProfileCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetProfileCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetProfileCmd.Flags().String("country", "", "Country where instrument is traded, e.g., `United States` or `US`")

	GetProfileCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetProfileCmd)

	GetSplitsCmd.Flags().String("symbol", "", "Symbol ticker of instrument. For preferred stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct")

	GetSplitsCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetSplitsCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetSplitsCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetSplitsCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetSplitsCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetSplitsCmd.Flags().String("country", "", "Country where instrument is traded, e.g., `United States` or `US`")

	flagx.Register(GetSplitsCmd, "range", twelvedata.AllowedRangeSplitsEnumEnumValues, "Range of data to be returned")

	GetSplitsCmd.Flags().String("start-date", "", "The starting date for data selection. Format `2006-01-02`")

	GetSplitsCmd.Flags().String("end-date", "", "The ending date for data selection. Format `2006-01-02`")

	GetSplitsCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetSplitsCmd)

	GetSplitsCalendarCmd.Flags().String("symbol", "", "Symbol ticker of instrument. For preferred stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct")

	GetSplitsCalendarCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetSplitsCalendarCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetSplitsCalendarCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetSplitsCalendarCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetSplitsCalendarCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetSplitsCalendarCmd.Flags().String("country", "", "Country where instrument is traded, e.g., `United States` or `US`")

	GetSplitsCalendarCmd.Flags().String("start-date", "", "The starting date (inclusive) for filtering split events in the calendar. Format `2006-01-02`")

	GetSplitsCalendarCmd.Flags().String("end-date", "", "The ending date (inclusive) for filtering split events in the calendar. Format `2006-01-02`")

	GetSplitsCalendarCmd.Flags().Int64("outputsize", 0, "Number of data points to retrieve. Supports values in the range from `1` to `500`. Default `100` when no date parameters are set, otherwise set to maximum")

	GetSplitsCalendarCmd.Flags().String("page", "", "Page number")

	GetSplitsCalendarCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetSplitsCalendarCmd)

	GetStatisticsCmd.Flags().String("symbol", "", "Symbol ticker of instrument. For preferred stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct")

	GetStatisticsCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetStatisticsCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetStatisticsCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetStatisticsCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetStatisticsCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetStatisticsCmd.Flags().String("country", "", "Country where instrument is traded, e.g., `United States` or `US`")

	GetStatisticsCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetStatisticsCmd)

	PressReleasesListParametersCmd.Flags().String("symbol", "", "Symbol ticker of instrument. For preferred stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct")

	PressReleasesListParametersCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	PressReleasesListParametersCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	PressReleasesListParametersCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	PressReleasesListParametersCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	PressReleasesListParametersCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	PressReleasesListParametersCmd.Flags().String("start-date", "", "Begin date for filtering items. Returns press releases with release date on or after this date. Format `2025-12-24T02:07:00`")

	PressReleasesListParametersCmd.Flags().String("end-date", "", "End date for filtering items. Returns press releases with release date on or before this date. Format `2025-12-24T02:07:00`")

	PressReleasesListParametersCmd.Flags().String("timezone", "", "Time zone for date filtering. Default is the identifier time zone.")

	PressReleasesListParametersCmd.Flags().String("language", "", "Comma-separated list of languages to filter press releases by language.")

	PressReleasesListParametersCmd.Flags().Int64("outputsize", 0, "Number of latest press releases returned. Only used if no data range is specified. Maximum value is `10`. type: number")

	PressReleasesListParametersCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(PressReleasesListParametersCmd)
}
