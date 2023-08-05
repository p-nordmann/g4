package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Canvas provides a view that draws an image.
type Canvas struct {
	width, height   int
	backgroundColor lipgloss.Color
	pixels          [][]lipgloss.Color
}

// Returns a new CanvasView object of requested size.
func NewCanvas(width, height int, backgroundColor lipgloss.Color) *Canvas {
	pixels := make([][]lipgloss.Color, height)
	for i := range pixels {
		pixels[i] = make([]lipgloss.Color, width)
	}
	return &Canvas{
		width:           width,
		height:          height,
		pixels:          pixels,
		backgroundColor: backgroundColor,
	}
}

// Draws a patch on the canvas.
func (canvas *Canvas) DrawPatch(x, y int, patch [][]lipgloss.Color) {
	for i := range patch {
		for j := range patch[i] {
			if x+i < canvas.height && y+j < canvas.width {
				canvas.pixels[x+i][y+j] = patch[i][j]
			}
		}
	}
}

type tallPixel struct {
	top, bottom lipgloss.Color
}

// TODO: use string builders.
func toString(line []tallPixel, backgroundColor lipgloss.Color) string {
	var empty lipgloss.Color
	var s = ""
	for _, tp := range line {
		style := lipgloss.NewStyle().Background(backgroundColor).Foreground(backgroundColor)
		if tp.bottom != empty {
			style = style.Background(tp.bottom)
		}
		if tp.top != empty {
			style = style.Foreground(tp.top)
		}
		s += style.Render("▀")
	}
	return s
}

// preRender builds a list of lines of tall pixels from the pixels data.
//
// This is useful because each line in the console is 2-pixel tall. We need to group lines 2-by-2 in
// order to draw the proper image in the console.
func (canvas *Canvas) preRender() [][]tallPixel {

	// If height is odd, we must add an empty line at the top.
	pixels := canvas.pixels
	if canvas.height%2 == 1 {
		pixels = make([][]lipgloss.Color, canvas.height+1)
		for i := range pixels {
			pixels[i] = make([]lipgloss.Color, canvas.width)
			if i > 0 {
				copy(pixels[i], canvas.pixels[i-1])
			}
		}
	}

	// Group lines of pixels two-by-two using an array of tallPixels.
	tallPixels := make([][]tallPixel, len(pixels)/2)
	for i := range tallPixels {
		tallPixels[i] = make([]tallPixel, canvas.width)
		for j := range tallPixels[i] {
			tallPixels[i][j] = tallPixel{top: pixels[2*i][j], bottom: pixels[2*i+1][j]}
		}
	}

	return tallPixels
}

func (canvas *Canvas) View() string {
	tallPixels := canvas.preRender()
	lines := make([]string, len(tallPixels))
	for k := range tallPixels {
		lines[k] = toString(tallPixels[k], canvas.backgroundColor)
	}
	return strings.Join(lines, "\n")
}
