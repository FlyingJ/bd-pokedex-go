package main

import (
	"bufio"
	"fmt"
	"os"
	"bd-pokedex-go/internal/api"
	"bd-pokedex-go/internal/repl"
)

func main() {
	// create user input scanner
	scanner := bufio.NewScanner(os.Stdin)
	// configure our pokedex
	c := api.NewConfig()

	for ;; {
		fmt.Print("Pokedex > ")
		if ok := scanner.Scan(); !ok {
			fmt.Println("Error: input scan error")
		}

		input := repl.CleanInput(scanner.Text())
		keyword := input[0]
		if command, exists := repl.GetCommands()[keyword]; exists {
			if err := command.Callback(&c); err != nil {
				fmt.Errorf("Error: %v", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}