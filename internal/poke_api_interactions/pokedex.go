package pokeapiinteractions

import (
	"fmt"
	"math/rand"
)

type PokemonResponse struct {
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Types  []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
}

var pokedex = make(map[string]PokemonResponse)

func Catch(pokemonName string) error {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonName)
	pokemon, err := Get[PokemonResponse](url)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	if rand.Intn(2) == 0 {
		fmt.Printf("Oh no! %s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("Congratulations! You caught a %s!\n", pokemon.Name)

	pokedex[pokemon.Name] = *pokemon

	return nil
}

func Inspect(pokemonName string) error {
	if pokemon, found := pokedex[pokemonName]; found {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)

		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf("- %s\n", t.Type.Name)
		}

		fmt.Println("Stats:")
		for _, s := range pokemon.Stats {
			fmt.Printf("- %s: %d\n", s.Stat.Name, s.BaseStat)
		}
	} else {
		fmt.Printf("You don't have a %s in your pokedex. Catch it first!\n", pokemonName)
	}
	return nil
}

func Pokedex() error {
	if len(pokedex) == 0 {
		fmt.Println("Your pokedex is empty. Catch some pokemon!")
		return nil
	}

	fmt.Println("Your Pokedex:")

	for name := range pokedex {
		fmt.Printf("- %s\n", name)
	}

	return nil
}
