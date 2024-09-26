package pokeapi

import (
	"log"
	"net/http"
)

type MockPokeApiClient struct {
	http.Client
}

func NewMockPokeApiClient() *MockPokeApiClient {
	return &MockPokeApiClient{}
}

func (m *MockPokeApiClient) HealthCheck() error {
	log.Print("Executing a mock health check!")
	return nil
}

func (m *MockPokeApiClient) Pokedex() (Pokedex, error) {
	pokemonEntries := []PokemonEntries{
		{
			252,
			PokemonSpecies{
				"treeko",
				"fake/url",
			},
		},
		{
			3,
			PokemonSpecies{
				"pikachu",
				"fake/url",
			},
		},
	}

	return Pokedex{
		ID:             1,
		Name:           "hoenn",
		PokemonEntries: pokemonEntries,
	}, nil
}
