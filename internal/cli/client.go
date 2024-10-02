package cli

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hugoaguirre/pokedex-cli/internal/client/pokeapi"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title string
	desc  string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func getPokedex() (pokeapi.PokedexData, error) {
	pokeApiClient, err := pokeapi.NewClient()
	if err != nil {
		return pokeapi.PokedexData{}, fmt.Errorf("unable to init poke api client due err: %v", err)
	}

	return pokeApiClient.Pokedex()
}

func Start() {
	pokedex, err := getPokedex()
	if err != nil {
		fmt.Println("unable to init poke-api client: %v", err)
		os.Exit(1)
	}

	items := make([]list.Item, 0)
	for i := 0; i < len(pokedex.PokemonEntries); i++ {
		items = append(items, item{
			title: pokedex.PokemonEntries[i].PokemonSpecies.Name,
			desc:  pokedex.Name,
		})
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = fmt.Sprintf("%s - Pokedex", pokedex.Name)

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program: ", err)
		os.Exit(1)
	}
}
