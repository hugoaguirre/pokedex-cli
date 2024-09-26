package pokeapi

import (
	"encoding/json"
	"log"
)

func (c *client) HealthCheck() error {
	return nil
}

func (c *client) Pokedex() (Pokedex, error) {
	log.Print("Checking pokedex")
	url := c.config.PokeApi.PokedexUrl + "/hoenn"
	res, err := c.Get(url)
	if err != nil {
		log.Fatalf("error obtaining pokedex: %v", err)
		return Pokedex{}, nil
	}
	defer res.Body.Close()

	var pokedex Pokedex
	err = json.NewDecoder(res.Body).Decode(&pokedex)
	if err != nil {
		log.Fatalf("error parsing pokedex response: %v", err)
		return Pokedex{}, nil
	}

	return pokedex, nil
}
