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

// TODO: add a test for FromString and documentation.
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
	redString := strings.ReplaceAll(
		strings.ReplaceAll(s, "r", "x"),
		"y",
		"1",
	)
	b.redBits, err = bitboardFromString(redString)
	if err != nil {
		return b, fmt.Errorf("error parsing red bits: %w", err)
	}

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
