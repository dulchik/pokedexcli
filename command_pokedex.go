package main

import "fmt"

func commandPokedex(c *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, p := range c.caughtPokemon {
		fmt.Printf(" - %s\n", p.Name)
	}
	return nil
}
