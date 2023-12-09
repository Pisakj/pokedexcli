package main

import "fmt"

func callbackHelp(cfg *config) error {
	fmt.Println("=====================================")
	fmt.Println("--- Welcome to pokedex help menu! ---")
	fmt.Println("--- Here are available commands:  ---")

	commands := getCommands()

	for _, cmd := range commands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}
