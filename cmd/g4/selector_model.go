package main

import (
	"g4"
	"g4/bitsim"

	tea "github.com/charmbracelet/bubbletea"
)

type SelectorModel struct {
	game         bitsim.Game
	selectedMove g4.Move
}

func (m SelectorModel) Init() tea.Cmd {
	return nil
}

type UpdatePossibleMoves []g4.Move

type LockedMove g4.Move
type SelectedMove g4.Move

func (m SelectorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		// TODO: dynamic keys depending on the moves.
		case "enter", "space":
			if m.selectedMove != (g4.Move{}) {
				sendCmd, err := p2pService.sendMove(m.selectedMove) // TODO: should validate move?
				if err != nil {
					return m, handleError(err)
				}
				return m, sendCmd
			}
		}
	}
	return m, nil
}

func (m SelectorModel) View() string {
	return "selector has no physical reality"
}
