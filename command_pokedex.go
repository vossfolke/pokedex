package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	pokedex := cfg.pokeapiClient.ShowPokedex()
	fmt.Println("Your Pokedex: ")
	if len(pokedex) == 0 {
		fmt.Println("No Entries... Go catch some Pokemon!")
		return nil
	}
	for i := 0; i < len(pokedex); i++ {
		fmt.Printf("  - %s\n", pokedex[i])
	}
	return nil
}
