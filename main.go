package main

import (
	"fmt"
	"os"

	"github.com/Masterminds/sprig/v3"
	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/cli/go-gh/v2/pkg/template"
	"github.com/spf13/cobra"
)

func main() {
	var forceColor bool

	root := &cobra.Command{
		Use:        "gh-tpl [template]",
		Short:      "Render JSON using a go template.",
		Args:       cobra.ExactArgs(1),
		ArgAliases: []string{"template"},
		RunE: func(c *cobra.Command, args []string) error {
			ios := iostreams.System()
			if forceColor {
				ios.SetColorEnabled(true)
			}

			t := template.New(c.OutOrStdout(), ios.TerminalWidth(), ios.ColorEnabled()).Funcs(sprig.FuncMap())

			if err := t.Parse(args[0]); err != nil {
				return fmt.Errorf("error parsing template: %w", err)
			}

			if err := t.Execute(c.InOrStdin()); err != nil {
				return fmt.Errorf("error executing template: %w", err)
			}

			return nil
		},
	}

	root.Flags().BoolVar(&forceColor, "color", false, "Force color output")

	if err := root.Execute(); err != nil {
		root.PrintErrln(err)
		os.Exit(1)
	}
}
