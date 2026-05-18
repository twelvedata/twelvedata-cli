package commands

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"

	"github.com/twelvedata/twelvedata-cli/internal/auth"
)

// completionMarker tags the snippet `twelvedata completion install` writes into a
// shell profile so we don't append a second copy on re-runs.
const completionMarker = "# twelvedata-cli shell completion"

var completionShells = []string{"bash", "zsh", "fish", "powershell"}

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate or install the shell completion script",
	Long: `Generate or install the shell completion script for twelvedata.

Use a per-shell subcommand to print the script to stdout — pipe it into
your shell, or save it where the shell expects completion files.

` + "`twelvedata completion install [shell]`" + ` writes the script and wires it
into the shell's profile in one step. The shell is auto-detected from
$SHELL when omitted.`,
	Example: strings.TrimSpace(`
  twelvedata completion install               # auto-detect $SHELL
  twelvedata completion install bash          # explicit
  source <(twelvedata completion bash)        # one-shot, current shell only
  twelvedata completion fish > ~/.config/fish/completions/twelvedata.fish`),
}

var completionBashCmd = &cobra.Command{
	Use:                   "bash",
	Short:                 "Generate the bash completion script",
	DisableFlagsInUseLine: true,
	Long: `Generate the bash completion script.

Source it for the current shell:

  source <(twelvedata completion bash)

Or add it permanently:

  # Linux
  echo 'eval "$(twelvedata completion bash)"' >> ~/.bashrc

  # macOS
  echo 'eval "$(twelvedata completion bash)"' >> ~/.bash_profile

Requires bash-completion v2.`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		return rootCmd.GenBashCompletionV2(cmd.OutOrStdout(), !noDesc(cmd))
	},
}

var completionZshCmd = &cobra.Command{
	Use:                   "zsh",
	Short:                 "Generate the zsh completion script",
	DisableFlagsInUseLine: true,
	Long: `Generate the zsh completion script.

Source it for the current shell:

  source <(twelvedata completion zsh)

Or save it where compinit will pick it up:

  twelvedata completion zsh > "${fpath[1]}/_td"

Then start a new shell.`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		if noDesc(cmd) {
			return rootCmd.GenZshCompletionNoDesc(cmd.OutOrStdout())
		}
		return rootCmd.GenZshCompletion(cmd.OutOrStdout())
	},
}

var completionFishCmd = &cobra.Command{
	Use:                   "fish",
	Short:                 "Generate the fish completion script",
	DisableFlagsInUseLine: true,
	Long: `Generate the fish completion script.

Source it for the current shell:

  twelvedata completion fish | source

Or save it where fish loads completions automatically:

  twelvedata completion fish > ~/.config/fish/completions/twelvedata.fish`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		return rootCmd.GenFishCompletion(cmd.OutOrStdout(), !noDesc(cmd))
	},
}

var completionPowerShellCmd = &cobra.Command{
	Use:                   "powershell",
	Short:                 "Generate the PowerShell completion script",
	DisableFlagsInUseLine: true,
	Long: `Generate the PowerShell completion script.

Load it for the current session:

  twelvedata completion powershell | Out-String | Invoke-Expression

Or persist it in your PowerShell profile:

  Add-Content $PROFILE 'twelvedata completion powershell | Out-String | Invoke-Expression'`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		if noDesc(cmd) {
			return rootCmd.GenPowerShellCompletion(cmd.OutOrStdout())
		}
		return rootCmd.GenPowerShellCompletionWithDesc(cmd.OutOrStdout())
	},
}

var completionInstallCmd = &cobra.Command{
	Use:   "install [bash|zsh|fish|powershell]",
	Short: "Install completion into your shell profile",
	Long: `Write the completion script and wire it into the shell's profile.

If no shell is given, twelvedata auto-detects from $SHELL; on an interactive
terminal it falls back to prompting. Already-installed completions are
detected via a marker line and not duplicated.`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		shell, err := resolveCompletionShell(args)
		if err != nil {
			return err
		}
		return installCompletion(cmd.OutOrStdout(), shell)
	},
}

func noDesc(cmd *cobra.Command) bool {
	v, _ := cmd.Flags().GetBool("no-descriptions")
	return v
}

func resolveCompletionShell(args []string) (string, error) {
	if len(args) == 1 {
		s := strings.ToLower(args[0])
		if slices.Contains(completionShells, s) {
			return s, nil
		}
		return "", fmt.Errorf("unsupported shell %q (want one of: %s)", args[0], strings.Join(completionShells, ", "))
	}
	if detected := detectShell(); detected != "" {
		return detected, nil
	}
	if !auth.IsInteractive() {
		return "", errors.New("could not detect shell; pass one of: " + strings.Join(completionShells, ", "))
	}
	opts := make([]huh.Option[string], 0, len(completionShells))
	for _, s := range completionShells {
		opts = append(opts, huh.NewOption(s, s))
	}
	var v string
	if err := huh.NewSelect[string]().
		Title("Which shell do you use?").
		Options(opts...).
		Value(&v).
		Run(); err != nil {
		return "", err
	}
	return v, nil
}

