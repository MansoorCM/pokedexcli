package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, name *string) error {
	if name == nil {
		return errors.New("enter pokemon name to catch")
	}

	_, ok := cfg.pokedex[*name]
	if ok {
		fmt.Printf("%s already caught\n", *name)
		return nil
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(*name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", *name)
	n := max(700, pokemon.BaseExperience+10)
	randomNum := rand.IntN(n)
	if randomNum > pokemon.BaseExperience {
		fmt.Printf("%s was caught!\n", *name)
		cfg.pokedex[*name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", *name)
	}

	return nil
}
