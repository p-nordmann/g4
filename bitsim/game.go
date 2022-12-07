package bitsim

import (
	"fmt"
	"g4"
)

// FromString builds a new Game instance with the board in requested position.
//
// `s` is a description of the board. Use bitsim.StartingPosition for usual starting position.
// `color` denotes the player with the move. Usually g4.Yellow goes first.
func FromString(s string, color g4.Color) (Game, error) {
	switch color {
	case g4.Red:
	case g4.Yellow:
	default:
		return Game{}, fmt.Errorf("unexpected color: %v", color)
	}
	board, err := boardFromString(s)
	if err != nil {
		return Game{}, fmt.Errorf("failed to parse board string: %w", err)
	}
	return Game{
		board: board,
		color: color,
	}, nil
}

type Game struct {
	board Board
	color g4.Color
}

func (g Game) ToArray() [8][8]g4.Color {
	return g.board.ToArray()
}

func (g Game) String() string {
	return g.board.String()
}

// Generate computes the list of possible moves from a given position.
func (g Game) Generate() ([]g4.Move, error) {
	var moves []g4.Move

	// Look for connect 4s.
	hasYellowConnect4 := g.board.hasYellowConnect4()
	hasRedConnect4 := g.board.hasRedConnect4()
	if hasYellowConnect4 && hasRedConnect4 {
		return moves, g4.ErrorGameOver(g4.Empty) // Draw.
	} else if hasYellowConnect4 {
		return moves, g4.ErrorGameOver(g4.Yellow)
	} else if hasRedConnect4 {
		return moves, g4.ErrorGameOver(g4.Red)
	}

	// Check whether board is full.
	heights := g.board.heights()
	if heights[0]+heights[1]+heights[2]+heights[3]+heights[4]+heights[5]+heights[6]+heights[7] == 64 {
		return moves, g4.ErrorGameOver(g4.Empty) // Draw.
	}

	// Tilt moves.
	moves = append(moves, g4.TiltMove(g4.LEFT), g4.TiltMove(g4.DOWN), g4.TiltMove(g4.RIGHT))

	// Token moves.
	for column, height := range heights {
		if height < 8 {
			moves = append(moves, g4.TokenMove(g.color, column))
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
		g.board = g.board.RotateLeft(times).ApplyGravity()
	case g4.Token:
		if move.ColumnIdx < 0 || move.ColumnIdx >= 8 {
			return g, g4.ErrorInvalidMove{}
		}
		if g.board.heights()[move.ColumnIdx] == 8 {
			return g, g4.ErrorInvalidMove{}
		}
		g.board = g.board.AddToken(move.ColumnIdx, g.color)
	default:
		return g, g4.ErrorInvalidMove{}
	}

	// Switch colors.
	if g.color == g4.Red {
		g.color = g4.Yellow
	} else {
		g.color = g4.Red
	}

	return g, nil
}
