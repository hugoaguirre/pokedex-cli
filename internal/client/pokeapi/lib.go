package pokeapi

import (
	"net/http"

	"github.com/hugoaguirre/pokedex-cli/internal/config"
)

type Client interface {
	HealthCheck() error
	Pokedex() (Pokedex, error)
	// TODO: define more functions
}

type client struct {
	http.Client

	config config.Config
}

func NewClient(c config.Config) (Client, error) {
	return &client{
		config: c,
	}, nil
}
