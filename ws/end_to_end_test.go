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

// SendMove sends the selected move to the opponent.
func asyncSendMove(move g4.Move, ch g4.Channel) chan error {
	errCh := make(chan error)
	go func() {
		errCh <- ch.SendMove(move)
	}()
	return errCh
}

// ReadMove waits for the oponent to send a move and receives it.
func asyncReadMove(ch g4.Channel) (chan g4.Move, chan error) {
	moveCh := make(chan g4.Move)
	errCh := make(chan error)
	go func() {
		move, err := ch.ReadMove()
		if err != nil {
			errCh <- err
		}
		moveCh <- move
	}()
	return moveCh, errCh
}

func TestChannel(t *testing.T) {
	// Create two channels.
	ch1 := ws.New(ws.ChannelConfig{
		DialTimeout:  10 * time.Millisecond,
		ServeTimeout: 5 * time.Second,
		Address:      "localhost:8080",
	})
	ch2 := ws.New(ws.ChannelConfig{
		DialTimeout:  10 * time.Millisecond,
		ServeTimeout: 5 * time.Second,
		Address:      "localhost:8081",
	})

	// ConnectWait both channels.
	errCh1 := asyncConnectWait(context.Background(), "ws://localhost:8081", ch1) // TODO: address management
	time.Sleep(20 * time.Millisecond)
	errCh2 := asyncConnectWait(context.Background(), "ws://localhost:8080", ch2)

	err1 := <-errCh1
	err2 := <-errCh2

	if err1 != nil || err2 != nil {
		t.Errorf("errors in connectWait:\nch1: %v\nch2: %v", err1, err2)
	}

	// Prepare some data.
	m1 := g4.TokenMove(g4.Red, 8)
	m2 := g4.TiltMove(g4.LEFT)
	m3 := g4.TokenMove(g4.Yellow, 1)

	// Send a move from ch1.
	errCh1 = asyncSendMove(m1, ch1)
	moveCh2, errCh2 := asyncReadMove(ch2)

	err1 = <-errCh1
	if err1 != nil {
		t.Errorf("ch1.SendMove returned an error: %v", err1)
	}
	select {
	case err := <-errCh2:
		t.Errorf("ch2.ReadMove returned an error: %v", err)
	case move := <-moveCh2:
		if move != m1 {
			t.Errorf("ch2.ReadMove returned %v but expected %v", move, m1)
		}
	}

	// Send a move from ch2.
	moveCh1, errCh1 := asyncReadMove(ch1)
	errCh2 = asyncSendMove(m2, ch2)

	select {
	case err := <-errCh1:
		t.Errorf("ch1.ReadMove returned an error: %v", err)
	case move := <-moveCh1:
		if move != m2 {
			t.Errorf("ch1.ReadMove returned %v but expected %v", move, m2)
		}
	}
	err2 = <-errCh2
	if err2 != nil {
		t.Errorf("ch2.SendMove returned an error: %v", err2)
	}

	// Send a move from ch1 again.
	errCh1 = asyncSendMove(m3, ch1)
	moveCh2, errCh2 = asyncReadMove(ch2)

	err1 = <-errCh1
	if err1 != nil {
		t.Errorf("ch1.SendMove returned an error: %v", err1)
	}
	select {
	case err := <-errCh2:
		t.Errorf("ch2.ReadMove returned an error: %v", err)
	case mp := <-moveCh2:
		if mp != m3 {
			t.Errorf("ch2.ReadMove returned %v but expected %v", mp, m3)
		}
	}

	// Close channels.
	err1 = ch1.Close()
	err2 = ch2.Close()

	if err1 != nil || err2 != nil {
		t.Errorf("errors in Close:\nch1: %v\nch2: %v", err1, err2)
	}
}
