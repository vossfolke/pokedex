package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {
	pkmn, err := cfg.pokeapiClient.InspectPokemon(args[0])
	if err != nil {
		fmt.Printf("Name: %s\n", args[0])
		fmt.Println("Stats: unknown")
		fmt.Println("Types: unknown")
		return err
	}
	fmt.Printf("Name: %s\n", pkmn.Name)
	fmt.Printf("Height: %d\n", pkmn.Height)
	fmt.Printf("Weight: %d\n", pkmn.Weight)
	fmt.Println("Stats:")
	for _, stat := range pkmn.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range pkmn.Types {
		fmt.Printf("  - %v\n", typ.Type.Name)
	}
	return nil
}
