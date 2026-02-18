package repl

import (
	"testing"
)

func TestCommandHelp(t *testing.T) {
	if err := commandHelp(); err != nil {
		t.Errorf("Error: commandHelp returned non-nil: %v", err)
	}
}

func TestGetCommands(t *testing.T) {
	commands := GetCommands()

	if _, exists := commands["bogus"]; exists {
		t.Errorf("Error: bogus command should not exist")
	}

	if helpCommand, exists := commands["help"]; !exists {
		t.Errorf("Error: help command should exist")
	} else {
		if err := helpCommand.Callback(); err != nil {
			t.Errorf("Error: help command callback did not return nil")
		}
	}
}