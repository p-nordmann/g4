package g4_test

import (
	"g4"
	"testing"
)

func TestErrorInvalidMove(t *testing.T) {
	k := 0
	want := "invalid move"
	err := g4.ErrorInvalidMove{}
	got := err.Error()
	if got != want {
		t.Errorf("example %d: got '%s' but want '%s'", k, got, want)
	}
}

func TestShorthands(t *testing.T) {
	got := []g4.Move{
		g4.TiltMove(g4.Yellow, g4.LEFT),
		g4.TiltMove(g4.Red, g4.DOWN),
		g4.TokenMove(g4.Yellow, 3),
		g4.TokenMove(g4.Red, 0),
	}
	want := []g4.Move{
		{Type: g4.Tilt, Color: g4.Yellow, Direction: g4.LEFT},
		{Type: g4.Tilt, Color: g4.Red, Direction: g4.DOWN},
		{Type: g4.Token, Color: g4.Yellow, Column: 3},
		{Type: g4.Token, Color: g4.Red, Column: 0},
	}
	for k := range got {
		if got[k] != want[k] {
			t.Errorf("example %d: got %v but want %v", k, got[k], want[k])
		}
	}
}
