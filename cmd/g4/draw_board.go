package main

import (
	"g4"
	"g4/bitsim"

	"github.com/charmbracelet/lipgloss"
)

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

type boardSize struct {
	tokenSize int
	stride    int
}

func fitBoard(spaceX, spaceY int) boardSize {
	// Compute dimensions.
	// If dimension is odd, decrement by one because it might not fit into the screen otherwise.
	height := min(spaceX, 2*spaceY)
	if height%2 == 1 {
		height--
	}

	// Compute the stride and square's sizes.
	// The stride will be responsive, and the size of squares will be basically 1/8th of what's left.
	var stride int
	if height >= 62 {
		stride = 2
	} else if height >= 23 {
		stride = 1
	}
	tokenSize := (height - 7*stride) / 8

	return boardSize{
		tokenSize: tokenSize,
		stride:    stride,
	}
}

func drawBoard(board bitsim.Board, s boardSize) string {

	size := 8*s.tokenSize + 7*s.stride

	// Get the board's array representation.
	array := toArray(board)

	// Draw the board on a CanvasView.
	canvas := NewCanvas(size, size, dark)
	for i := range array {
		for j := range array[i] {
			// Draw a nice square between holes.
			if s.stride > 0 && i > 0 && j > 0 {
				canvas.DrawPatch(
					s.tokenSize*i+s.stride*(i-1),
					s.tokenSize*j+s.stride*(j-1),
					makeSquaredPatch(s.stride, pinker),
				)
			}

			// If there is a token, draw a circle of correct color.
			// TODO: use a map, and if array[i][j] is anything that exists in the map, draw the corresponding color.
			switch array[i][j] {
			case g4.Yellow:
				canvas.DrawPatch(
					(s.tokenSize+s.stride)*i,
					(s.tokenSize+s.stride)*j,
					makeCircularPatch(s.tokenSize, yellow),
				)
			case g4.Red:
				canvas.DrawPatch(
					(s.tokenSize+s.stride)*i,
					(s.tokenSize+s.stride)*j,
					makeCircularPatch(s.tokenSize, red),
				)
			}
		}
	}

	// Return the rendered canvas.
	return canvas.View()
}

func toArray(b bitsim.Board) (array [8][8]g4.Color) {
	var col, row int
	for _, c := range b.String() {
		switch c {
		case '|':
			row = 0
			col++
		case 'r':
			array[7-row][col] = g4.Red
			row++
		case 'y':
			array[7-row][col] = g4.Yellow
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
			return
		}
	}
	return
}

func makeSquaredPatch(size int, col lipgloss.Color) [][]lipgloss.Color {
	patch := make([][]lipgloss.Color, size)
	for i := range patch {
		patch[i] = make([]lipgloss.Color, size)
		for j := range patch[i] {
			patch[i][j] = col
		}
	}
	return patch
}

func makeCircularPatch(size int, col lipgloss.Color) [][]lipgloss.Color {
	radius := float64(size) / 2
	patch := make([][]lipgloss.Color, size)
	for i := range patch {
		patch[i] = make([]lipgloss.Color, size)
		for j := range patch[i] {
			if (float64(i)+0.5-radius)*(float64(i)+0.5-radius)+(float64(j)+0.5-radius)*(float64(j)+0.5-radius) <= radius*radius {
				patch[i][j] = col
			}
		}
	}
	return patch
}
