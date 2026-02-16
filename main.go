package main

import (
	"bufio"
	"fmt"
	"os"
	"bd-pokedex-go/repl"
)

type cliCommand struct {
	name		string
	description string
	callback	func() error
}

map[string]cliCommand{
	"exit": {
		name:		 "exit",
		description: "Exit the Pokedex",
		callback:	 commandExit,
	},
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
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
		

	}
}