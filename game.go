package g4

// Game is supposed to hold the data about the current game state.
//
// It can generate a list of possible move in its current state and it can be updated using moves.
//
// It can be queried for display information.
//
// TODO: display information.
type Game interface {
	// Apply performs a move from a game state.
	// It returns the new game state and a generator for next possible moves.
	Apply(Move) (Game, error)
	Generate() ([]Move, error)
}

type Board interface {
	// RotateLeft applies `times` left rotations on the board.
	//
	// It does not make the token drop according to new gravity.
	RotateLeft(times int) Board

	// ApplyGravity makes the token drop according to gravity.
	ApplyGravity() Board

	// AddToken adds a token on top of requested column.
	AddToken(column int, color Color) Board

	// HasYellowConnect4 returns whether the board has a connect 4 with yellow tokens.
	HasYellowConnect4() bool

	// HasRedConnect4 returns whether the board has a connect 4 with red tokens.
	HasRedConnect4() bool

	// Heights returns a list of heights for all the columns.
	Heights() [8]int
}
