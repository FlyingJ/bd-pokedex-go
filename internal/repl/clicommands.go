package repl

import (
    "errors"
    "fmt"
    "os"
)

type cliCommand struct {
    name        string
    description string
    callback    func() error
}

var CLICommandRegistry = map[string]cliCommand{
    "exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    commandExit,
    },
}

func commandExit() error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return errors.New("Error: no idea how to get here")
}