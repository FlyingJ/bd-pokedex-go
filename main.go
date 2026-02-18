package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	r "bd-pokedex-go/internal/repl"
)

func main() {
	// create user input scanner
	scanner := bufio.NewScanner(os.Stdin)

	for ;; {
		fmt.Print("Pokedex > ")
		if ok := scanner.Scan(); !ok {
			fmt.Println("Error: input scan error")
		}

		input := r.CleanInput(scanner.Text())
		keyword := input[0]
		cliCommands := r.GetCliCommands()
		var ok bool
		command, ok = cliCommands[keyword]
		if ok {
			if err := command.callback(); err != nil {
				fmt.Errorf("Error: %v", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}