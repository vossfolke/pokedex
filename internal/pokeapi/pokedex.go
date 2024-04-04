package pokeapi

import (
	"fmt"
	"sync"
)

type Pokedex struct {
	pkmn map[string]Pokemon
	mu   *sync.Mutex
}

func (p *Pokedex) NewPokedex() Pokedex {
	pokedex := Pokedex{pkmn: make(map[string]Pokemon), mu: &sync.Mutex{}}
	return pokedex
}

func (p *Pokedex) Add(name string, pokemon Pokemon) string {
	p.mu.Lock()
	defer p.mu.Unlock()
	_, ok := p.pkmn[name]
	if ok {
		return fmt.Sprintf("%s is already in your Pokedex...", pokemon.Name)
	}
	p.pkmn[name] = pokemon
	return fmt.Sprintf("%s is added to your Pokedex!", pokemon.Name)
}
