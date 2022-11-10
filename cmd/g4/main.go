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

package main

import (
	"flag"
	"fmt"
	"g4/cmd/g4/frontend"

	tea "github.com/charmbracelet/bubbletea"
)

var port int
var url string

func init() {
	flag.IntVar(&port, "port", 8080, "port to listen to")
	flag.Parse()
	url = flag.Arg(0)
}

// TODO: rough draft/outline of main game loop
func playLoop(url string, port int) error {

	// Communication channel between backend and frontend.
	// comm := make(chan interface{})

	// Something like starting the frontend.

	p := tea.NewProgram(frontend.New(url, port), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func main() {
	playLoop(url, port)
}
