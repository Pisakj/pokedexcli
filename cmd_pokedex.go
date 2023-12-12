package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {

	if len(cfg.caughtPokemons) == 0 {
		return fmt.Errorf("your pokedex is empty")
	}

	fmt.Println("Your pokedex:")
	for _, pokemon := range cfg.caughtPokemons {
		fmt.Printf("- %s\n", pokemon.Name)
	}

	return nil
}
