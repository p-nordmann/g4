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
//   - computes bits with a gap immediately below them
//   - drop one square
//   - iterate (8 times)
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
