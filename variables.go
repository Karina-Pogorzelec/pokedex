package main
import "github.com/Karina-Pogorzelec/pokedex/internal/pokeapi"

type Config struct {
	pokeClient  pokeapi.Client
	nextURL     *string
	previousURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

var registry map[string]cliCommand