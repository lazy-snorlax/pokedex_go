package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	// Check cache for result
	cacheEntry, exists := c.cache.Get(url)
	if exists {
		pokemonResp := Pokemon{}
		// fmt.Println("Pulling from cache")
		err := json.Unmarshal(cacheEntry, &pokemonResp)
		if err != nil {
			return Pokemon{}, nil
		}
		return pokemonResp, nil
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

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonResp := Pokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, dat)

	return pokemonResp, nil
}
