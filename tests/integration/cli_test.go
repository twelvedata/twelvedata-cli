//go:build integration

// Package integration runs the built `twelvedata` binary against the live Twelve
// Data API. There is no sandbox tier — calls hit production with a real key.
//
// Run:
//
//	export TWELVEDATA_API_KEY=...
//	go test -tags=integration -v -timeout=10m ./tests/integration/...
//
// One test per endpoint, sequential, with a small delay between calls (see
// delayMs) to stay friendly with the API rate limit. Mirrors the Go SDK's
// integration suite in shape and field-level assertions.
package integration

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

const (
	symbolStock      = "AAPL"
	symbolEtf        = "SPY"
	symbolForex      = "EUR/USD"
	symbolCrypto     = "BTC/USD"
	symbolMutualFund = "VFINX"
	micCode          = "XNYS"
	startDate        = "2025-01-01"
	endDate          = "2025-01-31"
	timezone         = "UTC"
	currencyAmount   = "100"
	intervalDay      = "1day"
	outputsize       = "10"
	delayMs          = 150 * time.Millisecond
)

// binaryPath is the absolute path to the CLI binary used by every test in this
// package. Populated by TestMain — either from $TWELVEDATA_CLI_BIN or by
// building cmd/twelvedata into a temp dir once per test run.
var binaryPath string

func TestMain(m *testing.M) {
	if v := os.Getenv("TWELVEDATA_CLI_BIN"); v != "" {
		binaryPath = v
		os.Exit(m.Run())
	}

	dir, err := os.MkdirTemp("", "twelvedata-cli-it-")
	if err != nil {
		fmt.Fprintln(os.Stderr, "mkdir temp:", err)
		os.Exit(2)
	}
	defer os.RemoveAll(dir)

	binName := "twelvedata"
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}
	binaryPath = filepath.Join(dir, binName)

	repoRoot, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		fmt.Fprintln(os.Stderr, "abs repo root:", err)
		os.Exit(2)
	}
	cmd := exec.Command("go", "build", "-o", binaryPath, "./cmd/twelvedata")
	cmd.Dir = repoRoot
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "go build:", err)
		os.Exit(2)
	}

	os.Exit(m.Run())
}

// runCLI runs the CLI with --raw, asserts the process exited 0, and returns
// stdout. A delayMs sleep is scheduled via t.Cleanup so the next test starts
// after a rate-limit-friendly pause. Mirrors the SDK's setupClient cleanup.
func runCLI(t *testing.T, args ...string) []byte {
	t.Helper()
	args = append(args, "--raw")
	cmd := exec.Command(binaryPath, args...)
	cmd.Env = append(os.Environ(),
		"TWELVEDATA_NO_UPDATE_NOTIFIER=1",
		"TERM=dumb",
	)
	out, err := cmd.Output()
	t.Cleanup(func() { time.Sleep(delayMs) })
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			t.Fatalf("twelvedata %v: exit=%d\nstdout:\n%s\nstderr:\n%s",
				args, ee.ExitCode(), out, ee.Stderr)
		}
		t.Fatalf("twelvedata %v: %v\nstdout:\n%s", args, err, out)
	}
	return out
}

// requireAPIKey skips the test when no key is provided. Useful when running the
// integration suite without credentials to exercise only the non-network paths
// (which live in auth_test.go).
func requireAPIKey(t *testing.T) {
	t.Helper()
	if os.Getenv("TWELVEDATA_API_KEY") == "" {
		t.Skip("TWELVEDATA_API_KEY not set; skipping live-API test")
	}
}

// mustJSON unmarshals out into v or fails the test, including the raw body in
// the failure message for debugging.
func mustJSON(t *testing.T, out []byte, v any) {
	t.Helper()
	if err := json.Unmarshal(out, v); err != nil {
		t.Fatalf("json unmarshal: %v; body=%s", err, out)
	}
}

// --- MarketData ---

