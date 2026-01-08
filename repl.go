package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var registry map[string]cliCommand

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
	}
}

func commandExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")

	for _, cmd := range registry {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		text := scanner.Text()
		cleanedText := cleanInput(text)
		if len(cleanedText) == 0 {
    		continue
		}

		command, exists := registry[cleanedText[0]]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}


		
		//fmt.Printf("Your command was: %s\n", cleanedText[0])
	}
}


func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	return strings.Fields(lowerText)
}
