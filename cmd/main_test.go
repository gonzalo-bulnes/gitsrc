package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestCheckArgs(t *testing.T) {
	t.Run("rejects invalid arguments", func(t *testing.T) {
		knownInvalid := []struct {
			Description string
			Args        []string
		}{
			{"missing command or option", []string{"gitsrc"}},
			{"unknown command", []string{"gitsrc", "what?"}},
			{"too many arguments for '--version' option", []string{"gitsrc", "--version", "now"}},
			{"too many arguments for 'update' command", []string{"gitsrc", "update", "mutiple", "people"}},
			{"not enough arguments for 'update' command", []string{"gitsrc", "update"}},
		}

		for _, tc := range knownInvalid {
			t.Run(tc.Description, func(t *testing.T) {
				if err := checkArgs(tc.Args); err == nil {
					t.Errorf("Expected '%s' to be invalid arguments, but were found valid", tc.Args)
				}
			})
		}
	})

	t.Run("accepts valid arguments", func(t *testing.T) {
		knownInvalid := []struct {
			Description string
			Args        []string
		}{
			{"'--version' option", []string{"gitsrc", "--version"}},
			{"'update' command", []string{"gitsrc", "update", "gile"}},
		}

		for _, tc := range knownInvalid {
			t.Run(tc.Description, func(t *testing.T) {
				if err := checkArgs(tc.Args); err != nil {
					t.Errorf("Expected '%s' to be valid arguments, but were found invalid", tc.Args)
				}
			})
		}
	})
}

func TestErrorMessages(t *testing.T) {
	t.Run("with invalid arguments error", func(t *testing.T) {
		var message bytes.Buffer
		fail(&message, &argumentsError{}, 23)
		if !strings.Contains(message.String(), "Usage:") {
			t.Errorf("Expected usage instructon to be printed, got: %s", message.String())
		}
	})

	t.Run("with any other error", func(t *testing.T) {
		var message bytes.Buffer
		fail(&message, fmt.Errorf("Unspecificied error"), 23)
		if !strings.Contains(message.String(), "error 23") {
			t.Errorf("Expected error to be reported as such, got: %s", message.String())
		}
	})
}
