package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonArea(name string) (PokemonInArea, error) {
	url := baseURL + "/location-area/" + name

	if entry, ok := c.cache.Get(url); ok {
		pkmnAreaResp := PokemonInArea{}
		err := json.Unmarshal(entry, &pkmnAreaResp)
		if err != nil {
			return PokemonInArea{}, err
		}
		return pkmnAreaResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonInArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInArea{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonInArea{}, err
	}
	pkmnAreaResp := PokemonInArea{}
	err = json.Unmarshal(data, &pkmnAreaResp)
	if err != nil {
		return PokemonInArea{}, err
	}

	c.cache.Add(url, data)
	return pkmnAreaResp, nil
}
