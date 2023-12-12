package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommands struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		output := cleanInput(input)

		if len(output) == 0 {
			continue
		}

		commandName := output[0]
		fmt.Println(output)
		args := []string{}
		if len(output) > 1 {
			args = output[1:]
		}
		availCommands := getCommands()
		command, ok := availCommands[commandName]

		if !ok {
			fmt.Println("invalid command")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}

	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

func getCommands() map[string]cliCommands {
	return map[string]cliCommands{
		"help": {
			name:        "help",
			description: "Prints help menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit from pokedex",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "Next page of available locations",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Previous page of available location",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore {locationArea}",
			description: "Check the specific location",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon name}",
			description: "Attempt to catch a choosen pokemon",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect {pokemon name}",
			description: "Allowed to inspect a pokemon which is in your pokedex",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Check your pokemon list in your pokedex",
			callback:    callbackPokedex,
		},
	}
}
