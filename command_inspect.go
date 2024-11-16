package main

import "fmt"

func commandInspect(c *Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a pokemon name to inspect")
	}
	if pokemon, ok := c.Pokedex[args[0]]; ok {
		fmt.Printf("Name: %v\nHeight: %v\nWeight: %v\n", pokemon.Name, pokemon.Height, pokemon.Weight)

		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("\t- %v: %v\n", stat.Stat.Name, stat.BaseStat)
		}

		fmt.Println("Types:")
		for _, pokeType := range pokemon.Types {
			fmt.Printf("\t- %v\n", pokeType.Type.Name)
		}
	} else {
		fmt.Println("You have not caught that pokemon yet")
	}
	return nil
}
