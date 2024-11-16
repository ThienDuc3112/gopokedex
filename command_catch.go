package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("please provide a pokemon name to catch")
	}
	if _, ok := config.Pokedex[args[0]]; ok {
		fmt.Printf("%v already in your pokedex\n", args[0])
		return nil
	}
	pokemon, err := config.PokeClient.GetPokemon(args[0])
	if err != nil {
		return fmt.Errorf("error getting pokemon: %v", err)
	}

	catch := false

	if pokemon.BaseExperience < 100 {
		catch = true
	} else if pokemon.BaseExperience < 200 {
		catch = rand.Intn(2) == 0
	} else {
		catch = rand.Intn(2) == 0
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", args[0])
	if catch {
		config.Pokedex[args[0]] = pokemon
		fmt.Printf("%v was caught!\n", args[0])
	} else {
		fmt.Printf("%v escaped!\n", args[0])
	}

	return nil
}
