package main

import (
	"errors"
	"fmt"
	"os"
)

func commandHelp(cfg *config) error {
	fmt.Printf("\nWelcome to the Pokedex Help Menu\n")
	fmt.Printf("Here are the available commands:\n")

	commandList := createCommandList()
	for _, command := range commandList {
		fmt.Printf(" - %s: %s\n", command.name, command.description)
	}
	fmt.Printf("\n")

	return nil
}

func commandExit(cfg *config) error {
	fmt.Println("Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(cfg *config) error {

	response, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Printf("\nLocation Areas:\n")

	for _, area := range response.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	fmt.Printf("\n")

	cfg.nextLocationAreaURL = response.Next
	cfg.previousLocationAreaURL = response.Previous
	return nil
}

func commandMapBack(cfg *config) error {

	if cfg.previousLocationAreaURL == nil {
		return errors.New("you're on the first page")
	}

	response, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Printf("\nLocation Areas:\n")

	for _, area := range response.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	fmt.Printf("\n")

	cfg.nextLocationAreaURL = response.Next
	cfg.previousLocationAreaURL = response.Previous
	return nil
}
