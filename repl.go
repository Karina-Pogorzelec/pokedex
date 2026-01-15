package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"time"
	"github.com/Karina-Pogorzelec/pokedex/internal/pokeapi"
)

func startRepl() {
	cfg := &Config{
		pokeClient: pokeapi.NewClient(5 * time.Second, 5 * time.Second,),
	} 

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
			err := command.callback(cfg, cleanedText[1:])
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}


func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	return strings.Fields(lowerText)
}
