package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MansoorCM/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	pokedex          map[string]pokeapi.Pokemon
	nextLocationsUrl *string
	prevLocationsUrl *string
}

func startRepl(config *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print("pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := strings.Split(scanner.Text(), " ")
		arg1 := input[0]
		arg2 := ""
		if len(input) > 1 {
			arg2 = input[1]
		}
		command, ok := commands[arg1]
		if !ok {
			fmt.Println("Enter valid command")
		} else {
			err := command.callback(config, &arg2)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input", err)
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get the next 20 pokemon locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous 20 pokemon locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Get the names of pokemons in the region",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch the pokemon of the given name",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "see details about an already caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all pokemons that were caught.",
			callback:    commandPokedex,
		},
	}
}
