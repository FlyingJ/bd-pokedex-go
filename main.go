package main

import (
	"bufio"
	"fmt"
	"os"
	"bd-pokedex-go/internal/repl"
)

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