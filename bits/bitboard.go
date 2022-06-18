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
	"errors"
	"math/bits"
)

type bitboard uint64

const (
	one        bitboard = 1
	northMask  bitboard = ^(one | one<<8 | one<<16 | one<<24 | one<<32 | one<<40 | one<<48 | one<<56)
	north2Mask bitboard = ^(^northMask | (^northMask << 1))
	north3Mask bitboard = ^(^northMask | (^northMask << 1) | (^northMask << 2))
	southMask  bitboard = ^(^northMask << 7)
)

// bitboardFromString returns a bitboard built from a description string.
//
// Format is the following:
// col1|col2|col3|...|col8
// Where each column is an alternation of ´x´ and integers, ´x´ denoting a 1 and an integer denoting a sequence of 0.
//
// NB: multiple integers one after the other is also valid.
// example: x7|... is equivalent to x1123|... or x43|...
func bitboardFromString(s string) (bitboard, error) {
	var b bitboard
	var col, row int

	for _, c := range s {
		switch c {
		case '|':
			if row != 8 {
				return 0, errors.New("invalid number of rows")
			}
			row = 0
			col++
		case 'x':
			b |= one << (row + 8*col)
			row++
		case '1':
			row += 1
		case '2':
			row += 2
		case '3':
			row += 3
		case '4':
			row += 4
		case '5':
			row += 5
		case '6':
			row += 6
		case '7':
			row += 7
		case '8':
			row += 8
		default:
			return 0, errors.New("invalid character")
		}
	}

	if col != 7 {
		return 0, errors.New("invalid number of columns")
	}
	return b, nil
}

func (b bitboard) North() bitboard {
	return (b << 1) & northMask
}

func (b bitboard) North2() bitboard {
	return (b << 2) & north2Mask
}

func (b bitboard) North3() bitboard {
	return (b << 3) & north3Mask
}

func (b bitboard) West() bitboard {
	return b >> 8
}

func (b bitboard) South() bitboard {
	return (b >> 1) & southMask
}

func (b bitboard) East() bitboard {
	return b << 8
}

func (b bitboard) East2() bitboard {
	return b << 16
}

func (b bitboard) East3() bitboard {
	return b << 24
}

func (b bitboard) NorthWest() bitboard {
	return (b >> 7) & northMask
}

func (b bitboard) NorthWest2() bitboard {
	return (b >> 14) & north2Mask
}

func (b bitboard) NorthWest3() bitboard {
	return (b >> 21) & north3Mask
}

func (b bitboard) NorthEast() bitboard {
	return (b << 9) & northMask
}

func (b bitboard) NorthEast2() bitboard {
	return (b << 18) & north2Mask
}

func (b bitboard) NorthEast3() bitboard {
	return (b << 27) & north3Mask
}

// HasConnect4 returns whether the bitboard has a connect 4 pattern.
// The pattern can occur horizontally, vertically or diagonally.
func (b bitboard) HasConnect4() bool {
	v4 := b & b.North() & b.North2() & b.North3()
	h4 := b & b.East() & b.East2() & b.East3()
	ld4 := b & b.NorthWest() & b.NorthWest2() & b.NorthWest3()
	rd4 := b & b.NorthEast() & b.NorthEast2() & b.NorthEast3()
	return (v4 | h4 | ld4 | rd4) != 0
}

// Count returns the number of 1 in the bitboard.
func (b bitboard) Count() int {
	return bits.OnesCount64(uint64(b))
}

// getColumn returns the byte corresponding to a column.
func (b bitboard) getColumn(column int) byte {
	return byte(b >> (column * 8))
}

// RotateLeft rotates the bitboard 90 degrees left.
//
// It proceeds column by column, using rotationLookup (lazy dev 😁).
func (b bitboard) RotateLeft() bitboard {
	var rotatedBitboard bitboard
	for column := 0; column < 8; column++ {
		rotatedBitboard |= rotationLookup[b.getColumn(column)] << column
	}
	return rotatedBitboard
}
