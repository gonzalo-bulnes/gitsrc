// Package gitsrc provides primitives to manage co-author declarations in Git commit templates.
package gitsrc

import (
	"fmt"
	"io"

	"github.com/gonzalo-bulnes/gitsrc/git"
)

const version = "0.1.0" // adheres to semantic versioning

// GetGitConnector returns a Git CLI instance.
func GetGitConnector() git.Connector {
	return git.NewCLI()
}

// Update pulls the latest changes from the default remote.
func Update(cli git.Connector, out, errors io.Writer) error {
	return fmt.Errorf("not yet implemented")
}

// Version prints the package version.
func Version(out, errors io.Writer) error {
	fmt.Fprintf(out, fmt.Sprintf("gitsrc version %s\n", version))
	return nil
}
