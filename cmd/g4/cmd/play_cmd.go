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
	waiting "g4/frontend/pages/waiting"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

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

		p := tea.NewProgram(waiting.InitialWaitingModel(url), tea.WithAltScreen())
		if err := p.Start(); err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	},
}

// TODO: rough draft/outline of main game loop
// func playLoop() {

// 	// Communication channel between backend and frontend.
// 	comm := make(chan interface{})

// 	// Something like starting the frontend.
// 	go func() {
// 		front := startFrontend()
// 		for msg := range comm {
// 			front.SendMessage(msg)
// 		}
// 	}()

// 	// Something like this to start the backend.
// 	go func() {
// 		back := startBackend()

// 		// TODO outline
// 	}()
// }
