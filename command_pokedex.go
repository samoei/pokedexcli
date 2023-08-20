package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {

	fmt.Println("All Pokemon Caught:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("Name: %s\n", pokemon.Name)
	}

	return nil
}
