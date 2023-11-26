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
	callback    func() error
}

func startRepl() {
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
		availCommands := getCommands()
		command, ok := availCommands[commandName]

		if !ok {
			fmt.Println("invalid command")
			continue
		}

		command.callback()

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
	}
}
