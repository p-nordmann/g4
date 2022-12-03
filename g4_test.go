package g4_test

import (
	"g4"
	"testing"
)

func TestErrorGameOver(t *testing.T) {
	examples := []struct {
		in  g4.Color
		out string
	}{
		{
			in:  g4.Yellow,
			out: "game is over - yellow wins",
		},
		{
			in:  g4.Red,
			out: "game is over - red wins",
		},
		{
			in:  g4.Empty,
			out: "game is over - draw",
		},
		{
			in:  245,
			out: "game is over - draw",
		},
	}
	for k, ex := range examples {
		err := g4.ErrorGameOver(ex.in)
		got := err.Error()
		if got != ex.out {
			t.Errorf("example %d: got '%s' but want '%s'", k, got, ex.out)
		}
	}
}

func TestErrorInvalidMove(t *testing.T) {
	k := 0
	want := "invalid move"
	err := g4.ErrorInvalidMove{}
	got := err.Error()
	if got != want {
		t.Errorf("example %d: got '%s' but want '%s'", k, got, want)
	}
}

func TestMoveMaker(t *testing.T) {
	got := []g4.Move{
		g4.Base().Tilt().Gravity(g4.UP),
		g4.Base().Tilt().Gravity(g4.DOWN),
		g4.Base().Token().Color(g4.Yellow).Column(3),
		g4.Base().Token().Color(g4.Red).Column(0),
	}
	want := []g4.Move{
		{Type: g4.Tilt, Direction: g4.UP},
		{Type: g4.Tilt, Direction: g4.DOWN},
		{Type: g4.Token, Col: g4.Yellow, ColumnIdx: 3},
		{Type: g4.Token, Col: g4.Red, ColumnIdx: 0},
	}
	for k := range got {
		if got[k] != want[k] {
			t.Errorf("example %d: got %v but want %v", k, got[k], want[k])
		}
	}
}

func TestShorthands(t *testing.T) {
	got := []g4.Move{
		g4.TiltMove(g4.UP),
		g4.TiltMove(g4.DOWN),
		g4.TokenMove(g4.Yellow, 3),
		g4.TokenMove(g4.Red, 0),
	}
	want := []g4.Move{
		{Type: g4.Tilt, Direction: g4.UP},
		{Type: g4.Tilt, Direction: g4.DOWN},
		{Type: g4.Token, Col: g4.Yellow, ColumnIdx: 3},
		{Type: g4.Token, Col: g4.Red, ColumnIdx: 0},
	}
	for k := range got {
		if got[k] != want[k] {
			t.Errorf("example %d: got %v but want %v", k, got[k], want[k])
		}
	}
}
