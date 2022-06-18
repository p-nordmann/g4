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
