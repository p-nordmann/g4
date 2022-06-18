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

import (
	"g4"
)

// RotateLeft applies `times` left rotations on the board.
//
// It does not make the token drop according to new gravity.
func (b Board) RotateLeft(times int) g4.Board {
	for k := 0; k < times%4; k++ {
		b.yellowBits = b.yellowBits.RotateLeft()
		b.redBits = b.redBits.RotateLeft()
	}
	return b
}

// ApplyGravity makes the token drop according to gravity.
//
// Uses the naive approach:
// 	- computes bits with a gap immediately below them
// 	- drop one square
// 	- iterate (8 times)
func (b Board) ApplyGravity() g4.Board {
	for k := 0; k < 8; k++ {
		gaps := ^(b.yellowBits | b.redBits)
		yellowDrop := gaps & b.yellowBits.South()
		b.yellowBits = (b.yellowBits ^ yellowDrop.North()) | yellowDrop
		redDrop := gaps & b.redBits.South()
		b.redBits = (b.redBits ^ redDrop.North()) | redDrop
	}
	return b
}

// AddToken adds a token on top of requested column.
func (b Board) AddToken(column int, color g4.Color) g4.Board {
	height := b.Heights()[column]
	if height < 8 {
		switch color {
		case g4.Yellow:
			b.yellowBits |= one << (height + column*8)
		case g4.Red:
			b.redBits |= one << (height + column*8)
		}
	}
	return b
}
