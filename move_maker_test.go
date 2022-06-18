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
