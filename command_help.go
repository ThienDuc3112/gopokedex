package main

import (
	"fmt"
)

func commandHelp(*Config, []string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, c := range getCommand() {
		fmt.Printf("%v: %v\n", c.name, c.description)
	}
	return nil
}
