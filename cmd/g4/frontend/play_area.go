package frontend

import (
	"g4"
	"image"
	"image/color"

	ti "github.com/knipferrc/teacup/image"
)

const (
	squareWidth = 6
	stride      = 2
)

func paintRect(x, y int, w, h int, col color.RGBA, img *image.RGBA) {
	for i := x; i < x+w; i++ {
		for j := y; j < y+h; j++ {
			img.SetRGBA(i, j, col)
		}
	}
}

func arrayToImage(array [8][8]g4.Color, squareWidth int) image.Image {
	rect := image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{8*squareWidth + 7*stride, 8*squareWidth + 7*stride},
	}
	img := image.NewRGBA(rect)
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			var col color.RGBA
			switch array[x][y] {
			case g4.Red:
				col = color.RGBA{R: 230, A: 255}
			case g4.Yellow:
				col = color.RGBA{R: 200, G: 200, A: 255}
			default:
				continue
			}
			paintRect((squareWidth+stride)*x, (squareWidth+stride)*y+1, squareWidth, squareWidth-2, col, img)
			paintRect((squareWidth+stride)*x+1, (squareWidth+stride)*y, squareWidth-2, squareWidth, col, img)
			paintRect((squareWidth+stride)*x+1, (squareWidth+stride)*y+1, squareWidth-2, squareWidth-2, col, img)
		}
	}

	for x := 1; x < 8; x++ {
		for y := 1; y < 8; y++ {
			col := color.RGBA{R: 30, G: 30, B: 30, A: 255}
			paintRect((squareWidth+stride)*x-stride, (squareWidth+stride)*y-stride, stride, stride, col, img)
		}
	}
	return img
}

type playArea struct {
	Board     [8][8]g4.Color
	Direction g4.Direction
}

func (m playArea) View() string {
	img := arrayToImage(m.Board, squareWidth)
	return ti.ToString(8*squareWidth+7*stride, img)
}
