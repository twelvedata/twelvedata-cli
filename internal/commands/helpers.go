package commands

import (
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
