package pokeapi

type PokemonSpecies struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonEntries struct {
	PokemonSpecies PokemonSpecies `json:"pokemon_species"`
	EntryNumber    int            `json:"entry_number"`
}

type Pokedex struct {
	Name           string           `json:"name"`
	PokemonEntries []PokemonEntries `json:"pokemon_entries"`
	ID             int              `json:"id"`
}

type Sprites struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Stats struct {
	Stat     Stat `json:"stat"`
	BaseStat int  `json:"base_stat"`
}

type Type struct {
	Name string `json:"name"`
}

type Types struct {
	Type []Type `json:"type"`
	Slot int    `json:"slot"`
}

type PokemonInfo struct {
	Name   string  `json:"name"`
	Stats  []Stats `json:"stats"`
	Height int     `json:"height"`
	ID     int     `json:"id"`
	Weight int     `json:"weight"`
}
