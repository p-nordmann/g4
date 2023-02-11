package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func handleError(err error) tea.Cmd {
	if p2pService.ch != nil {
		p2pService.ch.Close()
	}
	return func() tea.Msg { return err }
}

type ErrorModel struct {
	err error
}

func (m ErrorModel) View() string {
	return m.err.Error() + "\n" + "Press 'q' to quit."
}

func (m ErrorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if isQuitMessage(msg) {
		return m, tea.Quit
	}
	return m, nil
}

func (m ErrorModel) Init() tea.Cmd {
	return nil
}
