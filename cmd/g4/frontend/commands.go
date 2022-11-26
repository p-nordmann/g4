package frontend

import (
	"context"
	"fmt"
	"g4"
	"g4/ws"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// TODO: maybe need a frontend struct to hold the channel and provide commands factories.

// connect(url, port) is a command for connecting to the peer.
func connect(url string, port int) tea.Cmd {
	return func() tea.Msg {
		channel := ws.New(ws.ChannelConfig{
			DialTimeout:  1 * time.Second,
			ServeTimeout: 60 * time.Second,
			Address:      fmt.Sprintf(":%d", port),
		})
		err := channel.ConnectWait(context.Background(), url) // TODO: shouldn't have to specify the protocol.
		if err != nil {
			panic(fmt.Errorf("error connecting through channel: %w", err))
		}
		return channel
	}
}

func chooseColor(ch g4.Channel) tea.Cmd {
	return func() tea.Msg {
		color, _ := ch.ChooseColor()
		return color
	}
}

func watchChannel(channel g4.Channel) tea.Cmd {
	return func() tea.Msg {
		// TODO: wait for channel to be closed and return.
		return nil
	}
}

func sendMove(ch g4.Channel, move g4.Move) tea.Cmd {
	return func() tea.Msg {
		ch.SendMove(move)
		return move
	}
}
func receiveMove(ch g4.Channel) tea.Cmd {
	return func() tea.Msg {
		move, _ := ch.ReadMove()
		return move
	}
}

func closeChannel() tea.Msg {
	// TODO: close channel
	return nil
}

func dumpGame() tea.Msg {
	// TODO: dump game listing somewhere.
	return nil
}
