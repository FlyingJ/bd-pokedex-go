package repl_test

import (
	"bd-pokedex-go/internal/repl"
	"testing"
)

func TestGetCommands(t *testing.T) {
	commands := repl.GetCommands()

	// ensure a 'bogus' command does not exist
	commandName := "bogus"
	if _, exists := commands[commandName]; exists {
		t.Errorf("command %s should not exist", commandName)
	}
	// ensure a legitimate command does exist
	commandName = "help"
	if _, exists := commands[commandName]; !exists {
		t.Errorf("command %s should exist", commandName)
	}
}

func TestCommandHelp(t *testing.T) {
	commandName := "help"
	if err := repl.GetCommands()[commandName].Callback(); err != nil {
		t.Errorf("command %s returned non-nil: %v", commandName, err)
	}
}