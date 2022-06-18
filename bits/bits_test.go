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

package bits_test

import (
	b "g4/bits"
	"testing"
)

func TestHeights(t *testing.T) {
	examples := []struct {
		in  string
		out [8]int
	}{
		{
			in:  "8|8|8|8|8|ryr5|r7|y7",
			out: [8]int{0, 0, 0, 0, 0, 3, 1, 1},
		},
		{
			in:  "8|ry6|8|8|rry5|8|8|8",
			out: [8]int{0, 2, 0, 0, 3, 0, 0, 0},
		},
		{
			in:  "8|8|8|8|rrryr3|ryyyyyr1|r7|yr6",
			out: [8]int{0, 0, 0, 0, 5, 7, 1, 2},
		},
		{
			in:  "8|yr6|yr6|yrrr4|yyr5|8|8|8",
			out: [8]int{0, 2, 2, 4, 3, 0, 0, 0},
		},
		{
			in:  "8|8|8|8|8|rryyyy2|r7|r7",
			out: [8]int{0, 0, 0, 0, 0, 6, 1, 1},
		},
		{
			in:  "8|ryryryry|8|8|rry5|8|8|8",
			out: [8]int{0, 8, 0, 0, 3, 0, 0, 0},
		},
	}
	for k, ex := range examples {
		in, err := b.FromString(ex.in)
		if err != nil {
			t.Errorf("example %d: FromString returned an error %v", k, err)
		}
		if in.Heights() != ex.out {
			t.Errorf("example %d: got %v but want %v", k, in.Heights(), ex.out)
		}
	}
}

func TestHasYellowConnect4(t *testing.T) {
	examples := []struct {
		in  string
		out bool
	}{
		{
			in:  "8|8|rryyyyrr|rrrr4|yyy5|y7|8|8",
			out: true,
		},
		{
			in:  "8|ry6|ry6|yyrr4|ry6|8|8|8",
			out: true,
		},
		{
			in:  "8|8|ryry4|yyyrrr2|ryrr4|y7|r7|8",
			out: true,
		},
		{
			in:  "y7|ry6|rry5|rrry4|yyy5|8|8|8",
			out: true,
		},
		{
			in:  "8|8|ryry4|yyyrrr2|yrrr4|y7|r7|8",
			out: false,
		},
		{
			in:  "y7|ry6|rry5|8|yyy5|8|8|8",
			out: false,
		},
	}
	for k, ex := range examples {
		b, _ := b.FromString(ex.in)
		if b.HasYellowConnect4() != ex.out {
			t.Errorf("example %d: got %v but want %v", k, b.HasYellowConnect4(), ex.out)
		}
	}
}

func TestHasRedConnect4(t *testing.T) {
	examples := []struct {
		in  string
		out bool
	}{
		{
			in:  "8|8|rryyyyrr|rrrr4|yyy5|y7|8|8",
			out: true,
		},
		{
			in:  "8|ry6|ry6|ryyr4|ry6|8|8|8",
			out: true,
		},
		{
			in:  "8|8|ryrr4|yyryrr2|yryy4|r7|r7|8",
			out: true,
		},
		{
			in:  "y7|ry6|rrr5|ryrr4|yyyr4|8|8|8",
			out: true,
		},
		{
			in:  "8|8|ryry4|yyyrrr2|yrrr4|y7|r7|8",
			out: false,
		},
		{
			in:  "y7|ry6|rry5|8|yyy5|8|8|8",
			out: false,
		},
	}
	for k, ex := range examples {
		b, _ := b.FromString(ex.in)
		if b.HasRedConnect4() != ex.out {
			t.Errorf("example %d: got %v but want %v", k, b.HasRedConnect4(), ex.out)
		}
	}
}
