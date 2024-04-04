package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {

	pokemon, err := cfg.pokeapiClient.CatchPokemon(args[0])
	if err != nil {
		return err
	}

	// add catch func
	maxBaseExp := 608.0

	chance := (1 - (float64(pokemon.BaseExperience))/maxBaseExp) * rand.Float64()

	fmt.Printf("A wild %s appeared!\n", args[0])
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	if chance < 0.5 {
		fmt.Printf("%s escaped!\n", args[0])
	} else {
		fmt.Printf("%s was caught!\n", args[0])
		
	}
	 return nil
}
