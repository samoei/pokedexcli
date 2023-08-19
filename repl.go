package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
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

		availableCommands := getCliCommands()

		cmd, ok := availableCommands[commandName]

		if !ok {
			fmt.Println("Invalid command!")
			continue
		}

		cmd.callback()

	}
}

type cliCommands struct {
	name        string
	description string
	callback    func() error
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
			description: "Displays some of the locations",
			callback:    callbackMap,
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
