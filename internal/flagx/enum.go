// Package flagx is a small set of Cobra/pflag value types that aren't in
// upstream pflag. The most important one is Enum: a flag value constrained to a
// set of allowed strings, with shell completion and a marker the schema dumper
// uses for agent discovery.
package flagx

import (
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Enum is a pflag.Value bound to a fixed set of allowed strings. Set rejects
// values outside the allow-list; the resulting parse error becomes a usage
// error (exit code 2) at the CLI layer.
type Enum struct {
	value   string
	allowed []string
}

// AllowedValuer is implemented by Enum so the schema dumper can introspect the
// allowed set and surface it to agents in `td schema`. Any flag value that
// wants to advertise an enum to agents can satisfy this interface.
type AllowedValuer interface {
	Allowed() []string
}

// New constructs an Enum from any string-typed enum slice (e.g. the SDK's
// AllowedIntervalEnumEnumValues). The constraint T ~string lets us accept the
// SDK's typed enums directly without per-type conversion at every call site.
func New[T ~string](allowed []T) *Enum {
	s := make([]string, len(allowed))
	for i, v := range allowed {
		s[i] = string(v)
	}
	return &Enum{allowed: s}
}

func (e *Enum) String() string    { return e.value }
func (e *Enum) Type() string      { return "string" }
func (e *Enum) Allowed() []string { return e.allowed }

func (e *Enum) Set(v string) error {
	if slices.Contains(e.allowed, v) {
		e.value = v
		return nil
	}
	return fmt.Errorf("must be one of: %s", strings.Join(e.allowed, ", "))
}

// Register wires an enum flag onto cmd in one call: declares the flag with
// validation and registers a shell-completion handler that suggests the
// allowed values.
func Register[T ~string](cmd *cobra.Command, name string, allowed []T, usage string) {
	cmd.Flags().Var(New(allowed), name, usage)
	_ = cmd.RegisterFlagCompletionFunc(name, completion(allowed))
}

func completion[T ~string](allowed []T) func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
	s := make([]string, len(allowed))
	for i, v := range allowed {
		s[i] = string(v)
	}
	return func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return s, cobra.ShellCompDirectiveNoFileComp
	}
}

// EnumValues returns the allowed list for an enum flag, or nil if the named
// flag was not declared via this package. Used by the schema dumper.
func EnumValues(f *pflag.Flag) []string {
	if av, ok := f.Value.(AllowedValuer); ok {
		return av.Allowed()
	}
	return nil
}
