package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}
	areaName := args[0]
	resp, err := cfg.pokeapiClient.GetLocationArea(areaName)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemon in: %s\n", resp.Name)

	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
