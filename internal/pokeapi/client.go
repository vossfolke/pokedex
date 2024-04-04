package pokeapi

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

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

func (c *Client) InspectPokemon(name string) (*Pokemon, error) {
	pkmn, ok := c.pokedex.pkmn[name]
	if !ok {
		return &Pokemon{}, errors.New(fmt.Sprintf("No Informations... you have to catch %s", name))
	}
	return &pkmn, nil
}

func (c *Client) ShowPokedex() []string {
	pokemon := make([]string, 0, len(c.pokedex.pkmn))
	for pkmn := range c.pokedex.pkmn {
		pokemon = append(pokemon, pkmn)
	}
	return pokemon
}
