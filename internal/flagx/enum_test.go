package flagx

import (
	"slices"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// stringPlain is here so we can exercise the `~string` type constraint without
// pulling in the SDK's enum types.
type stringPlain string

func TestEnum_Set_Valid(t *testing.T) {
	e := New([]stringPlain{"1min", "5min", "1day"})
	if err := e.Set("5min"); err != nil {
		t.Fatalf("Set(\"5min\") returned %v", err)
	}
	if got := e.String(); got != "5min" {
		t.Errorf("String() = %q, want \"5min\"", got)
	}
	if got := e.Type(); got != "string" {
		t.Errorf("Type() = %q, want \"string\"", got)
	}
}

func TestEnum_Set_Invalid(t *testing.T) {
	e := New([]stringPlain{"asc", "desc"})
	err := e.Set("sideways")
	if err == nil {
		t.Fatal("Set on a non-allowed value must return an error")
	}
	for _, want := range []string{"asc", "desc"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("error should list allowed values, missing %q in %q", want, err)
		}
	}
}

func TestEnum_Allowed(t *testing.T) {
	e := New([]stringPlain{"a", "b", "c"})
	got := e.Allowed()
	want := []string{"a", "b", "c"}
	if !slices.Equal(got, want) {
		t.Errorf("Allowed() = %v, want %v", got, want)
	}
}

func TestRegister_AddsFlag(t *testing.T) {
	cmd := &cobra.Command{Use: "t"}
	Register(cmd, "interval", []stringPlain{"1min", "5min"}, "candle interval")

	flag := cmd.Flags().Lookup("interval")
	if flag == nil {
		t.Fatal("Register should declare --interval on the command")
	}
	if flag.Usage != "candle interval" {
		t.Errorf("usage = %q, want \"candle interval\"", flag.Usage)
	}
	if err := flag.Value.Set("1min"); err != nil {
		t.Errorf("Set(1min) returned %v", err)
	}
	if err := flag.Value.Set("99min"); err == nil {
		t.Error("Set on disallowed value should fail")
	}
}

func TestCompletion_EmitsAllowedValues(t *testing.T) {
	f := completion([]stringPlain{"1min", "5min"})
	got, directive := f(&cobra.Command{}, nil, "")
	if !slices.Equal(got, []string{"1min", "5min"}) {
		t.Errorf("completion candidates = %v, want [1min 5min]", got)
	}
	if directive != cobra.ShellCompDirectiveNoFileComp {
		t.Errorf("directive = %v, want NoFileComp", directive)
	}
}

func TestEnumValues_OnEnumFlag(t *testing.T) {
	cmd := &cobra.Command{Use: "t"}
	Register(cmd, "order", []stringPlain{"asc", "desc"}, "")
	flag := cmd.Flags().Lookup("order")
	if !slices.Equal(EnumValues(flag), []string{"asc", "desc"}) {
		t.Errorf("EnumValues() = %v, want [asc desc]", EnumValues(flag))
	}
}

func TestEnumValues_OnPlainFlag(t *testing.T) {
	fs := pflag.NewFlagSet("t", pflag.ContinueOnError)
	fs.String("symbol", "", "")
	flag := fs.Lookup("symbol")
	if EnumValues(flag) != nil {
		t.Errorf("EnumValues on a plain string flag must return nil, got %v", EnumValues(flag))
	}
}
