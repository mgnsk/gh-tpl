package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/cli/go-gh/v2/pkg/jq"
	"github.com/cli/go-gh/v2/pkg/template"
	sprig "github.com/go-task/slim-sprig/v3"
	"github.com/spf13/cobra"
)

func main() {
	var (
		forceColor bool
		jqScript   string
	)

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

			var jsonReader io.Reader

			if jqScript != "" {
				buf := &bytes.Buffer{}

				if err := jq.Evaluate(c.InOrStdin(), buf, jqScript); err != nil {
					return fmt.Errorf("error executing jq expression: %w", err)
				}

				jsonReader = buf
			} else {
				jsonReader = c.InOrStdin()
			}

			if err := t.Execute(jsonReader); err != nil {
				return fmt.Errorf("error executing template: %w", err)
			}

			return nil
		},
	}

	root.Flags().BoolVar(&forceColor, "color", false, "Force color output")
	root.Flags().StringVar(&jqScript, "jq", "", "Filter JSON input using a jq expression")

	if err := root.Execute(); err != nil {
		root.PrintErrln(err)
		os.Exit(1)
	}
}
