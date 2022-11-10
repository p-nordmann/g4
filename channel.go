package g4

import (
	"context"
)

type Channel interface {

	// SendMove sends the selected move to the opponent.
	SendMove(move Move) error // TODO: too complex to send and receive position, just do the bare minimum...

	// ReadMove waits for the oponent to send a move and receives it.
	ReadMove() (Move, error)

	// ConnectWait tries to connect to the provided URL.
	//
	// In case the attempt to connect timeouts, it will wait for a client
	// from the corresponding URL.
	ConnectWait(ctx context.Context, urlStr string) error

	// Close is necessary for closing the inner connections.
	Close() error

	// ChooseColor communicates with the peer to assign each player a color.
	ChooseColor() (Color, error)
}
