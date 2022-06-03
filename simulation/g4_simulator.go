package simulation

import "g4"

// state describes a board state.
//
// It is not aware of concepts of player and move.
// It is just a snapshot of the game at a given time.
type state struct {
	board   [g4.RowCount]g4.Column
	gravity g4.Direction
}

type generator struct {
	pos state
}

// Generate computes the list of possible moves from a given position.
func (g generator) Generate() ([]g4.Move, error) {
	return nil, nil
}
