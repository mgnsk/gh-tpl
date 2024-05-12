package main

import (
	"fmt"
	"os"

	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/cli/go-gh/v2/pkg/template"
)

func main() {
	if len(os.Args) != 2 {
		os.Stderr.WriteString("expected template argument")
		os.Exit(1)
	}

	ios, _, _, _ := iostreams.Test()
	t := template.New(os.Stdout, ios.TerminalWidth(), ios.ColorEnabled())

	if err := t.Parse(os.Args[1]); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("error parsing template: %s\n", err.Error()))
		os.Exit(1)
	}

	if err := t.Execute(os.Stdin); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("error executing template: %s\n", err.Error()))
		os.Exit(1)
	}
}
