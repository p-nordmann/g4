package main

import (
	"g4"
	"g4/bitsim"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// TODO: use a customizable color map.

const (
	yellow = "#aaaa00"
	red    = "#dd0000"
	grey   = "#7e1e5e"
	dark   = "#0a0a0a" // TODO: transparent background for canvas.
)

// BoardModel is a model for displaying a g4 board.
type BoardModel struct {
	width, height int
	game          bitsim.Game
}

func (m BoardModel) Init() tea.Cmd {
	return nil
}

func (m BoardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(Size); ok {
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m BoardModel) View() string {

	// Compute dimensions.
	// If dimension is odd, decrement by one because it might not fit into the screen otherwise.
	height := min(m.width, 2*m.height)
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

	// Center the canvas vertically and horizontally.
	canvasSize := 8*tokenSize + 7*stride
	if canvasSize%2 == 1 {
		canvasSize++
	}
	verticalPadding := m.height - canvasSize/2
	horizontalPadding := m.width - canvasSize
	return lipgloss.NewStyle().
		MarginTop(verticalPadding / 2).
		MarginBottom(verticalPadding - verticalPadding/2).
		MarginLeft(horizontalPadding / 2).
		Render(m.draw(tokenSize, stride))
}

func (m BoardModel) draw(tokenSize, stride int) string {

	size := 8*tokenSize + 7*stride

	// Get the board's array representation.
	array := toArray(m.game.Board)

	// Draw the board on a CanvasView.
	canvas := NewCanvas(size, size, lipgloss.Color(dark))
	for i := range array {
		for j := range array[i] {
			// Draw a nice square between holes.
			if stride > 0 && i > 0 && j > 0 {
				canvas.DrawPatch(
					tokenSize*i+stride*(i-1),
					tokenSize*j+stride*(j-1),
					makeSquaredPatch(stride, lipgloss.Color(grey)),
				)
			}

			// If there is a token, draw a circle of correct color.
			// TODO: use a map, and if array[i][j] is anything that exists in the map, draw the corresponding color.
			switch array[i][j] {
			case g4.Yellow:
				canvas.DrawPatch(
					(tokenSize+stride)*i,
					(tokenSize+stride)*j,
					makeCircularPatch(tokenSize, lipgloss.Color(yellow)),
				)
			case g4.Red:
				canvas.DrawPatch(
					(tokenSize+stride)*i,
					(tokenSize+stride)*j,
					makeCircularPatch(tokenSize, lipgloss.Color(red)),
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
