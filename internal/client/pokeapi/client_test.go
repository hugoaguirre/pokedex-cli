package pokeapi

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestMockHealthCheck(t *testing.T) {
	mock := NewMockPokeApiClient()

	mock.HealthCheck()
}

func TestMockPokedex(t *testing.T) {
	mock := NewMockPokeApiClient()

	pokedex, err := mock.Pokedex()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, pokedex.ID, 1)
	assert.Equal(t, pokedex.Name, "hoenn")
	assert.Assert(t, len(pokedex.PokemonEntries) > 0)
	assert.Assert(t, pokedex.PokemonEntries[0].EntryNumber > 0)
	assert.Assert(t, pokedex.PokemonEntries[0].PokemonSpecies.URL != "")
	assert.Assert(t, pokedex.PokemonEntries[0].PokemonSpecies.Name != "")
}
