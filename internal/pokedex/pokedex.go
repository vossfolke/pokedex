package pokedex

import (
	"fmt"

	"github.com/vossfolke/pokedex/internal/pokeapi"
)

type Pokedex struct {
	pkmn map[string]pokeapi.Pokemon
}

func (p *Pokedex) NewPokedex() Pokedex {
	pokedex := Pokedex{pkmn: make(map[string]pokeapi.Pokemon)}
	return pokedex
}

func (p *Pokedex) Add(name string, pokemon pokeapi.Pokemon) string {
	_, ok := p.pkmn[name]
	if ok {
		return fmt.Sprintf("%s is already in your Pokedex...")
	}
	p.pkmn[name] = pokemon
	return fmt.Sprintf("%s is added to your Pokedex!")
}
