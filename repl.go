package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cmdCli struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Pokedex >")
		reader.Scan()
		words := cleanInput(reader.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback()
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
			description: "help text",
			callback:    commandHelp,
		},
	}
}
