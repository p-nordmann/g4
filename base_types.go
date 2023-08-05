package g4

// Color describes the color of a token.
type Color byte

const (
	Empty Color = iota
	Yellow
	Red
)

// Direction describes the direction of gravity.
type Direction int

const (
	UP Direction = iota
	LEFT
	DOWN
	RIGHT
)

// MoveType indicates the kind of a move.
//
// A move can be one of:
//   - Tilt: changing the direction of the gravity.
//   - Token: placing a new token on top of one column.
type MoveType byte

const (
	Token MoveType = iota
	Tilt
)

// TODO: proper naming for attributes.
type Move struct {
	// Type indicates whether the move is a tilt or a token.
	Type MoveType
	// Direction indicates the new direction for a tilt move.
	Direction Direction
	// ColumnIdx indicates the column that was played for a token move.
	ColumnIdx int
	// Col indicates the color that was played for a token move.
	Col Color
}

func (m Move) Token() Move {
	m.Type = Token
	return m
}

func (m Move) Tilt() Move {
	m.Type = Tilt
	return m
}

func (m Move) Column(column int) Move {
	m.ColumnIdx = column
	return m
}

func (m Move) Color(color Color) Move {
	m.Col = color
	return m
}

func (m Move) Gravity(gravity Direction) Move {
	m.Direction = gravity
	return m
}

func TokenMove(color Color, column int) Move {
	return Move{}.Token().Color(color).Column(column)
}

func TiltMove(color Color, direction Direction) Move {
	return Move{}.Color(color).Tilt().Gravity(direction)
}

type ErrorGameOver Color

func (err ErrorGameOver) Error() string {
	switch Color(err) {
	case Yellow:
		return "game is over - yellow wins"
	case Red:
		return "game is over - red wins"
	default:
		return "game is over - draw"
	}
}

type ErrorInvalidMove struct{}

func (err ErrorInvalidMove) Error() string {
	return "invalid move"
}
