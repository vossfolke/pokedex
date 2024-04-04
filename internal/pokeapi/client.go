package pokeapi

import (
	"net/http"
	"time"
	"sync"

	"github.com/vossfolke/pokedex/internal/pokecache"
)

// Client -
type Client struct {
	pokedex    Pokedex
	cache      pokecache.Cache
	httpClient http.Client
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		pokedex: Pokedex{pkmn: make(map[string]Pokemon), mu: &sync.Mutex{}},
		cache:   pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
