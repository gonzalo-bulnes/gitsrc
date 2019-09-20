package gitsrc

import (
	"bytes"
	"path/filepath"
	"regexp"
	"testing"
)

// semverRegexp matches the most common semantinc version strings.
// See https://semver.org
const semverRegexp = `\s\d+.\d+.\d+(?:-(?:alpha|beta|rc)\d*)?\s`

func TestCmd(t *testing.T) {
	t.Run("Version()", func(t *testing.T) {

		var out, error bytes.Buffer
		err := Version(&out, &error)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !regexp.MustCompile(semverRegexp).Match(out.Bytes()) {
			t.Errorf("Expected (semantic) version to be put out, was not")
		}
	})
}

type MockGitConnector struct {
	CloneGlobal bool
	CloneError  error
	PullError   error
	FetchError  error
	CloneCount  int
	PullCount   int
	FetchCount  int
}

func (m *MockGitConnector) Clone(remote string) error {
	m.CloneCount++
	return m.CloneError
}
func (m *MockGitConnector) Pull() error {
	m.PullCount++
	return m.PullError
}
func (m *MockGitConnector) Fetch() error {
	m.FetchCount++
	return m.FetchError
}

func full(path string) string {
	return filepath.Join("testdata", path)
}
