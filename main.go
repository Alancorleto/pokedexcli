package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	pokeapiinteractions "github.com/alancorleto/pokedexcli/internal/poke_api_interactions"
	pokecache "github.com/alancorleto/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
	config      *config
}

type config struct {
	Next     string
	Previous string
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
			config:      &config{},
		},
		"help": {
			name:        "help",
			description: "Show help message",
			callback:    commandHelp,
			config:      &config{},
		},
		"map": {
			name:        "map",
			description: "Show locations",
			callback:    pokeapiinteractions.Map,
			config:      &config{},
		},
		"mapb": {
			name:        "mapb",
			description: "Show previous locations",
			callback:    pokeapiinteractions.Mapb,
			config:      &config{},
		},
	}
}

func main() {
	pokecache.NewCache(5 * time.Second)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)
		if len(words) > 0 {
			processCommand(words[0])
		}
	}
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	if trimmed == "" {
		return []string{}
	}
	words := strings.Fields(trimmed)
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return words
}

func processCommand(command string) {
	if cmd, exists := commands[command]; exists {
		cmd.callback()
	} else {
		fmt.Printf("Unknown command: %s. Type 'help' for a list of commands.\n", command)
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, cmd := range commands {
		fmt.Printf("  %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
