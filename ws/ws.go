// Package ws provides a communication channel implementation using websockets.
package ws

import (
	"context"
	"fmt"
	"g4"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// TODO: address and port must be defined somewhere else.
const DefaultAddress = ":8080"

// TODO: documentation.

type ChannelConfig struct {
	DialTimeout  time.Duration
	ServeTimeout time.Duration
	Address      string
}

var DefautConfig = ChannelConfig{
	DialTimeout:  5 * time.Second,
	ServeTimeout: time.Minute,
	Address:      DefaultAddress,
}

type Channel struct {
	config    ChannelConfig
	connector Connector
	conn      Conn
}

func New(config ChannelConfig) *Channel {
	return &Channel{
		config: config,
		connector: GorillaDialer{
			websocket.DefaultDialer,
			&http.Server{Addr: config.Address},
		},
		conn: nil,
	}
}

func NewWithConnector(config ChannelConfig, connector Connector) *Channel {
	return &Channel{
		config:    config,
		connector: connector,
		conn:      nil,
	}
}

// TODO: address management: should not be necessary to input ws:// and port this way.
func (ch *Channel) ConnectWait(ctx context.Context, urlStr string) error {
	// Try to dial urlStr first.
	ctxDial, cancel := context.WithTimeout(ctx, ch.config.DialTimeout)
	defer cancel()

	conn, err := ch.connector.DialContext(ctxDial, urlStr)
	if err == nil {
		ch.conn = conn
		return nil
	}

	// Dial failed, try to serve for urlStr.
	ctxServe, cancel := context.WithTimeout(ctx, ch.config.ServeTimeout)
	defer cancel()

	conn, err = ch.connector.ServeContext(ctxServe, urlStr)
	if err == nil {
		ch.conn = conn
		return nil
	}

	// Error happened, return error.
	return fmt.Errorf("unable to connect to '%s'", urlStr)
}

// TODO: check that connection is correctly established.
func (ch *Channel) SendMove(move g4.Move) error {
	err := ch.conn.WriteJSON(move)
	if err != nil {
		return fmt.Errorf("error writing to connection: %w", err)
	}
	return nil
}

// TODO: check that connection is correctly established.
func (ch *Channel) ReadMove() (g4.Move, error) {
	move := g4.Move{}
	err := ch.conn.ReadJSON(&move)
	if err != nil {
		return move, fmt.Errorf("error reading from connection: %w", err)
	}
	return move, nil
}

func (ch *Channel) Close() error {
	if ch.conn == nil {
		return fmt.Errorf("connection not opened")
	}
	err := ch.conn.Close()
	if err != nil {
		return fmt.Errorf("error closing the connection: %w", err)
	}
	return nil
}
