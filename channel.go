/*
G4 is an open-source board game inspired by the popular game of connect-4.
Copyright (C) 2022  Pierre-Louis Nordmann

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package g4

import (
	"context"
)

// Position provides a standard way to describe a position.
type Position struct {
	BoardStr      string
	Direction     Direction
	ColorWithMove Color
}

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
}
