package commands

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/twelvedata/twelvedata-go/twelvedata"

	"github.com/twelvedata/twelvedata-cli/internal/output"
)

// wantsCSV reports whether --output csv is set. ResolveFormat has already run
// successfully in rootCmd.PersistentPreRunE by the time RunE executes, so an
// error here would be a programming bug — surface it via panic rather than
// silently falling back to JSON and masking the typo.
func wantsCSV(cmd *cobra.Command) bool {
	f, err := output.ResolveFormat(cmd)
	if err != nil {
		panic(err)
	}
	return f == output.FormatCSV
}

// csvFormat returns a pointer to FORMATENUM_CSV when --output csv is set, nil
// otherwise. Per-command code uses it like:
//
//	if f := commands.csvFormat(cmd); f != nil { req = req.Format(*f) }
func csvFormat(cmd *cobra.Command) *twelvedata.FormatEnum {
	if !wantsCSV(cmd) {
		return nil
	}
	v := twelvedata.FORMATENUM_CSV
	return &v
}

// joinChunks joins items with sep, breaking to lineSep after every perLine
// items so a long list wraps cleanly in the help block.
func joinChunks(items []string, perLine int, sep, lineSep string) string {
	var b strings.Builder
	for i, s := range items {
		if i > 0 {
			if i%perLine == 0 {
				b.WriteString(lineSep)
			} else {
				b.WriteString(sep)
			}
		}
		b.WriteString(s)
	}
	return b.String()
}
