package g4

import "fmt"

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

func (m Move) String() string {
	if m.Type == Token {
		return fmt.Sprint(m.ColumnIdx)
	}
	if m.Type == Tilt {
		switch m.Direction {
		case UP:
			return "UP"
		case LEFT:
			return "LEFT"
		case RIGHT:
			return "RIGHT"
		case DOWN:
			return "DOWN"

		}
	}
	return "Unknown"
}

func Base() Move {
	return Move{}
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
	return Base().Token().Color(color).Column(column)
}

func TiltMove(direction Direction) Move {
	return Base().Tilt().Gravity(direction)
}
