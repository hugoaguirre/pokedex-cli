package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	// pokemon list
	// pokemon sprites ?
	err error
}

type (
	errMsg struct{ err error }
)

// error message interface on the message
func (e errMsg) Error() string { return e.err.Error() }

func checkServer() tea.Msg {
	return nil
}

func (m model) Init() tea.Cmd {
	return checkServer
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() string {
	return ""
}

func main() {
	p := tea.NewProgram(model{})
	if _, err := p.Run(); err != nil {
		fmt.Printf("init error: %v", err)
		os.Exit(1)
	}
}