func detectShell() string {
	if sh := os.Getenv("SHELL"); sh != "" {
		switch filepath.Base(sh) {
		case "bash":
			return "bash"
		case "zsh":
			return "zsh"
		case "fish":
			return "fish"
		}
	}
	if os.Getenv("PSModulePath") != "" {
		return "powershell"
	}
	return ""
}

func installCompletion(out io.Writer, shell string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("locate home dir: %w", err)
	}
	switch shell {
	case "bash":
		return installEvalLine(out, bashProfile(home), `eval "$(twelvedata completion bash)"`)
	case "powershell":
		return installEvalLine(out, powershellProfile(home), `twelvedata completion powershell | Out-String | Invoke-Expression`)
	case "zsh":
		return installZsh(out, home)
	case "fish":
		return installFish(out, home)
	}
	return fmt.Errorf("unsupported shell: %s", shell)
}

func bashProfile(home string) string {
	if runtime.GOOS == "darwin" {
		return filepath.Join(home, ".bash_profile")
	}
	return filepath.Join(home, ".bashrc")
}

func powershellProfile(home string) string {
	if runtime.GOOS == "windows" {
		return filepath.Join(home, "Documents", "PowerShell", "Microsoft.PowerShell_profile.ps1")
	}
	return filepath.Join(home, ".config", "powershell", "Microsoft.PowerShell_profile.ps1")
}

func installEvalLine(out io.Writer, path, line string) error {
	if existing, err := os.ReadFile(path); err == nil {
		if bytes.Contains(existing, []byte(completionMarker)) {
			fmt.Fprintf(out, "Completion already installed in %s\n", path)
			return nil
		}
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := fmt.Fprintf(f, "\n%s\n%s\n", completionMarker, line); err != nil {
		return err
	}
	fmt.Fprintf(out, "Completion installed in %s\n", path)
	fmt.Fprintln(out, "Restart your shell to activate.")
	return nil
}

func installZsh(out io.Writer, home string) error {
	completionDir := filepath.Join(home, ".zsh", "completions")
	if err := os.MkdirAll(completionDir, 0o755); err != nil {
		return err
	}
	target := filepath.Join(completionDir, "_td")
	var buf bytes.Buffer
	if err := rootCmd.GenZshCompletion(&buf); err != nil {
		return err
	}
	if err := os.WriteFile(target, buf.Bytes(), 0o644); err != nil {
		return err
	}
	fmt.Fprintf(out, "Completion written to %s\n", target)

	profile := filepath.Join(home, ".zshrc")
	if existing, err := os.ReadFile(profile); err == nil && bytes.Contains(existing, []byte(completionDir)) {
		fmt.Fprintf(out, "fpath entry already present in %s\n", profile)
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(profile), 0o755); err != nil {
		return err
	}
	f, err := os.OpenFile(profile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := fmt.Fprintf(f, "\n%s\nfpath=(%s $fpath)\nautoload -Uz compinit && compinit\n", completionMarker, completionDir); err != nil {
		return err
	}
	fmt.Fprintf(out, "Added fpath entry to %s\n", profile)
	fmt.Fprintln(out, "Restart your shell to activate.")
	return nil
}

func installFish(out io.Writer, home string) error {
	completionDir := filepath.Join(home, ".config", "fish", "completions")
	if err := os.MkdirAll(completionDir, 0o755); err != nil {
		return err
	}
	target := filepath.Join(completionDir, "twelvedata.fish")
	var buf bytes.Buffer
	if err := rootCmd.GenFishCompletion(&buf, true); err != nil {
		return err
	}
	if err := os.WriteFile(target, buf.Bytes(), 0o644); err != nil {
		return err
	}
	fmt.Fprintf(out, "Completion written to %s\n", target)
	fmt.Fprintln(out, "Open a new fish shell to activate.")
	return nil
}

func init() {
	// Replace Cobra's auto-registered completion command so we can hang the
	// `install` subcommand off it and own the Long descriptions.
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	for _, c := range []*cobra.Command{completionBashCmd, completionZshCmd, completionFishCmd, completionPowerShellCmd} {
		c.Flags().Bool("no-descriptions", false, "Disable completion descriptions")
	}
	completionCmd.AddCommand(completionBashCmd, completionZshCmd, completionFishCmd, completionPowerShellCmd, completionInstallCmd)
	rootCmd.AddCommand(completionCmd)
}
