package main

import (
	"time"

	"github.com/vossfolke/pokedex/internal/pokeapi"

)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
