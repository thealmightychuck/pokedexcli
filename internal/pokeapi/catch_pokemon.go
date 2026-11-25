package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

type Captured_Pokemon struct {
	Pokemon_name string `json:"name"`
	Exp          int    `json:"base_experience"`
	Height int 	`json:"height"`
	Weight int `json:"weight"`
	Stats []BaseStat `json:"stats"`
	Type []Types `json:"types"`
}
// curl -s https://pokeapi.co/api/v2/pokemon/pidgey | jq '.stats[].stat.name'
type BaseStat struct {
	Stat_value int `json:"base_stat"`
	Stat_name struct {
		Name string `json:"name"`
	} `json:"stat"`
}

// curl -s https://pokeapi.co/api/v2/pokemon/pidgey | jq '.types[].type.name'
type Types struct {
	Type_name struct {
		Name string `json:"name"`
	} `json:"type"`
}


func (c *Client) CatchPokemon(name string) (Captured_Pokemon, error) {
	baseURL := "https://pokeapi.co/api/v2/pokemon/"
	url := baseURL + name

	// Check the cache
	val, ok := c.cache.Get(url)
	if ok {
		pokemonResp := Captured_Pokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return Captured_Pokemon{}, err
		}
		return pokemonResp, nil
	}

	// Start request
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return Captured_Pokemon{}, err
	}
	defer resp.Body.Close()

	var response Captured_Pokemon
	if resp.StatusCode > 299 {
		return Captured_Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Captured_Pokemon{}, fmt.Errorf("unable to read response body: %v", err)
	}

	// serial raw byte data into go structs
	if err := json.Unmarshal(body, &response); err != nil {
		return Captured_Pokemon{}, err
	}
	c.cache.Add(url, body)
	return response, nil
}
