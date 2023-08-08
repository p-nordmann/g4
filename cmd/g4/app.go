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
	yellow  = lipgloss.Color("#dbdb00")
	red     = lipgloss.Color("#dd0000")
	pink    = lipgloss.Color("#7e1e5e")
	pinker  = lipgloss.Color("#F25D94")
	light   = lipgloss.Color("#b0b0b0")
	lighter = lipgloss.Color("#e0e0e0")
	dark    = lipgloss.Color("#0a0a0a")
)

// TODO make it simple
type ConnectionStatus int
type GameStatus int

const (
	connecting ConnectionStatus = iota
	connected
	closed
	inProgress GameStatus = iota
	draw
	yellowWins
	redWins
	suspended
)

type AppModel struct {
	width, height int
	keyHandler    KeyHandler

	spec string

	connStatus ConnectionStatus
	listening  bool
	gameStatus GameStatus

	game    bitsim.Game
	myColor g4.Color

	modalContent string
	modalHover   bool

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
	if err == nil {
		return nil
	}
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

// TODO do not allow to make moves if game not in progress
func (app AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// Modal takes precedence.
	if app.modalContent != "" {
		return handleModalEvents(app, msg)
	}

	switch msg := msg.(type) {

	case error:
		app.debug = msg.Error()
		if p2pService.ch != nil {
			p2pService.ch.Close()
		}
		app.connStatus = closed
		app.gameStatus = suspended
		app.modalContent = "Error occured:\n" + msg.Error()
		return app, nil

	case tea.WindowSizeMsg:
		app.height = msg.Height
		app.width = msg.Width
		return app, nil

	case ConnectionSuccessful:
		cmd, err := p2pService.chooseColor()
		if err != nil {
			return app, handleError(err)
		}
		return app, cmd

	case ColorFound:
		app.myColor = g4.Color(msg)
		app.connStatus = connected
		if app.myColor == g4.Red {
			app.modalContent = "Game on!\nYou play the red pieces."
		}
		if app.myColor == g4.Yellow {
			app.modalContent = "Game on!\nYou play the yellow pieces."
		}

	case g4.Move:
		game, err := app.game.Apply(g4.Move(msg))
		if err != nil {
			app.debug = "err:" + err.Error()
			return app, handleError(err)
		}
		app.game = game
		app.listening = false

	case tea.KeyMsg:
		combo := app.keyHandler.handle(msg.String())
		app.debug = combo
		switch combo {

		case "quit":
			if p2pService.ch != nil {
				p2pService.ch.Close()
			}
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
	if app.connStatus == connected &&
		app.gameStatus == inProgress &&
		app.myColor != app.game.Mover &&
		!app.listening {
		cmd, err := p2pService.receiveMove()
		if err != nil {
			return app, handleError(err)
		}
		app.listening = true
		return app, cmd
	}

	return app, nil
}

func (app AppModel) View() string {
	if app.width == 0 || app.height == 0 {
		return ""
	}

	var mainSection string

	if app.modalContent != "" {
		mainSection = viewModal(app.modalContent, app.modalHover)
	} else {
		rightPanel := lipgloss.NewStyle().Padding(1).Render(viewKeymap(app)) // TODO responsive right panel
		rightPanelWidth := lipgloss.Width(rightPanel)
		mainSection = lipgloss.JoinHorizontal(
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
		)
	}

	return lipgloss.JoinVertical(
		lipgloss.Center,
		lipgloss.
			NewStyle().
			Height(app.height-1).
			AlignVertical(lipgloss.Center).
			Render(mainSection),
		viewStatusBar(app),
	)
}

func viewStatusBar(app AppModel) string {

	style := lipgloss.NewStyle().Width(app.width).
		PaddingLeft(1).PaddingRight(1).
		Background(light).Foreground(dark)

	spans := []string{}

	if app.debug != "" {
		spans = append(spans, "debug: "+app.debug)
	}

	switch app.connStatus {
	case connecting:
		spans = append(spans, "Connecting...")
	case connected:
		spans = append(spans, "Connected")
	case closed:
		spans = append(spans, "Disconnected")
	}

	switch app.gameStatus {
	case inProgress:
		if app.connStatus != connected {
			break
		}
		if app.myColor == g4.Yellow {
			spans = append(spans, "Game on, you play yellow")
		} else {
			spans = append(spans, "Game on, you play red")
		}
		if app.myColor == app.game.Mover {
			spans = append(spans, "Your move")
		} else {
			spans = append(spans, "Opponent's move")
		}
	case draw:
		spans = append(spans, "Game Over > Draw")
	case yellowWins:
		spans = append(spans, "Game Over > Yellow wins")
	case redWins:
		spans = append(spans, "Game Over > Red wins")
	case suspended:
		spans = append(spans, "Game Over > Suspended")
	}

	return style.Render(clipStr(strings.Join(spans, " | "), app.width-2))
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
