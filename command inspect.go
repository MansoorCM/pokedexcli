package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, name *string) error {
	if name == nil || *name == "" {
		return errors.New("please enter pokemon name to inspect")
	}
	pokemon, ok := cfg.pokedex[*name]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	if len(pokemon.Stats) > 0 {
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
	}
	if len(pokemon.Types) > 0 {
		fmt.Println("Types:")
		for _, pType := range pokemon.Types {
			fmt.Printf("  - %s\n", pType.Type.Name)
		}
	}
	return nil
}
