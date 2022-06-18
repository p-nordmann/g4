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

// func TestApplyCorrectMoves(t *testing.T) {
// 	examples := []struct {
// 		before string
// 		color  g4.Color
// 		moves  []g4.Move
// 		after  string
// 	}{
// 		{
// 			before: "UP ///////",
// 			color:  g4.Yellow,
// 			moves: []g4.Move{
// 				g4.TokenMove(g4.Yellow, 0),
// 				g4.TokenMove(g4.Red, 0),
// 				g4.TokenMove(g4.Yellow, 0),
// 			},
// 			after: "UP yry///////",
// 		},
// 		{
// 			before: "UP ///////",
// 			color:  g4.Yellow,
// 			moves: []g4.Move{
// 				g4.TokenMove(g4.Yellow, 0),
// 				g4.TokenMove(g4.Red, 1),
// 				g4.TokenMove(g4.Yellow, 2),
// 				g4.TokenMove(g4.Red, 3),
// 				g4.TiltMove(g4.LEFT),
// 			},
// 			after: "LEFT ryry///////",
// 		},
// 		{
// 			before: "UP rr/y/r/yy////",
// 			color:  g4.Yellow,
// 			moves: []g4.Move{
// 				g4.TiltMove(g4.LEFT),
// 			},
// 			after: "LEFT yryr/yr//////",
// 		},
// 		{
// 			before: "UP rr/y/r/yy////",
// 			color:  g4.Yellow,
// 			moves: []g4.Move{
// 				g4.TiltMove(g4.RIGHT),
// 			},
// 			after: "RIGHT //////ry/ryry",
// 		},
// 	}
// 	for k, ex := range examples {
// 		board, err := simulation.FromString(ex.before)
// 		if err != nil {
// 			t.Errorf("example %d: error in FromString (before): %v", k, err)
// 		}
// 		var game g4.Game
// 		game, err = simulation.FromBoard(board, ex.color)
// 		if err != nil {
// 			t.Errorf("example %d: error in FromBoard: %v", k, err)
// 		}
// 		for _, move := range ex.moves {
// 			game, _, err = game.Apply(move)
// 			if err != nil {
// 				t.Errorf("example %d: error in Apply: %v", k, err)
// 			}
// 		}
// 		want, err := simulation.FromString(ex.after)
// 		if err != nil {
// 			t.Errorf("example %d: error in FromString (after): %v", k, err)
// 		}
// 		got := game.(simulation.Game).Board()
// 		if got != want {
// 			t.Errorf("example %d: wrong board after game moves: got %v but wanted %v", k, got, want)
// 		}
// 	}
// }

// func TestApplyInvalidMoves(t *testing.T) {
// 	examples := []struct {
// 		board string
// 		color g4.Color
// 		move  g4.Move
// 		err   error
// 	}{
// 		{
// 			board: "UP ///////",
// 			color: g4.Yellow,
// 			move:  g4.TiltMove(g4.UP),
// 			err:   g4.ErrorInvalidMove{},
// 		},
// 		{
// 			board: "UP ryryryry///////",
// 			color: g4.Yellow,
// 			move:  g4.TokenMove(g4.Yellow, 0),
// 			err:   g4.ErrorInvalidMove{},
// 		},
// 	}
// 	for k, ex := range examples {
// 		board, err := simulation.FromString(ex.board)
// 		if err != nil {
// 			t.Errorf("example %d: error in FromString: %v", k, err)
// 		}
// 		game, err := simulation.FromBoard(board, ex.color)
// 		if err != nil {
// 			t.Errorf("example %d: error in FromBoard: %v", k, err)
// 		}
// 		_, _, err = game.Apply(ex.move)
// 		if err != ex.err {
// 			t.Errorf("example %d: incorrect error: got %v but want %v", k, err, ex.err)
// 		}
// 	}
// }

// TODO(Pierre-Louis): test perft.
