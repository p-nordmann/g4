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
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

// TODO: refactor selects at the end of tests.

// Utility for dialing to wsHandler.
func dialWS(urlStr string) chan error {
	chanErr := make(chan error)
	go func() {
		ws, _, err := websocket.DefaultDialer.Dial(urlStr, nil)
		if err != nil {
			chanErr <- err
			return
		}
		ws.Close()
	}()
	return chanErr
}

// Utility for running ServeContext and returning output through channels.
func runServeContext(ctx context.Context, gorillaDialer GorillaDialer) (chan Conn, chan error) {
	chanConn := make(chan Conn)
	chanServeErr := make(chan error)
	go func() {
		conn, err := gorillaDialer.ServeContext(ctx, "localhost")
		if err != nil {
			chanServeErr <- err
		}
		chanConn <- conn
	}()
	return chanConn, chanServeErr
}

func TestServeContext(t *testing.T) {

	url := "localhost:5678"
	mockServer := &MockServer{
		server: &http.Server{Addr: url},
		cancel: make(chan context.Context),
	}
	gorillaDialer := GorillaDialer{
		server: mockServer,
	}

	t.Run("when dialed, should open a valid connection and shut server down", func(t *testing.T) {

		// Prepare mockServer.
		mockServer.countListenAndServe = 0
		mockServer.countShutdown = 0
		mockServer.mustCrash = false

		// Run gorillaDialer.ServeContext.
		chanConn, chanServeErr := runServeContext(context.Background(), gorillaDialer)

		// Dial server.
		chanErr := dialWS("ws://" + url)

		// Check that we correctly receive connection through channel.
		select {
		case conn := <-chanConn:
			time.Sleep(50 * time.Millisecond) // Leave time to shutdown. TODO: find a way to wait for shutdown.
			if conn == nil {
				t.Error("received <nil> but wanted a live connection")
			}
			if mockServer.countShutdown != 1 || mockServer.countListenAndServe != 1 {
				t.Errorf("invalid number of calls: expected (1, 1) but got (%d, %d)",
					mockServer.countShutdown, mockServer.countListenAndServe)
			}
		case err := <-chanServeErr:
			t.Errorf("serving error: %v", err)
		case err := <-chanErr:
			t.Errorf("dialing error: %v", err)
		}
	})

	t.Run("when context timeouts, should return an error and shut server down", func(t *testing.T) {

		// Prepare mockServer.
		mockServer.countListenAndServe = 0
		mockServer.countShutdown = 0
		mockServer.mustCrash = false

		// Run gorillaDialer.ServeContext with cancelable context.
		ctx, cancel := context.WithCancel(context.Background())
		chanConn, chanServeErr := runServeContext(ctx, gorillaDialer)

		// Cancel context.
		cancel()

		// Check that server returns an error.
		select {
		case <-chanConn:
			t.Error("server should not create connection")
		case err := <-chanServeErr:
			time.Sleep(50 * time.Millisecond) // Leave time to shutdown. TODO: find a way to wait for shutdown.
			if err == nil {
				t.Error("ServeContext should have returned an error but it was <nil>")
			}
			if mockServer.countShutdown != 1 || mockServer.countListenAndServe != 1 {
				t.Errorf("invalid number of calls: expected (1, 1) but got (%d, %d)",
					mockServer.countShutdown, mockServer.countListenAndServe)
			}
		}

	})

	t.Run("when server fails, should return an error", func(t *testing.T) {

		// Prepare mockServer.
		mockServer.countListenAndServe = 0
		mockServer.countShutdown = 0
		mockServer.mustCrash = true

		// Run gorillaDialer.ServeContext.
		chanConn, chanServeErr := runServeContext(context.Background(), gorillaDialer)

		// Check that server returns an error.
		select {
		case <-chanConn:
			t.Error("server should not create connection")
		case err := <-chanServeErr:
			if err == nil {
				t.Error("ServeContext should have returned an error but it was <nil>")
			}
			if mockServer.countShutdown != 0 || mockServer.countListenAndServe != 1 {
				t.Errorf("invalid number of calls: expected (0, 1) but got (%d, %d)",
					mockServer.countShutdown, mockServer.countListenAndServe)
			}
		}

	})
}

func TestGorillaDialer(t *testing.T) {

	gorillaDialer := GorillaDialer{}

	t.Run("calling DialContext with nil dialer should return an error", func(t *testing.T) {
		_, err := gorillaDialer.DialContext(context.Background(), "localhost")
		if err == nil {
			t.Error("expected error but got <nil>")
		}
	})

	t.Run("calling ServeContext with nil dialer should return an error", func(t *testing.T) {
		_, err := gorillaDialer.ServeContext(context.Background(), "localhost")
		if err == nil {
			t.Error("expected error but got <nil>")
		}
	})
}

func TestWsHandler(t *testing.T) {

	// Utility for retrieving handler and launching server.
	makeServer := func(remoteAddr string) (string, chan *websocket.Conn, func()) {
		chanWS := make(chan *websocket.Conn)
		handler := wsHandler(remoteAddr, chanWS)
		server := httptest.NewServer(http.HandlerFunc(handler))
		return server.URL, chanWS, server.Close
	}

	t.Run("dialing to wsHandler should create a valid websocket connection", func(t *testing.T) {

		// Start server.
		url, chanWS, close := makeServer("localhost")
		defer close()

		// Connect to the server in goroutine.
		chanErr := dialWS("ws" + strings.TrimPrefix(url, "http"))

		// Check that we correctly receive connection through channel.
		select {
		case conn := <-chanWS:
			if conn == nil {
				t.Error("received <nil> but wanted a live connection")
			}
		case err := <-chanErr:
			t.Errorf("dialing error: %v", err)
		}
	})

	t.Run("dialing to wsHandler from a wrong address should create a bad handshake", func(t *testing.T) {

		// Start server.
		url, chanWS, close := makeServer("127.0.0.2")
		defer close()

		// Connect to the server in goroutine.
		chanErr := dialWS("ws" + strings.TrimPrefix(url, "http"))

		// Check that we get a bad handshake error.
		select {
		case <-chanWS:
			t.Error("should not receive a connection from handler")
		case err := <-chanErr:
			if err == nil {
				t.Error("received <nil> error from dialer but should be bad handshake")
			}
		}
	})

	t.Run("dialing with bad handshake should not prevent from dialing after", func(t *testing.T) {

		// Start server.
		url, chanWS, close := makeServer("127.0.0.1")
		defer close()

		// Connect to the server in goroutine.
		chanErr := dialWS(url)

		// Check that we get a bad handshake error.
		select {
		case <-chanWS:
			t.Error("should not receive a connection from handler")
		case err := <-chanErr:
			if err == nil {
				t.Error("received <nil> error from dialer but should be bad handshake")
			}
		}

		// Connect with correct url.
		chanErr = dialWS("ws" + strings.TrimPrefix(url, "http"))

		// Check that we correctly receive connection through channel.
		select {
		case conn := <-chanWS:
			if conn == nil {
				t.Error("received <nil> but wanted a live connection")
			}
		case err := <-chanErr:
			t.Errorf("dialing error: %v", err)
		}
	})
}
