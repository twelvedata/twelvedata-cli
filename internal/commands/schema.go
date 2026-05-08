package commands

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/twelvedata/twelvedata-cli/internal/flagx"
)

// schemaNode is the agent-facing description of a command. Stable shape: agents
// can introspect the entire CLI tree as JSON without scraping --help text.
type schemaNode struct {
	Name        string       `json:"name"`
	Path        string       `json:"path"`
	Short       string       `json:"short,omitempty"`
	Long        string       `json:"long,omitempty"`
	Aliases     []string     `json:"aliases,omitempty"`
	Flags       []schemaFlag `json:"flags,omitempty"`
	Subcommands []schemaNode `json:"subcommands,omitempty"`
}

type schemaFlag struct {
	Name      string   `json:"name"`
	Shorthand string   `json:"shorthand,omitempty"`
	Type      string   `json:"type"`
	Default   string   `json:"default,omitempty"`
	Usage     string   `json:"usage,omitempty"`
	Required  bool     `json:"required,omitempty"`
	Inherited bool     `json:"inherited,omitempty"`
	Enum      []string `json:"enum,omitempty"`
}

var schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "Dump the full command tree as JSON for agent discovery",
	RunE: func(cmd *cobra.Command, args []string) error {
		root := buildSchema(rootCmd, "")
		enc := json.NewEncoder(cmd.OutOrStdout())
		enc.SetIndent("", "  ")
		return enc.Encode(root)
	},
}

func buildSchema(c *cobra.Command, parentPath string) schemaNode {
	path := c.Name()
	if parentPath != "" {
		path = parentPath + " " + c.Name()
	}
	node := schemaNode{
		Name:    c.Name(),
		Path:    path,
		Short:   c.Short,
		Long:    c.Long,
		Aliases: c.Aliases,
	}

	collect := func(fs *pflag.FlagSet, inherited bool) {
		fs.VisitAll(func(f *pflag.Flag) {
			required := false
			if anns, ok := f.Annotations[cobra.BashCompOneRequiredFlag]; ok {
				for _, v := range anns {
					if v == "true" {
						required = true
					}
				}
			}
			node.Flags = append(node.Flags, schemaFlag{
				Name:      f.Name,
				Shorthand: f.Shorthand,
				Type:      f.Value.Type(),
				Default:   f.DefValue,
				Usage:     f.Usage,
				Required:  required,
				Inherited: inherited,
				Enum:      flagx.EnumValues(f),
			})
		})
	}
	collect(c.LocalFlags(), false)
	collect(c.InheritedFlags(), true)

	for _, sub := range c.Commands() {
		if sub.Hidden || sub.Name() == "help" {
			continue
		}
		node.Subcommands = append(node.Subcommands, buildSchema(sub, path))
	}
	return node
}

func init() {
	rootCmd.AddCommand(schemaCmd)
}
