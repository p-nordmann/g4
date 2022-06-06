package g4

import "testing"

func TestMoveMaker(t *testing.T) {
	got := []Move{
		Base().Tilt().Gravity(UP),
		Base().Tilt().Gravity(DOWN),
		Base().Token().Color(Yellow).Column(3),
		Base().Token().Color(Red).Column(0),
	}
	want := []Move{
		{moveType: Tilt, gravity: UP},
		{moveType: Tilt, gravity: DOWN},
		{moveType: Token, color: Yellow, column: 3},
		{moveType: Token, color: Red, column: 0},
	}
	for k := range got {
		if got[k] != want[k] {
			t.Errorf("example %d: got %v but want %v", k, got[k], want[k])
		}
	}
}

func TestShorthands(t *testing.T) {
	got := []Move{
		TiltMove(UP),
		TiltMove(DOWN),
		TokenMove(Yellow, 3),
		TokenMove(Red, 0),
	}
	want := []Move{
		{moveType: Tilt, gravity: UP},
		{moveType: Tilt, gravity: DOWN},
		{moveType: Token, color: Yellow, column: 3},
		{moveType: Token, color: Red, column: 0},
	}
	for k := range got {
		if got[k] != want[k] {
			t.Errorf("example %d: got %v but want %v", k, got[k], want[k])
		}
	}
}
