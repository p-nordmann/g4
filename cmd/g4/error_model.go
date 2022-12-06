package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func handleError(err error) tea.Cmd {
	p2pService.ch.Close()
	return func() tea.Msg { return err }
}

type errorModel struct {
	err      error
	position string
}

func (m errorModel) View() string {
	return m.err.Error() + "\n" + m.position + "\n" + "Press 'q' to quit."
}

func (m errorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if isQuitMessage(msg) {
		return m, tea.Quit
	}
	return m, nil
}

func (m errorModel) Init() tea.Cmd {
	return nil
}
