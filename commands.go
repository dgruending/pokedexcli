package main

import (
	"fmt"
	"os"

	"github.com/dgruending/pokedexcli/internal/api"
)

type configuration struct {
	previous string
	next     string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*configuration) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Lists the next 20 Locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the previous 20 Locations",
			callback:    commandMapB,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandExit(cfg *configuration) error {
	fmt.Println("Exiting Pokedex!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *configuration) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")

	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

func commandMap(cfg *configuration) error {
	locations, next, _, err := api.GetLocations(cfg.next)
	if err != nil {
		return err
	}
	cfg.previous = cfg.next
	cfg.next = next
	printLocations(locations)
	return nil
}

func commandMapB(cfg *configuration) error {
	if cfg.previous == "" {
		fmt.Println("No previous page of Locations")
		return fmt.Errorf("Call without previous URL")
	}
	locations, next, prev, err := api.GetLocations(cfg.previous)
	if err != nil {
		return err
	}
	cfg.next = next
	cfg.previous = prev
	printLocations(locations)
	return nil
}

func printLocations(locations []api.Location) {
	for _, location := range locations {
		fmt.Printf("%s\n", location.Name)
	}
}
