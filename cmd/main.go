// Command gitsrc allows to set and manage co-author declarations
// for seamless gitsrcing sessions using Git and Github.
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/gonzalo-bulnes/gitsrc"
)

type config struct {
	gitsrc string
}

type argumentsError struct{}

func (e *argumentsError) Error() string {
	return `Usage:

gitsrc update <file>
Example:
  gitsrc update all.csv

gitsrc stop
`
}

func main() {
	err := checkArgs(os.Args)
	if err != nil {
		os.Exit(fail(os.Stderr, err, 0))
	}

	git := gitsrc.GetGitConnector()

	switch os.Args[1] {
	case "--version":
		_ = gitsrc.Version(os.Stdout, os.Stderr)
		os.Exit(0)
	case "update":
		err = gitsrc.Update(git, os.Stdout, os.Stderr)
		if err != nil {
			os.Exit(fail(os.Stderr, err, 21))
		}
		os.Exit(0)
	default:
	}
}

func checkArgs(args []string) error {
	if len(args) == 3 && args[1] == "update" {
		return nil
	}
	if len(args) == 2 && args[1] == "--version" {
		return nil
	}
	return &argumentsError{}
}

func fail(errors io.Writer, err error, code int) int {
	if e, ok := err.(*argumentsError); ok {
		fmt.Fprint(errors, e)
		return code
	}
	var version bytes.Buffer
	gitsrc.Version(&version, errors)

	fmt.Fprintf(errors, `
Oh no! You might have found a bug in gitsrc!

Please open an issue mentioning (error %d) and maybe we can pair to fix it : )
https://github.com/gonzalo-bulnes/gitsrc/issues

%serror: %v
`, code, version.String(), err)
	return code
}
