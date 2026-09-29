package api

import (
    "context"
    "encoding/json"
    "fmt"
    "io"
    "net/http"

    "bd-pokedex-go/internal/pokecache"
)

// PokeAPI returns a few reusable response shapes across many endpoints.
type APIResource struct {
    URL string `json:"url"`
}

type NamedAPIResource struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}

type APIResourceList[T any] struct {
    Count    int    `json:"count"`
    Next     string `json:"next"`
    Previous string `json:"previous"`
    Results  []T    `json:"results"`
}

func PrintNames(resources []NamedAPIResource) {
    for _, r := range resources {
        fmt.Println(r.Name)
    }
}

type Client struct {
    Cfg        *Config
    HttpClient *http.Client
}

type Endpoint struct {
    Named bool
    Path  string
}

var endpoints = map[string]Endpoint{
    "location":      {Path: "/location", Named: true},
    "location-area": {Path: "/location-area", Named: true},
    "pokemon":       {Path: "/pokemon", Named: true},
    "region":        {Path: "/region", Named: true},
}

func NewClient(config *Config) *Client {
    return &Client{
        Cfg:        config,
        HttpClient: http.DefaultClient,
    }
}

func GetJSON[T any](ctx context.Context, cache *pokecache.Cache, client *Client, url string) (*T, error) {
	bytes, ok := cache.Get(url)
	if ok {
		fmt.Println("CACHE HIT:", url)
	} else {
		fmt.Println("CACHE MISS:", url)
		// and we fetch
    	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	    if err != nil {
	        return nil, fmt.Errorf("create request: %w", err)
	    }

	    res, err := client.HttpClient.Do(req)
	    if err != nil {
	        return nil, fmt.Errorf("perform request: %w", err)
	    }
	    defer res.Body.Close()

	    if res.StatusCode > 299 {
	        return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	    }

	    if bytes, err = io.ReadAll(res.Body); err != nil {
	    	return nil, fmt.Errorf("unable to read response body in fetch: %w", err)
	    }
	    cache.Add(url, bytes)
	}

	var val T
	if err := json.Unmarshal(bytes, &val); err != nil {
		return nil, fmt.Errorf("decode response body: %w", err)
	}
    return &val, nil
}

/*
REFACTOR THESE
*/
func (c *Client) ListLocationAreas(ctx context.Context, cache *pokecache.Cache) error {
    res, err := GetJSON[APIResourceList[NamedAPIResource]](ctx, cache, c, c.Cfg.Next)
    if err != nil {
        return fmt.Errorf("error retrieving json: %w", err)
    }
    c.Cfg.Previous = res.Previous
    c.Cfg.Next = res.Next
    PrintNames(res.Results)
    return nil
}
/*
REFACTOR THESE
*/
func (c *Client) ListLocationAreasBack(ctx context.Context, cache *pokecache.Cache) error {
    res, err := GetJSON[APIResourceList[NamedAPIResource]](ctx, cache, c, c.Cfg.Previous)
    if err != nil {
        return fmt.Errorf("error retrieving json: %w", err)
    }
    c.Cfg.Previous = res.Previous
    c.Cfg.Next = res.Next
    PrintNames(res.Results)
    return nil
}
/*
REFACTOR THESE
*/
