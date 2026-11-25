package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {

	//fmt.Printf("This command will inspect pokemon that you've caught")

	val, ok := cfg.pokedex[args[0]]
	if ok {
		fmt.Printf("Name: %v\n Height: %v\n Weight: %v\n Stats:\n",
			args[0], val.Height, val.Weight)
		for _, stat := range val.Stats {
			fmt.Printf("\t%v: %v\n", stat.Stat_name.Name, stat.Stat_value)
		}
		fmt.Printf("Types:\n")
		for _, pokemonType := range val.Type {
			fmt.Printf("\t%v", pokemonType.Type_name.Name)
		}
	}

	if !ok {
		fmt.Printf("\nyou have not caught that pokemon: %v \n", args[0])
	}

	//fmt.Printf("\nyou have not caught that pokemon %v", args[0])
	return nil
}
