package main

import "fmt"

// Decoder will find json key "next"
// tag tell it to use "Next" field

func commandMap(cfg *config, args ...string) error {
	// Use the apiClient to make a call to the pokeapi URL
	// Once I recieve the response as json. Unmarshall the data into a struct.
	resp, err := cfg.client.GetPokeAPI(cfg.next)
	cfg.next = resp.Next
	cfg.previous = resp.Previous
	
	if err != nil {
		return err
	}
	
	for _, item := range resp.Results {
		fmt.Println(item.Name)
	}

	
	//fmt.Printf("%v, %v", resp, err)
	//fmt.Printf("Running map command\n")
	return nil
}
