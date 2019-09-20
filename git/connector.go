package git

// Connector allows to interact with the Git command-line interface.
type Connector interface {
	Clone(remote string) error
	Pull() error
	Fetch() error
}
