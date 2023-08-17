package main

import "fmt"

func callbackHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	commands := getCliCommands()

	for _, cmd := range commands {
		fmt.Printf("%v: %v", cmd.name, cmd.description)
		fmt.Println("")
	}
	fmt.Println("")
	return nil
}
