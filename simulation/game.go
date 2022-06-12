package simulation

import (
	"fmt"
	"g4"
)

// TODO:   Board implementation must go.
//		 + Must use interface g4.Board when created.

type Game struct {
	board            g4.Board
	colorWithTheMove g4.Color
	gravity          g4.Direction
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
//
// TODO: errors game over.
func (g Game) Generate() ([]g4.Move, error) {
	var moves []g4.Move

	// Tilt moves.
	if g.gravity != g4.UP {
		moves = append(moves, g4.TiltMove(g4.UP))
	}
	if g.gravity != g4.LEFT {
		moves = append(moves, g4.TiltMove(g4.LEFT))
	}
	if g.gravity != g4.DOWN {
		moves = append(moves, g4.TiltMove(g4.DOWN))
	}
	if g.gravity != g4.RIGHT {
		moves = append(moves, g4.TiltMove(g4.RIGHT))
	}

	// Token moves.
	headMask := uint64(0b11111111 << (8 * 7))
	for col := 0; col < 8; col++ {
		if uint64(g.board.board[col])&headMask == 0 {
			moves = append(moves, g4.TokenMove(g.colorWithTheMove, col))
		}
	}

	return moves, nil
}

func (g Game) Apply(move g4.Move) (g4.Game, error) {
	return nil, nil
}
