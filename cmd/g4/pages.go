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

import tea "github.com/charmbracelet/bubbletea"

type Page int

const (
	Home Page = iota
	Waiting
	Game
	Exit
)

// TODO: actual application model.

func (p Page) Next() Page {
	switch p {
	case Home:
		return Waiting
	case Waiting:
		return Game
	case Game:
		return Exit
	default:
		return Exit
	}
}

func (p Page) Init() tea.Cmd {
	return nil
}

func (p Page) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if p == Exit {
		return Exit, tea.Quit
	}

	// Process message.
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return Exit, tea.Quit
		case "tab":
			return p.Next(), nil
		}
	}

	// Default don't update.
	return p, nil
}

func (p Page) View() string {
	switch p {
	case Home:
		return "Welcome to G4!"
	case Waiting:
		return "Waiting for a game..."
	case Game:
		return "Now we're playing!"
	default:
		return "Bye!"
	}
}
