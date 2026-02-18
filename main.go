package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"bd-pokedex-go/internal/repl"
)

type cliCommand struct {
    name        string
    description string
    callback    func() error
}

var cliCommands = map[string]cliCommand{
    "exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    commandExit,
    },
    "help": {
    	name:		 "help",
    	description: "Displays a help message",
    	callback:	 commandHelp,
    },
}

func commandExit() error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return errors.New("Error: no idea how to get here")
}

func print_usage() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for key, value := range cliCommands {
		fmt.Println("%s: %s", key, value.description)
	}
	return nil
}

func commandHelp() error {
	return print_usage()
}

func main() {
	// create user input scanner
	scanner := bufio.NewScanner(os.Stdin)

	for ;; {
		fmt.Print("Pokedex > ")
		if ok := scanner.Scan(); !ok {
			fmt.Println("Error: input scan error")
		}

		input := repl.CleanInput(scanner.Text())
		keyword := input[0]
		var command cliCommand
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