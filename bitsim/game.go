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

// Generate computes the list of possible moves from a given position.
func (g Game) Generate() ([]g4.Move, error) {
	var moves []g4.Move

	// Look for connect 4s.
	hasYellowConnect4 := g.Board.hasYellowConnect4()
	hasRedConnect4 := g.Board.hasRedConnect4()
	if hasYellowConnect4 && hasRedConnect4 {
		return moves, g4.ErrorGameOver(g4.Empty) // Draw.
	} else if hasYellowConnect4 {
		return moves, g4.ErrorGameOver(g4.Yellow)
	} else if hasRedConnect4 {
		return moves, g4.ErrorGameOver(g4.Red)
	}

	// Check whether board is full.
	heights := g.Board.heights()
	if heights[0]+heights[1]+heights[2]+heights[3]+heights[4]+heights[5]+heights[6]+heights[7] == 64 {
		return moves, g4.ErrorGameOver(g4.Empty) // Draw.
	}

	// Tilt moves.
	moves = append(
		moves,
		g4.TiltMove(g.Mover, g4.LEFT),
		g4.TiltMove(g.Mover, g4.DOWN),
		g4.TiltMove(g.Mover, g4.RIGHT),
	)

	// Token moves.
	for column, height := range heights {
		if height < 8 {
			moves = append(moves, g4.TokenMove(g.Mover, column))
		}
	}

	return moves, nil
}

// Apply performs a move from a game state.
func (g Game) Apply(move g4.Move) (Game, error) {
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
		if move.ColumnIdx < 0 || move.ColumnIdx >= 8 {
			return g, g4.ErrorInvalidMove{}
		}
		if g.Board.heights()[move.ColumnIdx] == 8 {
			return g, g4.ErrorInvalidMove{}
		}
		g.Board = g.Board.AddToken(move.ColumnIdx, g.Mover)
	default:
		return g, g4.ErrorInvalidMove{}
	}

	// Switch Mover.
	if g.Mover == g4.Red {
		g.Mover = g4.Yellow
	} else {
		g.Mover = g4.Red
	}

	return g, nil
}
