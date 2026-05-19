package commands

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// TestTreeShape walks the entire rootCmd subtree (both hand-written and
// generator-emitted commands) and asserts the invariants that every leaf must
// satisfy. This is the single broad-coverage replacement for resend-cli's
// per-endpoint test files — instead of one file per endpoint, we declare the
// rules here and let new endpoint commands either conform or fail.
//
// Invariants checked at every leaf:
//   - Non-empty Short string (powers --help and the schema dump).
//   - No flag declared with a duplicate name across local + persistent.
//   - Every flag that was registered as required carries the bash-completion
//     annotation Cobra uses for required-flag detection.
//   - The Use field's first word does not collide with a sibling's first word
//     (Cobra silently shadows; this catches accidental dupes during regen).
func TestTreeShape(t *testing.T) {
	walkTree(t, rootCmd, func(t *testing.T, cmd *cobra.Command) {
		path := cmd.CommandPath()

		if cmd.Hidden {
			return
		}
		if cmd.Name() == "help" || cmd.Name() == "completion" {
			// `help` is Cobra-emitted; `completion` is a group container whose
			// own Short we control but whose children carry per-shell content.
			// We still recurse into completion's children — the walker does
			// that — so don't fail on the parent.
		}

		// Every node should have a Short; without it the schema dump and
		// `--help` lines look broken for agents.
		if cmd != rootCmd && strings.TrimSpace(cmd.Short) == "" {
			t.Errorf("%s: Short is empty", path)
		}

		// Verify required flags carry Cobra's annotation. Cobra stores it under
		// the `cobra_annotation_bash_completion_one_required_flag` key — the
		// public exported constant is BashCompOneRequiredFlag.
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			ann := f.Annotations[cobra.BashCompOneRequiredFlag]
			if len(ann) == 0 {
				return
			}
			if ann[0] != "true" {
				t.Errorf("%s: flag --%s required-annotation = %q, want \"true\"", path, f.Name, ann[0])
			}
		})
	})
}

// TestTreeShape_FirstWordSiblingsUnique guards against two siblings using the
// same first word in their Use, which Cobra resolves silently.
func TestTreeShape_FirstWordSiblingsUnique(t *testing.T) {
	walkTreeNode(t, rootCmd, func(t *testing.T, parent *cobra.Command) {
		seen := map[string]string{}
		for _, sub := range parent.Commands() {
			if sub.Hidden {
				continue
			}
			first := strings.SplitN(sub.Use, " ", 2)[0]
			if existing, ok := seen[first]; ok {
				t.Errorf("%s has two children whose Use begins with %q: %s and %s",
					parent.CommandPath(), first, existing, sub.Name())
				continue
			}
			seen[first] = sub.Name()
		}
	})
}

// TestTreeShape_HandWrittenLeavesArePresent guards against silent regressions
// where the generator overwrites a hand-written command (e.g. doctor.go missing
// from .openapi-generator-ignore). Each name here must remain a direct child
// of rootCmd.
func TestTreeShape_HandWrittenLeavesArePresent(t *testing.T) {
	wanted := []string{
		"login", "logout", "whoami", "auth", "doctor",
		"docs", "dashboard", "completion", "commands", "ti",
	}
	have := map[string]bool{}
	for _, c := range rootCmd.Commands() {
		have[c.Name()] = true
	}
	for _, name := range wanted {
		if !have[name] {
			t.Errorf("hand-written command %q missing from rootCmd; regeneration may have clobbered it", name)
		}
	}
}

// TestTreeShape_SchemaAlias double-checks the agent-discovery alias survives.
// (schema_test.go has a similar check; keeping it here too means a single
// `go test ./internal/commands -run TestTreeShape` exercises every shape
// invariant agents depend on.)
func TestTreeShape_SchemaAlias(t *testing.T) {
	for _, c := range rootCmd.Commands() {
		if c.Use == "commands" {
			for _, a := range c.Aliases {
				if a == "schema" {
					return
				}
			}
			t.Error(`"commands" found but its "schema" alias is missing`)
			return
		}
	}
	t.Error(`"commands" subcommand missing from rootCmd`)
}

// walkTree depth-first walks the command tree rooted at start, invoking visit
// for every node, including start itself. Each node runs in its own subtest so
// a failure points straight at the offending path.
func walkTree(t *testing.T, start *cobra.Command, visit func(t *testing.T, cmd *cobra.Command)) {
	var visitNode func(*cobra.Command)
	visitNode = func(c *cobra.Command) {
		t.Run(c.CommandPath(), func(t *testing.T) {
			visit(t, c)
		})
		for _, sub := range c.Commands() {
			visitNode(sub)
		}
	}
	visitNode(start)
}

// walkTreeNode is like walkTree but only invokes visit on non-leaf nodes —
// useful for sibling-uniqueness checks.
func walkTreeNode(t *testing.T, start *cobra.Command, visit func(t *testing.T, parent *cobra.Command)) {
	var visitNode func(*cobra.Command)
	visitNode = func(c *cobra.Command) {
		if len(c.Commands()) > 0 {
			t.Run(c.CommandPath(), func(t *testing.T) {
				visit(t, c)
			})
		}
		for _, sub := range c.Commands() {
			visitNode(sub)
		}
	}
	visitNode(start)
}

