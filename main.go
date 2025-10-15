package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

func main() {
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
