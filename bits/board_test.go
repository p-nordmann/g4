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

package bits

import "testing"

func TestFromString(t *testing.T) {
	one := bitboard(1)
	ones8 := bitboard(255)
	examples := []struct {
		in  string
		out Board
	}{
		{
			in: "8|8|8|8|8|8|8|8",
			out: Board{
				yellowBits: 0,
				redBits:    0,
			},
		},
		{
			in: "y7|8|8|8|8|8|8|8",
			out: Board{
				yellowBits: one,
				redBits:    0,
			},
		},
		{
			in: "r6r|8|8|8|8|yyyyyyyy|8|8",
			out: Board{
				yellowBits: ones8 << 40,
				redBits:    one | one<<7,
			},
		},
		{
			in: "8|8|8|8|rrryr3|ryyyyyr1|r7|yr6",
			out: Board{
				yellowBits: (one<<3)<<32 | // column 4
					(one<<1|one<<2|one<<3|one<<4|one<<5)<<40 | // column 5
					(one)<<56, // column 7
				redBits: (one|one<<1|one<<2|one<<4)<<32 | // column 4
					(one|one<<6)<<40 | // column 5
					(one)<<48 | // column 6
					(one<<1)<<56, // column 7
			},
		},
	}
	for k, ex := range examples {
		if out, err := FromString(ex.in); err != nil || out != ex.out {
			t.Errorf("example %d: got (%v, %v) but want (%v, %v)", k, out, err, ex.out, nil)
		}
	}
}

func TestFromStringError(t *testing.T) {
	examples := []string{
		"y7|ry6|r7|8|8|8|8",
		"yrryyrry|8|8|8|8|rrryyyyryr|8|8",
		"zry5|8|8|8|8|8|8|8",
	}
	for k, ex := range examples {
		if out, err := FromString(ex); err == nil {
			t.Errorf("example %d: got (%v, %v) but expected error", k, out, err)
		}
	}
}
