package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	// BaseExperience: higher for harder to catch pokemon,
	// lower for easier to catch pokemon

	// threshold is arbitrary number below which
	// we "catch" a pokemon
	const threshold = 50
	randomNumber := rand.Intn(pokemon.BaseExperience)
	fmt.Println(pokemon.BaseExperience, randomNumber, threshold)
	if randomNumber > threshold {
		fmt.Printf("Failed to catch %s!\n\n", pokemonName)
		return nil
	}

	fmt.Printf("%s pokemon was caught!\n\n", pokemonName)

	return nil
}

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Goodbye!")
	os.Exit(0)
	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area specified")
	}
	locationAreaName := args[0]

	locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}
	fmt.Printf("Pokemon in %s:\n", locationArea.Name)

	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	fmt.Printf("\n")

	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Printf("\nWelcome to the Pokedex Help Menu\n")
	fmt.Printf("Here are the available commands:\n")

	commandList := createCommandList()
	for _, command := range commandList {
		fmt.Printf(" - %s: %s\n", command.name, command.description)
	}
	fmt.Printf("\n")

	return nil
}

func commandMap(cfg *config, args ...string) error {

	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Printf("\nLocation Areas:\n")

	for _, area := range res.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	fmt.Printf("\n")

	cfg.nextLocationAreaURL = res.Next
	cfg.previousLocationAreaURL = res.Previous
	return nil
}

func commandMapBack(cfg *config, args ...string) error {

	if cfg.previousLocationAreaURL == nil {
		return errors.New("oops! you're on the first page")
	}

	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Printf("\nLocation Areas:\n")

	for _, area := range res.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	fmt.Printf("\n")

	cfg.nextLocationAreaURL = res.Next
	cfg.previousLocationAreaURL = res.Previous
	return nil
}
