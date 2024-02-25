package main

import (
	"fmt"
)

func main() {
	commandList := createCommandList()

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Enter 'help' to see a list of commands")
	fmt.Println(commandList)
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func createCommandList() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
