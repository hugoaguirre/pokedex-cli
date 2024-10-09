package cli

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hugoaguirre/pokedex-cli/internal/client/pokeapi"
	"github.com/muesli/termenv"
)

var (
	AccentColor       = lipgloss.ANSIColor(termenv.ANSIYellow)
	docStyle          = lipgloss.NewStyle().Margin(1, 2)
	ItemStyle         = lipgloss.NewStyle().PaddingLeft(1).BorderStyle(lipgloss.HiddenBorder()).BorderLeft(true)
	SelectedItemStyle = ItemStyle.Foreground(AccentColor).Bold(true).BorderStyle(lipgloss.ThickBorder()).BorderForeground(AccentColor)
)

type item struct {
	title string
	desc  string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	render := ItemStyle.Render
	if index == m.Index() {
		render = func(s ...string) string {
			return SelectedItemStyle.Render(strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, render(i.Title()))
}

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
		fmt.Printf("unable to init poke-api client: %v", err)
		os.Exit(1)
	}

	items := make([]list.Item, 0)
	for i := 0; i < len(pokedex.PokemonEntries); i++ {
		items = append(items, item{
			title: capitalize(pokedex.PokemonEntries[i].PokemonSpecies.Name),
		})
	}

	m := model{list: list.New(items, itemDelegate{}, 0, 0)}
	m.list.Title = fmt.Sprintf("%s - Pokedex", capitalize(pokedex.Name))
	// m.list.SetShowTitle(false)
	m.list.Paginator.Type = paginator.Arabic

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program: ", err)
		os.Exit(1)
	}
}
