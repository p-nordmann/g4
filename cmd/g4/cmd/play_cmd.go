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

package cmd

import (
	"fmt"
	"g4/cmd/g4/frontend"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var port int

func init() {
	playCmd.Flags().IntVar(&port, "port", 8080, "port to listen to")
}

// playCmd provides the play command: let's play a game of G4!
//
// TODO
var playCmd = &cobra.Command{
	Use:   "play address",
	Short: "Let's play a game of G4!",
	Long:  ``,
	Args:  cobra.ExactValidArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var url string = args[0]
		return playLoop(url, port)
	},
}

// func RunBackend(url string, port int, front *frontend.Frontend) {
// 	// Try to connect wait.
// 	channel := ws.New(ws.ChannelConfig{
// 		DialTimeout:  1 * time.Second,
// 		ServeTimeout: 10 * time.Second,
// 		Address:      fmt.Sprintf("localhost:%d", port),
// 	})
// 	err := channel.ConnectWait(context.Background(), url) // TODO: shouldn't have to specify the protocol.
// 	if err != nil {
// 		front.ShowError(fmt.Errorf("error connecting through channel: %w", err))
// 		return
// 	}

// 	// Play the game.
// 	front.StartPlaying()
// 	time.Sleep(5 * time.Second)

// 	// Game is over.
// 	front.GameOver()
// }

// TODO: rough draft/outline of main game loop
func playLoop(url string, port int) error {

	// Communication channel between backend and frontend.
	// comm := make(chan interface{})

	// Something like starting the frontend.

	p := tea.NewProgram(frontend.New(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
