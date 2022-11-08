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
func (ch *Channel) SendMoveAndPosition(mp g4.MoveAndPosition) error {
	err := ch.conn.WriteJSON(mp)
	if err != nil {
		return fmt.Errorf("error writing to connection: %w", err)
	}
	return nil
}

// TODO: check that connection is correctly established.
func (ch *Channel) ReadMoveAndPosition() (g4.MoveAndPosition, error) {
	mp := g4.MoveAndPosition{}
	err := ch.conn.ReadJSON(&mp)
	if err != nil {
		return mp, fmt.Errorf("error reading from connection: %w", err)
	}
	return mp, nil
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
