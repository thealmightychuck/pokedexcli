package pokeapi

// response data from get request formatted into a struct example
//"count": 1539,
//  "next": "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20",
//  "previous": null,
//  "results": [
//    {
//      "name": "canalave-city-area",
//     "url": "https://pokeapi.co/api/v2/location-area/1/"
//    },

type LocationAreasResponse struct {
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationName `json:"results"`
}
