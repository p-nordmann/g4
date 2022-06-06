package g4

// MoveType indicates the kind of a move.
//
// A move can be one of:
//	- Tilt: changing the direction of the gravity.
//	- Token: placing a new token on top of one column.
type MoveType byte

const (
	Token MoveType = iota
	Tilt
)

type Move struct {
	// MoveType indicates whether the move is a tilt or a token.
	moveType MoveType
	// Gravity indicates the new direction for a tilt move.
	gravity Direction
	// Column indicates the column that was played for a token move.
	column int
	// Color indicates the color that was played for a token move.
	color Color
}

func Base() Move {
	return Move{}
}

func (m Move) Token() Move {
	m.moveType = Token
	return m
}

func (m Move) Tilt() Move {
	m.moveType = Tilt
	return m
}

func (m Move) Column(column int) Move {
	m.column = column
	return m
}

func (m Move) Color(color Color) Move {
	m.color = color
	return m
}

func (m Move) Gravity(gravity Direction) Move {
	m.gravity = gravity
	return m
}

func TokenMove(color Color, column int) Move {
	return Base().Token().Color(color).Column(column)
}

func TiltMove(direction Direction) Move {
	return Base().Tilt().Gravity(direction)
}
