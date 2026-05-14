package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/twelvedata/twelvedata-cli/internal/auth"
	"github.com/twelvedata/twelvedata-cli/internal/flagx"
	"github.com/twelvedata/twelvedata-cli/internal/output"
)

// oneRequiredAnnotation mirrors the unexported constant in cobra/flag_groups.go.
// Stable across cobra v1.x.
const oneRequiredAnnotation = "cobra_annotation_one_required"

// promptMissingFlags fills in missing required flags interactively before
// Cobra's own ValidateRequiredFlags / ValidateFlagGroups runs. In raw mode or
// off a TTY this is a no-op — Cobra still emits the usual "required flag(s)"
// usage error.
//
// Two annotation kinds are handled:
//
//  1. Single required flags marked via MarkFlagRequired.
//  2. One-required groups marked via MarkFlagsOneRequired — the user picks
//     which flag in the group to supply, then is prompted for its value.
//
// Collected values are applied with Flags().Set() so the flag is marked
// Changed, satisfying downstream validation.
func promptMissingFlags(cmd *cobra.Command) error {
	if !shouldPrompt(cmd) {
		return nil
	}
	if err := promptSingleRequired(cmd); err != nil {
		return err
	}
	return promptOneRequiredGroups(cmd)
}

func shouldPrompt(cmd *cobra.Command) bool {
	if output.IsRaw(cmd) {
		return false
	}
	return auth.IsInteractive()
}

func promptSingleRequired(cmd *cobra.Command) error {
	var missing []*pflag.Flag
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		ann, ok := f.Annotations[cobra.BashCompOneRequiredFlag]
		if !ok || len(ann) == 0 || ann[0] != "true" {
			return
		}
		if !f.Changed {
			missing = append(missing, f)
		}
	})
	for _, f := range missing {
		if err := promptFlagValue(cmd, f); err != nil {
			return err
		}
	}
	return nil
}

func promptOneRequiredGroups(cmd *cobra.Command) error {
	// group key (space-joined flag names) -> any member already Changed?
	groups := map[string]bool{}
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		for _, g := range f.Annotations[oneRequiredAnnotation] {
			if _, ok := groups[g]; !ok {
				groups[g] = false
			}
			if f.Changed {
				groups[g] = true
			}
		}
	})
	for g, satisfied := range groups {
		if satisfied {
			continue
		}
		names := strings.Fields(g)
		chosen, err := pickGroupFlag(names)
		if err != nil {
			return err
		}
		f := cmd.Flags().Lookup(chosen)
		if f == nil {
			return fmt.Errorf("unknown flag %q in one-required group", chosen)
		}
		if err := promptFlagValue(cmd, f); err != nil {
			return err
		}
	}
	return nil
}

func pickGroupFlag(names []string) (string, error) {
	opts := make([]huh.Option[string], 0, len(names))
	for _, n := range names {
		opts = append(opts, huh.NewOption(n, n))
	}
	var v string
	err := huh.NewSelect[string]().
		Title("Pick one of the required options").
		Options(opts...).
		Value(&v).
		Run()
	if err != nil {
		return "", err
	}
	return v, nil
}

// promptFlagValue dispatches on the flag's type: enums use a Select, bools a
// Confirm, everything else a free-form Input whose result is fed through
// Flags().Set() so int/float parsing errors surface naturally.
func promptFlagValue(cmd *cobra.Command, f *pflag.Flag) error {
	desc := strings.TrimSpace(f.Usage)
	title := "Value for --" + f.Name

	if allowed := flagx.EnumValues(f); len(allowed) > 0 {
		opts := make([]huh.Option[string], 0, len(allowed))
		for _, a := range allowed {
			opts = append(opts, huh.NewOption(a, a))
		}
		var v string
		err := huh.NewSelect[string]().
			Title(title).
			Description(desc).
			Options(opts...).
			Value(&v).
			Run()
		if err != nil {
			return err
		}
		return cmd.Flags().Set(f.Name, v)
	}

	if f.Value.Type() == "bool" {
		var v bool
		err := huh.NewConfirm().
			Title(title).
			Description(desc).
			Affirmative("true").
			Negative("false").
			Value(&v).
			Run()
		if err != nil {
			return err
		}
		if v {
			return cmd.Flags().Set(f.Name, "true")
		}
		return cmd.Flags().Set(f.Name, "false")
	}

	var v string
	err := huh.NewInput().
		Title(title).
		Description(desc).
		Validate(func(s string) error {
			if strings.TrimSpace(s) == "" {
				return errors.New("value is required")
			}
			return nil
		}).
		Value(&v).
		Run()
	if err != nil {
		return err
	}
	return cmd.Flags().Set(f.Name, strings.TrimSpace(v))
}
