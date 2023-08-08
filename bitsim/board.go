package bitsim

import (
	"fmt"
	"g4"
	"strconv"
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
	// StartingPosition represents the empty g4 board.
	StartingPosition string = "8|8|8|8|8|8|8|8"
)

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

type Board struct {
	yellowBits bitboard
	redBits    bitboard
}

// String returns the string representation of the board.
func (b Board) String() string {
	var s strings.Builder
	for col := 0; col < 8; col++ {
		void := 0
		for row := 0; row < 8; row++ {
			mask := one << (row + 8*col)
			if b.redBits&mask != 0 {
				if void > 0 {
					s.WriteString(strconv.Itoa(void))
					void = 0
				}
				s.WriteString("r")
			} else if b.yellowBits&mask != 0 {
				if void > 0 {
					s.WriteString(strconv.Itoa(void))
					void = 0
				}
				s.WriteString("y")
			} else {
				void++
			}
		}
		if void > 0 {
			s.WriteString(strconv.Itoa(void))
		}
		if col < 7 {
			s.WriteString("|")
		}
	}
	return s.String()
}

// Returns the total number of tokens on the board.
func (b Board) count() int {
	return b.yellowBits.count() + b.redBits.count()
}

// heights returns a list of heights for all the columns.
func (b Board) heights() [8]int {
	return [8]int{
		((b.yellowBits | b.redBits) & col0Mask).count(),
		((b.yellowBits | b.redBits) & col1Mask).count(),
		((b.yellowBits | b.redBits) & col2Mask).count(),
		((b.yellowBits | b.redBits) & col3Mask).count(),
		((b.yellowBits | b.redBits) & col4Mask).count(),
		((b.yellowBits | b.redBits) & col5Mask).count(),
		((b.yellowBits | b.redBits) & col6Mask).count(),
		((b.yellowBits | b.redBits) & col7Mask).count(),
	}
}

// hasYellowConnect4 returns whether the board has a connect 4 with yellow tokens.
func (b Board) hasYellowConnect4() bool {
	return b.yellowBits.hasConnect4()
}

// hasRedConnect4 returns whether the board has a connect 4 with red tokens.
func (b Board) hasRedConnect4() bool {
	return b.redBits.hasConnect4()
}

// RotateLeft applies `times` left rotations on the board.
//
// It does not make the token drop according to new gravity.
func (b Board) RotateLeft(times int) Board {
	for k := 0; k < times%4; k++ {
		b.yellowBits = b.yellowBits.rotateLeft()
		b.redBits = b.redBits.rotateLeft()
	}
	return b
}

// ApplyGravity makes the token drop according to gravity.
//
// Uses the naive approach:
//   - computes bits with a gap immediately below them
//   - drop one square
//   - iterate (8 times)
func (b Board) ApplyGravity() Board {
	for k := 0; k < 8; k++ {
		gaps := ^(b.yellowBits | b.redBits)
		yellowDrop := gaps & b.yellowBits.south()
		b.yellowBits = (b.yellowBits ^ yellowDrop.north()) | yellowDrop
		redDrop := gaps & b.redBits.south()
		b.redBits = (b.redBits ^ redDrop.north()) | redDrop
	}
	return b
}

// AddToken adds a token on top of requested column.
func (b Board) AddToken(column int, color g4.Color) Board {
	height := b.heights()[column]
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