func TestCLITimeSeries(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "time-series", "--symbol", symbolStock, "--interval", intervalDay, "--outputsize", outputsize)
	var resp struct {
		Status string           `json:"status"`
		Values []map[string]any `json:"values"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Values) == 0 {
		t.Fatal("empty values")
	}
	t.Logf("%+v", resp.Values[0])
}

func TestCLIExchangeRate(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "exchange-rate", "--symbol", symbolForex)
	var resp struct {
		Rate float64 `json:"rate"`
	}
	mustJSON(t, out, &resp)
	if resp.Rate == 0 {
		t.Fatal("expected non-zero rate")
	}
	t.Logf("%+v", resp)
}

func TestCLIPrice(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "price", "--symbol", symbolStock)
	var resp struct {
		Price string `json:"price"`
	}
	mustJSON(t, out, &resp)
	if resp.Price == "" {
		t.Fatal("expected non-empty price")
	}
	t.Logf("%+v", resp)
}

func TestCLIQuote(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "quote", "--symbol", symbolStock)
	var resp struct {
		Symbol string `json:"symbol"`
	}
	mustJSON(t, out, &resp)
	if resp.Symbol == "" {
		t.Fatal("expected non-empty symbol")
	}
	t.Logf("%+v", resp)
}

func TestCLIEod(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "eod", "--symbol", symbolStock)
	var resp struct {
		Close string `json:"close"`
	}
	mustJSON(t, out, &resp)
	if resp.Close == "" {
		t.Fatal("expected non-empty close")
	}
	t.Logf("%+v", resp)
}

func TestCLICurrencyConversion(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "currency-conversion", "--symbol", symbolForex, "--amount", currencyAmount)
	var resp struct {
		Rate float64 `json:"rate"`
	}
	mustJSON(t, out, &resp)
	if resp.Rate == 0 {
		t.Fatal("expected non-zero rate")
	}
	t.Logf("%+v", resp)
}

func TestCLITimeSeriesCross(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "time-series-cross", "--base", "EUR", "--quote", "USD", "--interval", intervalDay, "--outputsize", outputsize)
	var resp struct {
		Values []map[string]any `json:"values"`
	}
	mustJSON(t, out, &resp)
	if len(resp.Values) == 0 {
		t.Fatal("empty values")
	}
	t.Logf("%+v", resp.Values[0])
}

// --- Fundamentals ---

func TestCLIIpoCalendar(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "ipo-calendar", "--start-date", startDate, "--end-date", endDate)
	if len(out) == 0 {
		t.Fatal("empty response")
	}
	// Response is a free-form array/object; assert valid JSON.
	var any any
	mustJSON(t, out, &any)
}

func TestCLIEarningsCalendar(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "earnings-calendar")
	var resp struct {
		Status   string `json:"status"`
		Earnings any    `json:"earnings"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if resp.Earnings == nil {
		t.Fatal("expected earnings data")
	}
}

func TestCLIDividendsCalendar(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "dividends-calendar", "--symbol", symbolStock)
	var any any
	mustJSON(t, out, &any)
}

func TestCLISplitsCalendar(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "splits-calendar", "--symbol", symbolStock, "--start-date", startDate, "--end-date", endDate)
	var any any
	mustJSON(t, out, &any)
}

func TestCLIEarnings(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "earnings", "--symbol", symbolStock)
	var resp struct {
		Status   string           `json:"status"`
		Earnings []map[string]any `json:"earnings"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Earnings) == 0 {
		t.Fatal("expected earnings data")
	}
	t.Logf("%+v", resp.Earnings[0])
}

func TestCLIDividends(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "dividends", "--symbol", symbolStock)
	var any any
	mustJSON(t, out, &any)
}

func TestCLISplits(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "splits", "--symbol", symbolStock)
	var any any
	mustJSON(t, out, &any)
}

func TestCLIBalanceSheet(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "balance-sheet", "--symbol", symbolStock)
	var any any
	mustJSON(t, out, &any)
}

func TestCLIBalanceSheetConsolidated(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "balance-sheet-consolidated", "--symbol", symbolStock)
	var any any
	mustJSON(t, out, &any)
}

func TestCLICashFlow(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "cash-flow", "--symbol", symbolStock)
	var any any
	mustJSON(t, out, &any)
}

func TestCLICashFlowConsolidated(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "cash-flow-consolidated", "--symbol", symbolStock)
	var any any
	mustJSON(t, out, &any)
}

func TestCLIIncomeStatement(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "income-statement", "--symbol", symbolStock)
	var any any
	mustJSON(t, out, &any)
}

func TestCLIIncomeStatementConsolidated(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "income-statement-consolidated", "--symbol", symbolStock)
	var any any
	mustJSON(t, out, &any)
}

func TestCLIProfile(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "profile", "--symbol", symbolStock)
	var resp struct {
		Symbol string `json:"symbol"`
	}
	mustJSON(t, out, &resp)
	if resp.Symbol == "" {
		t.Fatal("expected non-empty symbol")
	}
	t.Logf("%+v", resp)
}

func TestCLIStatistics(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "statistics", "--symbol", symbolStock)
	var any any
	mustJSON(t, out, &any)
}

func TestCLIMarketCap(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "market-cap", "--symbol", symbolStock)
	var any any
	mustJSON(t, out, &any)
}

func TestCLIKeyExecutives(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "key-executives", "--symbol", symbolStock)
	var any any
	mustJSON(t, out, &any)
}

func TestCLILogo(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "logo", "--symbol", symbolStock)
	var any any
	mustJSON(t, out, &any)
}

func TestCLIPressReleases(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "press-releases", "--symbol", symbolStock)
	var resp struct {
		Status string `json:"status"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
}

