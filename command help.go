package main

import (
	"fmt"
)

func commandHelp() error {
	commands := getCommands()
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range commands {
		fmt.Println(command.name, ": ", command.description)
	}
	fmt.Println()
	return nil
}
