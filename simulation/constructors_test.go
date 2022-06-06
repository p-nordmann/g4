package simulation

import (
	"fmt"
	"g4"
	"testing"
)

var exampleBoard = [g4.ColumnCount]g4.Column{
	g4.Column(g4.Yellow) | g4.Column(g4.Red)<<8,
	g4.Column(g4.Red),
	0,
	0,
	0,
	0,
	0,
	0,
}

func TestCorrectBoardFromString(t *testing.T) {

	examples := []struct {
		in  string
		out Board
	}{
		{
			in: "UP yr/r//////",
			out: Board{
				gravity: g4.UP,
				board:   exampleBoard,
			},
		},
		{
			in: "DOWN yr/r//////",
			out: Board{
				gravity: g4.DOWN,
				board:   exampleBoard,
			},
		},
		{
			in: "LEFT yr/r//////",
			out: Board{
				gravity: g4.LEFT,
				board:   exampleBoard,
			},
		},
		{
			in: "RIGHT yr/r//////",
			out: Board{
				gravity: g4.RIGHT,
				board:   exampleBoard,
			},
		},
		{
			in: "UP yy/rr/yryr/rrrr/yrry/r/y/r",
			out: Board{
				gravity: g4.UP,
				board: [g4.ColumnCount]g4.Column{
					g4.Column(g4.Yellow) | g4.Column(g4.Yellow)<<8,
					g4.Column(g4.Red) | g4.Column(g4.Red)<<8,
					g4.Column(g4.Yellow) | g4.Column(g4.Red)<<8 | g4.Column(g4.Yellow)<<16 | g4.Column(g4.Red)<<24,
					g4.Column(g4.Red) | g4.Column(g4.Red)<<8 | g4.Column(g4.Red)<<16 | g4.Column(g4.Red)<<24,
					g4.Column(g4.Yellow) | g4.Column(g4.Red)<<8 | g4.Column(g4.Red)<<16 | g4.Column(g4.Yellow)<<24,
					g4.Column(g4.Red),
					g4.Column(g4.Yellow),
					g4.Column(g4.Red),
				},
			},
		},
	}
	for k, ex := range examples {
		if out, err := FromString(ex.in); err != nil || out != ex.out {
			t.Errorf("example %d: error FromString", k)
		}
	}
}

func TestInvalidBoardFromString(t *testing.T) {
	examples := []string{
		"",
		"UP DOWN yr/r//////",
		"yr/r//////",
		"UP yr/r/////",
		" yr/r//////r/",
		"UP yr/r/////ryryryryr",
	}
	for k, ex := range examples {
		if _, err := FromString(ex); err == nil {
			t.Errorf("example %d: expected error, got <nil> (%s)", k, ex)
		}
	}
}

func TestFromBoard(t *testing.T) {
	examples := []struct {
		in    string
		color g4.Color
		out   Game
		err   error
	}{
		{
			in:    "UP yr/r//////",
			color: g4.Yellow,
			out: Game{
				board: Board{
					gravity: g4.UP,
					board:   exampleBoard,
				},
				colorWithTheMove: g4.Yellow,
			},
			err: nil,
		},
		{
			in:    "UP yr/r//////",
			color: g4.Empty,
			out:   Game{},
			err:   fmt.Errorf("unexpected color: %v", g4.Empty),
		},
		{
			in:    "UP yr/r//////",
			color: g4.Color(255),
			out:   Game{},
			err:   fmt.Errorf("unexpected color: %v", g4.Color(255)),
		},
	}
	for k, ex := range examples {
		board, err := FromString(ex.in)
		if err != nil {
			t.Errorf("example %d: error in FromString: %v", k, err)
		}
		out, err := FromBoard(board, ex.color)
		if err != nil && ex.err != nil {
			if err.Error() != ex.err.Error() {
				t.Errorf("example %d: invalid error: got %v but want %v", k, err, ex.err)
			}
		} else if err != nil || ex.err != nil {
			t.Errorf("example %d: invalid error: got %v but want %v", k, err, ex.err)
		}
		if out != ex.out {
			t.Errorf("example %d: invalid Game", k)
		}
	}
}
