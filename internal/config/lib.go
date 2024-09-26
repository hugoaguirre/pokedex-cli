package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	PokeApi PokeApiConf
}

func LoadConfig() (Config, error) {
	var config Config
	err := envconfig.Process("pokeapi", &config)
	return config, err
}

type PokeApiConf struct {
	PokedexUrl string `envconf:"pokedex_url"`
	PokeApi    string `envconf:"pokemon_api_url"`
	Region     string `envconf:"pokemon_region" default:"hoenn"`
}
