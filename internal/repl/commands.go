package repl

import (
    "context"
    "fmt"
    "os"

    "bd-pokedex-go/internal/api"
)

type Command struct {
    Name        string
    Description string
    Callback    func(*api.Config) error
}

func commandExit(c *api.Config) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandHelp(c *api.Config) error {
    fmt.Println()
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    fmt.Println()
    for _, command := range GetCommands() {
        fmt.Printf("%s: %s\n", command.Name, command.Description)
    }
    return nil
}

func commandMap(c *api.Config) error {
    client := api.NewClient(c)
    return client.ListLocationAreas(context.Background())
}

func commandMapb(c *api.Config) error {
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
