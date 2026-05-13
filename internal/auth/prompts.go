package auth

import (
	"errors"
	"os"

	"github.com/charmbracelet/huh"
	"golang.org/x/term"
)

// ErrNotInteractive is returned by prompt helpers when stdin/stdout is not a
// TTY. Callers translate this into a usage error pointing at the missing flag.
var ErrNotInteractive = errors.New("not running in an interactive terminal")

// IsInteractive reports whether both stdin and stdout are TTYs.
func IsInteractive() bool {
	return term.IsTerminal(int(os.Stdin.Fd())) && term.IsTerminal(int(os.Stdout.Fd()))
}

// PromptAPIKey reads a key in masked form. Errors with ErrNotInteractive on a
// non-TTY.
func PromptAPIKey() (string, error) {
	if !IsInteractive() {
		return "", ErrNotInteractive
	}
	var v string
	err := huh.NewInput().
		Title("Enter your Twelve Data API key").
		EchoMode(huh.EchoModePassword).
		Value(&v).
		Run()
	if err != nil {
		return "", err
	}
	return v, nil
}

// SelectProfile shows a single-select list of profile names. Invalid names are
// hinted but selectable, so users can rename legacy entries via `auth rename`.
func SelectProfile(prompt string, profiles []ProfileInfo) (string, error) {
	if !IsInteractive() {
		return "", ErrNotInteractive
	}
	if len(profiles) == 0 {
		return "", errors.New("no profiles configured")
	}
	opts := make([]huh.Option[string], 0, len(profiles))
	for _, p := range profiles {
		label := p.Name
		if p.Active {
			label += " (active)"
		}
		if ValidateProfileName(p.Name) != nil {
			label += " (invalid name)"
		}
		opts = append(opts, huh.NewOption(label, p.Name))
	}
	var v string
	err := huh.NewSelect[string]().
		Title(prompt).
		Options(opts...).
		Value(&v).
		Run()
	if err != nil {
		return "", err
	}
	return v, nil
}

// ConfirmDestructive shows a yes/no confirmation. Default is no.
func ConfirmDestructive(prompt string) (bool, error) {
	if !IsInteractive() {
		return false, ErrNotInteractive
	}
	var v bool
	err := huh.NewConfirm().
		Title(prompt).
		Affirmative("Yes").
		Negative("No").
		Value(&v).
		Run()
	if err != nil {
		return false, err
	}
	return v, nil
}

// PromptText asks for free-form text with optional validation.
func PromptText(prompt, placeholder string, validate func(string) error) (string, error) {
	if !IsInteractive() {
		return "", ErrNotInteractive
	}
	var v string
	in := huh.NewInput().
		Title(prompt).
		Placeholder(placeholder).
		Value(&v)
	if validate != nil {
		in = in.Validate(validate)
	}
	if err := in.Run(); err != nil {
		return "", err
	}
	return v, nil
}
