package p2p_test

import (
	"context"
	"g4/p2p"
	"testing"
	"time"
)

func ConnectAsync(ctx context.Context, ch *p2p.Channel) chan error {
	errCh := make(chan error)
	go func() {
		errCh <- ch.Open(ctx)
	}()
	return errCh
}

func TestChannelError(t *testing.T) {
	t.Run("reading from closed channel should fail", func(t *testing.T) {
		ch, _ := p2p.New("1234:localhost:5678")
		err := ch.ReadJSON(&struct{}{})
		if err == nil {
			t.Error("expected error but got <nil>")
		}
	})

	t.Run("writing to closed channel should fail", func(t *testing.T) {
		ch, _ := p2p.New("1234:localhost:5678")
		err := ch.WriteJSON(struct{}{})
		if err == nil {
			t.Error("expected error but got <nil>")
		}
	})

	t.Run("closing closed channel should fail", func(t *testing.T) {
		ch, _ := p2p.New("1234:localhost:5678")
		err := ch.Close()
		if err == nil {
			t.Error("expected error but got <nil>")
		}
	})

	t.Run("opening opened channel should fail", func(t *testing.T) {
		// Create two channels.
		ch1, _ := p2p.New("2345:localhost:5432")
		errCh1 := ConnectAsync(context.Background(), ch1)

		ch2, _ := p2p.New("5432:localhost:2345")
		time.Sleep(500 * time.Millisecond)
		errCh2 := ConnectAsync(context.Background(), ch2)

		// Wait for termination.
		err1 := <-errCh1
		err2 := <-errCh2
		if err1 != nil || err2 != nil {
			t.Errorf("errors in opening channels: %v, %v", err1, err2)
		}

		// Try to connect again.
		errCh1 = ConnectAsync(context.Background(), ch1)
		errCh2 = ConnectAsync(context.Background(), ch2)
		err1 = <-errCh1
		err2 = <-errCh2
		if err1 == nil || err2 == nil {
			t.Errorf("expected both errors but got %v, %v", err1, err2)
		}

		// Close channels.
		err1 = ch1.Close()
		err2 = ch2.Close()
		if err1 != nil || err2 != nil {
			t.Errorf("errors in Close:\nch1: %v\nch2: %v", err1, err2)
		}
	})
}

func TestChannel(t *testing.T) {

	t.Run("should Open and Close", func(t *testing.T) {
		// Create two channels.
		ch1, _ := p2p.New("1234:localhost:5678")
		errCh1 := ConnectAsync(context.Background(), ch1)

		ch2, _ := p2p.New("5678:localhost:1234")
		errCh2 := ConnectAsync(context.Background(), ch2)

		// Wait for termination.
		err1 := <-errCh1
		err2 := <-errCh2
		if err1 != nil || err2 != nil {
			t.Errorf("errors in opening channels: %v, %v", err1, err2)
		}

		// Close channels.
		err1 = ch1.Close()
		err2 = ch2.Close()
		if err1 != nil || err2 != nil {
			t.Errorf("errors in Close:\nch1: %v\nch2: %v", err1, err2)
		}
	})

	t.Run("should Read and Write", func(t *testing.T) {
		// Create two channels.
		ch1, _ := p2p.New("1111:localhost:2222")
		errCh1 := ConnectAsync(context.Background(), ch1)

		ch2, _ := p2p.New("2222:localhost:1111")
		errCh2 := ConnectAsync(context.Background(), ch2)

		// Wait for termination.
		err1 := <-errCh1
		err2 := <-errCh2
		if err1 != nil || err2 != nil {
			t.Errorf("errors in opening channels: %v, %v", err1, err2)
		}

		// Write to channel 1 and read from 2.
		type Car struct {
			Constructor string
			Year        int
		}
		car1 := Car{
			Constructor: "Renault",
			Year:        2022,
		}
		ch1.WriteJSON(car1)
		var car2 Car
		ch2.ReadJSON(&car2)

		if car1 != car2 {
			t.Errorf("expected %v but got %v", car1, car2)
		}

		// Close channels.
		err1 = ch1.Close()
		err2 = ch2.Close()
		if err1 != nil || err2 != nil {
			t.Errorf("errors in Close:\nch1: %v\nch2: %v", err1, err2)
		}
	})
}
