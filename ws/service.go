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

// Server is intented to fit http.Server.
type Server interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}
