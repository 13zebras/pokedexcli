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

		commandList := createCommandList()

		command, ok := commandList[commandName]
		if !ok {
			fmt.Printf("Invalid command: %v\nTry again\n\n", command)
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func createCommandList() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
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
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func cleanInput(input string) []string {
	lowerCaseInput := strings.ToLower(input)
	words := strings.Fields(lowerCaseInput)
	return words
}
