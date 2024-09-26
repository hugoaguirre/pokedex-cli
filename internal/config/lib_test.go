package config

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestDefaultRegion(t *testing.T) {
	config, err := LoadConfig()
	assert.NilError(t, err)
	assert.Equal(t, config.PokeApi.Region, "hoenn")
	assert.Assert(t, config.PokeApi.PokedexUrl != "")
}
