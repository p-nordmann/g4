package main

import (
	"context"
	"g4"
	"g4/bitsim"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	yellow  = lipgloss.Color("#999900")
	red     = lipgloss.Color("#dd0000")
	pink    = lipgloss.Color("#7e1e5e")
	dark    = lipgloss.Color("#0a0a0a")
	light   = lipgloss.Color("#b0b0b0")
	lighter = lipgloss.Color("#d0d0d0")
)

type AppModel struct {
	width, height int
	keyHandler    KeyHandler

	spec      string
	connected bool

	game    bitsim.Game
	myColor g4.Color
	waiting bool

	debug string
}

func (app AppModel) Init() tea.Cmd {
	cmd, err := p2pService.connect(context.Background(), app.spec)
	if err != nil {
		return handleError(err)
	}
	return cmd
}

func handleError(err error) tea.Cmd {
	if p2pService.ch != nil {
		p2pService.ch.Close()
	}
	return func() tea.Msg { return err }
}

func contains(target g4.Move, moves []g4.Move) bool {
	for _, move := range moves {
		if target == move {
			return true
		}
	}
	return false
}

func (app AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case error:
		app.debug = msg.Error()
		return app, nil

	case tea.WindowSizeMsg:
		app.height = msg.Height
		app.width = msg.Width
		return app, nil

	case ConnectionSuccessful:
		app.connected = true
		cmd, err := p2pService.chooseColor()
		if err != nil {
			return app, handleError(err)
		}
		return app, cmd

	case ColorFound:
		app.myColor = g4.Color(msg)

	case g4.Move:
		game, err := app.game.Apply(g4.Move(msg))
		if err != nil {
			app.debug = "err:" + err.Error()
			return app, handleError(err)
		}
		app.game = game
		app.waiting = false

	case tea.KeyMsg:
		combo := app.keyHandler.handle(msg.String())
		app.debug = combo
		switch combo {

		case "quit":
			return app, tea.Quit

		case ":1", ":2", ":3", ":4", ":5", ":6", ":7", ":8", ":left", ":down", ":right":
			// We generate the move with myColor and test it against legal moves.
			move := makeMove(combo, app.myColor)
			legalMoves, _ := app.game.Generate()
			if !contains(move, legalMoves) {
				return app, nil
			}

			// If move is legal it means it is our turn.
			cmd, err := p2pService.sendMove(move)
			if err != nil {
				return app, handleError(err)
			}
			return app, cmd

		}
	}

	// Make sure to try receiving moves.
	if app.connected && app.myColor != g4.Empty && app.myColor != app.game.Mover && !app.waiting {
		cmd, err := p2pService.receiveMove()
		if err != nil {
			return app, handleError(err)
		}
		app.waiting = true
		return app, cmd
	}

	return app, nil
}

func (app AppModel) View() string {
	if app.width == 0 || app.height == 0 {
		return ""
	}

	status := []string{}
	if app.connected {
		status = append(status, "Connected")
	} else {
		status = append(status, "Waiting for peer...")
	}
	if app.myColor == g4.Red {
		status = append(status, "you play red")
	} else if app.myColor == g4.Yellow {
		status = append(status, "you play yellow")
	}
	if app.myColor == app.game.Mover {
		status = append(status, "your turn")
	} else if app.waiting {
		status = append(status, "opponent turn")
	}

	hStyle := lipgloss.NewStyle().Bold(true).Foreground(light)
	pStyle := lipgloss.NewStyle().PaddingLeft(1).MarginBottom(1).Foreground(lighter)
	controls := lipgloss.JoinVertical(
		lipgloss.Left,
		hStyle.Render("Token moves"),
		pStyle.Render(":1 :2 :3 :4 :5 :6 :7 :8"),
		hStyle.Render("Tilt moves"),
		pStyle.Render(":left :down :right"),
		hStyle.Render("Quit"),
		pStyle.Render(":q or ctrl+c"),
	)

	rightPanel := lipgloss.NewStyle().Padding(1).Render(controls)
	if app.debug != "" {
		rightPanel = lipgloss.JoinVertical(
			lipgloss.Left,
			rightPanel,
			lipgloss.NewStyle().Background(lipgloss.Color(pink)).Render("debug: "+app.debug),
		)
	}
	rightPanelWidth := lipgloss.Width(rightPanel)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.
			NewStyle().
			Height(app.height-1).
			AlignVertical(lipgloss.Center).
			Render(
				lipgloss.JoinHorizontal(
					lipgloss.Center,
					lipgloss.
						NewStyle().
						Width(app.width-rightPanelWidth).
						Align(lipgloss.Center).
						Render(
							drawBoard(
								app.game.Board,
								fitBoard(app.width-rightPanelWidth, app.height-1),
							),
						),
					rightPanel,
				),
			),
		drawStatusBar(strings.Join(status, " | "), app.width),
	)
}

func drawStatusBar(msg string, width int) string {
	return lipgloss.
		NewStyle().
		Width(width).
		PaddingLeft(1).
		PaddingRight(1).
		Background(light).
		Foreground(dark).
		Render(
			clipStr(msg, width-2),
		)
}

// TODO make a proper overflow with lipgloss styles
func clipStr(s string, width int) string {
	for lipgloss.Width(s) > width {
		s = s[:len(s)-1]
	}
	for lipgloss.Width(s) < width {
		s = s + " "
	}
	return s
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
		return g4.Base()
	}
}
