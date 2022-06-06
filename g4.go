//

package g4

const (
	RowCount = 8
)

// Color describes the color of a token.
//
// There are really only two colors available but
// as we work with 64-bit machines we might as well use memory.
type Color byte

const (
	Empty Color = iota
	Yellow
	Red
)

// Column describes an 8 token-long column.
//
// Its size is (conveniently) 8*size(Color).
type Column uint64

// Direction describes the direction of gravity.
type Direction int

const (
	UP Direction = iota
	LEFT
	DOWN
	RIGHT
)
