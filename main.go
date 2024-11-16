package main

import (
	"time"

	"github.com/ThienDuc3112/gopokedex.git/pokedex/internal/pokeapi"
)

func main() {
	config := Config{
		PokeClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
		Next:       nil,
		Previous:   nil,
		Pokedex:    make(map[string]pokeapi.Pokemon),
	}

	startRepl(&config)
}
