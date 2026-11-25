package main

import (
	"github.com/thealmightychuck/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

// config struct to store command type information to avoid more memory allocation, by recalling the getCommands function.
// interacted with by passing around memory address
type config struct {
	commands map[string]cliCommand
	client   *pokeapi.Client
	next     *string
	previous *string
	pokedex map[string]pokeapi.Captured_Pokemon
}

// - Every callback receives that exact same pointer, giving them access to the shared data without copying anything.
// - The map contains reference info, with the keys being strings and the values being the struct.
// - the functions in the struct contain the memory address locations. They're not executed as they're not wrapped in parathesis.
// - Callback is a value which references the memory address where the function exists.
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Gets Next PokeAPI location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Gets Previous PokeAPI location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Display all pokemon found in an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "This command will allow you to view pokemon stats",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "This command will display all pokemon you've caught",
			callback:    commandPokedex,
		},

	}
}
