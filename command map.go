package main

import (
	"fmt"
)

func commandMap(cfg *config) error {

	locationsresp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsUrl)
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
