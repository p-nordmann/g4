package frontend

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

func (m selectorArea) Update(ch g4.Channel, msg tea.Msg) (selectorArea, tea.Cmd) {
	if m.Disabled {
		return m, nil
	}
	var cmd tea.Cmd
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
				cmd = tea.Batch(sendMove(ch, m.PossibleMoves[m.SelectedMove]), receiveMove(ch))
			}
		}
	}
	return m, cmd
}
