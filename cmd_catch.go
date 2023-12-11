package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("no pokemon was provided")
	}

	pokemon := args[0]

	resp, err := cfg.pokeapiClient.GetPokemon(pokemon)
	if err != nil {
		return err
	}

	const threshold = 50
	randNum := rand.Intn(resp.BaseExperience)
	fmt.Println(resp.BaseExperience, randNum, threshold)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemon)
	}

	cfg.caughtPokemons[pokemon] = resp
	fmt.Printf("%s got catched and added to pokedex\n", pokemon)

	return err
}
