package commands

import "github.com/spf13/cobra"

// tiCmd is the parent command for technical indicators. The generator emits
// one command file per OpenAPI tag and routes the technical_indicator tag's
// operations under tiCmd via the api.mustache template's x-cli-is-ti branch.
// Static so it survives regeneration; protected by .openapi-generator-ignore.
var tiCmd = &cobra.Command{
	Use:   "ti",
	Short: "Technical indicators",
}

func init() {
	rootCmd.AddCommand(tiCmd)
}
