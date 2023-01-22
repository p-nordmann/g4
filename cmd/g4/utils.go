package main

import tea "github.com/charmbracelet/bubbletea"

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

// isQuitMessage returns true if the message signals to quit the program.
func isQuitMessage(msg tea.Msg) bool {
	if key, ok := msg.(tea.KeyMsg); ok {
		switch key.String() {
		case "ctrl+c", "q":
			return true
		default:
			return false
		}
	}
	return false
}
