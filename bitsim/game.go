package bitsim

import (
	"g4"
)

// Game holds the state of the game and provides an interface to make moves.
type Game struct {

	// Board holds the state of the board.
	Board Board

	// Mover denotes the player with the move.
	Mover g4.Color
}

// Returns an error if game is over.
func (g Game) Validate() error {
	hasYellowConnect4 := g.Board.hasYellowConnect4()
	hasRedConnect4 := g.Board.hasRedConnect4()

	if hasYellowConnect4 && hasRedConnect4 {
		return g4.Draw{}
	}
	if hasYellowConnect4 {
		return g4.YellowWins{}
	}
	if hasRedConnect4 {
		return g4.RedWins{}
	}

	if g.Board.count() == 64 {
		return g4.Draw{}
	}

	return nil
}

// Generate computes the list of possible moves from a given position.
func (g Game) Generate() ([]g4.Move, error) {
	var moves []g4.Move

	// Check that game is still live.
	if err := g.Validate(); err != nil {
		return moves, err
	}

	// Tilt moves.
	moves = append(
		moves,
		g4.TiltMove(g.Mover, g4.LEFT),
		g4.TiltMove(g.Mover, g4.DOWN),
		g4.TiltMove(g.Mover, g4.RIGHT),
	)

	// Token moves.
	for column, height := range g.Board.heights() {
		if height < 8 {
			moves = append(moves, g4.TokenMove(g.Mover, column))
		}
	}

	return moves, nil
}

// Apply performs a move from a game state.
func (g Game) Apply(move g4.Move) (Game, error) {

	// Check that game is still live.
	if err := g.Validate(); err != nil {
		return g, err
	}

	switch t := move.Type; t {

	case g4.Tilt:
		var times int
		switch move.Direction {
		case g4.LEFT:
			times = 1
		case g4.DOWN:
			times = 2
		case g4.RIGHT:
			times = 3
		default:
			return g, g4.ErrorInvalidMove{}
		}
		g.Board = g.Board.RotateLeft(times).ApplyGravity()

	case g4.Token:
		if move.Column < 0 || move.Column >= 8 {
			return g, g4.ErrorInvalidMove{}
		}
		if g.Board.heights()[move.Column] == 8 {
			return g, g4.ErrorInvalidMove{}
		}
		g.Board = g.Board.AddToken(move.Column, g.Mover)

	default:
		return g, g4.ErrorInvalidMove{}

	}

	// Switch Mover.
	if g.Mover == g4.Red {
		g.Mover = g4.Yellow
	} else {
		g.Mover = g4.Red
	}

	return g, g.Validate()
}
