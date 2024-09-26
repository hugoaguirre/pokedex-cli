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

func (m *MockPokeApiClient) Pokedex() (PokedexData, error) {
	pokemonEntries := []PokemonEntries{
		{
			PokemonSpecies{
				"treeko",
				"fake/url",
			},
			252,
		},
		{
			PokemonSpecies{
				"pikachu",
				"fake/url",
			},
			3,
		},
	}

	return PokedexData{
		ID:             1,
		Name:           "hoenn",
		PokemonEntries: pokemonEntries,
	}, nil
}
