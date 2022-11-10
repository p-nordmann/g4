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