// --- Analysis ---

func TestCLIAnalystRatingsLight(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "analyst-ratings-light", "--symbol", symbolStock)
	var resp struct {
		Status string `json:"status"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
}

func TestCLIAnalystRatingsUsEquities(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "analyst-ratings-us-equities", "--symbol", symbolStock)
	var resp struct {
		Status string `json:"status"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
}

func TestCLIPriceTarget(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "price-target", "--symbol", symbolStock)
	var resp struct {
		Status string `json:"status"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
}

func TestCLIRecommendations(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "recommendations", "--symbol", symbolStock)
	var resp struct {
		Status string `json:"status"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
}

func TestCLIEarningsEstimate(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "earnings-estimate", "--symbol", symbolStock)
	var resp struct {
		Status           string           `json:"status"`
		EarningsEstimate []map[string]any `json:"earnings_estimate"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.EarningsEstimate) == 0 {
		t.Fatal("expected earnings_estimate data")
	}
}

func TestCLIEpsRevisions(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "eps-revisions", "--symbol", symbolStock)
	var resp struct {
		Status      string           `json:"status"`
		EpsRevision []map[string]any `json:"eps_revision"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.EpsRevision) == 0 {
		t.Fatal("expected eps_revision data")
	}
}

func TestCLIEpsTrend(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "eps-trend", "--symbol", symbolStock)
	var resp struct {
		Status   string           `json:"status"`
		EpsTrend []map[string]any `json:"eps_trend"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.EpsTrend) == 0 {
		t.Fatal("expected eps_trend data")
	}
}

func TestCLIGrowthEstimates(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "growth-estimates", "--symbol", symbolStock)
	var resp struct {
		Status string `json:"status"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
}

func TestCLIRevenueEstimate(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "revenue-estimate", "--symbol", symbolStock)
	var resp struct {
		Status          string           `json:"status"`
		RevenueEstimate []map[string]any `json:"revenue_estimate"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.RevenueEstimate) == 0 {
		t.Fatal("expected revenue_estimate data")
	}
}

// --- ETFs ---

func TestCLIEtfsList(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "etfs-list", "--symbol", symbolEtf)
	var resp struct {
		Status string `json:"status"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
}

func TestCLIEtfsType(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "etfs-type")
	var resp struct {
		Status string         `json:"status"`
		Result map[string]any `json:"result"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Result) == 0 {
		t.Fatal("expected non-empty result")
	}
}

func TestCLIEtfsFamily(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "etfs-family")
	var resp struct {
		Status string         `json:"status"`
		Result map[string]any `json:"result"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Result) == 0 {
		t.Fatal("expected non-empty result")
	}
}

func TestCLIEtfsWorld(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "etfs-world", "--symbol", symbolEtf)
	var resp struct {
		Status string `json:"status"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
}

func TestCLIEtfsWorldSummary(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "etfs-world-summary", "--symbol", symbolEtf)
	var resp struct {
		Status string `json:"status"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
}

func TestCLIEtfsWorldComposition(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "etfs-world-composition", "--symbol", symbolEtf)
	var resp struct {
		Status string `json:"status"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
}

func TestCLIEtfsWorldPerformance(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "etfs-world-performance", "--symbol", symbolEtf)
	var resp struct {
		Status string `json:"status"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
}

