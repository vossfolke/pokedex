package main

import (
	"fmt"
)

func commandExplore(cfg *config, name *string) error {

	pkmnAreaResp, err := cfg.pokeapiClient.PokemonArea(*name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s\n", *name)
	fmt.Println("Found Pokemon:")
	for _, pkmn := range pkmnAreaResp.PokemonEncounters {
		fmt.Printf("- %s\n", pkmn.Pokemon.Name)
	}
	return nil
}
