package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	//fmt.Printf("This command will Catch pokemon\n")
	if len(args) == 0 {
		return errors.New("A pokemon is required to pass to commandCatch\n")
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", args[0])
	resp, err := cfg.client.CatchPokemon(args[0])

	if err != nil {
		return err
	}


	roll := rand.Intn(700)
	if roll > resp.Exp {

		fmt.Printf("Captured %v\n", resp.Pokemon_name)
		// resp is the entire Captured_Pokemon struct.
		cfg.pokedex[resp.Pokemon_name] = resp
	} else {
		fmt.Printf("%v escaped\n", args[0])
	}
	return nil

}
