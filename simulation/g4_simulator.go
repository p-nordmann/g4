package simulation

import (
	"fmt"
	"g4"
)

// Board describes a board state.
//
// It is not aware of concepts of player and move.
// It is just a snapshot of the game at a given time.
type Board struct {
	board   [g4.RowCount]g4.Column
	gravity g4.Direction
}

func FromString(board string) (Board, error) {
	return Board{}, nil
}

type Game struct {
	board            Board
	colorWithTheMove g4.Color
}

func FromBoard(board Board, color g4.Color) (Game, error) {
	switch color {
	case g4.Red:
	case g4.Yellow:
	default:
		return Game{}, fmt.Errorf("unexpected color: %v", color)
	}
	return Game{
		board:            board,
		colorWithTheMove: color,
	}, nil
}

func (g Game) Board() Board {
	return g.board
}

// Generate computes the list of possible moves from a given position.
func (g Game) Generate() ([]g4.Move, error) {
	return nil, nil
}

func (g Game) Apply(move g4.Move) (g4.Game, g4.Generator, error) {
	return nil, nil, nil
}
