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
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// wsHandler returns the handler corresponding to urlStr and sending to chanWS.
func wsHandler(urlStr string, chanWS chan *websocket.Conn) func(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}
	return func(w http.ResponseWriter, r *http.Request) {
		if !areUrlEqual(r.RemoteAddr, urlStr) {
			return
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		chanWS <- conn
	}
}

var handler func(w http.ResponseWriter, r *http.Request)

func handlerWrapper(w http.ResponseWriter, r *http.Request) {
	if handler == nil {
		return
	}
	handler(w, r)
}

// GorillaDialer encapsulates a dialer from gorilla/websocket to comply with the Dialer interface.
type GorillaDialer struct {
	dialer *websocket.Dialer
	server Server
}

// WARNING: untested.
func (gorillaDialer GorillaDialer) DialContext(ctx context.Context, urlStr string) (Conn, error) {

	// Avoid nil dialer.
	if gorillaDialer.dialer == nil {
		return nil, errors.New("<nil> dialer")
	}

	conn, _, err := gorillaDialer.dialer.DialContext(ctx, urlStr, nil)
	return conn, err
}

func (gorillaDialer GorillaDialer) ServeContext(ctx context.Context, urlStr string) (Conn, error) {

	// Avoid nil server.
	if gorillaDialer.server == nil {
		return nil, errors.New("<nil> server")
	}

	// Channels for feedback.
	chanWS := make(chan *websocket.Conn)
	chanErr := make(chan error)

	// Set handler up.
	// TODO: cleaner handler. With the net.http package we can only have one handler for /.
	//   This is why we use the following workaround. A cleaner way should be designed instead.
	if handler == nil {
		http.HandleFunc("/", handlerWrapper)
	}
	handler = wsHandler(urlStr, chanWS)

	// Run server.
	go func() {
		chanErr <- gorillaDialer.server.ListenAndServe()
	}()

	// Handle feedback.
	select {
	case err := <-chanErr:
		return nil, fmt.Errorf("server error: %w", err)
	case <-ctx.Done():
		gorillaDialer.server.Shutdown(context.Background())
		<-chanErr
		return nil, fmt.Errorf("context done: %w", ctx.Err())
	case conn := <-chanWS:
		go func() {
			// This will take until the websocket closes to shutdown.
			gorillaDialer.server.Shutdown(context.Background())
			<-chanErr
		}()
		return conn, nil
	}
}
