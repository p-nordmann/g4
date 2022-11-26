package g4

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

// Direction describes the direction of gravity.
//
// TODO: only use direction information for rotations. We always consider the gravity to be downward.
type Direction int

const (
	UP Direction = iota
	LEFT
	DOWN
	RIGHT
)