func TestCLIEtfsWorldRisk(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "etfs-world-risk", "--symbol", symbolEtf)
	var resp struct {
		Status string `json:"status"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
}

// --- Mutual funds ---

func TestCLIMutualFundsList(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "mutual-funds-list", "--symbol", symbolMutualFund)
	var resp struct {
		Status string `json:"status"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
}

func TestCLIMutualFundsType(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "mutual-funds-type")
	var resp struct {
		Status string         `json:"status"`
		Result map[string]any `json:"result"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Result) == 0 {
		t.Fatal("expected non-empty result")
	}
}

func TestCLIMutualFundsFamily(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "mutual-funds-family")
	var resp struct {
		Status string         `json:"status"`
		Result map[string]any `json:"result"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Result) == 0 {
		t.Fatal("expected non-empty result")
	}
}

func TestCLIMutualFundsWorld(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "mutual-funds-world", "--symbol", symbolMutualFund)
	var resp struct {
		Status string `json:"status"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
}

func TestCLIMutualFundsWorldSummary(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "mutual-funds-world-summary", "--symbol", symbolMutualFund)
	var resp struct {
		Status string `json:"status"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
}

func TestCLIMutualFundsWorldComposition(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "mutual-funds-world-composition", "--symbol", symbolMutualFund)
	var resp struct {
		Status string `json:"status"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
}

func TestCLIMutualFundsWorldPerformance(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "mutual-funds-world-performance", "--symbol", symbolMutualFund)
	var resp struct {
		Status string `json:"status"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
}

func TestCLIMutualFundsWorldRisk(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "mutual-funds-world-risk", "--symbol", symbolMutualFund)
	var resp struct {
		Status string `json:"status"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
}

func TestCLIMutualFundsWorldRatings(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "mutual-funds-world-ratings", "--symbol", symbolMutualFund)
	var resp struct {
		Status string `json:"status"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
}

func TestCLIMutualFundsWorldPurchaseInfo(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "mutual-funds-world-purchase-info", "--symbol", symbolMutualFund)
	var resp struct {
		Status string `json:"status"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
}

func TestCLIMutualFundsWorldSustainability(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "mutual-funds-world-sustainability", "--symbol", symbolMutualFund)
	var resp struct {
		Status string `json:"status"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
}

// --- Reference data ---

func TestCLIStocks(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "stocks", "--symbol", symbolStock)
	var resp struct {
		Status string `json:"status"`
		Data   []any  `json:"data"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Data) == 0 {
		t.Fatal("expected non-empty data")
	}
}

func TestCLICryptocurrencies(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "cryptocurrencies", "--symbol", symbolCrypto)
	var resp struct {
		Status string `json:"status"`
		Data   []any  `json:"data"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Data) == 0 {
		t.Fatal("expected non-empty data")
	}
}

func TestCLICryptocurrencyExchanges(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "cryptocurrency-exchanges")
	var resp struct {
		Status string `json:"status"`
		Data   []any  `json:"data"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Data) == 0 {
		t.Fatal("expected non-empty data")
	}
}

func TestCLIForexPairs(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "forex-pairs", "--symbol", symbolForex)
	var resp struct {
		Status string `json:"status"`
		Data   []any  `json:"data"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Data) == 0 {
		t.Fatal("expected non-empty data")
	}
}

func TestCLIExchanges(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "exchanges")
	var resp struct {
		Status string `json:"status"`
		Data   []any  `json:"data"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Data) == 0 {
		t.Fatal("expected non-empty data")
	}
}

func TestCLIExchangeSchedule(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "exchange-schedule", "--mic-code", micCode)
	var resp struct {
		Data []any `json:"data"`
	}
	mustJSON(t, out, &resp)
	if len(resp.Data) == 0 {
		t.Fatal("expected non-empty data")
	}
}

func TestCLICountries(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "countries")
	var resp struct {
		Data []any `json:"data"`
	}
	mustJSON(t, out, &resp)
	if len(resp.Data) == 0 {
		t.Fatal("expected non-empty data")
	}
}

func TestCLIIntervals(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "intervals")
	var resp struct {
		Status string `json:"status"`
		Data   []any  `json:"data"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Data) == 0 {
		t.Fatal("expected non-empty data")
	}
}

func TestCLIMarketState(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "market-state")
	var resp []any
	mustJSON(t, out, &resp)
	if len(resp) == 0 {
		t.Fatal("expected non-empty array")
	}
}

