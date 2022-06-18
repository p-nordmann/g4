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
	"fmt"
	"strings"
)

const (
	col0Mask bitboard = 255
	col1Mask bitboard = col0Mask << 8
	col2Mask bitboard = col1Mask << 8
	col3Mask bitboard = col2Mask << 8
	col4Mask bitboard = col3Mask << 8
	col5Mask bitboard = col4Mask << 8
	col6Mask bitboard = col5Mask << 8
	col7Mask bitboard = col6Mask << 8
)

type Board struct {
	yellowBits bitboard
	redBits    bitboard
}

// FromString returns a board built from a description string.
//
// Format is the following:
// col1|col2|col3|...|col8
// Where each column is an alternation of 'y', 'r' and integers,
// 'y' and 'r' respectively denoting a yellow or red token and
// an integer denoting a sequence of 0.
//
// NB: multiple integers one after the other is also valid.
// example: y7|... is equivalent to y1123|... or r43|...
func FromString(s string) (b Board, err error) {
	// Parse yellow bits.
	yellowString := strings.ReplaceAll(
		strings.ReplaceAll(s, "r", "1"),
		"y",
		"x",
	)
	b.yellowBits, err = bitboardFromString(yellowString)
	if err != nil {
		return b, fmt.Errorf("error parsing yellow bits: %w", err)
	}

	// Parse red bits.
	// NB: no need to check for errors, as it can never happen here.
	redString := strings.ReplaceAll(
		strings.ReplaceAll(s, "r", "x"),
		"y",
		"1",
	)
	b.redBits, _ = bitboardFromString(redString)

	return
}

// Heights returns a list of heights for all the columns.
func (b Board) Heights() [8]int {
	return [8]int{
		((b.yellowBits | b.redBits) & col0Mask).Count(),
		((b.yellowBits | b.redBits) & col1Mask).Count(),
		((b.yellowBits | b.redBits) & col2Mask).Count(),
		((b.yellowBits | b.redBits) & col3Mask).Count(),
		((b.yellowBits | b.redBits) & col4Mask).Count(),
		((b.yellowBits | b.redBits) & col5Mask).Count(),
		((b.yellowBits | b.redBits) & col6Mask).Count(),
		((b.yellowBits | b.redBits) & col7Mask).Count(),
	}
}

// HasYellowConnect4 returns whether the board has a connect 4 with yellow tokens.
func (b Board) HasYellowConnect4() bool {
	return b.yellowBits.HasConnect4()
}

// HasRedConnect4 returns whether the board has a connect 4 with red tokens.
func (b Board) HasRedConnect4() bool {
	return b.redBits.HasConnect4()
}
