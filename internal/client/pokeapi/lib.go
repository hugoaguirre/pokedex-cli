package pokeapi

import (
	"log"
	"net/http"

	"github.com/hugoaguirre/pokedex-cli/internal/config"
)

type Client interface {
	HealthCheck() error
	Pokedex() (PokedexData, error)
	// TODO: define more functions
}

type client struct {
	http.Client

	config config.Config
}

func NewClient() (Client, error) {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("unable to get config for client: %v", err)
	}
	return &client{
		config: c,
	}, nil
}
