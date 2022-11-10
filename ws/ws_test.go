package ws_test

import (
	"context"
	"errors"
	"g4"
	"g4/ws"
	"testing"
)

var exampleMove = g4.TokenMove(g4.Red, 8)

func TestCloseErrorOnNilConn(t *testing.T) {
	ch := ws.NewWithConnector(ws.DefautConfig, &MockConnector{})
	err := ch.Close()
	if err == nil {
		t.Error("expected error but got <nil>")
	}
}

func TestCloseSuccessful(t *testing.T) {
	conn := &MockConn{NextError: nil}
	connector := &MockConnector{
		DialContextReturnValue: ConnError{
			conn: conn,
			err:  nil,
		},
	}
	ch := ws.NewWithConnector(ws.DefautConfig, connector)

	err := ch.ConnectWait(context.Background(), "")
	if err != nil {
		t.Errorf("ConnectWait not supposed to break: got err=%v", err)
	}

	err = ch.Close()
	if err != nil {
		t.Errorf("error in Close: want <nil> but got %v", err)
	}
}

func TestCloseError(t *testing.T) {
	conn := &MockConn{NextError: errors.New("error closing")}
	connector := &MockConnector{
		DialContextReturnValue: ConnError{
			conn: conn,
			err:  nil,
		},
	}
	ch := ws.NewWithConnector(ws.DefautConfig, connector)

	err := ch.ConnectWait(context.Background(), "")
	if err != nil {
		t.Errorf("ConnectWait not supposed to break: got err=%v", err)
	}

	err = ch.Close()
	if err == nil {
		t.Error("expected error but got <nil>")
	}
}

func TestConnectWaitDialFailure(t *testing.T) {
	connector := &MockConnector{
		DialContextReturnValue: ConnError{
			conn: nil,
			err:  errors.New("connection timeout"),
		},
		ServeContextReturnValue: ConnError{
			conn: nil,
			err:  errors.New("connection timeout"),
		},
	}
	ch := ws.NewWithConnector(ws.DefautConfig, connector)

	err := ch.ConnectWait(context.Background(), "")
	if err == nil {
		t.Error("expected error but got <nil>")
	}

	if connector.DialContextCallCount != 1 || connector.ServeContextCallCount != 1 {
		t.Errorf("unexpected number of calls: want (1, 1) but got (%d, %d)",
			connector.DialContextCallCount, connector.ServeContextCallCount)
	}
}

func TestConnectWaitDialSuccess(t *testing.T) {
	connector := &MockConnector{
		DialContextReturnValue: ConnError{
			conn: nil,
			err:  nil,
		},
		ServeContextReturnValue: ConnError{
			conn: nil,
			err:  errors.New("connection timeout"),
		},
	}
	ch := ws.NewWithConnector(ws.DefautConfig, connector)

	err := ch.ConnectWait(context.Background(), "")
	if err != nil {
		t.Errorf("error in ConnectWait: %v", err)
	}

	if connector.DialContextCallCount != 1 || connector.ServeContextCallCount != 0 {
		t.Errorf("unexpected number of calls: want (1, 0) but got (%d, %d)",
			connector.DialContextCallCount, connector.ServeContextCallCount)
	}
}

func TestConnectServeSuccess(t *testing.T) {
	connector := &MockConnector{
		DialContextReturnValue: ConnError{
			conn: nil,
			err:  errors.New("connection timeout"),
		},
		ServeContextReturnValue: ConnError{
			conn: nil,
			err:  nil,
		},
	}
	ch := ws.NewWithConnector(ws.DefautConfig, connector)

	err := ch.ConnectWait(context.Background(), "")
	if err != nil {
		t.Errorf("error in ConnectWait: %v", err)
	}

	if connector.DialContextCallCount != 1 || connector.ServeContextCallCount != 1 {
		t.Errorf("unexpected number of calls: want (1, 1) but got (%d, %d)",
			connector.DialContextCallCount, connector.ServeContextCallCount)
	}
}

func TestReadMoveError(t *testing.T) {
	conn := &MockConn{}
	connector := &MockConnector{
		DialContextReturnValue: ConnError{
			conn: conn,
			err:  nil,
		},
	}
	ch := ws.NewWithConnector(ws.DefautConfig, connector)

	err := ch.ConnectWait(context.Background(), "")
	if err != nil {
		t.Errorf("ConnectWait not supposed to break: got err=%v", err)
	}

	// Setup ReadJSON return values.
	conn.NextError = errors.New("nothing to read")
	_, err = ch.ReadMove()
	if err == nil {
		t.Error("expected error but got <nil>")
	}
}

func TestReadMoveSuccess(t *testing.T) {
	conn := &MockConn{}
	connector := &MockConnector{
		DialContextReturnValue: ConnError{
			conn: conn,
			err:  nil,
		},
	}
	ch := ws.NewWithConnector(ws.DefautConfig, connector)

	err := ch.ConnectWait(context.Background(), "")
	if err != nil {
		t.Errorf("ConnectWait not supposed to break: got err=%v", err)
	}

	// Setup ReadJSON return values.
	conn.NextError = nil
	conn.WriteJSON(exampleMove)
	conn.NextJSON = conn.LastJSON

	// Call ReadMove and test return values.
	mp, err := ch.ReadMove()
	if err != nil {
		t.Errorf("error in ReadMove: %v", err)
	}
	if mp != exampleMove {
		t.Errorf("wrong return value for ReadMove: want %v but got %v",
			exampleMove, mp)
	}
}

func TestSendMoveError(t *testing.T) {
	conn := &MockConn{}
	connector := &MockConnector{
		DialContextReturnValue: ConnError{
			conn: conn,
			err:  nil,
		},
	}
	ch := ws.NewWithConnector(ws.DefautConfig, connector)

	err := ch.ConnectWait(context.Background(), "")
	if err != nil {
		t.Errorf("ConnectWait not supposed to break: got err=%v", err)
	}

	// Setup WriteJSON return values.
	conn.NextError = errors.New("nothing to write")
	err = ch.SendMove(exampleMove)
	if err == nil {
		t.Error("expected error but got <nil>")
	}
}

func TestSendMoveSuccess(t *testing.T) {
	conn := &MockConn{}
	connector := &MockConnector{
		DialContextReturnValue: ConnError{
			conn: conn,
			err:  nil,
		},
	}
	ch := ws.NewWithConnector(ws.DefautConfig, connector)

	err := ch.ConnectWait(context.Background(), "")
	if err != nil {
		t.Errorf("ConnectWait not supposed to break: got err=%v", err)
	}

	// Setup ReadJSON return values.
	conn.NextError = nil
	conn.WriteJSON(exampleMove)
	expectedJSON := conn.LastJSON
	conn.LastJSON = nil

	// Call ReadMove and test return values.
	err = ch.SendMove(exampleMove)
	if err != nil {
		t.Errorf("error in SendMove: %v", err)
	}
	if string(conn.LastJSON) != string(expectedJSON) {
		t.Errorf("wrong JSON sent: want %v but got %v",
			expectedJSON, conn.LastJSON)
	}
}
