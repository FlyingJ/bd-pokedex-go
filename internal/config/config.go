package config

type Config struct {
    Next string
    Previous string
}

func (c *Config) Init() error {
	c.Next = "https://pokeapi.co/api/v2/location-areas/"
	c.Previous = nil
	return nil
}