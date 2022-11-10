package ws_test

import (
	"context"
	"encoding/json"
	"g4/ws"
)

type MockConn struct {

	// NextError will be returned as error by the next call to a method of MockConn.
	NextError error

	// NextJSON will be unmarshaled by the next call to ReadJSON.
	NextJSON []byte

	// LastJSON holds the JSON-marshaled value retrieved from the last call to WriteJSON.
	LastJSON []byte
}

func (conn *MockConn) Close() error {
	return conn.NextError
}

func (conn *MockConn) ReadJSON(v interface{}) error {
	if conn.NextError != nil {
		return conn.NextError
	}
	err := json.Unmarshal(conn.NextJSON, v)
	if err != nil {
		panic(err)
	}
	return nil
}

func (conn *MockConn) WriteJSON(v interface{}) error {
	if conn.NextError != nil {
		return conn.NextError
	}
	var err error
	conn.LastJSON, err = json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return nil
}

type ConnError struct {
	conn ws.Conn
	err  error
}

type MockConnector struct {
	DialContextReturnValue  ConnError
	DialContextCallCount    int
	ServeContextReturnValue ConnError
	ServeContextCallCount   int
}

func (connector *MockConnector) DialContext(ctx context.Context, urlStr string) (ws.Conn, error) {
	connector.DialContextCallCount++
	return connector.DialContextReturnValue.conn, connector.DialContextReturnValue.err
}

func (connector *MockConnector) ServeContext(ctx context.Context, urlStr string) (ws.Conn, error) {
	connector.ServeContextCallCount++
	return connector.ServeContextReturnValue.conn, connector.ServeContextReturnValue.err
}
