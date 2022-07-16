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
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// TODO: address and port
const Address = "localhost:8080"

// GorillaDialer encapsulates a dialer from gorilla/websocket to comply with the Dialer interface.
type GorillaDialer struct {
	dialer *websocket.Dialer
}

// WARNING: untested.
func (gorillaDialer GorillaDialer) DialContext(ctx context.Context, urlStr string) (Conn, error) {
	// Necessary for casting *websocket.Conn to Conn interface.
	conn, _, err := gorillaDialer.dialer.DialContext(ctx, urlStr, nil)
	return conn, err
}

// WARNING: untested.
// TODO: find a way to test this or it will be painful...
func (gorillaDialer GorillaDialer) ServeContext(ctx context.Context, urlStr string) (Conn, error) {
	server := http.Server{Addr: Address}
	upgrader := websocket.Upgrader{}
	chanWS := make(chan *websocket.Conn)
	chanErr := make(chan error)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: checking should call a specialized function. String comparison is not quite enough.
		if r.RemoteAddr != urlStr {
			return
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		chanWS <- conn
	})

	// Listen and serve in goroutine to be able to select from context.
	go func() {
		chanErr <- server.ListenAndServe()
	}()

	// Wait for something to happen.
	select {
	case err := <-chanErr:
		return nil, fmt.Errorf("server error: %w", err)
	case <-ctx.Done():
		server.Shutdown(context.Background())
		<-chanErr // Empty error chan.
		return nil, fmt.Errorf("context done: %w", ctx.Err())
	case conn := <-chanWS:
		server.Shutdown(context.Background())
		<-chanErr // Empty error chan.
		return conn, nil
	}
}
