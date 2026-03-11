package repl

import (
    poke "bd-pokedex-go/internal/pokeapi"
    "fmt"
    "net/http"
    "os"
)

type Command struct {
    Name        string
    Description string
    Callback    func() error
}

func commandExit() error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandHelp() error {
    fmt.Println()
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    fmt.Println()
    for _, command := range GetCommands() {
        fmt.Printf("%s: %s\n", command.Name, command.Description)
    }
    return nil
}

func commandMap() error {
    return fmt.Errorf("not implemented")
}

func commandMapb() error {
    return fmt.Errorf("not implemented")
}

func GetCommands() map[string]Command {
    return map[string]Command{
        "exit": {
            Name:        "exit",
            Description: "Exit the Pokedex",
            Callback:    commandExit,
        },
        "help": {
            Name:        "help",
            Description: "Displays a help message",
            Callback:    commandHelp,
        },
        "map": {
            Name:        "map",
            Description: "Displays the next page of locations",
            Callback:    commandMap,
        },
        "mapb": {
            Name:        "mapb",
            Description: "Displays the previous page of locations",
            Callback:    commandMapb,
        },
    }
}
