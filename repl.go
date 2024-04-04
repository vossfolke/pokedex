package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vossfolke/pokedex/internal/pokeapi"
)

type cmdCli struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Pokedex >")
		reader.Scan()
		words := cleanInput(reader.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue

		} else {
			fmt.Println("unknown command")
			continue
		}

	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]cmdCli {
	return map[string]cmdCli{
		"help": {
			name:        "help",
			description: "Informations about the usage.",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore <location>",
			description: "Explore the Pokemons in an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Catching a Pokemon by chance",
			callback:    commandCatch,
		},
	}
}
