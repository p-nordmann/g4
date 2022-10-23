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

package pages

import (
	"os"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"

	"fmt"
)

type WaitingModel struct {
	spinner spinner.Model
	url     string
}

// var (
// 	centeredParagraph = lipgloss.NewStyle().
// 		Align(lipgloss.Center).Margin(0, 1).
// 		BorderStyle(lipgloss.NormalBorder())
// )

func InitialWaitingModel(url string) WaitingModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return WaitingModel{spinner: s, url: url}
}

func (m WaitingModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m WaitingModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		default:
			return m, nil
		}

	case error:
		fmt.Println("an error occurred")
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

var docStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Align(lipgloss.Center)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (m WaitingModel) View() string {

	width, height, _ := term.GetSize(int(os.Stdout.Fd()))

	if width > 0 {
		docStyle = docStyle.Width(min(width, 100))
	}
	if height > 0 {
		docStyle = docStyle.Height(min(height, 20))
	}

	return docStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			lipgloss.JoinHorizontal(
				lipgloss.Center,
				m.spinner.View(),
				fmt.Sprintf("Waiting for <%s>", m.url),
			),
			"press q to quit",
		),
	)
}
