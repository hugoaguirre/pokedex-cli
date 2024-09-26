package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	PokeApi PokeApiConf
}

func LoadConfig() (Config, error) {
	var config Config
	err := envconfig.Process("pa", &config)
	return config, err
}

type PokeApiConf struct {
	PokedexUrl string `envconfig:"pokedex_url"`
	PokeApi    string `envconfig:"pokemon_api_url"`
	Region     string `envconfig:"pokemon_region" default:"hoenn"`
}
