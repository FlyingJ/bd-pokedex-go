package repl_test

import (
	"testing"
	"time"

	"bd-pokedex-go/internal/api"
	"bd-pokedex-go/internal/pokecache"
	"bd-pokedex-go/internal/repl"
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
	config := api.NewConfig()
	reapTime := 5 * time.Second
	cache := pokecache.NewCache(reapTime)
	commandName := "help"
	if err := repl.GetCommands()[commandName].Callback(config, cache); err != nil {
		t.Errorf("command %s returned non-nil: %v", commandName, err)
	}
}
