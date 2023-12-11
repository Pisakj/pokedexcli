package main

import (
	"time"

	"github.com/pisakj/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationURL     *string
	previousLocationURL *string
	caughtPokemons      map[string]pokeapi.Pokemon
}

func main() {

	cfg := config{
		pokeapiClient:  pokeapi.NewClient(time.Hour),
		caughtPokemons: make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cfg)
}
