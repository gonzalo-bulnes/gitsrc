// Package git provides primitives to manage the Git commit.template configuration.
package git

import (
	"os/exec"
)

// CLI allows to operate the system's Git command-line interface.
type CLI struct {
	clone func(remote string) error
	fetch func() error
	pull  func() error
}

// NewCLI returns a new CLI that wraps the Git command-line interface
func NewCLI() *CLI {
	return &CLI{
		clone: _clone,
		fetch: _fetch,
		pull:  _pull,
	}
}

// Clone clones a repository from a remote.
func (cli *CLI) Clone(remote string) error {
	return cli.clone(remote)
}

// Fetch removes local Git commit template configuration.
func (cli *CLI) Fetch() (err error) {
	return cli.fetch()
}

// Pull configures Git locally to use a commit template.
func (cli *CLI) Pull() error {
	return cli.pull()
}

func _clone(remote string) (err error) {
	cmd := exec.Command("git", "clone", remote)
	return cmd.Run()

}

func _fetch() error {
	cmd := exec.Command("git", "fetch", "origin")
	return cmd.Run()
}

func _pull() error {
	cmd := exec.Command("git", "pull", "origin", "master")
	return cmd.Run()
}
