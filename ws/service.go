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

package ws

import (
	"context"
)

// Conn is intended to represent the functions needed from a websocket connection.
//
// It is designed so gorilla/websocket's *Conn implements this interface.
type Conn interface {
	Close() error
	ReadJSON(v interface{}) error
	WriteJSON(v interface{}) error
}

// Connector is intended to represent the functions needed to dial or serve through websockets.
type Connector interface {
	DialContext(ctx context.Context, urlStr string) (Conn, error)
	ServeContext(ctx context.Context, urlStr string) (Conn, error)
}
