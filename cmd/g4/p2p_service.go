package main

import (
	"context"
	"errors"
	"fmt"
	"g4"
	"g4/p2p"
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	// maxColorTries denotes the limit to the number of attempts at choosing a color.
	maxColorTries = 100
)

// service provides factories for p2p commands.
//
// Such a command will itself always return either a success message or an error.
// The main model should treat error with care and act accordingly.
type P2PService struct {
	ch *p2p.Channel
	r  *rand.Rand
}

// p2pService is a global instance of P2PService.
//
// We assume that it is thread-safe.
var p2pService = &P2PService{
	r: rand.New(rand.NewSource(time.Now().UnixNano())),
}

// connect builds a command that opens the p2p connection with opponent.
func (s *P2PService) connect(ctx context.Context, descr string) (tea.Cmd, error) {
	if s.ch != nil {
		return nil, errors.New("channel already created")
	}
	ch, err := p2p.New(descr)
	if err != nil {
		return nil, fmt.Errorf("error creating channel: %w", err)
	}
	s.ch = ch
	return func() tea.Msg {
		err := ch.Open(ctx)
		if err != nil {
			return err
		}
		return ConnectionSuccessful{}
	}, nil
}

type ConnectionSuccessful struct{}

// chooseColor builds a command that tries to find a common color with the peer.
//
// It will repeatedly choose red or yellow at random, send it and wait for an answer from the peer.
// When both colors are different it will terminate.
func (s *P2PService) chooseColor() (tea.Cmd, error) {
	if s.ch == nil {
		return nil, errors.New("channel has not been created")
	}
	return func() tea.Msg {

		colors := [2]g4.Color{
			g4.Red,
			g4.Yellow,
		}

		for tries := 0; tries < maxColorTries; tries++ {
			color := colors[s.r.Intn(2)]
			err := s.ch.WriteJSON(color)
			if err != nil {
				return fmt.Errorf("error writing to peer: %w", err)
			}

			var peerColor g4.Color
			err = s.ch.ReadJSON(&peerColor)
			if err != nil {
				return fmt.Errorf("error reading from peer: %w", err)
			}

			if color != peerColor {
				return ColorFound(color)
			}
		}

		return errors.New("reach max number of tries")
	}, nil
}

type ColorFound g4.Color

// sendMove builds a command that sends a move to the peer.
func (s *P2PService) sendMove(move g4.Move) (tea.Cmd, error) {
	if s.ch == nil {
		return nil, errors.New("channel has not been created")
	}
	return func() tea.Msg {
		err := s.ch.WriteJSON(move)
		if err != nil {
			return err
		}
		return move
	}, nil
}

// receiveMove builds a command that receives a move from the peer.
func (s *P2PService) receiveMove() (tea.Cmd, error) {
	if s.ch == nil {
		return nil, errors.New("channel has not been created")
	}
	return func() tea.Msg {
		var move g4.Move
		err := s.ch.ReadJSON(&move)
		if err != nil {
			return err
		}
		return move
	}, nil
}
