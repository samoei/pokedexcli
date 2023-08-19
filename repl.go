package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()

		text := scanner.Text()

		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		args := []string{}

		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCliCommands()

		cmd, ok := availableCommands[commandName]

		if !ok {
			fmt.Println("Invalid command!")
			continue
		}

		err := cmd.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}

	}
}

type cliCommands struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCliCommands() map[string]cliCommands {
	return map[string]cliCommands{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next page locations",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous page locations",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "Lists the pokemon in a location area",
			callback:    callbackExplore,
		},
		"exit": {
			name:        "exit",
			description: "Exits the pokedex",
			callback:    callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lower := strings.ToLower(str)
	words := strings.Fields(lower)
	return words
}
