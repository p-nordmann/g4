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
	"g4"
	"g4/ws"
	"testing"
	"time"
)

func asyncConnectWait(ctx context.Context, URL string, ch g4.Channel) chan error {
	errCh := make(chan error)
	go func() {
		errCh <- ch.ConnectWait(ctx, URL)
	}()
	return errCh
}

// SendMoveAndPosition sends the selected move to the opponent.
func asyncSendMoveAndPosition(mp g4.MoveAndPosition, ch g4.Channel) chan error {
	errCh := make(chan error)
	go func() {
		errCh <- ch.SendMoveAndPosition(mp)
	}()
	return errCh
}

// ReadMoveAndPosition waits for the oponent to send a move and receives it.
func asyncReadMoveAndPosition(ch g4.Channel) (chan g4.MoveAndPosition, chan error) {
	mpCh := make(chan g4.MoveAndPosition)
	errCh := make(chan error)
	go func() {
		mp, err := ch.ReadMoveAndPosition()
		if err != nil {
			errCh <- err
		}
		mpCh <- mp
	}()
	return mpCh, errCh
}

func TestChannel(t *testing.T) {
	// Create two channels.
	ch1 := ws.New(ws.ChannelConfig{
		DialTimeout:  10 * time.Millisecond,
		ServeTimeout: 5 * time.Second,
	})
	ch2 := ws.New(ws.ChannelConfig{
		DialTimeout:  10 * time.Millisecond,
		ServeTimeout: 5 * time.Second,
	})

	// ConnectWait both channels.
	errCh1 := asyncConnectWait(context.Background(), "ws://localhost:8080", ch1) // TODO: address management
	time.Sleep(20 * time.Millisecond)
	errCh2 := asyncConnectWait(context.Background(), "ws://localhost:8080", ch2)

	err1 := <-errCh1
	err2 := <-errCh2

	if err1 != nil || err2 != nil {
		t.Errorf("errors in connectWait:\nch1: %v\nch2: %v", err1, err2)
	}

	// Prepare some data.
	mp1 := g4.MoveAndPosition{
		Move: g4.TokenMove(g4.Red, 8),
		Position: g4.Position{
			BoardStr:      "ry6|8|8|8|8|8|8|rryy4",
			Direction:     g4.UP,
			ColorWithMove: g4.Red,
		},
	}
	mp2 := g4.MoveAndPosition{
		Move: g4.TiltMove(g4.LEFT),
		Position: g4.Position{
			BoardStr:      "8|8|8|8|8|8|8|8",
			Direction:     g4.DOWN,
			ColorWithMove: g4.Red,
		},
	}
	mp3 := g4.MoveAndPosition{
		Move: g4.TokenMove(g4.Yellow, 1),
		Position: g4.Position{
			BoardStr:      "yy6|8|8|8|8|8|8|yyyy4",
			Direction:     g4.LEFT,
			ColorWithMove: g4.Red,
		},
	}

	// Send a move from ch1.
	errCh1 = asyncSendMoveAndPosition(mp1, ch1)
	mpCh2, errCh2 := asyncReadMoveAndPosition(ch2)

	err1 = <-errCh1
	if err1 != nil {
		t.Errorf("ch1.SendMoveAndPosition returned an error: %v", err1)
	}
	select {
	case err := <-errCh2:
		t.Errorf("ch2.ReadMoveAndPosition returned an error: %v", err)
	case mp := <-mpCh2:
		if mp != mp1 {
			t.Errorf("ch2.ReadMoveAndPosition returned %v but expected %v", mp, mp1)
		}
	}

	// Send a move from ch2.
	mpCh1, errCh1 := asyncReadMoveAndPosition(ch1)
	errCh2 = asyncSendMoveAndPosition(mp2, ch2)

	select {
	case err := <-errCh1:
		t.Errorf("ch1.ReadMoveAndPosition returned an error: %v", err)
	case mp := <-mpCh1:
		if mp != mp2 {
			t.Errorf("ch1.ReadMoveAndPosition returned %v but expected %v", mp, mp2)
		}
	}
	err2 = <-errCh2
	if err2 != nil {
		t.Errorf("ch2.SendMoveAndPosition returned an error: %v", err2)
	}

	// Send a move from ch1 again.
	errCh1 = asyncSendMoveAndPosition(mp3, ch1)
	mpCh2, errCh2 = asyncReadMoveAndPosition(ch2)

	err1 = <-errCh1
	if err1 != nil {
		t.Errorf("ch1.SendMoveAndPosition returned an error: %v", err1)
	}
	select {
	case err := <-errCh2:
		t.Errorf("ch2.ReadMoveAndPosition returned an error: %v", err)
	case mp := <-mpCh2:
		if mp != mp3 {
			t.Errorf("ch2.ReadMoveAndPosition returned %v but expected %v", mp, mp3)
		}
	}

	// Close channels.
	err1 = ch1.Close()
	err2 = ch2.Close()

	if err1 != nil || err2 != nil {
		t.Errorf("errors in Close:\nch1: %v\nch2: %v", err1, err2)
	}
}
