package config_test

import (
    "bd-pokedex-go/internal/config"
    "testing"
)

func TestInit(t *testing.T) {
    var config config.Config
    config.Init()

    if config.Next != "https://pokeapi.co/api/v2/location-areas/" {
        t.Errorf("config.Next is not set to call PokeAPI")
    }

    if config.Previous != nil {
        t.Errorf("config.Previous is not nil")
    }
}