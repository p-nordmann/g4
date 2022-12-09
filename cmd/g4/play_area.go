package main

import (
	"g4"
	"g4/bitsim"
	"image"
	"image/color"

	ti "github.com/knipferrc/teacup/image"
)

// TODO: build a tea.Model that can be initialized with size params and can draw boards.

const (
	squareWidth        = 6
	stride             = 2
	squareWidthPreview = 2
	stridePreview      = 1
)

func paintRect(x, y int, w, h int, col color.RGBA, img *image.RGBA) {
	for i := x; i < x+w; i++ {
		for j := y; j < y+h; j++ {
			img.SetRGBA(i, j, col)
		}
	}
}

func paintCircle(x, y int, w int, col color.RGBA, img *image.RGBA) {
	cX, cY := float64(2*x+w-1)/2, float64(2*y+w-1)/2
	for i := x; i < x+w; i++ {
		for j := y; j < y+w; j++ {
			if (float64(i)-cX)*(float64(i)-cX)+(float64(j)-cY)*(float64(j)-cY) < float64(w*w)/4 {
				img.SetRGBA(i, j, col)
			}
		}
	}
}

func toArray(b bitsim.Board) (array [8][8]g4.Color) {
	var col, row int
	for _, c := range b.String() {
		switch c {
		case '|':
			row = 0
			col++
		case 'r':
			array[col][7-row] = g4.Red
			row++
		case 'y':
			array[col][7-row] = g4.Yellow
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

func arrayToImage(array [8][8]g4.Color, squareWidth, stride int) image.Image {
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
			paintCircle((squareWidth+stride)*x, (squareWidth+stride)*y, squareWidth, col, img)
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
	board bitsim.Board
}

type previewArea struct {
	board bitsim.Board
}

func (m playArea) View() string {
	img := arrayToImage(toArray(m.board), squareWidth, stride)
	return ti.ToString(8*squareWidth+7*stride, img)
}

func (m previewArea) View() string {
	img := arrayToImage(toArray(m.board), squareWidthPreview, stridePreview)
	return ti.ToString(8*squareWidthPreview+7*stridePreview, img)
}
