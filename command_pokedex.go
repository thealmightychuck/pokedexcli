package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {

	fmt.Printf("Pokedex: \n")
	for _, pokemonName := range cfg.pokedex {
		fmt.Printf("- %v\n", pokemonName.Pokemon_name)
	}

	return nil
}
