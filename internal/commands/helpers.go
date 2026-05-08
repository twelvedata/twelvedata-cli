package commands

import (
	"github.com/spf13/cobra"
	"github.com/twelvedata/twelvedata-go/twelvedata"

	"github.com/twelvedata/twelvedata-cli/internal/output"
)

// enumValue reads the string value of a flag declared via flagx.Register.
// pflag's GetString does not work for custom Value types, so we go through
// Lookup().Value.String() instead.
func enumValue(cmd *cobra.Command, name string) string {
	f := cmd.Flags().Lookup(name)
	if f == nil {
		return ""
	}
	return f.Value.String()
}

func wantsCSV(cmd *cobra.Command) bool {
	f, err := output.ResolveFormat(cmd)
	if err != nil {
		return false
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

// flagChanged reports whether the named flag was explicitly set by the user.
// Used to decide whether to apply a setter to the SDK builder; unset flags must
// not become explicit empty/zero values on the wire.
func flagChanged(cmd *cobra.Command, name string) bool {
	return cmd.Flags().Changed(name)
}
