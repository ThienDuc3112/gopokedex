package main

import "fmt"

func commandExplore(c *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("missing name of location")
	}
	name := args[0]
	location, err := c.PokeClient.GetLocation(name)
	if err != nil {
		return fmt.Errorf("error getting location: %v", err)
	}

	fmt.Printf("Exploring %v...\n", name)
	fmt.Println("Found pokemon:")
	for _, v := range location.PokemonEncounters {
		fmt.Printf(" - %v\n", v.Pokemon.Name)
	}

	return nil
}
