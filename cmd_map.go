package main

import (
	"fmt"
	"log"

	"github.com/pisakj/pokedexcli/internal/pokeapi"
)

func callbackMap() error {

	pokeClient := pokeapi.NewClient()

	resp, err := pokeClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("--- List of available locations: ---")
	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}

	return nil
}
