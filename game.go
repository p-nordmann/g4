package g4

// TODO: Board interface.

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

type ErrorGameOver struct {
	Winner Color
}

func (err ErrorGameOver) Error() string {
	switch err.Winner {
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
