package p2p

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"regexp"
	"strconv"
	"sync"
	"time"
)

var re *regexp.Regexp = regexp.MustCompile(`^(\d+)\:(.*)$`)

// Channel is a peer-to-peer channel.
//
// One channel can only communicate with one pre-defined peer on a pre-defined port.
type Channel struct {
	// High level data.
	localPort int
	url       string

	// Connection and JSON utilities.
	in  net.Conn
	dec *json.Decoder
	out net.Conn
	enc *json.Encoder
}

// Returns a new *Channel from description.
//
// Expected format is:
//
//	port:url:port
func New(descr string) (*Channel, error) {
	matches := re.FindSubmatch([]byte(descr))
	if len(matches) != 3 {
		return nil, errors.New("invalid channel description string")
	}
	localPort, _ := strconv.Atoi(string(matches[1]))
	url := string(matches[2])
	return &Channel{
		localPort: localPort,
		url:       url,
	}, nil
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

// Tries to connect to `addr` with exponential backoff.
//
// TODO: use context.
func connectExponentialBackoff(addr string, minBackoff, maxBackoff int, maxTries int) (net.Conn, error) {
	backoff := minBackoff
	for tries := 0; tries < maxTries; tries++ {
		conn, err := net.DialTimeout("tcp", addr, time.Duration(minBackoff)*time.Millisecond)
		if err == nil {
			return conn, nil
		}
		time.Sleep(time.Duration(backoff) * time.Millisecond)
		backoff = min(2*backoff, maxBackoff)
	}
	return nil, errors.New("reached maxTries")
}

func (ch *Channel) Open(ctx context.Context) error {

	var errIn, errOut error
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		out, err := connectExponentialBackoff(ch.url, 500, 10000, 100)
		if err != nil {
			errOut = fmt.Errorf("error while dialing: %w", err)
			return
		}
		ch.out = out
	}()

	// TODO: make sure we only accept connection from peer with the right url.
	go func() {
		defer wg.Done()
		listener, err := net.Listen("tcp", ":"+strconv.Itoa(ch.localPort))
		if err != nil {
			errIn = fmt.Errorf("error while listening: %w", err)
			return
		}
		in, err := listener.Accept()
		if err != nil {
			errIn = fmt.Errorf("error while accepting: %w", err)
			return
		}
		ch.in = in
	}()

	wg.Wait()
	if errIn != nil {
		ch.Close()
		return fmt.Errorf("error making 'in' connection: %w", errIn)
	}
	if errOut != nil {
		ch.Close()
		return fmt.Errorf("error making 'out' connection: %w", errOut)
	}

	ch.enc = json.NewEncoder(ch.out)
	ch.dec = json.NewDecoder(ch.in)

	return nil
}

func (ch *Channel) Close() error {
	if ch.in == nil && ch.out == nil {
		return errors.New("channel not open")
	}
	ch.enc = nil
	ch.dec = nil
	if ch.in != nil {
		ch.in.Close()
	}
	if ch.out != nil {
		ch.out.Close()
	}
	return nil
}

func (ch *Channel) ReadJSON(v interface{}) error {
	if ch.in == nil || ch.dec == nil {
		return errors.New("channel not properly opened")
	}
	if err := ch.dec.Decode(v); err == io.EOF {
		return errors.New("channel seems to be closed") // TODO: detect closing with a proper API.
	} else if err != nil {
		return fmt.Errorf("error reading from peer: %w", err)
	}
	return nil
}

func (ch *Channel) WriteJSON(v interface{}) error {
	if ch.out == nil || ch.enc == nil {
		return errors.New("channel not properly opened")
	}
	if err := ch.enc.Encode(v); err != nil {
		return fmt.Errorf("error sending to peer: %w", err)
	}
	return nil
}
