package pokeapi

import (
	"encoding/json"
	"log"
)

func (c *client) HealthCheck() error {
	return nil
}

func (c *client) Pokedex() (PokedexData, error) {
	log.Print("Checking pokedex")
	url := c.config.PokeApi.PokedexUrl + "/" + c.config.PokeApi.Region
	res, err := c.Get(url)
	if err != nil {
		log.Fatalf("error obtaining pokedex: %v", err)
		return PokedexData{}, nil
	}
	defer res.Body.Close()

	var pokedex PokedexData
	err = json.NewDecoder(res.Body).Decode(&pokedex)
	if err != nil {
		log.Fatalf("error parsing pokedex response: %v", err)
		return PokedexData{}, nil
	}

	return pokedex, nil
}
