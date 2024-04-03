package main

import (
	"fmt"
	"time"

	"github.com/vossfolke/pokedex/internal/pokeapi"
)

func main() {
	fmt.Println("Starting Pokedex Client...")
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	fmt.Println("Setting up config...")
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
