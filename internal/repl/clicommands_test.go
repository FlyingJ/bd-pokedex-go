package repl

import (
	"testing"
)

func TestCommandHelp(t *testing.T) {
	if err := commandHelp(); err != nil {
		t.Errorf("Error: commandHelp returned non-nil: %v", err)
	}
}