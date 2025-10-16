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
	callback    func(...string) error
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Show locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show previous locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List the names of pokemon found at a specific location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon by name",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught pokemon by name",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught pokemon",
			callback:    commandPokedex,
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
			processCommand(words[0], words[1:]...)
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

func processCommand(command string, args ...string) {
	if cmd, exists := commands[command]; exists {
		err := cmd.callback(args...)
		if err != nil {
			fmt.Printf("Error executing command %s: %v\n", command, err)
		}
	} else {
		fmt.Printf("Unknown command: %s. Type 'help' for a list of commands.\n", command)
	}
}

func commandExit(args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, cmd := range commands {
		fmt.Printf("  %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(args ...string) error {
	return pokeapiinteractions.Map()
}

func commandMapb(args ...string) error {
	return pokeapiinteractions.Mapb()
}

func commandExplore(args ...string) error {
	return pokeapiinteractions.Explore(args[0])
}

func commandCatch(args ...string) error {
	return pokeapiinteractions.Catch(args[0])
}

func commandInspect(args ...string) error {
	return pokeapiinteractions.Inspect(args[0])
}

func commandPokedex(args ...string) error {
	return pokeapiinteractions.Pokedex()
}
