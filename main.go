package main

import (
	"log"

	"github.com/hugoaguirre/pokedex-cli/internal/client/pokeapi"
)

func main() {
	pokeApiClient, err := pokeapi.NewClient()
	if err != nil {
		log.Fatalf("unable to init pokeapi client: %v", err)
	}

	p, err := pokeApiClient.Pokedex()
	if err != nil {
		log.Fatalf("unable to get pokedex: %v", err)
	}

	log.Printf("%#v", p)
}
