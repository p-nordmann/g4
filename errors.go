package g4

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
