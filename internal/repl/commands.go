package repl

import (
    "context"
    "fmt"
    "os"

    "bd-pokedex-go/internal/api"
    "bd-pokedex-go/internal/pokecache"
)

type Command struct {
    Name        string
    Description string
    Callback    func(*api.Config, *pokecache.Cache) error
}

func commandExit(config *api.Config, cache *pokecache.Cache) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandHelp(config *api.Config, cache *pokecache.Cache) error {
    fmt.Println()
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    fmt.Println()
    for _, command := range GetCommands() {
        fmt.Printf("%s: %s\n", command.Name, command.Description)
    }
    return nil
}

func commandMap(config *api.Config, cache *pokecache.Cache) error {
    client := api.NewClient(config)
    return client.ListLocationAreas(context.Background(), cache)
}

func commandMapb(config *api.Config, cache *pokecache.Cache) error {
    if config.Previous == "" {
        fmt.Println("You are on the first page")
        return nil
    }
    client := api.NewClient(config)
    return client.ListLocationAreasBack(context.Background(), cache)
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
