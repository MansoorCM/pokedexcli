package main

import (
	"errors"
	"fmt"
)

func commandMapb(cfg *config, name *string) error {
	if cfg.prevLocationsUrl == nil {
		return errors.New("already on first page of result! cannot go back")
	}

	locationsresp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsUrl)
	if err != nil {
		return err
	}

	for _, loc := range locationsresp.Results {
		fmt.Println(loc.Name)
	}

	cfg.nextLocationsUrl = locationsresp.Next
	cfg.prevLocationsUrl = locationsresp.Previous
	return nil
}
