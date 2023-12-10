package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("no location was provided")
	}

	locationName := args[0]

	resp, err := cfg.pokeapiClient.GetLocationArea(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("--- Pokemons found in %s: ---\n", locationName)
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return err
}
