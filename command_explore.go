package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {

	pkmnAreaResp, err := cfg.pokeapiClient.PokemonArea(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring: %s\n", args[0])
	fmt.Println("Found Pokemon:")
	for _, pkmn := range pkmnAreaResp.PokemonEncounters {
		fmt.Printf("- %s\n", pkmn.Pokemon.Name)
	}
	return nil
}
