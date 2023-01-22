package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type StatusBarModel struct {
	width     int
	connected bool
}

func (m StatusBarModel) Init() tea.Cmd {
	return nil
}

func (m StatusBarModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(Size); ok {
		m.width = msg.Width
	}

	switch msg.(type) {
	case ConnectionSuccessful:
		m.connected = true
	}

	return m, nil
}

func (m StatusBarModel) View() string {
	s := ""
	for k := 0; k < m.width/2-4; k++ {
		s += "#"
	}
	if m.connected {
		s += "   OK    "
	} else {
		s += " Waiting "
	}
	for k := 0; k < m.width/2-4; k++ {
		s += "#"
	}
	return s
}
