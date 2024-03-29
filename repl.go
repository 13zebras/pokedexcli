package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {

	// Greeting
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Enter 'help' to see a list of commands\n\n")

	// scanner setup
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		text := scanner.Text()

		cleanedText := cleanInput(text)
		if len(cleanedText) == 0 {
			continue
		}

		commandName := cleanedText[0]
		args := []string{}
		if len(cleanedText) > 1 {
			args = cleanedText[1:]
		}

		commandList := createCommandList()

		command, ok := commandList[commandName]
		if !ok {
			fmt.Printf("Invalid command: %v\nTry again\n\n", command)
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func createCommandList() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			name:        "catch  {pokemon_name}",
			description: "Attempt to catch a pokemon and add it to your pokedex",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore  {location_area}",
			description: "Lists the pokemon in a location area",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"inspect": {
			name:        "catch  {pokemon_name}",
			description: "View information about a caught pokemon",
			callback:    commandInspect,
		},
		"map": {
			name:        "map",
			description: "List next page of location areas",
			callback:    commandMap,
		},
		"mapback": {
			name:        "mapback",
			description: "Previous page location areas",
			callback:    commandMapBack,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View all the pokemon in your pokedex",
			callback:    commandPokedex,
		},
	}
}

func cleanInput(input string) []string {
	lowerCaseInput := strings.ToLower(input)
	words := strings.Fields(lowerCaseInput)
	return words
}
