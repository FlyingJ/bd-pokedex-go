package api_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"bd-pokedex-go/internal/api"
	"bd-pokedex-go/internal/pokecache"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func TestGetJSON(t *testing.T) {
	cfg := api.Config{
		Next: "https://example.test/api/v2/location-area/",
	}

	client := api.NewClient(&cfg)
	client.HttpClient = &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			if r.Method != http.MethodGet {
				t.Fatalf("unexpected method: %s", r.Method)
			}

			if r.URL.String() != cfg.Next {
				t.Fatalf("unexpected URL: %s", r.URL.String())
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body: io.NopCloser(strings.NewReader(
					`{"count":1,"next":"","previous":"","results":[{"name":"canalave-city-area","url":"https://pokeapi.co/api/v2/location-area/1/"}]}`,
				)),
				Header: make(http.Header),
			}, nil
		}),
	}

	// prepare external request cache
	cacheTimeout := 5 * time.Second
	cache := pokecache.NewCache(cacheTimeout)

	err := client.ListLocationAreas(context.Background(), cache)
	if err != nil {
		t.Fatalf("ListLocationAreas returned error: %v", err)
	}

	if client.Cfg.Next != "" {
		t.Fatalf("client.Cfg.Next not updated: %s", client.Cfg.Next)
	}

	if client.Cfg.Previous != "" {
		t.Fatalf("client.Cfg.Previous not empty: %s", client.Cfg.Previous)
	}
}
