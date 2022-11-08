package frontend

import (
	"g4"
)

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
