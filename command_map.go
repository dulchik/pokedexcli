package main

import (
	"fmt"
	"errors"
)

func commandMapf(c *config, args ...string) error {
	res, err := c.pokeapiClient.ListLocations(c.nextLocationsURL)
	if err != nil {
		return err
	}

	c.nextLocationsURL = res.Next
	c.prevLocationsURL = res.Previous

	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(c *config, args ...string) error {
	if c.prevLocationsURL == nil {
		return errors.New("your're on the first page")
	}

	res, err := c.pokeapiClient.ListLocations(c.prevLocationsURL)
	if err != nil {
		return err
	}

	c.nextLocationsURL = res.Next
	c.prevLocationsURL = res.Previous

	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
