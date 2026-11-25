# pokedexcli
Boot Dev Project Building PokeDex cmdline tool in git

A command-line Pokédex built in Go that interacts with the [PokeAPI](https://pokeapi.co/) to explore Pokémon locations, inspect details, and catch wild Pokémon.

## Overview

This project is an interactive Read-Eval-Print Loop (REPL) application designed to fetch and display data from PokeAPI. To ensure fast performance and minimize unnecessary network requests, it includes an in-memory caching system with automatic entry expiration (reaping) running on background goroutines.

## Key Features

- **Interactive REPL**: A command-driven console interface with commands like `help`, `exit`, `map`, `mapb`, `explore`, `catch`, `inspect`, and `pokedex`.
- **Location Exploration**: Paginated navigation across various map areas and the ability to inspect wild Pokémon inhabiting those locations.
- **Catching & Pokedex Storage**: Attempt to catch Pokémon using a probability model based on base experience, storing successful captures in a local Pokédex map.
- **Custom In-Memory Cache**: Thread-safe caching mechanism implemented from scratch using Go's `sync.Mutex` and `time.Ticker` to automatically evict stale HTTP responses.

## Technical Highlights

- **Language**: Go
- **Concurrency**: Goroutines, channels, and mutexes for safe concurrent cache access and periodic cache cleanup.
- **Networking**: HTTP client interactions with JSON decoding and error handling.
- **Testing**: Unit tests covering cache operations and reap loops.
