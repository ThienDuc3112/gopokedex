package main

import "fmt"

func commandPokedex(c *Config, _ []string) error {
	fmt.Println("Your Pokedex:")
	for k := range c.Pokedex {
		fmt.Printf("\t- %v\n", k)
	}
	return nil
}
