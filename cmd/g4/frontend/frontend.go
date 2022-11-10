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
	return mainModel{
		url:  url,
		port: port,
		selector: selectorArea{
			SelectedMove: 3,
			Disabled:     false,
			PossibleMoves: []g4.Move{
				g4.TokenMove(g4.Red, 1),
				g4.TokenMove(g4.Red, 2),
				g4.TokenMove(g4.Red, 3),
				g4.TokenMove(g4.Red, 4),
				g4.TokenMove(g4.Red, 5),
				g4.TokenMove(g4.Red, 6),
				g4.TokenMove(g4.Red, 7),
				g4.TokenMove(g4.Red, 8),
				g4.TiltMove(g4.LEFT),
				g4.TiltMove(g4.DOWN),
				g4.TiltMove(g4.UP),
			},
		},
		listing: listingArea{
			History: []g4.Move{
				g4.TokenMove(g4.Yellow, 4),
				g4.TokenMove(g4.Red, 5),
				g4.TokenMove(g4.Yellow, 4),
				g4.TokenMove(g4.Red, 4),
				g4.TiltMove(g4.RIGHT),
			},
		},
		game: game,
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

	// When a move is played.
	case receivedMove:
		move := g4.Move(msg)
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
	}
	return m, nil
}

func (m mainModel) View() string {
	var s string
	s += lipgloss.JoinHorizontal(
		lipgloss.Top,
		sidePannelStyle.Render(m.listing.View()),
		playAreaStyle.Render(m.board.View()),
		sidePannelStyle.Render(m.selector.View()),
	)
	return s
}
