package frontend

import (
	"fmt"
	"g4"
	"g4/bits"
	"g4/simulation"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	sidePannelStyle = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("grey"))
	playAreaStyle = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("grey"))
	// listingStyle = lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color("#ffaa99"))
	// spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
)

type mainModel struct {
	// Display information.
	width, height int

	// Submodels.
	listing  listingArea
	board    playArea
	selector selectorArea

	// App-level information.
	url           string
	port          int
	colorWithMove g4.Color
	playerColor   g4.Color

	// Communication.
	ch   g4.Channel
	game g4.Game
}

func New(url string, port int) mainModel {
	board, _ := bits.FromString("8|8|8|8|8|8|8|8")
	game, _ := simulation.FromBoard(board, g4.Yellow, g4.UP)
	moves, _ := game.Generate()
	return mainModel{
		url:  url,
		port: port,
		selector: selectorArea{
			SelectedMove:  -1,
			Disabled:      true,
			PossibleMoves: moves,
		},
		listing: listingArea{},
		board: playArea{
			Direction: g4.UP,
		},
		game:          game,
		colorWithMove: g4.Yellow,
	}
}

func (m mainModel) Init() tea.Cmd {
	return connect(m.url, m.port)
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

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// Handle quit message.
	if isQuitMessage(msg) {
		return m, tea.Quit
	}

	// Handle window size.
	// TODO: on windows, use a ticker combined with an alternate way to fetch width.
	if size, ok := msg.(tea.WindowSizeMsg); ok {
		m.height = size.Height
		m.width = size.Width
		return m, nil
	}

	switch msg := msg.(type) {

	// On successful connection.
	case g4.Channel:
		m.ch = msg
		return m, chooseColor(m.ch)

	// When color is chosen (at the start).
	case g4.Color:
		m.playerColor = msg
		if msg == m.colorWithMove {
			m.selector.Disabled = false
		} else {
			return m, receiveMove(m.ch)
		}

	// When a move is played.
	case g4.Move:
		move := msg
		game, err := m.game.Apply(move)
		if err != nil {
			m.ch.Close()
			fmt.Println(err)
			return m, tea.Quit
		}
		m.game = game
		m.board.Board, m.board.Direction = game.ToArray()
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
			m.ch.Close()
			fmt.Println(err)
			return m, tea.Quit
		}

	default:
		// Handle move selection.
		// TODO: should not delegate channel to selector.
		// selector should be service-agnostic.
		selector, cmd := m.selector.Update(m.ch, msg)
		m.selector = selector
		return m, cmd
	}
	return m, nil
}

func (m mainModel) View() string {
	var s string
	s += lipgloss.JoinHorizontal(
		lipgloss.Top,
		sidePannelStyle.Render(m.listing.View()),
		playAreaStyle.Render(m.board.View()),
		lipgloss.JoinVertical(
			lipgloss.Center,
			m.board.Direction.String(),
			sidePannelStyle.Render(
				m.selector.View(),
			),
		),
	)
	return s
}
