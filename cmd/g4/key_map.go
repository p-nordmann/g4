package main

import (
	"g4"

	"github.com/charmbracelet/lipgloss"
)

func viewKeymap(app AppModel) string {
	hStyle := lipgloss.NewStyle().Bold(true).Foreground(light)
	pStyle := lipgloss.NewStyle().PaddingLeft(1).MarginBottom(1).Foreground(lighter)
	return lipgloss.JoinVertical(
		lipgloss.Left,
		hStyle.Render("Token moves"),
		pStyle.Render(":1 :2 :3 :4 :5 :6 :7 :8"),
		hStyle.Render("Tilt moves"),
		pStyle.Render(":left :down :right"),
		hStyle.Render("Quit"),
		pStyle.Render(":q or ctrl+c"),
	)
}

type KeyHandler struct {
	lastKey string
	keyMap  map[string]string
}

var defaultKeymap = map[string]string{
	": q":     "quit",
	"ctrl+c":  "quit",
	": 1":     ":1",
	": 2":     ":2",
	": 3":     ":3",
	": 4":     ":4",
	": 5":     ":5",
	": 6":     ":6",
	": 7":     ":7",
	": 8":     ":8",
	": left":  ":left",
	": down":  ":down",
	": right": ":right",
}

func (h *KeyHandler) handle(key string) string {
	// Two-key combos.
	combo, ok := h.keyMap[h.lastKey+" "+key]
	if ok {
		h.lastKey = ""
		return combo
	}

	// One-key combos.
	combo, ok = h.keyMap[key]
	if ok {
		h.lastKey = ""
		return combo
	}

	// No combo.
	h.lastKey = key
	return key
}

func makeMove(combo string, color g4.Color) g4.Move {
	switch combo {
	case ":1":
		return g4.TokenMove(color, 0)
	case ":2":
		return g4.TokenMove(color, 1)
	case ":3":
		return g4.TokenMove(color, 2)
	case ":4":
		return g4.TokenMove(color, 3)
	case ":5":
		return g4.TokenMove(color, 4)
	case ":6":
		return g4.TokenMove(color, 5)
	case ":7":
		return g4.TokenMove(color, 6)
	case ":8":
		return g4.TokenMove(color, 7)
	case ":left":
		return g4.TiltMove(color, g4.LEFT)
	case ":down":
		return g4.TiltMove(color, g4.DOWN)
	case ":right":
		return g4.TiltMove(color, g4.RIGHT)
	default:
		return g4.Move{}
	}
}
