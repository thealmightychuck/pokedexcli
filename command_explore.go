package main

import (
	"fmt"
	"errors"
)

func commandExplore(cfg *config, args ...string) error {
	fmt.Printf("This command will display pokemon in an area\n")
	if len(args) == 0 {
		return errors.New("A location arguement is required to pass to commandExplore")
	}
	
	name := args[0]
	
	resp, err := cfg.client.ExplorePokemon(name)

	if err != nil {
		return err
	}
	
	for _, item := range resp.PokemonEncounters {
		fmt.Println(item.Pokemon.Name)
	}
	return nil
}
