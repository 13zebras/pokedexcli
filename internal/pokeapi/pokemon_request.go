package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ////////////////////////////////////
// GetPokemon
// ////////////////////////////////////
func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName

	fullURL := baseURL + endpoint

	// fmt.Println("fullURL:", fullURL)
	// check cache for fullURL
	// if found, return cached data
	data, ok := c.cache.Get(fullURL)
	if ok {
		// fmt.Println("cache hit")
		pokemonName := Pokemon{}
		err := json.Unmarshal(data, &pokemonName)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonName, nil
	}

	fmt.Println("no cached values")

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)

	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)

	if err != nil {
		return Pokemon{}, err
	}

	// Add data to cache before returning
	c.cache.Add(fullURL, data)

	return pokemon, nil

}
