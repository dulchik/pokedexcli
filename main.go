package main

import (
	"time"

	"github.com/dulchik/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, time.Minute*5)
	c := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(c)
}
