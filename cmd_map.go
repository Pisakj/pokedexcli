package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config) error {

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	fmt.Println("--- List of available locations: ---")
	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}

	cfg.nextLocationURL = resp.Next
	cfg.previousLocationURL = resp.Previous

	return nil
}

func callbackMapb(cfg *config) error {
	if cfg.previousLocationURL == nil {
		return errors.New("you are on the first page, use map command instead")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationURL)
	if err != nil {
		return err
	}

	fmt.Println("--- List of available locations: ---")
	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}

	cfg.nextLocationURL = resp.Next
	cfg.previousLocationURL = resp.Previous

	return nil
}
