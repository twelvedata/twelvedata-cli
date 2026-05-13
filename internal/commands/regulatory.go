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

var GetDirectHoldersCmd = &cobra.Command{
	Use:   "direct-holders",
	Short: "Direct holders",
	Long:  "The direct holders endpoint provides detailed information about the number of shares directly held by individuals or entities as recorded in a company's official share registry. This data is essential for understanding the distribution of stock ownership within a company, helping users identify major shareholders and assess shareholder concentration.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.RegulatoryAPI.GetDirectHolders(cmd.Context())

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

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetEdgarFilingsArchiveCmd = &cobra.Command{
	Use:   "edgar-filings-archive",
	Short: "EDGAR fillings",
	Long:  "The EDGAR fillings endpoint provides access to a comprehensive collection of financial documents submitted to the SEC, including real-time and historical forms, filings, and exhibits. Users can retrieve detailed information about company disclosures, financial statements, and regulatory submissions, enabling them to access essential compliance and financial data directly from the SEC's EDGAR system.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.RegulatoryAPI.GetEdgarFilingsArchive(cmd.Context())

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

		if v, _ := cmd.Flags().GetString("form-type"); v != "" {
			req = req.FormType(v)
		}

		if v, _ := cmd.Flags().GetString("filled-from"); v != "" {
			req = req.FilledFrom(v)
		}

		if v, _ := cmd.Flags().GetString("filled-to"); v != "" {
			req = req.FilledTo(v)
		}

		if cmd.Flags().Changed("page") {
			v, _ := cmd.Flags().GetInt64("page")
			req = req.Page(v)
		}

		if cmd.Flags().Changed("page-size") {
			v, _ := cmd.Flags().GetInt64("page-size")
			req = req.PageSize(v)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetFundHoldersCmd = &cobra.Command{
	Use:   "fund-holders",
	Short: "Fund holders",
	Long:  "The fund holders endpoint provides detailed information about the proportion of a company's stock that is owned by mutual fund holders. It returns data on the number of shares held, the percentage of total shares outstanding, and the names of the mutual funds involved. This endpoint is useful for users looking to understand mutual fund investment in a specific company.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.RegulatoryAPI.GetFundHolders(cmd.Context())

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

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetInsiderTransactionsCmd = &cobra.Command{
	Use:   "insider-transactions",
	Short: "Insider transaction",
	Long:  "The insider transaction endpoint provides detailed data on trades executed by company insiders, such as executives and directors. It returns information including the insider's name, their role, the transaction type, the number of shares, the transaction date, and the price per share. This endpoint is useful for tracking insider activity and understanding potential insider sentiment towards a company's stock.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.RegulatoryAPI.GetInsiderTransactions(cmd.Context())

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

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetInstitutionalHoldersCmd = &cobra.Command{
	Use:   "institutional-holders",
	Short: "Institutional holders",
	Long:  "The institutional holders endpoint provides detailed information on the percentage and amount of a company's stock owned by institutional investors, such as pension funds, insurance companies, and investment firms. This data is essential for understanding the influence and involvement of large entities in a company's ownership structure.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.RegulatoryAPI.GetInstitutionalHolders(cmd.Context())

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

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

var GetTaxInfoCmd = &cobra.Command{
	Use:   "tax-info",
	Short: "Tax information",
	Long:  "The tax information endpoint provides detailed tax-related data for a specified financial instrument, including applicable tax rates and relevant tax codes. This information is essential for users needing to understand the tax implications associated with trading or investing in specific instruments.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := client.New(cmd)
		if err != nil {
			return err
		}

		req := api.RegulatoryAPI.GetTaxInfo(cmd.Context())

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

		if v, _ := cmd.Flags().GetString("exchange"); v != "" {
			req = req.Exchange(v)
		}

		if v, _ := cmd.Flags().GetString("mic-code"); v != "" {
			req = req.MicCode(v)
		}

		resp, httpResp, callErr := req.Execute()
		return output.Render(cmd, resp, httpResp, callErr)
	},
}

func init() {

	GetDirectHoldersCmd.Flags().String("symbol", "", "Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct")

	GetDirectHoldersCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetDirectHoldersCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetDirectHoldersCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetDirectHoldersCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetDirectHoldersCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetDirectHoldersCmd.Flags().String("country", "", "Country where instrument is traded, e.g., `United States` or `US`")

	GetDirectHoldersCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetDirectHoldersCmd)

	GetEdgarFilingsArchiveCmd.Flags().String("symbol", "", "The ticker symbol of an instrument for which data is requested")

	GetEdgarFilingsArchiveCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetEdgarFilingsArchiveCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetEdgarFilingsArchiveCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetEdgarFilingsArchiveCmd.Flags().String("exchange", "", "Filter by exchange name")

	GetEdgarFilingsArchiveCmd.Flags().String("mic-code", "", "Filter by market identifier code (MIC) under ISO 10383 standard")

	GetEdgarFilingsArchiveCmd.Flags().String("country", "", "Filter by country name or alpha code, e.g., `United States` or `US`")

	GetEdgarFilingsArchiveCmd.Flags().String("form-type", "", "Filter by form types, example `8-K`, `EX-1.1`")

	GetEdgarFilingsArchiveCmd.Flags().String("filled-from", "", "Filter by filled time from")

	GetEdgarFilingsArchiveCmd.Flags().String("filled-to", "", "Filter by filled time to")

	GetEdgarFilingsArchiveCmd.Flags().Int64("page", 0, "Page number")

	GetEdgarFilingsArchiveCmd.Flags().Int64("page-size", 0, "Number of records in response")

	GetEdgarFilingsArchiveCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetEdgarFilingsArchiveCmd)

	GetFundHoldersCmd.Flags().String("symbol", "", "Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct")

	GetFundHoldersCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetFundHoldersCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetFundHoldersCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetFundHoldersCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetFundHoldersCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetFundHoldersCmd.Flags().String("country", "", "Country where instrument is traded, e.g., `United States` or `US`")

	GetFundHoldersCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetFundHoldersCmd)

	GetInsiderTransactionsCmd.Flags().String("symbol", "", "The ticker symbol of an instrument for which data is requested, e.g., `AAPL`, `TSLA`. For preffered stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct")

	GetInsiderTransactionsCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetInsiderTransactionsCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetInsiderTransactionsCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetInsiderTransactionsCmd.Flags().String("exchange", "", "Exchange where instrument is traded, e.g., `Nasdaq`, `NSE`")

	GetInsiderTransactionsCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetInsiderTransactionsCmd.Flags().String("country", "", "Country where instrument is traded, e.g., United States or US.")

	GetInsiderTransactionsCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetInsiderTransactionsCmd)

	GetInstitutionalHoldersCmd.Flags().String("symbol", "", "Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct")

	GetInstitutionalHoldersCmd.Flags().String("figi", "", "Filter by financial instrument global identifier (FIGI). This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetInstitutionalHoldersCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetInstitutionalHoldersCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetInstitutionalHoldersCmd.Flags().String("exchange", "", "Exchange where instrument is traded")

	GetInstitutionalHoldersCmd.Flags().String("mic-code", "", "Market Identifier Code (MIC) under ISO 10383 standard")

	GetInstitutionalHoldersCmd.Flags().String("country", "", "Country where instrument is traded, e.g., `United States` or `US`")

	GetInstitutionalHoldersCmd.MarkFlagsOneRequired("symbol", "figi", "isin", "cusip")

	rootCmd.AddCommand(GetInstitutionalHoldersCmd)

	GetTaxInfoCmd.Flags().String("symbol", "", "The ticker symbol of an instrument for which data is requested, e.g., `SKYQ`, `AIRE`, `ALM:BME`, `HSI:HKEX`.")

	GetTaxInfoCmd.Flags().String("isin", "", "Filter by international securities identification number (ISIN). ISIN access is activating in the Data add-ons section")

	GetTaxInfoCmd.Flags().String("figi", "", "The FIGI of an instrument for which data is requested. This parameter is available on the Ultra plan (individual) and the Enterprise plan (business) and above.")

	GetTaxInfoCmd.Flags().String("cusip", "", "The CUSIP of an instrument for which data is requested. CUSIP access is activating in the Data add-ons section")

	GetTaxInfoCmd.Flags().String("exchange", "", "The exchange name where the instrument is traded, e.g., `Nasdaq`, `Euronext`")

	GetTaxInfoCmd.Flags().String("mic-code", "", "The Market Identifier Code (MIC) of the exchange where the instrument is traded, e.g., `XNAS`, `XLON`")

	GetTaxInfoCmd.MarkFlagsOneRequired("symbol", "isin", "figi", "cusip")

	rootCmd.AddCommand(GetTaxInfoCmd)
}
