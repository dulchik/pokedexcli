package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		if command, ok :=  getCommands()[commandName]; ok {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name 		string
	description 	string
	callback	func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:		"exit",
			description:	"Exit the Pokedex",
			callback:	commandExit,
		},
		"help": {
			name: 		"help",
			description:	"Displays a help message",
			callback:	commandHelp,
		},
		"map": {
			name: 		"map",
			description: 	"Displays the (next) names of 20 locaion areas in the Pokemon world",
			callback: 	commandMap,
		},
		"mapb": {
			name: 		"mapb",
			description: 	"Displays the (previous) names of 20 locaion areas in the Pokemon world",
			callback: 	commandMapb,
		},
	}
}
