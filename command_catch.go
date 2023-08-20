package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCacth(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon provided to be caught")
	}
	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokeman(pokemonName)
	if err != nil {
		return err
	}
	const threshold = 20
	randNum := rand.Intn(pokemon.BaseExperience)
	fmt.Println(pokemon.BaseExperience, threshold, randNum)
	if randNum > threshold {
		return fmt.Errorf("Sorry failed to catch %v", pokemon.Name)
	}

	cfg.caughtPokemon[pokemon.Name] = pokemon

	fmt.Printf("Great! You caught %s", pokemon.Name)
	return nil
}
