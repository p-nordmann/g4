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

package ws_test

import (
	"context"
	"errors"
	"g4"
	"g4/ws"
	"testing"
)

var exampleMP = g4.MoveAndPosition{
	Move: g4.TokenMove(g4.Red, 8),
	Position: g4.Position{
		BoardStr:      "ry6|8|8|8|8|8|8|rryy4",
		Direction:     g4.UP,
		ColorWithMove: g4.Red,
	},
}

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

func TestReadMoveAndPositionError(t *testing.T) {
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
	_, err = ch.ReadMoveAndPosition()
	if err == nil {
		t.Error("expected error but got <nil>")
	}
}

func TestReadMoveAndPositionSuccess(t *testing.T) {
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
	conn.WriteJSON(exampleMP)
	conn.NextJSON = conn.LastJSON

	// Call ReadMoveAndPosition and test return values.
	mp, err := ch.ReadMoveAndPosition()
	if err != nil {
		t.Errorf("error in ReadMoveAndPosition: %v", err)
	}
	if mp != exampleMP {
		t.Errorf("wrong return value for ReadMoveAndPosition: want %v but got %v",
			exampleMP, mp)
	}
}

func TestSendMoveAndPositionError(t *testing.T) {
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
	err = ch.SendMoveAndPosition(exampleMP)
	if err == nil {
		t.Error("expected error but got <nil>")
	}
}

func TestSendMoveAndPositionSuccess(t *testing.T) {
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
	conn.WriteJSON(exampleMP)
	expectedJSON := conn.LastJSON
	conn.LastJSON = nil

	// Call ReadMoveAndPosition and test return values.
	err = ch.SendMoveAndPosition(exampleMP)
	if err != nil {
		t.Errorf("error in SendMoveAndPosition: %v", err)
	}
	if string(conn.LastJSON) != string(expectedJSON) {
		t.Errorf("wrong JSON sent: want %v but got %v",
			expectedJSON, conn.LastJSON)
	}
}
