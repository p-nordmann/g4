package g4

type Color byte

const (
	Empty Color = iota
	Yellow
	Red
)

type MoveType byte

const (
	Token MoveType = iota // A move that places a new token on top of a column.
	Tilt                  // A move that changes the direction of gravity.
)

type Direction int

const (
	LEFT Direction = iota
	DOWN
	RIGHT
)

type Move struct {
	Color     Color
	Type      MoveType
	Direction Direction
	Column    int
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

type Draw struct{}
type YellowWins struct{}
type RedWins struct{}

func (err Draw) Error() string {
	return "draw"
}
func (err YellowWins) Error() string {
	return "yellow wins"
}
func (err RedWins) Error() string {
	return "red wins"
}

type ErrorInvalidMove struct{}

func (err ErrorInvalidMove) Error() string {
	return "invalid move"
}
