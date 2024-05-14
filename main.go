package main

import (
	"fmt"
	"os"

	"github.com/Masterminds/sprig/v3"
	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/cli/go-gh/v2/pkg/template"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: gh-tpl <template>\n")
		os.Exit(1)
	}

	ios, _, _, _ := iostreams.Test()
	t := template.New(os.Stdout, ios.TerminalWidth(), ios.ColorEnabled()).Funcs(sprig.FuncMap())

	if err := t.Parse(os.Args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "error parsing template: %s\n", err.Error())
		os.Exit(1)
	}

	if err := t.Execute(os.Stdin); err != nil {
		fmt.Fprintf(os.Stderr, "error executing template: %s\n", err.Error())
		os.Exit(1)
	}
}
