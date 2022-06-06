package simulation

import (
	"fmt"
	"g4"
	"strings"
)

// Board describes a board state.
//
// It is not aware of concepts of player and move.
// It is just a snapshot of the game at a given time.
type Board struct {
	board   [g4.ColumnCount]g4.Column
	gravity g4.Direction
}

func FromString(board string) (Board, error) {
	var b Board

	// Parse gravity.
	switch true {
	case strings.HasPrefix(board, "UP "):
		b.gravity = g4.UP
		board = board[3:]
	case strings.HasPrefix(board, "LEFT "):
		b.gravity = g4.LEFT
		board = board[5:]
	case strings.HasPrefix(board, "RIGHT "):
		b.gravity = g4.RIGHT
		board = board[6:]
	case strings.HasPrefix(board, "DOWN "):
		b.gravity = g4.DOWN
		board = board[5:]
	default:
		return Board{}, fmt.Errorf("board string '%s' is invalid: invalid gravity", board)
	}

	// Parse actual board.
	col := 0
	for k, row := 0, 0; k < len(board); k++ {
		if row > 8 {
			return Board{}, fmt.Errorf("board string '%s' is invalid: too many rows", board)
		}
		switch board[k] {
		case 'y':
			b.board[col] |= g4.Column(g4.Yellow) << (8 * row)
			row++
		case 'r':
			b.board[col] |= g4.Column(g4.Red) << (8 * row)
			row++
		case '/':
			col++
			row = 0
			if col > 7 {
				return Board{}, fmt.Errorf("board string '%s' is invalid: too many columns", board)
			}
		default:
			return Board{}, fmt.Errorf("board string '%s' is invalid: invalid character '%v'", board, board[k])
		}
	}
	if col < 7 {
		return Board{}, fmt.Errorf("board string '%s' is invalid: not enough columns", board)
	}

	return b, nil
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
//
// TODO: errors.
func (g Game) Generate() ([]g4.Move, error) {
	var moves []g4.Move

	// Tilt moves.
	if g.board.gravity != g4.UP {
		moves = append(moves, g4.TiltMove(g4.UP))
	}
	if g.board.gravity != g4.LEFT {
		moves = append(moves, g4.TiltMove(g4.LEFT))
	}
	if g.board.gravity != g4.DOWN {
		moves = append(moves, g4.TiltMove(g4.DOWN))
	}
	if g.board.gravity != g4.RIGHT {
		moves = append(moves, g4.TiltMove(g4.RIGHT))
	}

	// Token moves.
	headMask := uint64(0b11111111 << (8 * 7))
	for col := 0; col < g4.ColumnCount; col++ {
		if uint64(g.board.board[col])&headMask == 0 {
			moves = append(moves, g4.TokenMove(g.colorWithTheMove, col))
		}
	}

	return moves, nil
}

func (g Game) Apply(move g4.Move) (g4.Game, g4.Generator, error) {
	return nil, nil, nil
}
