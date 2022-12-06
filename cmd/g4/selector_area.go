package main

import (
	"g4"

	tea "github.com/charmbracelet/bubbletea"
)

// TODO: convert to model.
type selectorArea struct {
	PossibleMoves []g4.Move
	SelectedMove  int
	Disabled      bool
}

func (m selectorArea) View() string {
	if m.Disabled {
		return "Waiting for opponent..."
	}
	s := ""
	for k, move := range m.PossibleMoves {
		if k == m.SelectedMove {
			s += "[x] "
		} else {
			s += "[ ] "
		}
		s += move.String() + "\n"
	}
	return s
}

func (m selectorArea) Update(msg tea.Msg) (selectorArea, tea.Cmd) {
	if m.Disabled {
		return m, nil
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			m.SelectedMove--
			if m.SelectedMove < 0 {
				m.SelectedMove = len(m.PossibleMoves) - 1
			}
		case "down":
			m.SelectedMove++
			if m.SelectedMove >= len(m.PossibleMoves) {
				m.SelectedMove = 0
			}
		case "enter":
			if m.SelectedMove >= 0 {
				sendCmd, err := p2pService.sendMove(m.PossibleMoves[m.SelectedMove])
				if err != nil {
					return m, handleError(err)
				}
				return m, sendCmd
			}
		}
	}
	return m, nil
}
