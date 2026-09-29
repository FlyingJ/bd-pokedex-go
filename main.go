package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"bd-pokedex-go/internal/api"
	"bd-pokedex-go/internal/pokecache"
	"bd-pokedex-go/internal/repl"
)

func main() {
	// create user input scanner
	scanner := bufio.NewScanner(os.Stdin)
	// configure our pokedex
	config := api.NewConfig()
	// prepare external request cache
	cacheTimeout := 5 * time.Second
	cache := pokecache.NewCache(cacheTimeout)

	for ;; {
		fmt.Print("Pokedex > ")
		if ok := scanner.Scan(); !ok {
			fmt.Println("Error: input scan error")
		}

		input := repl.CleanInput(scanner.Text())
		keyword := input[0]
		if command, exists := repl.GetCommands()[keyword]; exists {
			if err := command.Callback(config, cache); err != nil {
				fmt.Errorf("Error: %v", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
