/*
G4 is an open-source board game inspired by the popular game of connect-4.
Copyright (C) 2022  Pierre-Louis Nordmann

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package simulation_test

import (
	"g4"
	"g4/bits"
	"g4/simulation"
	"testing"
)

func compareMoves(moves1 []g4.Move, moves2 []g4.Move) bool {
	counts := make(map[g4.Move]int)
	for _, move := range moves1 {
		counts[move]++
	}
	for _, move := range moves2 {
		counts[move]--
	}
	for _, count := range counts {
		if count != 0 {
			return false
		}
	}
	return true
}

func tiltMoves(directions []g4.Direction) []g4.Move {
	var moves []g4.Move
	for _, direction := range directions {
		moves = append(moves, g4.TiltMove(direction))
	}
	return moves
}

func tokenMoves(color g4.Color, columns []int) []g4.Move {
	var moves []g4.Move
	for _, column := range columns {
		moves = append(moves, g4.TokenMove(color, column))
	}
	return moves
}

func concatMoves(moves1 []g4.Move, moves2 []g4.Move) []g4.Move {
	var moves []g4.Move
	moves = append(moves, moves1...)
	moves = append(moves, moves2...)
	return moves
}

func TestUtils(t *testing.T) {
	// Test compareMoves.
	examples := []struct {
		in1 []g4.Move
		in2 []g4.Move
		out bool
	}{
		{
			in1: nil,
			in2: nil,
			out: true,
		},
		{
			in1: nil,
			in2: []g4.Move{},
			out: true,
		},
		{
			in1: nil,
			in2: []g4.Move{g4.TokenMove(g4.Yellow, 2)},
			out: false,
		},
		{
			in1: []g4.Move{g4.TokenMove(g4.Yellow, 2)},
			in2: []g4.Move{g4.TokenMove(g4.Yellow, 2)},
			out: true,
		},
		{
			in1: []g4.Move{g4.TokenMove(g4.Yellow, 2)},
			in2: []g4.Move{g4.TokenMove(g4.Yellow, 2), g4.TiltMove(g4.UP)},
			out: false,
		},
		{
			in1: []g4.Move{g4.TokenMove(g4.Yellow, 2), g4.TiltMove(g4.UP)},
			in2: []g4.Move{g4.TokenMove(g4.Yellow, 2), g4.TiltMove(g4.UP)},
			out: true,
		},
		{
			in1: []g4.Move{g4.TokenMove(g4.Yellow, 2), g4.TiltMove(g4.UP)},
			in2: []g4.Move{g4.TiltMove(g4.UP), g4.TokenMove(g4.Yellow, 2)},
			out: true,
		},
	}
	for k, ex := range examples {
		out12 := compareMoves(ex.in1, ex.in2)
		out21 := compareMoves(ex.in2, ex.in1)
		if out21 != ex.out || out12 != ex.out {
			t.Errorf("example %d: got %v/%v but want %v", k, out12, out21, ex.out)
		}
	}

	// Test tiltMoves and tokenMoves.
	got := tiltMoves([]g4.Direction{g4.UP, g4.LEFT})
	want := []g4.Move{
		g4.TiltMove(g4.UP),
		g4.TiltMove(g4.LEFT),
	}
	if !compareMoves(got, want) {
		t.Errorf("tiltMoves: got %v but want %v", got, want)
	}

	got = tokenMoves(g4.Red, []int{0, 1, 5, 6})
	want = []g4.Move{
		g4.TokenMove(g4.Red, 0),
		g4.TokenMove(g4.Red, 1),
		g4.TokenMove(g4.Red, 5),
		g4.TokenMove(g4.Red, 6),
	}
	if !compareMoves(got, want) {
		t.Errorf("tokenMoves: got %v but want %v", got, want)
	}

	// Test concatMoves.
	got = concatMoves(
		tiltMoves([]g4.Direction{g4.UP, g4.LEFT}),
		tokenMoves(g4.Red, []int{0, 1, 5, 6}),
	)
	want = []g4.Move{
		g4.TiltMove(g4.UP),
		g4.TiltMove(g4.LEFT),
		g4.TokenMove(g4.Red, 0),
		g4.TokenMove(g4.Red, 1),
		g4.TokenMove(g4.Red, 5),
		g4.TokenMove(g4.Red, 6),
	}
	if !compareMoves(got, want) {
		t.Errorf("concatMoves: got %v but want %v", got, want)
	}
}

func TestGenerate(t *testing.T) {
	examples := []struct {
		in        string
		color     g4.Color
		direction g4.Direction
		out       []g4.Move
		err       error
	}{
		{
			in:        "8|8|8|8|8|8|8|8",
			color:     g4.Yellow,
			direction: g4.UP,
			out: concatMoves(
				tiltMoves([]g4.Direction{g4.RIGHT, g4.DOWN, g4.LEFT}),
				tokenMoves(g4.Yellow, []int{0, 1, 2, 3, 4, 5, 6, 7}),
			),
			err: nil,
		},
		{
			in:        "ryryryry|8|8|8|8|8|8|8",
			color:     g4.Yellow,
			direction: g4.UP,
			out: concatMoves(
				tiltMoves([]g4.Direction{g4.RIGHT, g4.DOWN, g4.LEFT}),
				tokenMoves(g4.Yellow, []int{1, 2, 3, 4, 5, 6, 7}),
			),
			err: nil,
		},
		{
			in:        "8|8|8|8|8|8|8|8",
			color:     g4.Yellow,
			direction: g4.LEFT,
			out: concatMoves(
				tiltMoves([]g4.Direction{g4.UP, g4.RIGHT, g4.DOWN}),
				tokenMoves(g4.Yellow, []int{0, 1, 2, 3, 4, 5, 6, 7}),
			),
			err: nil,
		},
		{
			in:        "ryryryry|ryryryry|ryryryry|yryryryr|yryryryr|yryryryr|ryryryry|ryryryry",
			color:     g4.Yellow,
			direction: g4.LEFT,
			out:       nil,
			err:       g4.ErrorGameOver(g4.Empty),
		},
		{
			in:        "rrrr4|yryr4|8|8|8|8|8|8",
			color:     g4.Yellow,
			direction: g4.LEFT,
			out:       nil,
			err:       g4.ErrorGameOver(g4.Red),
		},
		{
			in:        "ryry4|ryyy4|r7|r7|8|8|8|8",
			color:     g4.Yellow,
			direction: g4.LEFT,
			out:       nil,
			err:       g4.ErrorGameOver(g4.Red),
		},
		{
			in:        "ryry4|yryy4|rrr5|yyyr4|8|8|8|8",
			color:     g4.Yellow,
			direction: g4.LEFT,
			out:       nil,
			err:       g4.ErrorGameOver(g4.Red),
		},
	}
	for k, ex := range examples {
		board, err := bits.FromString(ex.in)
		if err != nil {
			t.Errorf("example %d: error in FromString: %v", k, err)
		}
		game, err := simulation.FromBoard(board, ex.color, ex.direction)
		if err != nil {
			t.Errorf("example %d: error in FromBoard: %v", k, err)
		}
		out, err := game.Generate()
		if err != ex.err {
			t.Errorf("example %d: Generate; invalid error: got %v but want %v", k, err, ex.err)
		}
		if !compareMoves(out, ex.out) {
			t.Errorf("example %d: Generate; wrong output", k)
		}
	}
}

func TestApplyCorrectMoves(t *testing.T) {
	examples := []struct {
		in           string
		color        g4.Color
		direction    g4.Direction
		moves        []g4.Move
		out          string
		outColor     g4.Color
		outDirection g4.Direction
	}{
		{
			in:        "8|8|8|8|8|8|8|8",
			color:     g4.Yellow,
			direction: g4.UP,
			moves: []g4.Move{
				g4.TokenMove(g4.Yellow, 0),
				g4.TokenMove(g4.Red, 0),
				g4.TokenMove(g4.Yellow, 0),
			},
			out:          "yry5|8|8|8|8|8|8|8",
			outColor:     g4.Red,
			outDirection: g4.UP,
		},
		{
			in:        "8|8|8|8|8|8|8|8",
			color:     g4.Yellow,
			direction: g4.UP,
			moves: []g4.Move{
				g4.TokenMove(g4.Yellow, 0),
				g4.TokenMove(g4.Red, 1),
				g4.TokenMove(g4.Yellow, 2),
				g4.TokenMove(g4.Red, 3),
				g4.TiltMove(g4.LEFT),
			},
			out:          "ryry4|8|8|8|8|8|8|8",
			outColor:     g4.Red,
			outDirection: g4.LEFT,
		},
		{
			in:        "rr6|y7|r7|yy6|8|8|8|8",
			color:     g4.Yellow,
			direction: g4.UP,
			moves: []g4.Move{
				g4.TiltMove(g4.LEFT),
			},
			out:          "yryr4|yr6|8|8|8|8|8|8",
			outColor:     g4.Red,
			outDirection: g4.LEFT,
		},
		{
			in:        "rr6|y7|r7|yy6|8|8|8|8",
			color:     g4.Yellow,
			direction: g4.UP,
			moves: []g4.Move{
				g4.TiltMove(g4.RIGHT),
			},
			out:          "8|8|8|8|8|8|ry6|ryry4",
			outColor:     g4.Red,
			outDirection: g4.RIGHT,
		},
	}
	for k, ex := range examples {
		var game g4.Game
		board1, err := bits.FromString(ex.in)
		if err != nil {
			t.Errorf("example %d: error in FromString: %v", k, err)
		}
		game, err = simulation.FromBoard(board1, ex.color, ex.direction)
		if err != nil {
			t.Errorf("example %d: error in FromBoard (in): %v", k, err)
		}
		for i, move := range ex.moves {
			game, err = game.Apply(move)
			if err != nil {
				t.Errorf("example %d: move %d: error in Apply: %v", k, i, err)
			}
		}
		board2, err := bits.FromString(ex.out)
		if err != nil {
			t.Errorf("example %d: error in FromString: %v", k, err)
		}
		want, err := simulation.FromBoard(board2, ex.outColor, ex.outDirection)
		if err != nil {
			t.Errorf("example %d: error in FromBoard (out): %v", k, err)
		}
		got := game.(simulation.Game)
		if got != want {
			t.Errorf("example %d: wrong game state after game moves: got %v but wanted %v", k, got, want)
		}
	}
}

func TestApplyInvalidMoves(t *testing.T) {
	examples := []struct {
		in        string
		color     g4.Color
		direction g4.Direction
		move      g4.Move
		err       error
	}{
		{
			in:        "8|8|8|8|8|8|8|8",
			color:     g4.Yellow,
			direction: g4.UP,
			move:      g4.TiltMove(g4.UP),
			err:       g4.ErrorInvalidMove{},
		},
		{
			in:        "ryryryry|8|8|8|8|8|8|8",
			color:     g4.Yellow,
			direction: g4.UP,
			move:      g4.TokenMove(g4.Yellow, 0),
			err:       g4.ErrorInvalidMove{},
		},
		{
			in:        "8|8|8|8|8|8|8|8",
			color:     g4.Yellow,
			direction: g4.UP,
			move:      g4.TokenMove(g4.Yellow, 8), // NB: invalid column.
			err:       g4.ErrorInvalidMove{},
		},
		// TODO: add a move with invalid type.
	}
	for k, ex := range examples {
		board, err := bits.FromString(ex.in)
		if err != nil {
			t.Errorf("example %d: error in FromString: %v", k, err)
		}
		game, err := simulation.FromBoard(board, ex.color, ex.direction)
		if err != nil {
			t.Errorf("example %d: error in FromBoard (in): %v", k, err)
		}
		_, err = game.Apply(ex.move)
		if err != ex.err {
			t.Errorf("example %d: incorrect error: got %v but want %v", k, err, ex.err)
		}
	}
}

// TODO(Pierre-Louis): test perft.
