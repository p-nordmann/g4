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

type MoveType byte

const (
	Token MoveType = iota // A move that places a new token on top of a column.
	Tilt                  // A move that changes the direction of gravity.
)

type Move struct {
	Type      MoveType
	Direction Direction
	Column    int
	Color     Color
}

func TokenMove(color Color, column int) Move {
	return Move{
		Type:   Token,
		Column: column,
		Color:  color,
	}
}

func TiltMove(color Color, direction Direction) Move {
	return Move{
		Type:      Tilt,
		Direction: direction,
		Color:     color,
	}
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
