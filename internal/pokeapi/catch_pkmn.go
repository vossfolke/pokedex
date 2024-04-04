package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	if entry, ok := c.cache.Get(url); ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(entry, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)
	return pokemon, nil
}

func (c *Client) AddToPokedex(pkmn *Pokemon) string {
	return c.pokedex.Add(pkmn.Name, *pkmn)
}
