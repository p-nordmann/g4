package main

import (
	"context"
	"errors"
	"g4"
	"g4/bitsim"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// TODO: each sub-models should manage their own state, ex: each could have their own simulator,
//		or they should be provided with services (dependency injection?).

var (
	sidePannelStyle = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("grey"))
	playAreaStyle = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("grey"))
)

type mainModel struct {
	// Display information.
	width, height int

	// Submodels.
	listing  listingArea
	board    playArea
	selector selectorArea
	preview  previewArea

	// App-level information.
	descr         string
	colorWithMove g4.Color
	playerColor   g4.Color

	// Communication.
	game bitsim.Game
}

func NewFrontend(descr string) mainModel {
	board, _ := bitsim.FromString("8|8|8|8|8|8|8|8")
	game, _ := bitsim.FromBoard(board, g4.Yellow)
	moves, _ := game.Generate()
	return mainModel{
		descr: descr,
		selector: selectorArea{
			SelectedMove:  -1,
			Disabled:      true,
			PossibleMoves: moves,
		},
		listing:       listingArea{},
		board:         playArea{},
		preview:       previewArea{},
		game:          game,
		colorWithMove: g4.Yellow,
	}
}

func (m mainModel) Init() tea.Cmd {
	cmd, err := p2pService.connect(context.Background(), m.descr)
	if err != nil {
		return handleError(err)
	}
	return cmd
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

// NB: this is safe to pass m by reference here, because the outter Update function receives it
// by value.
// On the contrary, passing by value wouldn't work here because we cannot assign m in the outter
// Update function.
func (m *mainModel) updateMove(move g4.Move) error {
	game, err := m.game.Apply(move)
	if err != nil {
		return err
	}
	m.game = game
	m.board.Board = game.ToArray()
	if m.colorWithMove == g4.Yellow {
		m.colorWithMove = g4.Red
	} else {
		m.colorWithMove = g4.Yellow
	}
	m.listing.History = append(m.listing.History, move)
	m.listing.Waiting = !m.listing.Waiting
	m.selector.Disabled = !m.selector.Disabled
	m.selector.SelectedMove = -1
	m.selector.PossibleMoves, err = game.Generate()
	if err != nil {
		return err
	}
	m.preview.Board = m.game.ToArray()
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// Handle quit message.
	if isQuitMessage(msg) {
		return errorModel{
			err:      errors.New("SIGTERM"),
			position: m.game.String(),
		}, tea.Quit
	}

	// Handle error message.
	if err, ok := msg.(error); ok && err != nil {
		return errorModel{
			err:      err,
			position: m.game.String(),
		}, nil
	}

	// Handle window size.
	// TODO: on windows, use a ticker combined with an alternate way to fetch width.
	if size, ok := msg.(tea.WindowSizeMsg); ok {
		m.height = size.Height
		m.width = size.Width
		return m, nil
	}

	switch msg := msg.(type) {

	case ConnectionSuccessful:
		cmd, err := p2pService.chooseColor(context.Background())
		if err != nil {
			return m, handleError(err)
		}
		return m, cmd

	case ColorFound:
		color := g4.Color(msg)
		m.playerColor = color
		if color == m.colorWithMove {
			m.selector.Disabled = false
		} else {
			cmd, err := p2pService.receiveMove(context.Background())
			if err != nil {
				return m, handleError(err)
			}
			return m, cmd
		}

	case SentMove:
		err := m.updateMove(g4.Move(msg))
		if err != nil {
			return m, handleError(err)
		}
		cmd, err := p2pService.receiveMove(context.Background())
		if err != nil {
			return m, handleError(err)
		}
		return m, cmd

	case ReceivedMove:
		err := m.updateMove(g4.Move(msg))
		if err != nil {
			return m, handleError(err)
		}

	default:
		// Handle move selection.
		// TODO: selector should be service-agnostic?
		selector, cmd := m.selector.Update(msg)
		m.selector = selector

		// TODO: find a nice way to do this: here for instance if we pass a random message to selector,
		//	it will trigger the following.
		if !m.selector.Disabled && m.selector.SelectedMove >= 0 {
			game, _ := m.game.Apply(m.selector.PossibleMoves[m.selector.SelectedMove])
			m.preview.Board = game.ToArray()
		}
		return m, cmd
	}
	return m, nil
}

func (m mainModel) View() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		sidePannelStyle.Render(m.listing.View()),
		playAreaStyle.Render(m.board.View()),
		lipgloss.JoinVertical(
			lipgloss.Center,
			sidePannelStyle.Render(
				m.selector.View(),
			),
			sidePannelStyle.Render(
				m.preview.View(),
			),
		),
	)
}
