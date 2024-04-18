package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, name *string) error {
	if name == nil || *name == "" {
		return errors.New("enter location name to explore")
	}
	exploreResp, err := cfg.pokeapiClient.ListPokemons(*name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", *name)
	fmt.Println("Found Pokemon: ")
	for _, enc := range exploreResp.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil
}
