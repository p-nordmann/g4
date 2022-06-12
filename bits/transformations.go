package bits

import (
	"g4"
)

// RotateLeft applies `times` left rotations on the board.
//
// It does not make the token drop according to new gravity.
func (b Board) RotateLeft(times int) Board {
	return b
}

// ApplyGravity makes the token drop according to gravity.
func (b Board) ApplyGravity() Board {
	return b
}

// AddToken adds a token on top of requested column.
func (b Board) AddToken(column int, color g4.Color) Board {
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