func TestCLIEarliestTimestamp(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "earliest-timestamp", "--symbol", symbolStock, "--interval", intervalDay)
	var resp struct {
		Datetime string `json:"datetime"`
	}
	mustJSON(t, out, &resp)
	if resp.Datetime == "" {
		t.Fatal("expected non-empty datetime")
	}
}

func TestCLISymbolSearch(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "symbol-search", "--symbol", symbolStock)
	var resp struct {
		Status string `json:"status"`
		Data   []any  `json:"data"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Data) == 0 {
		t.Fatal("expected non-empty data")
	}
}

func TestCLIInstrumentType(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "instrument-type")
	var resp struct {
		Status string `json:"status"`
		Result []any  `json:"result"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Result) == 0 {
		t.Fatal("expected non-empty result")
	}
}

func TestCLICrossListings(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "cross-listings", "--symbol", symbolStock)
	var resp struct {
		Result any `json:"result"`
	}
	mustJSON(t, out, &resp)
	if resp.Result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestCLIBonds(t *testing.T) {
	// TODO: uncomment this when the API response content-type is fixed.
	// Matches the SDK's TestReferenceDataGetBonds skip.
	t.Skip("skipping bonds test due to incorrect API response content-type")
}

func TestCLICommodities(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "commodities")
	var resp struct {
		Status string `json:"status"`
		Data   []any  `json:"data"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Data) == 0 {
		t.Fatal("expected non-empty data")
	}
}

func TestCLIEtfs(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "etfs", "--symbol", symbolEtf)
	var resp struct {
		Status string `json:"status"`
		Data   []any  `json:"data"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Data) == 0 {
		t.Fatal("expected non-empty data")
	}
}

func TestCLIFunds(t *testing.T) {
	// TODO: uncomment this when the API response content-type is fixed.
	// Matches the SDK's TestReferenceDataGetFunds skip.
	t.Skip("skipping funds test due to incorrect API response content-type")
}

func TestCLITechnicalIndicators(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "technical-indicators")
	var resp struct {
		Status string         `json:"status"`
		Data   map[string]any `json:"data"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Data) == 0 {
		t.Fatal("expected non-empty data")
	}
}

// --- Regulatory ---

func TestCLIDirectHolders(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "direct-holders", "--symbol", symbolStock)
	var any any
	mustJSON(t, out, &any)
}

func TestCLIFundHolders(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "fund-holders", "--symbol", symbolStock)
	var any any
	mustJSON(t, out, &any)
}

func TestCLIInstitutionalHolders(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "institutional-holders", "--symbol", symbolStock)
	var any any
	mustJSON(t, out, &any)
}

func TestCLIInsiderTransactions(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "insider-transactions", "--symbol", symbolStock)
	var any any
	mustJSON(t, out, &any)
}

func TestCLITaxInfo(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "tax-info", "--symbol", symbolStock)
	var resp struct {
		Status string `json:"status"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
}

func TestCLIEdgarFilingsArchive(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "edgar-filings-archive", "--symbol", symbolStock)
	var any any
	mustJSON(t, out, &any)
}

// --- Technical indicators (RSI + MACD only, matching the SDK carve-out) ---

func TestCLITiRsi(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "ti", "rsi", "--symbol", symbolStock, "--interval", intervalDay, "--outputsize", outputsize, "--timezone", timezone)
	var resp struct {
		Status string           `json:"status"`
		Values []map[string]any `json:"values"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Values) == 0 {
		t.Fatal("empty values")
	}
	t.Logf("%+v", resp.Values[0])
}

func TestCLITiMacd(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "ti", "macd", "--symbol", symbolStock, "--interval", intervalDay, "--outputsize", outputsize, "--timezone", timezone)
	var resp struct {
		Status string           `json:"status"`
		Values []map[string]any `json:"values"`
	}
	mustJSON(t, out, &resp)
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Values) == 0 {
		t.Fatal("empty values")
	}
	t.Logf("%+v", resp.Values[0])
}

// --- Advanced ---

func TestCLIApiUsage(t *testing.T) {
	requireAPIKey(t)
	out := runCLI(t, "api-usage")
	var resp struct {
		Timestamp string `json:"timestamp"`
	}
	mustJSON(t, out, &resp)
	if resp.Timestamp == "" {
		t.Fatal("expected non-empty timestamp")
	}
	t.Logf("%+v", resp)
}
