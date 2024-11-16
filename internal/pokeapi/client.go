package pokeapi

import (
	"net/http"
	"time"

	pokecache "github.com/ThienDuc3112/gopokedex.git/pokedex/internal/pokeCache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout time.Duration, interval time.Duration) *Client {
	return &Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(interval),
	}
}
