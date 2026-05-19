package commands

import (
	"bytes"
	"encoding/json"
	"slices"
	"strings"
	"testing"

	"github.com/spf13/cobra"

	"github.com/twelvedata/twelvedata-cli/internal/flagx"
)

// makeTree builds a fixed Cobra tree we can introspect deterministically — no
// reliance on the generator-emitted commands which change as endpoints are
// added.
func makeTree() *cobra.Command {
	root := &cobra.Command{
		Use:   "tv",
		Short: "root short",
		Long:  "root long",
	}
	root.PersistentFlags().String("api-key", "", "the key")

	group := &cobra.Command{
		Use:     "market-data",
		Aliases: []string{"market"},
		Short:   "group short",
	}
	root.AddCommand(group)

	leaf := &cobra.Command{
		Use:   "price",
		Short: "leaf short",
		RunE:  func(*cobra.Command, []string) error { return nil },
	}
	leaf.Flags().StringP("symbol", "s", "", "ticker")
	_ = leaf.MarkFlagRequired("symbol")
	flagx.Register(leaf, "interval", []string{"1min", "5min"}, "candle interval")
	group.AddCommand(leaf)

	// A hidden command and a help command — both must be filtered out by
	// buildSchema.
	root.AddCommand(&cobra.Command{Use: "hidden", Hidden: true})
	root.AddCommand(&cobra.Command{Use: "help"})

	return root
}

func TestBuildSchema_RootShape(t *testing.T) {
	root := buildSchema(makeTree(), "")
	if root.Name != "tv" || root.Path != "tv" {
		t.Errorf("root name/path = %q/%q, want tv/tv", root.Name, root.Path)
	}
	if root.Short != "root short" || root.Long != "root long" {
		t.Errorf("root short/long not propagated: %+v", root)
	}
	if len(root.Subcommands) != 1 {
		t.Fatalf("hidden + help must be filtered: got %d subcommands (%v)", len(root.Subcommands), names(root.Subcommands))
	}
	if root.Subcommands[0].Name != "market-data" {
		t.Errorf("first subcommand = %q, want market-data", root.Subcommands[0].Name)
	}
}

func TestBuildSchema_PathsNested(t *testing.T) {
	root := buildSchema(makeTree(), "")
	group := root.Subcommands[0]
	if group.Path != "tv market-data" {
		t.Errorf("group path = %q, want \"tv market-data\"", group.Path)
	}
	if len(group.Aliases) == 0 || group.Aliases[0] != "market" {
		t.Errorf("group aliases not preserved: %v", group.Aliases)
	}
	if len(group.Subcommands) != 1 {
		t.Fatalf("expected 1 leaf, got %d", len(group.Subcommands))
	}
	leaf := group.Subcommands[0]
	if leaf.Path != "tv market-data price" {
		t.Errorf("leaf path = %q, want \"tv market-data price\"", leaf.Path)
	}
}

func TestBuildSchema_Flags(t *testing.T) {
	root := buildSchema(makeTree(), "")
	leaf := root.Subcommands[0].Subcommands[0]

	var symbol, interval, apiKey *schemaFlag
	for i := range leaf.Flags {
		f := &leaf.Flags[i]
		switch f.Name {
		case "symbol":
			symbol = f
		case "interval":
			interval = f
		case "api-key":
			apiKey = f
		}
	}

	if symbol == nil {
		t.Fatal("symbol flag missing from schema")
	}
	if symbol.Shorthand != "s" {
		t.Errorf("symbol shorthand = %q, want s", symbol.Shorthand)
	}
	if symbol.Type != "string" {
		t.Errorf("symbol type = %q, want string", symbol.Type)
	}
	if !symbol.Required {
		t.Error("MarkFlagRequired should set Required on the schema")
	}
	if symbol.Inherited {
		t.Error("local flag must not be marked inherited")
	}

	if interval == nil {
		t.Fatal("interval flag missing")
	}
	if len(interval.Enum) != 2 || interval.Enum[0] != "1min" || interval.Enum[1] != "5min" {
		t.Errorf("interval enum = %v, want [1min 5min]", interval.Enum)
	}

	if apiKey == nil {
		t.Fatal("inherited api-key flag missing")
	}
	if !apiKey.Inherited {
		t.Error("api-key was declared on root persistent flags; should be Inherited at the leaf")
	}
}

func TestBuildSchema_FiltersHiddenAndHelp(t *testing.T) {
	root := buildSchema(makeTree(), "")
	for _, sub := range root.Subcommands {
		if sub.Name == "hidden" {
			t.Error("hidden command must be filtered")
		}
		if sub.Name == "help" {
			t.Error("help command must be filtered")
		}
	}
}

func TestSchemaJSON_Encodable(t *testing.T) {
	// The schema is consumed by agents as JSON — make sure it round-trips
	// without errors and produces the documented shape (top-level "name",
	// nested "subcommands").
	root := buildSchema(makeTree(), "")
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetIndent("", "  ")
	if err := enc.Encode(root); err != nil {
		t.Fatalf("schema must be JSON-encodable: %v", err)
	}
	body := buf.String()
	for _, want := range []string{`"name"`, `"path"`, `"subcommands"`, `"flags"`} {
		if !strings.Contains(body, want) {
			t.Errorf("encoded schema missing %q in:\n%s", want, body)
		}
	}
}

// TestCommandsCmd_RegisteredOnRoot guards against accidentally deleting the
// `commands`/`schema` registration in init() — agents rely on it for
// discovery.
func TestCommandsCmd_RegisteredOnRoot(t *testing.T) {
	found := false
	for _, sub := range rootCmd.Commands() {
		if sub.Use == "commands" {
			found = true
			if !contains(sub.Aliases, "schema") {
				t.Error("`commands` should keep its `schema` alias for agent discovery")
			}
			break
		}
	}
	if !found {
		t.Error("`commands` subcommand should be registered on rootCmd")
	}
}

func names(nodes []schemaNode) []string {
	out := make([]string, len(nodes))
	for i, n := range nodes {
		out[i] = n.Name
	}
	return out
}

func contains(s []string, v string) bool {
	return slices.Contains(s, v)
}
