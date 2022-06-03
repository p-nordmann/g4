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
	// moveType indicates whether the move is a tilt or a token.
	moveType MoveType
	// gravity indicates the new direction for a tilt move.
	gravity Direction
	// column indicates the column that was played for a token move.
	column int
	// color indicates the color that was played for a token move.
	color Color
}

// Generator is the interface for generating a list of possible moves.
type Generator interface {
	Generate() ([]Move, error)
}

// Game is supposed to hold the data about the current game state.
//
// It can be updated using moves.
// It can be queried for display information.
// TODO: display information.
type Game interface {
	// Apply performs a move from a game state.
	// It returns the new game state and a generator for next possible moves.
	Apply(Move) (Game, Generator, error)
}
