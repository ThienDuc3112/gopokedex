package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ThienDuc3112/gopokedex.git/pokedex/internal/pokeapi"
)

type Config struct {
	PokeClient *pokeapi.Client
	Previous   *string
	Next       *string
	Pokedex    map[string]pokeapi.Pokemon
}

type CliCommand struct {
	name        string
	description string
	callback    func(*Config, []string) error
}

func cleanInput(input string) []string {
	return strings.Fields(strings.ToLower(input))
}

func getCommand() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore the location and list encountered pokemon(s)",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon and put it in your pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon in your pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all pokemons in your pokedex",
			callback:    commandPokedex,
		},
	}

}

func startRepl(config *Config) {
	reader := bufio.NewScanner(os.Stdin)
	commands := getCommand()
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		fmt.Println()
		c, exist := commands[words[0]]
		if !exist {
			fmt.Print("Invalid command\n\n")
			continue
		}

		err := c.callback(config, words[1:])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println()
	}
}
