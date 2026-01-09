package main

import (
	"os"
	"fmt"
)

func init() {
	registry = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:		"help",
			description:"Displays a help message",
			callback:	commandHelp,
		},
		"map": {
			name: 		"map",
			description:"List 20 location areas from the Pokemon world",
			callback:	commandMap,
		},
		"mapb": {
			name: 		"mapb",
			description:"List 20 previous location areas from the Pokemon world",
			callback:	commandMapb,
		},
	}
}

func commandExit(cfg *Config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range registry {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func commandMap(cfg *Config) error {
	locationsResp, err := cfg.pokeClient.ListLocations(cfg.nextURL)
	if err != nil {
		return err
	}

	cfg.nextURL = locationsResp.Next
	cfg.previousURL = locationsResp.Previous

    for _, loc := range locationsResp.Results {
        fmt.Println(loc.Name)
    }

    return nil
}

func commandMapb(cfg *Config) error {
	if cfg.previousURL == nil {
		fmt.Println("You're on the first page!")
		return nil
	}

	locationsResp, err := cfg.pokeClient.ListLocations(cfg.previousURL)
	if err != nil {
		return err
	}

	cfg.nextURL = locationsResp.Next
	cfg.previousURL = locationsResp.Previous

    for _, loc := range locationsResp.Results {
        fmt.Println(loc.Name)
    }

    return nil
}