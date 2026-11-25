package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

// json formatted struct on the pokemon_encounters array result from location-area/<area-name>
type PokemonEncounters struct {
	PokemonEncounters []PokemonIDs `json:"pokemon_encounters"`
}

type PokemonIDs struct {
	Pokemon Encounters `json:"pokemon"`
}

type Encounters struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
// Similar to Location_List.go. Uses the already built http client.
func (c *Client) ExplorePokemon(name string) (PokemonEncounters, error) {

	baseURL := "https://pokeapi.co/api/v2/"
	url := baseURL + "location-area/" + name


	// Check the Cache for the URL string before making the request. 
	// if url Exists, unmarshall pokemon encounters data
	val, ok := c.cache.Get(url)
	if ok {
		pokemonResp := PokemonEncounters{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return PokemonEncounters{}, err
		}
		return pokemonResp, nil
	}

	//Start request.
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return PokemonEncounters{}, err
	}
	defer resp.Body.Close()

	var response PokemonEncounters
	if resp.StatusCode > 299 {
		return PokemonEncounters{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonEncounters{}, fmt.Errorf("unable to read response body: %v", err)
	}
	
	// serial raw byte data into go structs
	if err := json.Unmarshal(body, &response); err != nil {
		return PokemonEncounters{}, err
	}

	c.cache.Add(url, body)
	return response, nil
}