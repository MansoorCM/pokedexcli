package main

import "fmt"

func commandPokedex(cfg *config, name *string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("Your pokedex is empty.")
		return nil
	}

	fmt.Println("Your pokedex:")
	for pokemon := range cfg.pokedex {
		fmt.Printf(" - %s\n", pokemon)
	}
	return nil
}
