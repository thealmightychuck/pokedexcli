package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

// LocationName struct which is a slice of Results. Each element hold it's own Name
type LocationName struct {
	Name string `json:"name"`
}

// Method to create the LocationsAreasResponse struct
// a Client can list location areas, given an optional page URL producing a response or an error
// Client is the actor, the response is the product
// used in map and mapb commands. pageURL is a pointer to string data from the cfg struct. 
// PageURL is passed even if it is empty
func (c *Client) GetPokeAPI(pageURL *string) (LocationAreasResponse, error) {

	// build and run request passing along the Next field
	url := "https://pokeapi.co/api/v2/location-area/"

	if pageURL != nil {
		url = *pageURL
	}

	// Check the Cache before making the request.
	val, ok := c.cache.Get(url)
	if ok {
		locationsResp := LocationAreasResponse{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationAreasResponse{}, err
		}
		//fmt.Printf("url was found in cache %v", url)
		return locationsResp, nil
	}

	// Start request after determining URL.
	//fmt.Printf("url was NOT found in cache %v", url)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer resp.Body.Close()

	var response LocationAreasResponse
	if resp.StatusCode > 299 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResponse{}, fmt.Errorf("unable to read response body: %v", err)
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return LocationAreasResponse{}, err
	}

	c.cache.Add(url, body)
	return response, nil
}
