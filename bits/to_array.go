package bits

import "g4"

func (b Board) ToArray() [8][8]g4.Color {
	var array [8][8]g4.Color
	// Beware the coordinates system is not the same as for bitboards.
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			mask := one << (row + 8*col)
			if b.redBits&mask != 0 {
				array[col][7-row] = g4.Red
			} else if b.yellowBits&mask != 0 {
				array[col][7-row] = g4.Yellow
			} else {
				array[col][7-row] = g4.Empty
			}
		}
	}
	return array
}
